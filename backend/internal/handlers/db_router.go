package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// DO NOT PUT THESE END POINTS IN PRODUCTION. THIS CLASS IS EXLUSIVELY FOR TESTING LMAO
type DBRouterConfig struct {
	database mongo.Database
	context  context.Context
}

func DBRouter(ctx context.Context, database mongo.Database) http.Handler {
	routerModel := DBRouterConfig{
		database: database,
		context:  ctx,
	}
	chi := chi.NewRouter()
	chi.Get("/users", routerModel.ClearUsersHandler)
	chi.Get("/all", routerModel.ClearAllHandler)

	return chi
}

/*
*
coll := client.Database("sample_mflix").Collection("movies")
filter := bson.D{{"runtime", bson.D{{"$gt", 800}}}}
// Deletes all documents that have a "runtime" value greater than 800
results, err := coll.DeleteMany(context.TODO(), filter)

	if err != nil {
		panic(err)
	}

func (*Collection) DeleteMany ¶

	An empty document (e.g. bson.D{}) should be used to delete all documents in the collection.
	GET

	type DeleteResult struct {
		DeletedCount int64 `bson:"n"` // The number of documents deleted.
	}
*/
func (config DBRouterConfig) ClearUsersHandler(res http.ResponseWriter, req *http.Request) {
	coll := config.database.Collection("users")
	filter := bson.D{}
	// Deletes all documents that have a "runtime" value greater than 800
	results, err := coll.DeleteMany(config.context, filter)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(err.Error()))
		fmt.Fprintln(res, err)
		return
	} else {
		numberDelete := results.DeletedCount
		res.WriteHeader(http.StatusOK)
		res.Write([]byte(fmt.Sprintf("ALL ENTRIES IN USER COLLECTION DELETED: %d", numberDelete)))
		log.Printf("USER COLLECTION EMPTIED - deleted a total of %d users", numberDelete)
		return
	}
}

func (config DBRouterConfig) ClearAllHandler(res http.ResponseWriter, req *http.Request) {
	config.ClearUsersHandler(res, req)

}
