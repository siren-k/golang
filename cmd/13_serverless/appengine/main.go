package main

import (
	"cloud.google.com/go/datastore"
	"context"
	"fmt"
	"golang/internal/13_serverless/appengine"
	"log"
	"net/http"
	"os"
)

func main() {
	ctx := context.Background()
	log.SetOutput(os.Stderr)

	// 프로덕션 환경에서 실행할 때 app.yml에 다음 코드를 설정한다.
	projectID := os.Getenv("GCLOUD_DATASET_ID")
	datastoreClient, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatal(err)
	}

	c := appengine.Controller{Store: datastoreClient}

	http.HandleFunc("/", c.Handle)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	log.Printf("Listening to port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
