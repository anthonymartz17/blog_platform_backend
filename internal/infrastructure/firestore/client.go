package firestore

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
)

func NewFirestoreClient(ctx context.Context, app *firebase.App) (*firestore.Client,error) {

	client, err := app.Firestore(ctx)
	
	if err != nil {
		return nil,fmt.Errorf("error initializing firestore: %w", err)
	}

	return client,nil
}


