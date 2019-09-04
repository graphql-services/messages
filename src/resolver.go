package src

import (
	"github.com/graphql-services/messages/gen"
	"github.com/novacloudcz/graphql-orm/events"
)

func New(db *gen.DB, ec *events.EventController) *gen.GeneratedResolver {
	resolver := gen.NewResolver(db, ec)

	// resolver.Handlers.CreateCompany = func(ctx context.Context, r *gen.GeneratedMutationResolver, input map[string]interface{}) (item *gen.Company, err error) {
	// 	return nil, fmt.Errorf("can't touch this!")
	// }

	return resolver
}
