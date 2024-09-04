package service

import (
	"context"
	repo "example.com/m/internal/repo"
	"example.com/m/internal/service"
	"example.com/m/internal/service/service/extended"
)

type serv struct {
	noteRepository repo.NoteRepo
}

func NewService(noteRepository repo.NoteRepo) service.Service {
	return &serv{noteRepository: noteRepository}
}

func (s *serv) Create(ctx context.Context, username string, password string) (int64, error) {
	id, err := s.Create(ctx, username, password)
	if err != nil {
		return id, err
	}
	return id, nil
}

func (s *serv) Get(ctx context.Context, id int64) (string, error) {
	user, err := s.noteRepository.Get(ctx, id)
	if err != nil {
		return "", err
	}
	role, err := extended.IsAdmin(ctx, user)
	if err != nil {
		return "", err
	}
	if role == 1 {
		return user.Username + "IS ADMIN", nil
	}
	return user.Username, nil
}

//func (s *serv) Update(ctx context.Context, username string, password string) error {
//	return nil
//}
//
//func (s *serv) Delete(ctx context.Context, username string) error {
//	return nil
//}
