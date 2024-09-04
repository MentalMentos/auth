package repo

import (
	"example.com/m/internal/model"
	repo "example.com/m/internal/repo/repoUser/model"
)

func FromRepoToUser(user *repo.RepoUser) *model.User {
	return &model.User{user.Id, user.Name, user.Password, user.Role, user.CreatedAt, user.UpdatedAt}
}
