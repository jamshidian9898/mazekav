package repository

import (
	"context"

	"mazekav/internal/entity"
)

type CommonBehaviourRepository[T entity.DBModel] interface {
	ByID(ctx context.Context, id uint) (T, error)
	ByFiled(ctx context.Context, field string, id uint) (T, error)
	Save(ctx context.Context, model *T) error
	// add more common behaviour
}

type UserRepository interface {
	CommonBehaviourRepository[entity.User]
	// add any spesefic function just for user
}
