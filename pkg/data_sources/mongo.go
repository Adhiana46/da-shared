package data_sources

import (
	"context"
	"fmt"
	"log/slog"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDb struct {
	*mongo.Database
	client *mongo.Client
}

var (
	mongoDBInstance MongoDb
	mongoDBOnce     sync.Once
)

func NewMongoDb(host, port, user, pass, dbname string) *MongoDb {
	mongoDBOnce.Do(func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		dsn := fmt.Sprintf(
			"mongodb://%s:%s",
			host,
			port,
		)
		opts := options.Client().ApplyURI(dsn)
		opts.SetAuth(options.Credential{
			Username: user,
			Password: pass,
		})

		conn, err := mongo.Connect(ctx, opts)
		if err != nil {
			slog.Error(err.Error(), slog.String("dsn", dsn))
			panic(0)
		}

		if err := conn.Ping(ctx, nil); err != nil {
			slog.Error(err.Error(), slog.String("dsn", dsn))
			panic(0)
		}

		slog.Info(fmt.Sprintf(
			"User '%s' successfully connected to MongoDB @'%s'",
			user,
			host,
		))

		db := conn.Database(dbname)

		mongoDBInstance = MongoDb{
			Database: db,
			client:   conn,
		}
	})

	return &mongoDBInstance
}

func (r *MongoDb) Close() error {
	return r.client.Disconnect(context.Background())
}
