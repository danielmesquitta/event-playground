package handler

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/danielmesquitta/event-playground/internal/domain/entity"
	"github.com/danielmesquitta/event-playground/internal/domain/usecase"
)

type UserCreatedHandler struct {
	sendWelcomeEmailUseCase *usecase.SendWelcomeEmail
}

func NewUserCreatedHandler(
	sendWelcomeEmailUseCase *usecase.SendWelcomeEmail,
) *UserCreatedHandler {
	return &UserCreatedHandler{
		sendWelcomeEmailUseCase: sendWelcomeEmailUseCase,
	}
}

func (h *UserCreatedHandler) Handle(
	ctx context.Context,
	message json.RawMessage,
) error {
	var user entity.User
	if err := json.Unmarshal(message, &user); err != nil {
		return fmt.Errorf("failed to unmarshal user: %w", err)
	}

	return h.sendWelcomeEmailUseCase.Execute(ctx, &user)
}

var _ Handler = &UserCreatedHandler{}
