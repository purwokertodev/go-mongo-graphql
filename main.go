package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/purwokertodev/go-mongo-graphql/config"
	"github.com/purwokertodev/go-mongo-graphql/internal/modules/profile/repository"

	"github.com/purwokertodev/go-mongo-graphql/internal/modules/profile/graphql/mutations"
	"github.com/purwokertodev/go-mongo-graphql/internal/modules/profile/graphql/queries"
	"github.com/purwokertodev/go-mongo-graphql/internal/modules/profile/graphql/schema"
)

func main() {
	fmt.Println("Go Mongo Db")

	db, err := config.GetMongoDB()

	if err != nil {
		fmt.Println(err)
	}

	profileRepository := repository.NewProfileRepositoryMongo(db, "profile")

	query := queries.New(profileRepository)
	mutation := mutations.New(profileRepository)

	profileSchema := schema.New(query, mutation)

	sc, err := profileSchema.GetSchema()

	if err != nil {
		fmt.Println(err)
	}

	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		result := profileSchema.ExecuteQuery(r.URL.Query().Get("query"), sc)
		json.NewEncoder(w).Encode(result)
	})

	http.ListenAndServe(":8080", nil)
}
