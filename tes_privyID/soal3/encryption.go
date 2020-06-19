package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

//Fungsi createHash akan mengambil frasa sandi atau key apa pun, lalu mengembalikan hash sebagai nilai heksadesimal.
//agar panjang key bisa menyesuaikan
func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func encrypt(data []byte, passphrase string) []byte {
	//membuat cipher blok baru berdasarkan frasa sandi hash
	//Setelah memiliki cipher blok, lalu dibungkus dalam Galois Counter Mode (GCM) dengan panjang standar nonce.
	block, _ := aes.NewCipher([]byte(createHash(passphrase)))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	//Sebelum membuat ciphertext, kita perlu membuat sebuah nonce.
	//panjang nonce yang dibuat ditentukan oleh GCM. 
	//nonce yang digunakan untuk dekripsi harus sama dengan nonce yang digunakan untuk enkripsi.
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	//prepending nonce agar nonce pada deskripsi sama dengan nonce enkripsi
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	//Parameter pertama dalam Seal adalah prefix value
	//Data terenkripsi akan ditambahkan padanya
	//mengembalikan nilai ciphertext
	return ciphertext
}

func decrypt(data []byte, passphrase string) []byte {
	//membuat cipher blok baru menggunakan frasa sandi hash.
	//lalu membungkus cipher blok dalam Mode Penghitung Galois dan mendapatkan ukuran nonce.
	key := []byte(createHash(passphrase))
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonceSize := gcm.NonceSize()
	//sebelumnya, mengawali data terenkripsi dengan nonce.
	//maka harus memisahkan data nonce dan data terenkripsi.
	// memisahkan nonce dan ciphertext, agar dapat mendekripsi data dan mengembalikannya sebagai plaintext
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}
	return plaintext
}

func main() {
	fmt.Println("Starting the application...")
	ciphertext := encrypt([]byte("Selamat Datang"), "1234privy5678")
	fmt.Printf("Encrypted: %x\n", ciphertext)
	plaintext := decrypt(ciphertext, "1234privy5678")
	fmt.Printf("Decrypted: %s\n", plaintext)
}