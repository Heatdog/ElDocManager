package user

import (
	"context"

	authServer "github.com/Heatdog/ElDocManager/backend/authServer/pkg/proto"
	cryptohash "github.com/Heatdog/ElDocManager/backend/mainServer/pkg/cryptoHash"

	logger "github.com/Heatdog/ElDocManager/backend/logger/app"
)

type UserService interface {
	SignIn(context context.Context, userLogin *UserSignIn, jwtKey string) (string, error)
	SignUp(context context.Context, userLogin *UserSignUp) error
	SignOut(context context.Context)
}

func NewUserService(logger *logger.Logger, repo UserRepository, authClient authServer.AuthServerClient) UserService {
	return &userService{
		logger:     logger,
		repo:       repo,
		authClient: authClient,
	}
}

type userService struct {
	logger     *logger.Logger
	repo       UserRepository
	authClient authServer.AuthServerClient
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

	res, err := s.authClient.CreateRefreshToken(context, &authServer.TokenCreateRequest{
		UserId: user.ID,
		Role:   string(user.Role),
	})

	if err != nil {
		s.logger.Errorf("token generation error: %s", err.Error())
		return "", err
	}

	return res.AccessToken, nil
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
