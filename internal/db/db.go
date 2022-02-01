package db

import (
	"context"
	"fmt"
	"github.com/volvinbur1/ethereum-transactions/internal/cmn"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

const (
	databaseName   = "ethereum"
	collectionName = "transactions"
)

type Manager struct {
	client *mongo.Client

	username, serverIp, port string
}

func New(username, serverIp, port string) *Manager {
	m := &Manager{
		username: username,
		serverIp: serverIp,
		port:     port,
	}
	ctx, ctxCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer ctxCancel()

	var err error
	m.client, err = mongo.Connect(ctx, options.Client().ApplyURI(fmt.Sprintf("mongodb://%s@%s:%s", m.username, m.serverIp, m.port)))
	if err != nil {
		log.Fatal(err)
	}

	return m
}

// Disconnect should be called before Manager release
func (m *Manager) Disconnect() {
	ctx, ctxCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer ctxCancel()

	err := m.client.Disconnect(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

func (m *Manager) AddTransaction(transaction cmn.Transaction) error {
	authColl := m.client.Database(databaseName).Collection(collectionName)

	ctx, ctxCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer ctxCancel()

	_, err := authColl.InsertOne(ctx, transaction)
	if err != nil {
		return err
	}

	return nil
}

func (m *Manager) GetAllTransactions() ([]cmn.Transaction, error) {
	authColl := m.client.Database(databaseName).Collection(collectionName)

	ctx, ctxCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer ctxCancel()

	cursor, err := authColl.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var transactions []cmn.Transaction
	if err = cursor.All(ctx, &transactions); err != nil {
		return nil, err
	}

	return transactions, nil
}

func (m *Manager) FindByField(fieldName, valueToSearch string) ([]cmn.Transaction, error) {
	authColl := m.client.Database(databaseName).Collection(collectionName)

	ctx, ctxCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer ctxCancel()

	filter := bson.D{
		{fieldName, bson.D{{"$eq", valueToSearch}}},
	}
	cursor, err := authColl.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var transactions []cmn.Transaction
	if err = cursor.All(ctx, &transactions); err != nil {
		return nil, err
	}

	return transactions, nil
}
