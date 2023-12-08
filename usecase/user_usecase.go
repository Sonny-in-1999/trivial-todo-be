package usecase

import (
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
	"trivial-todo-be/model"
	"trivial-todo-be/repository"
	"trivial-todo-be/validator"
)

type IUserUsecase interface {
	SignUp(user model.User) (model.UserResponse, error)
	Login(user model.User) (string, error)
}

type userUsecase struct {
	ur repository.IUserRepository
	uv validator.IUserValidator
}

func NewUserUsecase(ur repository.IUserRepository, uv validator.IUserValidator) IUserUsecase {
	return &userUsecase{ur, uv}
}

func (uu *userUsecase) SignUp(user model.User) (model.UserResponse, error) {
	if err := uu.uv.UserValidate(user); err != nil {
		return model.UserResponse{}, err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10) // 10 Rounds Salted Password Returned
	if err != nil {
		return model.UserResponse{}, err
	}

	newUser := model.User{Email: user.Email, Password: string(hash)}
	err = uu.ur.CreateUser(&newUser)
	if err != nil {
		return model.UserResponse{}, err
	}

	userRes := model.UserResponse{ID: newUser.ID, Email: newUser.Email}
	return userRes, nil
}

func (uu *userUsecase) Login(user model.User) (string, error) { // string: jwt 토큰 반환
	if err := uu.uv.UserValidate(user); err != nil {
		return "", err
	}

	existUser := model.User{}
	if err := uu.ur.GetUserByEmail(&existUser, user.Email); err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(existUser.Password), []byte(user.Password)); err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": existUser.ID,
		"exp":     time.Now().Add(time.Hour * 12).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
