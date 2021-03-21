package db

import (
    "context"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "log"
)

/* MongoCN is te connect object to database */
var MongoCN = ConnectDB()
var clientOptions = options.Client().ApplyURI("mongodb+srv://vito:vito@cluster0.5xd7z.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")

/* ConnectDb is the function for connect the DB */
func ConnectDB() *mongo.Client {
    client, err := mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        log.Fatal(err.Error())
        return client
    }
    err = client.Ping(context.TODO(), nil)
    if err != nil {
        log.Fatal(err.Error())
        return client
    }
    log.Println("Successfully connection with the DB!!!")
    return client
}

/* CheckConnection is the database Ping */
func CheckConnection() int {
    err := MongoCN.Ping(context.TODO(), nil)
    if err != nil {
        return 0
    }
    return 1
}