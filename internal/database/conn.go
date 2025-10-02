package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"

	"github.com/RifkyA911/management-inventory/pkg/config"
)

var MongoClient *mongo.Client
var MongoDB *mongo.Database

func InitMongo() {
	// Load config dari pkg/config
	cfg := config.LoadEnv()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Koneksi client
	client, err := mongo.Connect(options.Client().ApplyURI(cfg.MongoURI))
	if err != nil {
		log.Fatalf("❌ Gagal connect MongoDB: %v", err)
	}

	// Tes koneksi
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("❌ Gagal ping MongoDB: %v", err)
	}

	log.Println("✅ MongoDB connected")

	MongoClient = client
	MongoDB = client.Database(cfg.MongoDB)
}
