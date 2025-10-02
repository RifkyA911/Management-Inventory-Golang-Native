package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/RifkyA911/management-inventory/internal/database"
	"github.com/RifkyA911/management-inventory/models"
	"github.com/go-faker/faker/v4"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type BarangSeeder struct {
	collection *mongo.Collection
}

func NewBarangSeeder(db *mongo.Database) *BarangSeeder {
	return &BarangSeeder{
		collection: db.Collection("barangs"),
	}
}

func (s *BarangSeeder) Seed(count int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var barangs []interface{}
	kategoris := []string{"Elektronik", "Pakaian", "Makanan", "Minuman", "Peralatan", "Kesehatan", "Olahraga"}

	for i := 0; i < count; i++ {
		stok, _ := faker.RandomInt(10, 100)
		harga, _ := faker.RandomInt(10000, 1000000)

		barang := models.Barang{
			KodeBarang: fmt.Sprintf("BRG-%s", faker.UUIDDigit()),
			Nama:       faker.Word() + " " + faker.Word(),
			Kategori:   kategoris[i%len(kategoris)],
			Stok:       stok[0],
			Harga:      harga[0],
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		}
		barangs = append(barangs, barang)
	}

	result, err := s.collection.InsertMany(ctx, barangs)
	if err != nil {
		return fmt.Errorf("❌ gagal insert data: %v", err)
	}

	log.Printf("✅ Berhasil insert %d barang", len(result.InsertedIDs))
	return nil
}

func (s *BarangSeeder) Clear() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := s.collection.DeleteMany(ctx, map[string]interface{}{})
	if err != nil {
		return fmt.Errorf("❌ gagal hapus data: %v", err)
	}

	log.Printf("🗑️ Berhasil hapus %d barang", result.DeletedCount)
	return nil
}

func (s *BarangSeeder) Run(count int) error {
	log.Println("🧹 Clearing existing data...")
	if err := s.Clear(); err != nil {
		return err
	}

	log.Printf("🌱 Seeding %d barang...", count)
	if err := s.Seed(count); err != nil {
		return err
	}

	log.Println("🎉 Seeding completed!")
	return nil
}

func main() {
	database.InitMongo()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	defer database.MongoClient.Disconnect(ctx)

	barangSeeder := NewBarangSeeder(database.MongoDB)
	if err := barangSeeder.Clear(); err != nil {
		log.Fatalf("❌ Gagal flush data lama: %v", err)
	}

	if err := barangSeeder.Run(250); err != nil {
		log.Fatal(err)
	}
}
