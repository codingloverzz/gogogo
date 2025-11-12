package repository

import (
	"context"
	"myGoLearn/week02_gin_gorm/internal/domain"
	"myGoLearn/week02_gin_gorm/internal/repository/dao"

	"github.com/gin-gonic/gin"
)

type UserRepository struct {
	userDao *dao.UserDao
}

var (
	EmailDuplicateError = dao.EmailDuplicateError
	UserNotFoundError   = dao.RecordNotFoundError
)

func (ur *UserRepository) Create(ctx *gin.Context, user *domain.User) error {
	return ur.userDao.Insert(ctx, &dao.User{
		Email:    user.Email,
		Password: user.Password,
	})
}

func (ur *UserRepository) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	u, err := ur.userDao.FindByEmail(ctx, email)
	if err != nil {
		return domain.User{}, err
	}

	return toDomain(u), err
}

func (ur *UserRepository) UpdateById(ctx context.Context, u *domain.User) error {
	return ur.userDao.UpdateById(ctx, &dao.User{
		Id:      u.ID,
		AboutMe: u.AboutMe,
	})
}

func NewUserRepository(userDao *dao.UserDao) *UserRepository {
	return &UserRepository{
		userDao: userDao,
	}
}

func toDomain(user dao.User) domain.User {
	return domain.User{
		Email:    user.Email,
		Password: user.Password,
		ID:       user.Id,
	}
}
