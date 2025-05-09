package handler

import (
	"context"
	"encoding/json"
)

type Handler interface {
	Handle(ctx context.Context, message json.RawMessage) error
}
