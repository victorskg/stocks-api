package mongodb

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Connection struct {
	databaseName string
	host         string
	port         uint16
	client       *mongo.Client
}

func NewMongoConnection(databaseName string, host string, port uint16) *Connection {
	return &Connection{
		databaseName: databaseName,
		host:         host,
		port:         port,
	}
}

func (c *Connection) Collection(collectionName string) *mongo.Collection {
	c.connect()
	return c.client.Database(c.databaseName).Collection(collectionName)
}

func (c *Connection) connect() {
	if c.client != nil {
		clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%d", c.host, c.port))

		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			log.Fatal("Error trying to connect", err)
		}

		err = client.Ping(context.TODO(), nil)
		if err != nil {
			log.Fatal("Error trying to ping", err)
		}

		c.client = client
	}
}
