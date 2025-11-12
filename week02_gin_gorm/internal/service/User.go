package service

import (
	"context"
	"myGoLearn/week02_gin_gorm/internal/domain"
	"myGoLearn/week02_gin_gorm/internal/repository"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	rep *repository.UserRepository
}

var (
	EmailDuplicateError    = repository.EmailDuplicateError
	InvalidEmailOrPassword = repository.UserNotFoundError
)

func (us *UserService) Signup(ctx *gin.Context, u *domain.User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hash)

	return us.rep.Create(ctx, u)
}

func (us *UserService) Login(ctx context.Context, email string, password string) (domain.User, error) {
	u, err := us.rep.FindByEmail(ctx, email)
	if err == repository.UserNotFoundError {
		return domain.User{}, InvalidEmailOrPassword
	}
	if err != nil {
		return domain.User{}, err
	}
	//校验密码
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return domain.User{}, InvalidEmailOrPassword
	}
	return u, nil
}

func (us *UserService) Edit(ctx context.Context, u *domain.User) error {
	return us.rep.UpdateById(ctx, u)
}

func NewUserService(repository *repository.UserRepository) *UserService {
	return &UserService{
		rep: repository,
	}
}
