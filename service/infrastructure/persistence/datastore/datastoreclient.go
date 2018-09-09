package persistence

import (
	"context"
	"log"

	"cloud.google.com/go/datastore"
)

func DatastoreClient(ctx context.Context) *datastore.Client {
	client, err := datastore.NewClient(ctx, "cave-inventaire")
	if err != nil {
		log.Println("Datastore client error")
		log.Fatal(err)
	}
	return client
}
