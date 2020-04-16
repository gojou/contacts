package contacts

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
)

// Repo provides the hanger for database methods
type repo struct {
	db firestore.Client
}

// Initdb connects to the Firestore database
func initdb() (db firestore.Client, e error) {
	// Use the application default credentials

	ctx := context.Background()

	client, e := firestore.NewClient(ctx, "contacts-hvlst")

	defer client.Close()
	if e != nil {
		log.Fatalln(e)
		return *client, e
	}

	e = useDB(ctx, *client)
	if e != nil {
		log.Fatalf("Failed: %v", e)
		return
	}
	return
}

func useDB(ctx context.Context, db firestore.Client) (err error) {

	_, _, err = db.Collection("users").Add(ctx, map[string]interface{}{
		"first": "Ada",
		"last":  "Lovelace",
		"born":  1815,
	})
	if err != nil {
		log.Fatalf("Failed adding alovelace: %v", err)
		return
	}
	log.Println("added Ada")

	_, _, err = db.Collection("users").Add(ctx, map[string]interface{}{
		"first":  "Alan",
		"middle": "Mathison",
		"last":   "Turing",
		"born":   1912,
	})
	log.Println("added Alan")
	if err != nil {
		log.Fatalf("Failed adding aturing: %v", err)
		return
	}
	return

}
