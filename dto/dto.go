// Package dto defines shared Command and Event types used across 
// all services (command-service, domain-processor, projection-worker, query-service).
package dto

import "time"

// CommandType enumerates the types of commands supported.
type CommandType string

const (
	// CreateTodoCmd indicates a request to create a new Todo.
	CreateTodoCmd CommandType = "CreateTodo"
	// UpdateTodoCmd indicates a request to update an existing Todo.
	UpdateTodoCmd CommandType = "UpdateTodo"
	// DeleteTodoCmd indicates a request to delete a Todo.
	DeleteTodoCmd CommandType = "DeleteTodo"
)

// BaseCommand contains common fields for all commands.
type BaseCommand struct {
	Type CommandType `json:"type"`
	// Unique identifier for the Todo resource (empty for Create).
	ID string `json:"id,omitempty"`
}

// CreateTodoCommand carries the data needed to create a new Todo.
type CreateTodoCommand struct {
	BaseCommand
	Title string `json:"title"`
}

// UpdateTodoCommand carries data needed to update an existing Todo.
type UpdateTodoCommand struct {
	BaseCommand
	// Nullable fields allow partial updates.
	Title     *string `json:"title,omitempty"`
	Completed *bool   `json:"completed,omitempty"`
}

// DeleteTodoCommand carries the ID of the Todo to delete.
type DeleteTodoCommand struct {
	BaseCommand
}

// EventType enumerates the types of domain events emitted.
type EventType string

const (
	// TodoCreatedEvt signals a new Todo was created.
	TodoCreatedEvt EventType = "TodoCreated"
	// TodoUpdatedEvt signals an existing Todo was updated.
	TodoUpdatedEvt EventType = "TodoUpdated"
	// TodoDeletedEvt signals a Todo was deleted.
	TodoDeletedEvt EventType = "TodoDeleted"
)

// BaseEvent contains common fields for all events.
type BaseEvent struct {
	Type      EventType `json:"type"`
	ID        string    `json:"id"`
	Timestamp time.Time `json:"timestamp"`
}

// TodoCreatedEvent is emitted after a Todo is created.
type TodoCreatedEvent struct {
	BaseEvent
	// Title of the created Todo
	Title string `json:"title"`
}

// TodoUpdatedEvent is emitted after a Todo is updated.
type TodoUpdatedEvent struct {
	BaseEvent
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// TodoDeletedEvent is emitted after a Todo is deleted.
type TodoDeletedEvent struct {
	BaseEvent
}

// SearchRequest represents a client query for Todos.
type SearchRequest struct {
	Query string `json:"query"`
}

// SearchResult represents a single Todo returned by a search.
type SearchResult struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}

// ListTodosResponse is returned by the query-service for GET /todos.
type ListTodosResponse struct {
	Todos []SearchResult `json:"todos"`
}

// Notification carries asynchronous notifications to the frontend via SSE or WS.
type Notification struct {
	EventType EventType   `json:"event_type"`
	Payload   interface{} `json:"payload"`
}
