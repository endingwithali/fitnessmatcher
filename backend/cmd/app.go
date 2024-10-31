package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/endingwithali/fitnessapp/backend/internal/handlers"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// os.Setenv("SESSION_SECRET", os.Getenv("SESSION_SECRET"))

	ctx := context.Background()
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// connect to MONGO DB
	mongoClient := connectToDB(ctx)

	// close DB connection when main() stops running
	defer func() { // TODO: does this need to be in main given it runs on failure of app, or can it be held in the connectionFunction?
		if err := mongoClient.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	mongoDatabase := mongoClient.Database("fitnessapp")
	chiRouter := chi.NewRouter()
	chiRouter.Use(middleware.Logger)

	chiRouter.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi"))
	})

	chiRouter.Mount("/auth", handlers.AuthRouter(ctx, *mongoDatabase))
	chiRouter.Mount("/workout", handlers.WorkoutRouter(ctx, *mongoDatabase))
	chiRouter.Mount("/db", handlers.DBRouter(ctx, *mongoDatabase))

	log.Println("listening on localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", chiRouter))

	// http.ListenAndServe(":3000", chiRouter)
	// log.Println("Serving at localhost:3000")
}

func connectToDB(ctx context.Context) *mongo.Client {
	mongoURI := os.Getenv("MONGODB_URI")

	//create "options" used by mongo to connect to DB instance
	options := options.Client().ApplyURI(mongoURI)
	options = options.SetMaxPoolSize(80) //manually set connection limit to db  so we dont exceed based on tier usage

	mongoClient, err := mongo.Connect(ctx, options)
	if err != nil {
		log.Fatalf("failed to connect to database \n ERR: \n %s", err)
	}
	log.Println("Connection to MongoDB successful")

	return mongoClient
}
