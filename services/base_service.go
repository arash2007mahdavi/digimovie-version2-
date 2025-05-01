package services

import (
	"context"
	"database/sql"
	"digimovie/src/common"
	"digimovie/src/database"
	"digimovie/src/logging"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type BaseService[T any, Tc any, Tu any, Tr any] struct {
	Database *gorm.DB
	Logger logging.Logger
}

func NewBaseService[T any, Tc any, Tu any, Tr any]() *BaseService[T, Tc, Tu, Tr] {
	db := database.GetDB()
	if db == nil {
		panic("database connection is not initialized")
	}
	return &BaseService[T, Tc, Tu, Tr]{
		Database: db,
		Logger: logging.NewLogger(),
	}
}

func (s *BaseService[T, Tc, Tu, Tr]) Create(ctx context.Context, req *Tc) (*Tr, error) {
	model, err:= common.TypeComverter[T](req)
	if err != nil{
		return nil, fmt.Errorf("error in Comvertering models")
	}
	if s.Database == nil {
		panic("database is nil")
	}
	tx := s.Database.WithContext(ctx).Begin()
	err = tx.Create(model).Error
	if err != nil {
		tx.Rollback()
		s.Logger.Error(logging.Postgres, logging.Insert, "error in Creating new user", nil)
		return nil, fmt.Errorf("error in Creating new user")
	}
	tx.Commit()
	return common.TypeComverter[Tr](model)
}

func (s *BaseService[T, Tc, Tu, Tr]) Update(ctx context.Context, id int, req *Tu) (*Tr, error) {
	updateMap, err := common.TypeComverter[map[string]interface{}](req)
	if err != nil{
		return nil, fmt.Errorf("error in Comvertering models")
	}
	(*updateMap)["modified_by"] = &sql.NullInt64{Int64: int64(ctx.Value("Userid").(float64)), Valid: true}
	(*updateMap)["modified_at"] = sql.NullTime{Time: time.Now().UTC(), Valid: true}
	model := new(T)
	tx := s.Database.WithContext(ctx).Begin()
	err = tx.Model(model).
		Where("id = ? AND deleted_by is null", id).
		Updates(*updateMap).
		Error
	if err != nil {
		tx.Rollback()
		s.Logger.Error(logging.Postgres, logging.Insert, "error in Updating user", nil)
		return nil, fmt.Errorf("error in Updating user")
	}
	tx.Commit()
	res, _:= s.GetById(ctx, id)
	return res, nil
}

func (s *BaseService[T, Tc, Tu, Tr]) Delete(ctx context.Context, id int) (error) {
	deleteMap := map[string]interface{}{}
	(deleteMap)["deleted_by"] = &sql.NullInt64{Int64: int64(ctx.Value("Userid").(float64)), Valid: true}
	(deleteMap)["deleted_at"] = sql.NullTime{Time: time.Now().UTC(), Valid: true}
	if ctx.Value("Userid") == nil {
		return fmt.Errorf("deleter Userid error")
	}
	model := new(T)
	tx := s.Database.WithContext(ctx).Begin()
	cnt := tx.Model(model).
		Where("id = ? AND deleted_by is null", id).
		Updates(deleteMap).RowsAffected
	if cnt == 0 {
		tx.Rollback()
		s.Logger.Error(logging.Postgres, logging.Insert, "this id doesnt exists", nil)
		return fmt.Errorf("this id doesnt exists")
	}
	tx.Commit()
	return nil
}

func (s *BaseService[T, Tc, Tu, Tr]) GetById(ctx context.Context, id int) (*Tr, error) {
	model := new(T)
	s.Database.Model(model).
		Where("id = ? AND deleted_by is null", id).
		First(model)
	return common.TypeComverter[Tr](model)
}