package user

import (
	cryptohash "github.com/Heatdog/ElDocManager/backend/mainServer/pkg/cryptoHash"
	jwt_auth "github.com/Heatdog/ElDocManager/backend/mainServer/pkg/jwtAuth"

	"context"

	"github.com/Heatdog/ElDocManager/backend/mainServer/pkg/logging"
)

type UserService interface {
	SignIn(context context.Context, userLogin *UserSignIn, jwtKey string) (string, error)
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

func (s *userService) SignIn(context context.Context, userLogin *UserSignIn, jwtKey string) (string, error) {
	s.logger.Info("sign in service started")
	user, err := s.repo.FindOneWithLogin(context, userLogin.Login)
	if err != nil {
		s.logger.Infof("user with login %s was not detected", user.Login)
		return "", err
	}

	if !cryptohash.VerifyHash([]byte(user.Password), userLogin.Password) {
		s.logger.Infof("incorrect password for user %s", userLogin.Login)
		return "", err
	}

	token, err := jwt_auth.GenerateToken(jwt_auth.TokenFields{
		ID:   user.ID,
		Role: string(user.Role),
	}, jwtKey)
	if err != nil {
		s.logger.Errorf("token generation error: %s", err.Error())
		return "", err
	}

	_, err = jwt_auth.GenerateRefreshToken()
	if err != nil {
		s.logger.Errorf("refresh token generation error: %s", err.Error())
	}
	return token, nil
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
