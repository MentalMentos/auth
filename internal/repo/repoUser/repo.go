package repoUser

import (
	"context"
	"errors"
	"example.com/m/internal/client/db"
	"example.com/m/internal/model"
	repo "example.com/m/internal/repo"
	converter "example.com/m/internal/repo/repoUser/converter"
	modelRepo "example.com/m/internal/repo/repoUser/model"
	"github.com/Masterminds/squirrel"
	sq "github.com/Masterminds/squirrel"
	"golang.org/x/crypto/bcrypt"
)

const (
	tableNAME = "User"

	idColumn        = "id"
	userColumn      = "user"
	passwordColumn  = "password"
	roleColumn      = "role"
	createdAtColumn = "created_at"
	updatedAtColumn = "updated_at"
)

var ErrUserNotFound = errors.New("user not found")

type Repo struct {
	db db.Client
}

func NewRepository(db db.Client) repo.NoteRepo {
	return &Repo{db: db}
}

func (r *Repo) Create(ctx context.Context, note *model.User) (int64, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(note.Password), bcrypt.DefaultCost)
	if err != nil {
		return note.Id, err
	}
	builder := squirrel.Insert(tableNAME).
		Columns(idColumn, userColumn, passwordColumn, roleColumn, createdAtColumn, updatedAtColumn).
		Values(note.Id, note.Username, hashedPassword, note.Role, note.CreatedAt, note.UpdatedAt).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "note_repository.Create",
		QueryRaw: query,
	}
	var id int64
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *Repo) Get(ctx context.Context, id int64) (*model.User, error) {
	builder := sq.Select(idColumn, userColumn, passwordColumn, roleColumn, createdAtColumn, updatedAtColumn).
		PlaceholderFormat(sq.Dollar).
		From(tableNAME).
		Where(sq.Eq{idColumn: id}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "note_repository.Get",
		QueryRaw: query,
	}

	var note modelRepo.RepoUser
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&note.Id, &note.Name, &note.Password, &note.Role, &note.CreatedAt, &note.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return converter.FromRepoToUser(&note), nil
}
