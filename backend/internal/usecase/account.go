package usecase

import (
	"context"

	"web-server-template/internal/global/utils/log"
	"web-server-template/internal/service/dao"
	"web-server-template/internal/service/orm/model"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AccountUsecase struct {
	db       *gorm.DB
	dao      *dao.Dao
	validate *validator.Validate
}

func NewAccountUsecase(
	dao *dao.Dao,
	db *gorm.DB,
	validate *validator.Validate,
) *AccountUsecase {
	return &AccountUsecase{
		dao:      dao,
		db:       db,
		validate: validate,
	}
}

type (
	CreateAccountIn struct {
		UserName    string `validate:"required"`
		Password    string `validate:"required"`
		DisplayName string `validate:"required"`
	}
)

func (u *AccountUsecase) CreateAccount(
	ctx context.Context,
	in *CreateAccountIn,
) (err error) {
	if err = u.validate.Struct(in); err != nil {
		return log.WithTrace(err)
	}

	u.db.Create(&model.Account{
		UserName:    in.UserName,
		Password:    in.Password,
		DisplayName: in.DisplayName,
	})
	return nil
}

type UpdateAccountIn struct {
	DisplayName string `validate:"required"`
}

func (u *AccountUsecase) UpdateAccount(
	ctx context.Context,
	in *UpdateAccountIn,
) (err error) {
	if err = u.validate.Struct(in); err != nil {
		return log.WithTrace(err)
	}
	return nil
}

type DeleteAccountIn struct {
	AccountID uuid.UUID `validate:"required"`
}

func (u *AccountUsecase) DeleteAccount(
	ctx context.Context,
	in *DeleteAccountIn,
) (err error) {
	if err = u.validate.Struct(in); err != nil {
		return log.WithTrace(err)
	}
	return nil
}

type DeActiveAccountIn struct {
	AccountID uuid.UUID `validate:"required"`
}

func (u *AccountUsecase) DeActiveAccount(
	ctx context.Context,
	in *DeActiveAccountIn,
) (err error) {
	if err = u.validate.Struct(in); err != nil {
		return log.WithTrace(err)
	}
	return nil
}

type ActiveAccountIn struct {
	AccountID uuid.UUID `validate:"required"`
}

func (u *AccountUsecase) ActiveAccount(
	ctx context.Context,
	in *ActiveAccountIn,
) (err error) {
	if err = u.validate.Struct(in); err != nil {
		return log.WithTrace(err)
	}
	return nil
}
