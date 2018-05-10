package schema

import (
	"fmt"

	"github.com/graphql-go/graphql"

	"github.com/purwokertodev/go-mongo-graphql/internal/modules/profile/graphql/mutations"
	"github.com/purwokertodev/go-mongo-graphql/internal/modules/profile/graphql/queries"
)

//ProfileSchema
type ProfileSchema struct {
	query    *queries.ProfileQuery
	mutation *mutations.ProfileMutation
}

//New
func New(query *queries.ProfileQuery, mutation *mutations.ProfileMutation) *ProfileSchema {
	return &ProfileSchema{query: query, mutation: mutation}
}

//GetSchema
func (s *ProfileSchema) GetSchema() (graphql.Schema, error) {
	return graphql.NewSchema(graphql.SchemaConfig{
		Query:    s.query.GetQuery(),
		Mutation: s.mutation.GetMutation(),
	})
}

//ExecuteQuery
func (s *ProfileSchema) ExecuteQuery(q string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: q,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}
