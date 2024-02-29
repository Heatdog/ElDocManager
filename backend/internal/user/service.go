package user

import (
	cryptohash "ElDocManager/pkg/cryptoHash"
	"ElDocManager/pkg/logging"
	"context"
)

type UserService interface {
	SignIn(context context.Context, userLogin *UserSignIn) string
	SignUp(context context.Context, userLogin *UserSignUp) error
	SignOut(context context.Context)
}

func NewUserService(logger *logging.Logger, repo UserRepository) UserService {
	return &userService{
		logger: logger,
		repo:   repo,
	}
}

type userService struct {
	logger *logging.Logger
	repo   UserRepository
}

func (s *userService) SignIn(context context.Context, userLogin *UserSignIn) string {
	return ""
}

func (s *userService) SignUp(context context.Context, userLogin *UserSignUp) error {
	s.logger.Info("sign up service started")
	hashedPassword, err := cryptohash.Hash(userLogin.Password)
	if err != nil {
		s.logger.Errorf("hash error %s", err.Error())
	}
	user := &User{
		Login:    userLogin.Login,
		Name:     userLogin.Name,
		Surname:  userLogin.Surname,
		Email:    userLogin.Email,
		Password: hashedPassword,
	}
	err = s.repo.Create(context, user)
	s.logger.Info("sign up service ended")
	return err
}

func (s *userService) SignOut(context context.Context) {

}
