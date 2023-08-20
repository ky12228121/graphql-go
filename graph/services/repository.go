package services

import (
	"context"
	"graphql-go/graph/db"
	"graphql-go/graph/model"
	"time"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

const (
	layout = "2006-01-02T15:04:05Z"
)

type repositoryService struct {
	exec boil.ContextExecutor
}

func (r *repositoryService) GetRepositoryByFullName(ctx context.Context, owner, name string) (*model.Repository, error) {
	repo, err := db.Repositories(
		qm.Select(
			db.RepositoryColumns.ID,        // レポジトリID
			db.RepositoryColumns.Name,      // レポジトリ名
			db.RepositoryColumns.Owner,     // レポジトリを所有しているユーザーのID
			db.RepositoryColumns.CreatedAt, // 作成日時
		),
		db.RepositoryWhere.Owner.EQ(owner),
		db.RepositoryWhere.Name.EQ(name),
	).One(ctx, r.exec)
	if err != nil {
		return nil, err
	}
	return convertRepository(repo), nil
}

func convertRepository(repo *db.Repository) *model.Repository {
	return &model.Repository{
		ID:        repo.ID,
		Owner:     &model.User{ID: repo.Owner},
		Name:      repo.Name,
		CreatedAt: stringToTime(repo.CreatedAt),
	}
}

func stringToTime(str string) time.Time {
	t, err := time.Parse(layout, str)
	if err != nil {
		return time.Time{}
	}
	return t
}
