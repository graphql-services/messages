package gen

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/novacloudcz/graphql-orm/resolvers"
	"github.com/vektah/gqlparser/ast"
)

type GeneratedQueryResolver struct{ *GeneratedResolver }

type QueryCommentHandlerOptions struct {
	ID     *string
	Q      *string
	Filter *CommentFilterType
}

func (r *GeneratedQueryResolver) Comment(ctx context.Context, id *string, q *string, filter *CommentFilterType) (*Comment, error) {
	opts := QueryCommentHandlerOptions{
		ID:     id,
		Q:      q,
		Filter: filter,
	}
	return r.Handlers.QueryComment(ctx, r.GeneratedResolver, opts)
}
func QueryCommentHandler(ctx context.Context, r *GeneratedResolver, opts QueryCommentHandlerOptions) (*Comment, error) {
	query := CommentQueryFilter{opts.Q}
	offset := 0
	limit := 1
	rt := &CommentResultType{
		EntityResultType: resolvers.EntityResultType{
			Offset: &offset,
			Limit:  &limit,
			Query:  &query,
			Filter: opts.Filter,
		},
	}
	qb := r.DB.Query()
	if opts.ID != nil {
		qb = qb.Where("comments.id = ?", *opts.ID)
	}

	var items []*Comment
	err := rt.GetItems(ctx, qb, "comments", &items)
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, &NotFoundError{Entity: "Comment"}
	}
	return items[0], err
}

type QueryCommentsHandlerOptions struct {
	Offset *int
	Limit  *int
	Q      *string
	Sort   []CommentSortType
	Filter *CommentFilterType
}

func (r *GeneratedQueryResolver) Comments(ctx context.Context, offset *int, limit *int, q *string, sort []CommentSortType, filter *CommentFilterType) (*CommentResultType, error) {
	opts := QueryCommentsHandlerOptions{
		Offset: offset,
		Limit:  limit,
		Q:      q,
		Sort:   sort,
		Filter: filter,
	}
	return r.Handlers.QueryComments(ctx, r.GeneratedResolver, opts)
}
func QueryCommentsHandler(ctx context.Context, r *GeneratedResolver, opts QueryCommentsHandlerOptions) (*CommentResultType, error) {
	_sort := []resolvers.EntitySort{}
	for _, s := range opts.Sort {
		_sort = append(_sort, s)
	}
	query := CommentQueryFilter{opts.Q}

	var selectionSet *ast.SelectionSet
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		if f.Field.Name == "items" {
			selectionSet = &f.Field.SelectionSet
		}
	}

	return &CommentResultType{
		EntityResultType: resolvers.EntityResultType{
			Offset:       opts.Offset,
			Limit:        opts.Limit,
			Query:        &query,
			Sort:         _sort,
			Filter:       opts.Filter,
			SelectionSet: selectionSet,
		},
	}, nil
}

type GeneratedCommentResultTypeResolver struct{ *GeneratedResolver }

func (r *GeneratedCommentResultTypeResolver) Items(ctx context.Context, obj *CommentResultType) (items []*Comment, err error) {
	err = obj.GetItems(ctx, r.DB.db, "comments", &items)
	return
}

func (r *GeneratedCommentResultTypeResolver) Count(ctx context.Context, obj *CommentResultType) (count int, err error) {
	return obj.GetCount(ctx, r.DB.db, &Comment{})
}
