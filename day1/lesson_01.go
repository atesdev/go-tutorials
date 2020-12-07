/*
	temel işlemler
*/
package main

// kullanacağımız paketleri import ediyoruz
import (
	"fmt"
)

// birden fazla değişken tanımlama
var (
	name, email string
)

// uygulamanın giriş noktası
func main() {
	name = "Hulusi kentmen"
	email = "hulusi@kentmen.com"

	fmt.Println("Hello " + name + ". Email: " + email)
}
