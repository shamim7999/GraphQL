package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Database         *mongo.Database
	CollectionBook   *mongo.Collection
	CollectionAuthor *mongo.Collection
	ClientOptions    *options.ClientOptions
	ClientBook       *mongo.Client
	ErrBook          error
	CtxBook          context.Context
	CtxAuthor        context.Context
	ErrAuthor        error
	ClientAuthor     *mongo.Client
)
