package firebase

import (
	"context"
	"fmt"
	"os"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

func NewFirestoreClient(ctx context.Context) (*firestore.Client,error) {

	serviceAccountPath:= os.Getenv("FIREBASE_CREDENTIALS")

	if serviceAccountPath == "" {
		return nil, fmt.Errorf("FIREBASE_CREDENTIALS environment variable not set")
	}
	

	creds, err := os.ReadFile(serviceAccountPath)

	if err != nil {
		return nil,fmt.Errorf("failed to read firebase credentials: %v", err)
	}

	opt := option.WithAuthCredentialsJSON(option.ServiceAccount,creds)

	app, err := firebase.NewApp(ctx, nil, opt)

	if err != nil {
		return nil,fmt.Errorf("error initializing firebase app: %w", err)
	}
	
	client, err := app.Firestore(ctx)
	
	if err != nil {
		return nil,fmt.Errorf("error initializing firestore: %w", err)
	}

	return client,nil
}



