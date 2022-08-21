package postgres

import (
	"context"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/hipzz/orm-practice/models"
	"github.com/hipzz/orm-practice/repository"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type userRepository struct {
	db *pgxpool.Pool
}

// Delete implements repository.User
func (u *userRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM users WHERE id = $1`
	if _, err := u.db.Exec(ctx, query, id); err != nil {
		return err
	}
	return nil
}

// Get implements repository.User
func (u *userRepository) Get(ctx context.Context) (users []models.User, err error) {
	query := `SELECT * FROM users`
	rows, err := u.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	if err := pgxscan.ScanAll(&users, rows); err != nil {
		return nil, err
	}
	return users, nil
}

// GetById implements repository.User
func (u *userRepository) GetById(ctx context.Context, id string) (user models.User, err error) {
	query := `SELECT * FROM users WHERE id = $1`
	if err := u.db.QueryRow(ctx, query, id).Scan(&user); err != nil {
		return models.User{}, err
	}
	return user, nil
}

// Save implements repository.User
func (u *userRepository) Save(ctx context.Context, users ...models.User) error {
	tx, err := u.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	query := `INSERT INTO users (id, name, email, phone, password, last_update)
				VALUES ($1, $2, $3, $4, $5, $6)`
	for _, user := range users {
		if _, err := tx.Exec(ctx, query, user.BuildInsetArgs()...); err != nil {
			tx.Rollback(ctx)
			return err
		}
	}
	return tx.Commit(ctx)
}

var _ repository.User = (*userRepository)(nil)

func New(pool *pgxpool.Pool) repository.User {
	return &userRepository{db: pool}
}
