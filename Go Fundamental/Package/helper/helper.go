package helper

// Package itu Folder
// Pakai huruf besar supaya bisa di akses dari semua folder
// Jika pakai huruf kecil , tidak dapat di import dari luar folder , hanya bisa di dalam satu folder
// Nama function wajib huruf besar
func SayHello(name string) string {

	return "Hello " + name
}

// Ini tidak bisa di import dari luar package
// Karena hurud depan nya huruf kecil
func sayGoodBye(name string) string {
	return "Good bya : " + name
}
