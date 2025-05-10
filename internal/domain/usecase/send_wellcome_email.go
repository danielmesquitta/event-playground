package usecase

import (
	"context"
	"fmt"

	"github.com/danielmesquitta/event-playground/internal/domain/entity"
)

type SendWelcomeEmail struct {
}

func NewSendWelcomeEmail() *SendWelcomeEmail {
	return &SendWelcomeEmail{}
}

func (u *SendWelcomeEmail) Execute(
	ctx context.Context,
	user *entity.User,
) error {
	fmt.Printf("Sending welcome email to %s\n", user.Email)

	return nil
}
