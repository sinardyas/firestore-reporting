package firestore

import (
	"context"
	"log"

	fstore "cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

type FirestoreProjectId string

type firestore struct {
	app *fstore.Client
}

var _ FirestoreEngine = (*firestore)(nil)

func NewFirestoreDB(pId FirestoreProjectId) (FirestoreEngine, error) {
	fs := &firestore{}

	ctx := context.Background()
	opt := option.WithCredentialsFile("./Tikung.json")

	var err error
	fs.app, err = fstore.NewClient(ctx, string(pId), opt)
	if err != nil {
		log.Fatalf("Firestore config error:%s\n", err)
	}

	return fs, nil
}

func (f *firestore) GetDB() *fstore.Client {
	return f.app
}

func (f *firestore) Close() {
	if f != nil {
		f.app.Close()
	}
}
