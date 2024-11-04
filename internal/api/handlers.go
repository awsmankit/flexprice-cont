package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/flexprice/flexprice/internal/kafka"
	"github.com/flexprice/flexprice/internal/logger"
	"github.com/flexprice/flexprice/internal/meter"
	"github.com/flexprice/flexprice/internal/storage"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	producer *kafka.Producer
	registry *meter.MeterRegistry
	storage  *storage.ClickHouseStorage
}

func NewHandler(producer *kafka.Producer, registry *meter.MeterRegistry, storage *storage.ClickHouseStorage) *Handler {
	return &Handler{
		producer: producer,
		registry: registry,
		storage:  storage,
	}
}

func (h *Handler) IngestEvent(c *gin.Context) {
	var event storage.Event
	if err := c.ShouldBindJSON(&event); err != nil {
		logger.Log.Error("Failed to bind JSON", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// if _, ok := h.registry.GetMeter(event.FeatureID); !ok {
	// 	logger.Log.Error("Invalid feature ID", "feature_id", event.FeatureID)
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid feature ID"})
	// 	return
	// }

	payload, err := json.Marshal(event)
	if err != nil {
		logger.Log.Error("Failed to marshal event", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process event"})
		return
	}

	if err := h.producer.PublishWithID("events", payload, event.ID); err != nil {
		logger.Log.Error("Failed to publish event", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process event"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "Event accepted for processing"})
}

func (h *Handler) GetUsage(c *gin.Context) {
	customerID := c.Query("customer_id")
	featureID := c.Query("feature_id")
	startTimeStr := c.Query("start_time")
	endTimeStr := c.Query("end_time")

	if customerID == "" || featureID == "" || startTimeStr == "" || endTimeStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required parameters"})
		return
	}

	startTime, err := time.Parse(time.RFC3339, startTimeStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start_time format"})
		return
	}

	endTime, err := time.Parse(time.RFC3339, endTimeStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end_time format"})
		return
	}

	meter, ok := h.registry.GetMeter(featureID)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid feature ID"})
		return
	}

	usage, err := h.storage.GetUsage(c, customerID, featureID, startTime, endTime, meter.AggregationType)
	if err != nil {
		logger.Log.Error("Failed to get usage", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get usage"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"usage": usage})
}
