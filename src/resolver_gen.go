package src

import (
	"github.com/graphql-services/messages/gen"
	"github.com/novacloudcz/graphql-orm/events"
)

func NewResolver(db *gen.DB, ec *events.EventController) *Resolver {
	handlers := gen.DefaultResolutionHandlers()
	return &Resolver{&gen.GeneratedResolver{Handlers: handlers, DB: db, EventController: ec}}
}

type Resolver struct {
	*gen.GeneratedResolver
}

type MutationResolver struct {
	*gen.GeneratedMutationResolver
}

type QueryResolver struct {
	*gen.GeneratedQueryResolver
}

func (r *Resolver) Mutation() gen.MutationResolver {
	return &MutationResolver{&gen.GeneratedMutationResolver{r.GeneratedResolver}}
}
func (r *Resolver) Query() gen.QueryResolver {
	return &QueryResolver{&gen.GeneratedQueryResolver{r.GeneratedResolver}}
}

type CommentResultTypeResolver struct {
	*gen.GeneratedCommentResultTypeResolver
}

func (r *Resolver) CommentResultType() gen.CommentResultTypeResolver {
	return &CommentResultTypeResolver{&gen.GeneratedCommentResultTypeResolver{r.GeneratedResolver}}
}
