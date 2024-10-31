package handlers

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/mongo"
)

type WorkoutRouterConfig struct {
	database mongo.Database
}

func WorkoutRouter(ctx context.Context, database mongo.Database) http.Handler {
	routerModel := WorkoutRouterConfig{
		database: database,
	}
	chi := chi.NewRouter()
	chi.Get("/", routerModel.GetWorkoutHandler)
	chi.Post("/", routerModel.PostWorkoutHandler)
	return chi
}

/*
*

	Get workout of a certain workoutID
	GET
*/
func (configs WorkoutRouterConfig) GetWorkoutHandler(res http.ResponseWriter, req *http.Request) {

	/*
	1. check if user is logged in
	 - if not return to home
	2. create mongodb search object
	3. poll db
	4. create return object
	5. return return object
	*/
	return
}

/*
*

	Creates a new workout for a user
	POST
*/
func (configs WorkoutRouterConfig) PostWorkoutHandler(res http.ResponseWriter, req *http.Request) {
	/*
	
	*/
	return
}
