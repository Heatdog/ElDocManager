package userDb

import (
	"context"

	"github.com/Heatdog/ElDocManager/backend/mainServer/internal/user"
	"github.com/Heatdog/ElDocManager/backend/mainServer/pkg/client/postgresql"
	"github.com/Heatdog/ElDocManager/backend/mainServer/pkg/logging"
)

type repository struct {
	client postgresql.Client
	logger *logging.Logger
}

func (r *repository) Create(ctx context.Context, user *user.User) error {
	r.logger.Info("SQL INSERT Users")
	q := `
		INSERT INTO Users 
			(login, name, surname, email, password)
		VALUES
			($1, $2, $3, $4, $5)
		RETURNING id
	`
	row := r.client.QueryRow(ctx, q, user.Login, user.Name, user.Surname, user.Email, user.Password)
	if err := row.Scan(&user.ID); err != nil {
		r.logger.Errorf("SQL error: %s", err.Error())
		return err
	}
	r.logger.Infof("Successful Users INSERT: {id:%s}", user.ID)
	return nil
}

// Delete implements user.Repository.
func (r *repository) Delete(ctx context.Context, id string) error {
	panic("unimplemented")
}

func (r *repository) FindAll(ctx context.Context) ([]user.User, error) {
	r.logger.Infof("SQL SELECT Users")
	q := `
		SELECT id, login, name, surname, email, role, created, is_confirmed 
		FROM Users
	`
	rows, err := r.client.Query(ctx, q)
	if err != nil {
		r.logger.Errorf("SQL error: %s", err.Error())
		return nil, err
	}

	users := make([]user.User, 0, 10)

	for rows.Next() {
		var userRow user.User

		err = rows.Scan(
			&userRow.ID, &userRow.Login, &userRow.Name, &userRow.Surname,
			&userRow.Email, &userRow.Role, &userRow.Created, &userRow.IsConfirmed)
		if err != nil {
			r.logger.Errorf("SQL error: %s", err.Error())
			return nil, err
		}

		users = append(users, userRow)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *repository) FindOneWithId(ctx context.Context, id string) (user.User, error) {
	r.logger.Infof("SQL SELECT Users by Id")
	q := `
		SELECT id, login, name, surname, email, role, created, is_confirmed 
		FROM Users
		WHERE id = $1
	`
	row := r.client.QueryRow(ctx, q, id)

	var userRow user.User

	if err := row.Scan(&userRow); err != nil {
		r.logger.Errorf("SQL error: %s", err.Error())
		return user.User{}, err
	}

	return userRow, nil
}

func (r *repository) FindOneWithLogin(ctx context.Context, login string) (user.User, error) {
	r.logger.Infof("SQL Select Users by Login")
	q := `
		SELECT id, login, name, surname, email, role, created, is_confirmed 
		FROM Users
		WHERE login = $1
	`
	row := r.client.QueryRow(ctx, q, login)

	var userRow user.User

	if err := row.Scan(&userRow); err != nil {
		r.logger.Errorf("SQL error: %s", err.Error())
		return user.User{}, err
	}

	return userRow, nil
}

// Update implements user.Repository.
func (r *repository) Update(ctx context.Context, user user.User) error {
	panic("unimplemented")
}

func NewUserRepository(client postgresql.Client, logger *logging.Logger) user.UserRepository {
	return &repository{
		client: client,
		logger: logger,
	}
}
