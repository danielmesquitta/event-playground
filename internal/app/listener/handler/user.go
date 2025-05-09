package handler

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/danielmesquitta/event-playground/internal/domain/entity"
)

type UserCreatedHandler struct {
}

func NewUserCreatedHandler() *UserCreatedHandler {
	return &UserCreatedHandler{}
}

func (h *UserCreatedHandler) Handle(
	ctx context.Context,
	message json.RawMessage,
) error {
	var user entity.User
	if err := json.Unmarshal(message, &user); err != nil {
		return fmt.Errorf("failed to unmarshal user: %w", err)
	}

	fmt.Printf("Received user: %+v\n", user)

	return nil
}

var _ Handler = &UserCreatedHandler{}
