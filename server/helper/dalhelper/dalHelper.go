package dalhelper

import (
	"context"
	"fmt"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gocraft/dbr"
	"github.com/onkarsutar/UserManagement/server/helper/confighelper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var sqlConnection *dbr.Connection
var connectionError error
var sqlOnce sync.Once

var clientInstance *mongo.Client
var clientInstanceError error
var mongoOnce sync.Once

// GetSQLConnection : Connect to SQL DB
func GetSQLConnection() (*dbr.Connection, error) {
	sqlOnce.Do(func() {
		connection, err := dbr.Open("mysql", confighelper.GetConfig("mySQLDSN"), nil)
		if err != nil {
			fmt.Println(err)
			connectionError = err
		}
		connection.SetMaxIdleConns(100)
		connection.SetMaxOpenConns(5000)
		duration := 3 * 24 * time.Hour
		connection.SetConnMaxLifetime(duration)
		sqlConnection = connection
	})
	return sqlConnection, connectionError
}

// GetMongoClient : Conects to Mongo DB
func GetMongoClient() (*mongo.Client, error) {
	mongoOnce.Do(func() {
		clientOptions := options.Client().ApplyURI(confighelper.GetConfig("mongoDSN"))
		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			clientInstanceError = err
		}
		err = client.Ping(context.TODO(), nil)
		if err != nil {
			clientInstanceError = err
		}
		clientInstance = client
	})
	return clientInstance, clientInstanceError
}
