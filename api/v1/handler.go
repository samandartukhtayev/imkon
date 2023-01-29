package v1

import (
	"errors"

	"github.com/samandartukhtayev/imkon/api/models"
	"github.com/samandartukhtayev/imkon/config"
	"github.com/samandartukhtayev/imkon/storage"
)

type handlerV1 struct {
	cfg *config.Config

	storage  storage.StorageI
	inMemory storage.InMemoryStorageI
}

var (
	ErrForbidden = errors.New("forbidden")
)

type HandlerV1Options struct {
	Cfg      *config.Config
	Storage  storage.StorageI
	InMemory storage.InMemoryStorageI
}

func New(options *HandlerV1Options) *handlerV1 {
	return &handlerV1{
		cfg:      options.Cfg,
		storage:  options.Storage,
		inMemory: options.InMemory,
	}
}

func ErrorResponse(err error) *models.ErrorResponse {
	return &models.ErrorResponse{
		Error: err.Error(),
	}
}
