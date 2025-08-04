# Learn Go REST API

Simple REST API untuk manajemen produk menggunakan Go dan HTTP server bawaan.

## Features

- ✅ GET /products - Lihat semua produk
- ✅ POST /products - Tambah produk baru
- ✅ PUT /products/{id} - Update produk
- ✅ DELETE /products/{id} - Hapus produk

## Quick Start

1. Clone repository
```bash
git clone https://github.com/Satriaa11/learn-go-project.git
cd learn-go-project
```

2. Jalankan server
```bash
go run main.go
```

3. Server berjalan di `http://localhost:8080`

## API Usage

### Get All Products
```bash
curl http://localhost:8080/products
```

### Create Product
```bash
curl -X POST http://localhost:8080/products \
  -H "Content-Type: application/json" \
  -d '{"name": "Product A", "price": 100}'
```

### Update Product
```bash
curl -X PUT http://localhost:8080/products/1 \
  -H "Content-Type: application/json" \
  -d '{"name": "Updated Product", "price": 150}'
```

### Delete Product
```bash
curl -X DELETE http://localhost:8080/products/1
```

## Tech Stack

- Go 1.21+
- HTTP Server (net/http)
- JSON encoding/decoding

## Learning Goals

Belajar dasar-dasar:
- HTTP handlers
- JSON marshaling/unmarshaling
- REST API patterns
- Error handling