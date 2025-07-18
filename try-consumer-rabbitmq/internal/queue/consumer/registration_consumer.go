package consumer

import (
	"context"

	"github.com/rahmatrdn/go-skeleton/entity"
	"github.com/rahmatrdn/go-skeleton/internal/helper"
)

type RegistrationQueue struct {
	ctx context.Context
}

type RegistrationConsumer interface {
	Process(payload map[string]interface{}) error
}

func NewRegistrationConsumer(
	ctx context.Context,
) RegistrationConsumer {
	return &RegistrationQueue{ctx}
}

func (l *RegistrationQueue) Process(payload map[string]interface{}) error {
	var params entity.Log
	params.LoadFromMap(payload)

	helper.Dump(params)

	return nil
}
