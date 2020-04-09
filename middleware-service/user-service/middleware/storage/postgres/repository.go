package postgres

import (
	"context"
	sqb "github.com/Masterminds/squirrel"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"user-service/core/model"
	"user-service/core/repo"
	"user-service/middleware/business"
	"user-service/util"
)

type repository struct {
	db *sqlx.DB
	table string
}

func newPostgresClient(url string) (*sqlx.DB, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	db, err := sqlx.ConnectContext(ctx,"pgx", url)
	if err != nil {
		return nil, errors.Wrap(err, "repository.newPostgresClient")
	}
	if err = db.PingContext(ctx); err != nil {
		return nil, errors.Wrap(err, "repository.newPostgresClient")
	}
	return db, nil
}

func NewPostgresRepository(connection, table string) (repo.UserRepository, error) {
	db, err := newPostgresClient(connection)
	if err != nil {
		return nil, errors.Wrap(err, "repository.NewPostgresRepo")
	}
	repo := &repository{
		db: db,
		table:  table,
	}
	return repo, nil
}

func (r *repository) FetchOne(code string) (*model.User, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	user := &model.User{}
	cmd, args, err := sqb.Select("*").From(r.table).
		Where(sqb.Eq{"UniqueId":code}).PlaceholderFormat(sqb.Dollar).ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "repository.User.FetchOne")
	}
	if err = r.db.GetContext(ctx, user, cmd, args); err != nil {
		return nil, errors.Wrap(err, "repository.User.FetchOne")
	}
	return user, nil
}

func (r *repository) Fetch(codes []string) ([]*model.User, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	users := make([]*model.User, 0)
	cmd, args, err := sqb.Select("*").From(r.table).
		Where(sqb.Eq{"UniqueId":codes}).PlaceholderFormat(sqb.Dollar).ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "repository.User.Fetch")
	}
	if err = r.db.SelectContext(ctx, &users, cmd, args); err != nil {
		return nil, errors.Wrap(err, "repository.User.Fetch")
	}
	if users == nil || len(users) == 0 {
		return nil, errors.Wrap(business.ErrUserNotFound, "repository.User.Fetch")
	}
	return users, nil
}

func (r *repository) FetchAll() ([]*model.User, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	users := make([]*model.User, 0)
	cmd, args, err := sqb.Select("*").From(r.table).PlaceholderFormat(sqb.Dollar).ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "repository.User.FetchAll")
	}
	if err = r.db.SelectContext(ctx, &users, cmd, args); err != nil {
		return nil, errors.Wrap(err, "repository.User.FetchAll")
	}
	if users == nil || len(users) == 0 {
		return nil, errors.Wrap(business.ErrUserNotFound, "repository.User.FetchAll")
	}
	return users, nil
}

func (r *repository) FetchQuery(query interface{}) ([]*model.User, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	users := make([]*model.User, 0)
	where := util.ToSqlWhere(query)
	cmd, args, err := sqb.Select("*").From(r.table).
		Where(where).PlaceholderFormat(sqb.Dollar).ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "repository.User.FetchQuery")
	}
	if err = r.db.SelectContext(ctx, &users, cmd, args); err != nil {
		return nil, errors.Wrap(err, "repository.User.FetchQuery")
	}
	if users == nil || len(users) == 0 {
		return nil, errors.Wrap(business.ErrUserNotFound, "repository.User.FetchQuery")
	}
	return users, nil
}

func (r *repository) Store(user *model.User) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cmd, args, err := sqb.Insert(r.table).SetMap(util.ToMap(user)).PlaceholderFormat(sqb.Dollar).ToSql()
	if err != nil {
		return errors.Wrap(err, "repository.User.Store")
	}
	if _, err := r.db.ExecContext(ctx, cmd, args); err != nil {
		return errors.Wrap(err, "repository.User.Store")
	}
	return nil
}

func (r *repository) Modify(user *model.User) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cmd, args, err := sqb.Update(r.table).SetMap(util.ToMap(user)).
		Where(sqb.Eq{"UniqueId":user.UniqueId}).PlaceholderFormat(sqb.Dollar).ToSql()
	if err != nil {
		return errors.Wrap(err, "repository.User.Store")
	}
	if _, err := r.db.ExecContext(ctx, cmd, args); err != nil {
		return errors.Wrap(err, "repository.User.Store")
	}
	return nil
}

func (r *repository) Replace(user *model.User) error {
	if err := r.Modify(user); err != nil {
		return errors.Wrap(err, "repository.User.Replace")
	}
	return nil
}

func (r *repository) Remove(code string) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cmd, args, err := sqb.Delete(r.table).Where(sqb.Eq{"UniqueId":code}).PlaceholderFormat(sqb.Dollar).ToSql()
	if err != nil {
		return errors.Wrap(err, "repository.User.Remove")
	}
	if _, err := r.db.ExecContext(ctx, cmd, args); err != nil {
		return errors.Wrap(err, "repository.User.Remove")
	}
	return nil
}

func (r *repository) RemoveObj(user *model.User) error {
	if err := r.Remove(user.UniqueId); err != nil {
		return errors.Wrap(err, "repository.User.RemoveObj")
	}
	return nil
}

func (r *repository) Count(query interface{}) (count int64, err error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	where := util.ToSqlWhere(query)
	cmd, args, err := sqb.Select("COUNT(*)").From(r.table).
		Where(where).PlaceholderFormat(sqb.Dollar).ToSql()
	if err != nil {
		return 0, errors.Wrap(err, "repository.User.FetchQuery")
	}
	if err = r.db.SelectContext(ctx, &count, cmd, args); err != nil {
		return 0, errors.Wrap(err, "repository.User.FetchQuery")
	}
	return
}