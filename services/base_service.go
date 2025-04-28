package services

import (
	"context"
	"digimovie/src/common"
	"digimovie/src/database"
	"digimovie/src/logging"

	"gorm.io/gorm"
)

type BaseService[T any, Tc any, Tu any, Tr any] struct {
	Database *gorm.DB
	Logger logging.Logger
}

func NewBaseService[T any, Tc any, Tu any, Tr any]() *BaseService[T, Tc, Tu, Tr] {
	return &BaseService[T, Tc, Tu, Tr]{
		Database: database.GetDB(),
		Logger: logging.NewLogger(),
	}
}

func (s *BaseService[T, Tc, Tu, Tr]) Create(ctx context.Context, req *Tc) (*Tr, error) {
	model, _:= common.TypeComverter[T](req)
	tx := s.Database.WithContext(ctx).Begin()
	err := tx.Create(model).Error
	if err != nil {
		tx.Rollback()
		s.Logger.Error(logging.Postgres, logging.Insert, err.Error(), nil)
		return nil, err
	}
	tx.Commit()
	return common.TypeComverter[Tr](model)
}