package storage

import (
	"errors"
	"lab03-backend/models"
	"sync"
)

// MemoryStorage implements in-memory storage for messages
type MemoryStorage struct {
	// TODO: Add mutex field for thread safety (sync.RWMutex)
	// TODO: Add messages field as map[int]*models.Message
	// TODO: Add nextID field of type int for auto-incrementing IDs
	mutex    sync.RWMutex
	messages map[int]*models.Message
	nextID   int
}

// NewMemoryStorage creates a new in-memory storage instance
func NewMemoryStorage() *MemoryStorage {
	// TODO: Return a new MemoryStorage instance with initialized fields
	// Initialize messages as empty map
	// Set nextID to 1
	memoryStorage := &MemoryStorage{
		messages: make(map[int]*models.Message),
		nextID:   1,
	}

	return memoryStorage
}

// GetAll returns all messages
func (ms *MemoryStorage) GetAll() []*models.Message {
	// TODO: Implement GetAll method
	// Use read lock for thread safety
	// Convert map values to slice
	// Return slice of all messages
	ms.mutex.RLock()

	messages := make([]*models.Message, 0, len(ms.messages))
	for _, v := range ms.messages {
		messages = append(messages, v)
	}

	ms.mutex.RUnlock()
	return messages
}

// GetByID returns a message by its ID
func (ms *MemoryStorage) GetByID(id int) (*models.Message, error) {
	// TODO: Implement GetByID method
	// Use read lock for thread safety
	// Check if message exists in map
	// Return message or error if not found
	ms.mutex.RLock()

	message, ok := ms.messages[id]

	ms.mutex.RUnlock()

	if !ok {
		return nil, ErrMessageNotFound
	}

	return message, nil
}

// Create adds a new message to storage
func (ms *MemoryStorage) Create(username, content string) (*models.Message, error) {
	// TODO: Implement Create method
	// Use write lock for thread safety
	// Get next available ID
	// Create new message using models.NewMessage
	// Add message to map
	// Increment nextID
	// Return created message
	ms.mutex.Lock()

	nextID := ms.nextID
	newMessage := models.NewMessage(nextID, username, content)

	if !newMessage {
		return nil, ErrMessageNotFound
	}

	ms.messages[nextID] = newMessage

	ms.nextID++

	ms.mutex.Unlock()
	return newMessage, nil
}

// Update modifies an existing message
func (ms *MemoryStorage) Update(id int, content string) (*models.Message, error) {
	// TODO: Implement Update method
	// Use write lock for thread safety
	// Check if message exists
	// Update the content field
	// Return updated message or error if not found
	ms.mutex.Lock()

	if id <= 0 {
		return nil, ErrInvalidID
	}

	if _, ok := ms.messages[id]; !ok {
		return nil, ErrMessageNotFound
	}

	ms.messages[id].Content = content

	ms.mutex.Unlock()
	return ms.messages[id], nil
}

// Delete removes a message from storage
func (ms *MemoryStorage) Delete(id int) error {
	// TODO: Implement Delete method
	// Use write lock for thread safety
	// Check if message exists
	// Delete from map
	// Return error if message not found
	ms.mutex.Lock()

	if id <= 0 {
		return ErrInvalidID
	}

	if _, ok := ms.messages[id]; !ok {
		return ErrMessageNotFound
	}

	delete(ms.messages, id)

	ms.mutex.Unlock()
	return nil
}

// Count returns the total number of messages
func (ms *MemoryStorage) Count() int {
	// TODO: Implement Count method
	// Use read lock for thread safety
	// Return length of messages map
	ms.mutex.RLock()

	counter := len(ms.messages)

	ms.mutex.RUnlock()
	return counter
}

// Common errors
var (
	ErrMessageNotFound = errors.New("message not found")
	ErrInvalidID       = errors.New("invalid message ID")
)
