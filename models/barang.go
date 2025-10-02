package models

import "time"

type Barang struct {
	ID         string    `bson:"_id,omitempty" json:"id"`
	KodeBarang string    `bson:"kode_barang" json:"kode_barang"`
	Nama       string    `bson:"nama" json:"nama"`
	Kategori   string    `bson:"kategori" json:"kategori"`
	Stok       int       `bson:"stok" json:"stok"`
	Harga      int       `bson:"harga" json:"harga"`
	CreatedAt  time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt  time.Time `bson:"updated_at" json:"updated_at"`
}
