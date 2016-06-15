package eventsourcing

import "golang.org/x/net/context"

const (
	repositoryKey = "repository"
)

func WithRepository(parent context.Context, repo Repository) context.Context {
	ctx := context.WithValue(parent, repositoryKey, repo)
	return ctx
}

func RepositoryInContext(ctx context.Context) Repository {
	tr := ctx.Value(repositoryKey)
	r, _ := tr.(Repository)
	return r
}
