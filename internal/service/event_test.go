package service

import (
	"context"
	"encoding/json"
	"sync"
	"testing"
	"time"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/flexprice/flexprice/internal/api/dto"
	"github.com/flexprice/flexprice/internal/domain/events"
	"github.com/flexprice/flexprice/internal/domain/meter"
	"github.com/flexprice/flexprice/internal/testutil"
	"github.com/flexprice/flexprice/internal/types"
	"github.com/stretchr/testify/suite"
)

type EventServiceSuite struct {
	suite.Suite
	ctx        context.Context
	service    *eventService
	store      *testutil.InMemoryEventStore
	broker     *testutil.InMemoryMessageBroker
	msgChannel chan *message.Message
}

func TestEventService(t *testing.T) {
	suite.Run(t, new(EventServiceSuite))
}

func (s *EventServiceSuite) SetupTest() {
	s.ctx = testutil.SetupContext()
	s.store = testutil.NewInMemoryEventStore()
	s.broker = testutil.NewInMemoryMessageBroker()
	s.service = NewEventService(s.broker, s.store, nil).(*eventService)

	// Setup message consumer
	s.msgChannel = s.broker.Subscribe()

	// Start consuming messages
	go func() {
		for msg := range s.msgChannel {
			var event events.Event
			if err := json.Unmarshal(msg.Payload, &event); err != nil {
				continue
			}
			_ = s.store.InsertEvent(s.ctx, &event)
		}
	}()
}

func (s *EventServiceSuite) TearDownTest() {
	s.broker.Close()
}

func (s *EventServiceSuite) TestCreateEvent() {
	testCases := []struct {
		name          string
		input         *dto.IngestEventRequest
		expectedError bool
		verify        func(wg *sync.WaitGroup)
	}{
		{
			name: "successful_event_creation",
			input: &dto.IngestEventRequest{
				EventID:            "test-1",
				ExternalCustomerID: "customer-1",
				EventName:          "api_request",
				Timestamp:          time.Now(),
				Properties: map[string]interface{}{
					"duration_ms": 150,
				},
			},
			expectedError: false,
			verify: func(wg *sync.WaitGroup) {
				wg.Add(1)
				go func() {
					defer wg.Done()
					time.Sleep(100 * time.Millisecond)
					s.True(s.broker.HasMessage("events", "test-1"))
					s.True(s.store.HasEvent("test-1"))
				}()
			},
		},
		{
			name: "missing_required_fields",
			input: &dto.IngestEventRequest{
				EventID: "test-2",
			},
			expectedError: true,
			verify: func(wg *sync.WaitGroup) {
				s.False(s.store.HasEvent("test-2"))
				s.False(s.broker.HasMessage("events", "test-2"))
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			var wg sync.WaitGroup

			err := s.service.CreateEvent(s.ctx, tc.input)
			if tc.expectedError {
				s.Error(err)
			} else {
				s.NoError(err)
			}

			tc.verify(&wg)
			wg.Wait()
		})
	}
}

func (s *EventServiceSuite) TestGetUsage() {
	// Setup test data
	testingEvents := []*dto.IngestEventRequest{
		{
			EventID:            "evt-1",
			ExternalCustomerID: "cust-1",
			EventName:          "api_request",
			Timestamp:          time.Now().Add(-1 * time.Hour),
			Properties: map[string]interface{}{
				"duration_ms": float64(100),
			},
		},
		{
			EventID:            "evt-2",
			ExternalCustomerID: "cust-1",
			EventName:          "api_request",
			Timestamp:          time.Now().Add(-30 * time.Minute),
			Properties: map[string]interface{}{
				"duration_ms": float64(150),
			},
		},
	}

	// Insert test events directly into store
	for _, evt := range testingEvents {
		event := events.NewEvent(
			evt.EventName,
			types.GetTenantID(s.ctx),
			evt.ExternalCustomerID,
			evt.Properties,
			evt.Timestamp,
			evt.EventID,
			evt.CustomerID,
			evt.Source,
		)
		err := s.store.InsertEvent(s.ctx, event)
		s.NoError(err)
	}

	testCases := []struct {
		name          string
		request       *dto.GetUsageRequest
		expectedValue float64
		expectedError bool
	}{
		{
			name: "count_all_events",
			request: &dto.GetUsageRequest{
				ExternalCustomerID: "cust-1",
				EventName:          "api_request",
				AggregationType:    string(types.AggregationCount),
				StartTime:          time.Now().Add(-2 * time.Hour),
				EndTime:            time.Now(),
			},
			expectedValue: 2,
			expectedError: false,
		},
		{
			name: "sum_duration_with_window_size",
			request: &dto.GetUsageRequest{
				ExternalCustomerID: "cust-1",
				EventName:          "api_request",
				PropertyName:       "duration_ms",
				AggregationType:    string(types.AggregationSum),
				StartTime:          time.Now().Add(-2 * time.Hour),
				EndTime:            time.Now(),
				WindowSize:         "HOUR",
			},
			expectedValue: 250,
			expectedError: false,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			result, err := s.service.GetUsage(s.ctx, tc.request)
			if tc.expectedError {
				s.Error(err)
				return
			}
			s.NoError(err)
			s.Equal(tc.expectedValue, result.Value)
		})
	}
}

func (s *EventServiceSuite) TestGetUsageByMeter() {
	// Setup test data for meter
	testMeter := &meter.Meter{
		ID:        "meter-1",
		TenantID:  "tenant-1",
		EventName: "api_request",
		Aggregation: meter.Aggregation{
			Type:  types.AggregationSum,
			Field: "duration_ms",
		},
		WindowSize: "DAY",
	}

	// Add the test meter to a mocked meter repository
	mockedMeterRepo := testutil.NewInMemoryMeterStore()
	err := mockedMeterRepo.CreateMeter(s.ctx, testMeter)
	s.NoError(err)

	// Setup the event service with the mocked meter repository
	s.service = NewEventService(s.broker, s.store, mockedMeterRepo).(*eventService)

	// Setup test events
	testingEvents := []*dto.IngestEventRequest{
		{
			EventID:            "evt-1",
			ExternalCustomerID: "cust-1",
			EventName:          "api_request",
			Timestamp:          time.Now().Add(-1 * time.Hour),
			Properties: map[string]interface{}{
				"duration_ms": float64(100),
			},
		},
	}

	for _, evt := range testingEvents {
		event := events.NewEvent(
			evt.EventName,
			types.GetTenantID(s.ctx),
			evt.ExternalCustomerID,
			evt.Properties,
			evt.Timestamp,
			evt.EventID,
			evt.CustomerID,
			evt.Source,
		)
		err := s.store.InsertEvent(s.ctx, event)
		s.NoError(err)
	}

	// Test cases for GetUsageByMeter
	testCases := []struct {
		name          string
		request       *dto.GetUsageByMeterRequest
		expectedValue float64
		expectedError bool
	}{
		{
			name: "sum_duration_with_window_size",
			request: &dto.GetUsageByMeterRequest{
				MeterID:            "meter-1",
				ExternalCustomerID: "cust-1",
				StartTime:          time.Now().Add(-2 * time.Hour),
				EndTime:            time.Now(),
			},
			expectedValue: 100, // Matches partition window "DAY"
			expectedError: false,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			result, err := s.service.GetUsageByMeter(s.ctx, tc.request)
			if tc.expectedError {
				s.Error(err)
				return
			}
			s.NoError(err)
			s.Equal(tc.expectedValue, result.Value)
		})
	}
}