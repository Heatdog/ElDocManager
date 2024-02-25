package user

import "context"

type UserService interface {
	SignIn(context context.Context, userLogin *UserSignIn) string
	SignUp(context context.Context, userLogin *UserSignUp) error
	SignOut(context context.Context)
}

func NewAuthService() UserService {
	return &userService{}
}

type userService struct {
}

func (s *userService) SignIn(context context.Context, userLogin *UserSignIn) string {
	return ""
}

func (s *userService) SignUp(context context.Context, userLogin *UserSignUp) error {
	return nil
}

func (s *userService) SignOut(context context.Context) {

}
