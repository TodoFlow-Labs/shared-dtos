package dto

import "time"

// ----------------------
// Commands
// ----------------------

type CommandType string

const (
	CreateTodoCmd CommandType = "CreateTodo"
	UpdateTodoCmd CommandType = "UpdateTodo"
	DeleteTodoCmd CommandType = "DeleteTodo"
)

type BaseCommand struct {
	Type   CommandType `json:"type"`
	ID     string      `json:"id,omitempty"`
	UserID string      `json:"user_id"`
}

type CreateTodoCommand struct {
	BaseCommand
	Title       string     `json:"title"`
	Description string     `json:"description,omitempty"`
	DueDate     *time.Time `json:"due_date,omitempty"`
	Priority    *int       `json:"priority,omitempty"` // 1 = High, 2 = Medium, 3 = Low
	Tags        []string   `json:"tags,omitempty"`
}

type UpdateTodoCommand struct {
	BaseCommand
	Title       *string    `json:"title,omitempty"`
	Description *string    `json:"description,omitempty"`
	Completed   *bool      `json:"completed,omitempty"`
	DueDate     *time.Time `json:"due_date,omitempty"`
	Priority    *int       `json:"priority,omitempty"`
	Tags        *[]string  `json:"tags,omitempty"`
}

type DeleteTodoCommand struct {
	BaseCommand
}

// ----------------------
// Events
// ----------------------

type EventType string

const (
	TodoCreatedEvt EventType = "TodoCreated"
	TodoUpdatedEvt EventType = "TodoUpdated"
	TodoDeletedEvt EventType = "TodoDeleted"
)

type BaseEvent struct {
	Type      EventType `json:"type"`
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Timestamp time.Time `json:"timestamp"`
}

type TodoCreatedEvent struct {
	BaseEvent
	Title       string     `json:"title"`
	Description string     `json:"description,omitempty"`
	DueDate     *time.Time `json:"due_date,omitempty"`
	Priority    *int       `json:"priority,omitempty"`
	Tags        []string   `json:"tags,omitempty"`
}

type TodoUpdatedEvent struct {
	BaseEvent
	Title       string     `json:"title"`
	Description string     `json:"description,omitempty"`
	Completed   bool       `json:"completed"`
	DueDate     *time.Time `json:"due_date,omitempty"`
	Priority    *int       `json:"priority,omitempty"`
	Tags        []string   `json:"tags,omitempty"`
}

type TodoDeletedEvent struct {
	BaseEvent
}

// ----------------------
// Queries & Responses
// ----------------------

type SearchRequest struct {
	Query string `json:"query"`
}

type SearchResult struct {
	ID          string     `json:"id"`
	UserID      string     `json:"user_id"`
	Title       string     `json:"title"`
	Description string     `json:"description,omitempty"`
	Completed   bool       `json:"completed"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DueDate     *time.Time `json:"due_date,omitempty"`
	Priority    *int       `json:"priority,omitempty"`
	Tags        []string   `json:"tags,omitempty"`
}

type ListTodosResponse struct {
	Todos []SearchResult `json:"todos"`
}

// ----------------------
// Notification
// ----------------------

type Notification struct {
	EventType EventType   `json:"event_type"`
	Payload   interface{} `json:"payload"`
}
