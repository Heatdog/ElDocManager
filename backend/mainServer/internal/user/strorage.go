package user

import "context"

type UserRepository interface {
	Create(ctx context.Context, user *User) error
	FindAll(ctx context.Context) ([]User, error)
	FindOneWithId(ctx context.Context, id string) (User, error)
	FindOneWithLogin(ctx context.Context, login string) (User, error)
	Update(ctx context.Context, user User) error
	Delete(ctx context.Context, id string) error
}
