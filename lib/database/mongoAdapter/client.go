package mongoAdapter

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/url"
	"restaurantAPI/lib/database"
)

type Client struct {
	database *mongo.Database
}

var _ database.Client = &Client{}

func (c Client) Disconnect() error {
	return c.database.Client().Disconnect(context.Background())
}

func NewClient(DatabaseURI url.URL) *Client {
	password, _ := DatabaseURI.User.Password()
	client, _ := mongo.Connect(context.Background(),
		options.
			Client().
			ApplyURI(DatabaseURI.String()).
			SetAuth(options.Credential{Username: DatabaseURI.User.Username(), Password: password}),
	)
	db := client.Database(DatabaseURI.Path)
	return &Client{database: db}
}

func GetCollection[T database.Entity](Client *Client, name database.CollectionName) Collection[T] {
	return Collection[T]{
		name:            name,
		mongoCollection: Client.database.Collection(string(name)),
	}
}
