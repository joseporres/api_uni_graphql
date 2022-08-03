package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joseporres/api_uni_graphql/graph/models"
	"github.com/joseporres/api_uni_graphql/graph/model"
	"github.com/joseporres/api_uni_graphql/graph"
	"github.com/joseporres/api_uni_graphql/graph/generated"
)

const defaultPort = "8080"


func main() {
	db := models.FetchConnection()
	db.AutoMigrate(&model.Record{},&model.Student{},&model.Course{})
	db.Close()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
