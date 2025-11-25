package postgres

import (
	"context"
	"go-clean-arch/internal/adapter/repository"
	"go-clean-arch/internal/core/domain"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

var _ repository.UserRepository = &Postgres{}

func (pg *Postgres) Save(ctx context.Context, user *domain.User) error {
	query := pg.db.QueryBuilder.Insert("public.user").
		Columns("id", "document", "name", "email", "age", "password", "created_at", "updated_at").
		Values(user.ID, user.Document, user.Name, user.Email, user.Age, user.Password, time.Now(), time.Now()).
		Suffix("RETURNING *")

	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}

	err = pg.db.QueryRow(ctx, sql, args...).Scan(
		&user.ID,
		&user.Document,
		&user.Name,
		&user.Email,
		&user.Age,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if errCode := pg.db.ErrorCode(err); errCode == "23505" {
			return domain.ErrConflictingData
		}
		return err
	}

	return nil
}

func (pg *Postgres) Get(ctx context.Context, id string) (*domain.User, error) {
	var user domain.User

	query := pg.db.QueryBuilder.Select("*").
		From("public.user").
		Where(sq.Eq{"id": id}).
		Limit(1)

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	err = pg.db.QueryRow(ctx, sql, args...).Scan(
		&user.ID,
		&user.Document,
		&user.Name,
		&user.Email,
		&user.Age,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, domain.ErrDataNotFound
		}
		return nil, err
	}

	return &user, nil
}

func (pg *Postgres) List(ctx context.Context, skip, limit uint64) ([]domain.User, error) {
	var user domain.User
	var users []domain.User

	query := pg.db.QueryBuilder.Select("*").
		From("public.user").
		OrderBy("id").
		Limit(limit).
		Offset(skip * limit)

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := pg.db.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&user.ID,
			&user.Document,
			&user.Name,
			&user.Email,
			&user.Age,
			&user.Password,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (pg *Postgres) Update(ctx context.Context, user *domain.User) (*domain.User, error) {
	query := pg.db.QueryBuilder.Update("public.user").
		Set("name", sq.Expr("COALESCE(?, name)", user.Name)).
		Set("age", sq.Expr("COALESCE(?, email)", user.Age)).
		Set("updated_at", time.Now()).
		Where(sq.Eq{"id": user.ID}).
		Suffix("RETURNING *")

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	err = pg.db.QueryRow(ctx, sql, args...).Scan(
		&user.ID,
		&user.Name,
		&user.Age,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if errCode := pg.db.ErrorCode(err); errCode == "23505" {
			return nil, domain.ErrConflictingData
		}
		return nil, err
	}

	return user, nil
}

func (pg *Postgres) Delete(ctx context.Context, id string) error {
	query := pg.db.QueryBuilder.Delete("public.user").
		Where(sq.Eq{"id": id})

	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}

	_, err = pg.db.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}

	return nil
}
