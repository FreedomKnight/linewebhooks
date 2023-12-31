package database

import (
    "context"
    "log"
    "fmt"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"

	"github.com/freedomknight/linewebhooks/configs"
	"github.com/freedomknight/linewebhooks/container"
)

var database *mongo.Database
var ctx = context.TODO()

// XXX: Try to find a way to make database pool in future.

func init() {
    config := configs.Get()
    config.SetConfigName("database")
    config.SetConfigType("yaml")
    err := config.ReadInConfig()
    if err != nil {
        panic(err.Error())
    }

    host, port, name := config.Get("mongodb.host"), config.Get("mongodb.port"), config.Get("mongodb.database")
    println("mongo listen on: ", fmt.Sprintf("mongodb://%s:%d/%s", host, port, name))
    clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%d/", host, port))

    client, err := mongo.Connect(ctx, clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    database = client.Database(fmt.Sprintf("%s", name))
    container.Provide(func () *mongo.Database {
        return database
    })
}

func Get() *mongo.Database {
    var d *mongo.Database

    container.Invoke(func (db *mongo.Database) {
        d = db
    })

    return d
}

func GetCtx() context.Context {
    return ctx
}

