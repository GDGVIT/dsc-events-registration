package main

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/GDGVIT/dsc-events-registration/api/middleware"
	"github.com/gorilla/handlers"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {

	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	godotenv.Load()
}

func DBConnect() (*mongo.Client, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("DATABASE_URI")))
	if err != nil {
		return nil, err
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	return client, err
}

func main() {

	// connect to DB
	client, err := DBConnect()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(client)

	r := httprouter.New()
	r.GET("/", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.WriteHeader(http.StatusOK)
		return
	})

	// HTTP(s) binding
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	timeout := time.Duration(10 * time.Second)

	if port == "" {
		port = "3000"
	}

	conn := host + ":" + port

	// middlewares
	mwCors := middleware.CorsEveryWhere(r)
	mwLogs := handlers.LoggingHandler(os.Stdout, mwCors)

	srv := &http.Server{
		ReadTimeout:  timeout,
		WriteTimeout: timeout,
		Addr:         conn,
		Handler:      mwLogs,
	}

	log.Printf("Server running on %s", conn)
	log.Fatal(srv.ListenAndServe())
}
