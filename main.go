package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv" // Tambahkan import strconv
)

// Representasi dari produk JSON
type Products struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func main() {
	// 1. buat route multiplexer
	mux := http.NewServeMux()

	// 3 tambahkan handler ke multiplexer
	mux.HandleFunc("GET /products", listProduct)
	mux.HandleFunc("POST /products", createProduct)
	mux.HandleFunc("PUT /products/{id}", updateProduct)
	mux.HandleFunc("DELETE /products/{id}", deleteProduct)

	// 4. buat server HTTP
	server := &http.Server{
		Addr:    ":8080", // alamat server/port
		Handler: mux,     // gunakan multiplexer sebagai handler
	}

	// 5. jalankan server
	server.ListenAndServe() // jalankan server
}

var database = map[int]Products{
	// 1: {ID: "1", Name: "Product A", Price: 100},
	// 2: {ID: "2", Name: "Product B", Price: 200},
	// 3: {ID: "3", Name: "Product C", Price: 300},
}

var lastID = 0

// 2. buat handler untuk route "/"
func listProduct(w http.ResponseWriter, r *http.Request) {
	// slice untuk response JSON
	var products []Products

	// iterasi database untuk mengisi slice products
	for _, v := range database {
		products = append(products, v)
	}

	// marshal digunakan untuk mengubah slice of Products/struct menjadi JSON
	data, err := json.Marshal(products)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		w.Write([]byte(`{"error": "Internal Server Error"}`))
		return // Tambahkan return untuk menghentikan eksekusi
	}

	// set header Content-Type untuk response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200) // kode status 200 OK
	w.Write(data)      // kirim data JSON sebagai response
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	bodyByte, err := io.ReadAll(r.Body)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400) // Bad Request
		w.Write([]byte(`{"error": "Bad Request"}`))
		return // Tambahkan return untuk menghentikan eksekusi
	}

	var products Products
	err = json.Unmarshal(bodyByte, &products)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400) // Bad Request
		w.Write([]byte(`{"error": "Bad Request"}`))
		return // Tambahkan return untuk menghentikan eksekusi
	}

	lastID++                           // increment lastID untuk ID produk baru
	products.ID = strconv.Itoa(lastID) // Gunakan strconv.Itoa untuk konversi int ke string
	database[lastID] = products        // simpan produk baru ke database

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201) // Created
	w.Write([]byte(`{"message": "Product created successfully"}`))

}

func updateProduct(w http.ResponseWriter, r *http.Request) {

	// Baca path value dari URL
	productID := r.PathValue("id")               // Ambil ID produk dari URL
	productIDint, err := strconv.Atoi(productID) // Konversi ID produk ke integer
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500) // Internal Server Error
		w.Write([]byte(`{"error": "Internal Server Error"}`))
		return // Tambahkan return untuk menghentikan eksekusi
	}

	// baca body request
	bodyByte, err := io.ReadAll(r.Body)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400) // Bad Request
		w.Write([]byte(`{"error": "Bad Request"}`))
		return // Tambahkan return untuk menghentikan eksekusi
	}

	// unmarshal konversi JSON ke struct Products
	var products Products
	err = json.Unmarshal(bodyByte, &products)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400) // Bad Request
		w.Write([]byte(`{"error": "Bad Request"}`))
		return // Tambahkan return untuk menghentikan eksekusi
	}

	products.ID = productID           // supaya ID produk tetap sama(tidak berubah)
	database[productIDint] = products // Update produk di database (Map)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200) // OK
	w.Write([]byte(`{"message": "Product updated successfully"}`))
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id") // Ambil ID produk dari URL

	productIDint, err := strconv.Atoi(productID) // Konversi ID produk ke integer
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500) // Internal Server Error
		w.Write([]byte(`{"error": "Internal Server Error"}`))
		return // Tambahkan return untuk menghentikan eksekusi
	}

	delete(database, productIDint) // Hapus produk dari database (Map)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200) // OK
	w.Write([]byte(`{"message": "Product deleted successfully"}`))
}
