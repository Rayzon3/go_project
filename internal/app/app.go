package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Rayzon3/go_project/internal/api"
	"github.com/Rayzon3/go_project/internal/store"
)

type Application struct {
	DB             *sql.DB
	Logger         *log.Logger
	WorkoutHandler *api.WorkoutHandler
}

func NewApplication() (*Application, error) {
	pgDB, err := store.Open()

	if err != nil {
		return nil, err
	}

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	//stores

	//handlers
	workoutHandler := api.NewWorkoutHandler()

	app := &Application{
		DB:             pgDB,
		Logger:         logger,
		WorkoutHandler: workoutHandler,
	}

	return app, nil
}

func (a *Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Status is OK \n")
}
