package store

import (
	"sync"
)

// TrackerStore holds the tracking IDs and their respective open counts
type TrackerStore struct {
	mu          sync.Mutex
	trackCounts map[string]int
}

// NewTrackerStore creates a new instance of TrackerStore
func NewTrackerStore() *TrackerStore {
	return &TrackerStore{
		trackCounts: make(map[string]int),
	}
}

// AddTrackingID adds a new tracking ID to the store
func (ts *TrackerStore) AddTrackingID(trackingID string) {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	// Ensure the tracking ID is initialized with 0 open count
	if _, exists := ts.trackCounts[trackingID]; !exists {
		ts.trackCounts[trackingID] = 0
	}
}

// IncrementOpenCount increments the open count for a given tracking ID
func (ts *TrackerStore) IncrementOpenCount(trackingID string) {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	// Check if tracking ID exists, if not initialize it
	if _, exists := ts.trackCounts[trackingID]; exists {
		ts.trackCounts[trackingID]++
	}
}

// GetOpenCount retrieves the open count for a given tracking ID
func (ts *TrackerStore) GetOpenCount(trackingID string) int {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	// Return the open count or 0 if tracking ID doesn't exist
	return ts.trackCounts[trackingID]
}
