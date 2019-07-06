package routing

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// NewRouter Create router and assign routes
func NewRouter() http.Handler {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGODB_CONN"))
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("needyd").Collection("jobs_indeed")
	jobsHandler := &JobsHandler{collection: collection}

	router := httprouter.New()
	router.GET("/", jobsHandler.index)
	router.GET("/jobs", jobsHandler.getJobs)

	route := newValidateRouteMiddleware(router, router)
	cors := cors.AllowAll().Handler(route)

	return cors
}
