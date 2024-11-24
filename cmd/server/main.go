package main

import (
	"context"
	"encoding/json"
	"time"

	"go.uber.org/fx"

	lambdaEvents "github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	_ "github.com/flexprice/flexprice/docs/swagger"
	"github.com/flexprice/flexprice/internal/api"
	v1 "github.com/flexprice/flexprice/internal/api/v1"
	"github.com/flexprice/flexprice/internal/clickhouse"
	"github.com/flexprice/flexprice/internal/config"
	"github.com/flexprice/flexprice/internal/domain/events"
	"github.com/flexprice/flexprice/internal/kafka"
	"github.com/flexprice/flexprice/internal/logger"
	"github.com/flexprice/flexprice/internal/postgres"
	"github.com/flexprice/flexprice/internal/repository"
	"github.com/flexprice/flexprice/internal/service"
	"github.com/flexprice/flexprice/internal/types"
	"github.com/gin-gonic/gin"
)

// @title FlexPrice API
// @version 1.0
// @description FlexPrice API Service
// @host localhost:8080
// @BasePath /v1
// @schemes http https
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func init() {
	// Set UTC timezone for the entire application
	time.Local = time.UTC
}

func main() {
	var opts []fx.Option
	opts = append(opts,
		fx.Provide(
			// Config
			config.NewConfig,

			// Logger
			logger.NewLogger,

			// DB
			postgres.NewDB,
			clickhouse.NewClickHouseStore,

			// Producers and Consumers
			kafka.NewProducer,
			kafka.NewConsumer,

			// Repositories
			repository.NewEventRepository,
			repository.NewMeterRepository,

			// Services
			service.NewMeterService,
			service.NewEventService,

			// Handlers
			provideHandlers,

			// Router
			provideRouter,
		),
		fx.Invoke(startServer),
	)

	app := fx.New(opts...)
	app.Run()
}

func provideHandlers(
	logger *logger.Logger,
	meterService service.MeterService,
	eventService service.EventService,
) api.Handlers {
	return api.Handlers{
		Events: v1.NewEventsHandler(eventService, logger),
		Meter:  v1.NewMeterHandler(meterService, logger),
	}
}

func provideRouter(handlers api.Handlers) *gin.Engine {
	return api.NewRouter(handlers)
}

func startServer(
	lc fx.Lifecycle,
	cfg *config.Configuration,
	r *gin.Engine,
	consumer kafka.MessageConsumer,
	eventRepo events.Repository,
	log *logger.Logger,
) {
	mode := cfg.Deployment.Mode

	switch mode {
	case types.ModeLocal:
		if consumer == nil {
			log.Fatal("Kafka consumer required for local mode")
		}
		startAPIServer(lc, r, cfg, log)
		startConsumer(lc, consumer, eventRepo, cfg, log)
	case types.ModeAPI:
		startAPIServer(lc, r, cfg, log)
	case types.ModeConsumer:
		if consumer == nil {
			log.Fatal("Kafka consumer required for consumer mode")
		}
		startConsumer(lc, consumer, eventRepo, cfg, log)
	case types.ModeAWSLambdaAPI:
		startAWSLambdaAPI(r)
	case types.ModeAWSLambdaConsumer:
		startAWSLambdaConsumer(eventRepo, log)
	default:
		log.Fatalf("Unknown deployment mode: %s", mode)
	}
}

func startAPIServer(
	lc fx.Lifecycle,
	r *gin.Engine,
	cfg *config.Configuration,
	log *logger.Logger,
) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := r.Run(cfg.Server.Address); err != nil {
					log.Fatalf("Failed to start server: %v", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Info("Shutting down server...")
			return nil
		},
	})
}

func startConsumer(
	lc fx.Lifecycle,
	consumer kafka.MessageConsumer,
	eventRepo events.Repository,
	cfg *config.Configuration,
	log *logger.Logger,
) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go consumeMessages(consumer, eventRepo, cfg.Kafka.Topic, log)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Info("Shutting down consumer...")
			return nil
		},
	})
}

func startAWSLambdaAPI(r *gin.Engine) {
	ginLambda := ginadapter.New(r)
	lambda.Start(ginLambda.ProxyWithContext)
}

func startAWSLambdaConsumer(eventRepo events.Repository, log *logger.Logger) {
	handler := func(ctx context.Context, kafkaEvent lambdaEvents.KafkaEvent) error {
		for _, record := range kafkaEvent.Records {
			for _, r := range record {
				log.Debugf("Processing record: topic=%s, partition=%d, offset=%d",
					r.Topic, r.Partition, r.Offset)

				// TODO decide the repository to use based on the event topic and properties
				// For now we will use the event repository from the events topic

				var event events.Event
				if err := json.Unmarshal([]byte(r.Value), &event); err != nil {
					log.Errorf("Failed to unmarshal event: %v, payload: %s", err, r.Value)
					continue // Skip invalid messages
				}

				if err := eventRepo.InsertEvent(ctx, &event); err != nil {
					log.Errorf("Failed to insert event: %v, event: %+v", err, event)
					// TODO: Handle error and decide if we should retry or send to DLQ
					continue
				}

				log.Infof("Successfully processed event: topic=%s, partition=%d, offset=%d",
					r.Topic, r.Partition, r.Offset)
			}
		}
		return nil
	}

	lambda.Start(handler)
}

func consumeMessages(consumer kafka.MessageConsumer, eventRepo events.Repository, topic string, log *logger.Logger) {
	messages, err := consumer.Subscribe(topic)
	if err != nil {
		log.Fatalf("Failed to subscribe to topic %s: %v", topic, err)
	}

	for msg := range messages {
		var event events.Event
		if err := json.Unmarshal(msg.Payload, &event); err != nil {
			log.Errorf("Failed to unmarshal event: %v, payload: %s", err, string(msg.Payload))
			msg.Ack() // Acknowledge invalid messages
			continue
		}

		log.Debugf("Starting to process event: %+v", event)

		if err := eventRepo.InsertEvent(context.Background(), &event); err != nil {
			log.Errorf("Failed to insert event: %v, event: %+v", err, event)
			// TODO: Handle error and decide if we should retry or send to DLQ
		}
		msg.Ack()
		log.Debugf("Successfully processed event: %+v", event)
	}
}
