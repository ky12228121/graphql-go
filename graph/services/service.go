package services

import (
	"context"
	"graphql-go/graph/model"

	"github.com/volatiletech/sqlboiler/boil"
)

type Services interface {
	UserService
	RepositoryService
	IssueService
	// issueテーブルを扱うIssueServiceなど、他のサービスインターフェースができたらそれらを追加していく
}

type services struct {
	*userService
	*repositoryService
	*issueService
	// issueテーブルを扱うissueServiceなど、他のサービス構造体ができたらフィールドを追加していく
}

type UserService interface {
	GetUserByID(ctx context.Context, id string) (*model.User, error)
	GetUserByName(ctx context.Context, name string) (*model.User, error)
}
type RepositoryService interface {
	GetRepositoryByFullName(ctx context.Context, owner, name string) (*model.Repository, error)
}
type IssueService interface {
	GetIssueByRepoAndNumber(ctx context.Context, repoID string, number int) (*model.Issue, error)
}

func New(exec boil.ContextExecutor) Services {
	return &services{
		userService:       &userService{exec: exec},
		repositoryService: &repositoryService{exec: exec},
		issueService:      &issueService{exec: exec},
	}
}
