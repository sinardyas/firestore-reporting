package firestore

import fstore "cloud.google.com/go/firestore"

type FirestoreEngine interface {
	GetDB() *fstore.Client
	Close()
}
