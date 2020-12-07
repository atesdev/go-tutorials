/*
	for, if, switch kullanımı
*/
package main

// kullanacağımız paketleri import ediyoruz
import (
	"fmt"
)

// uygulamanın giriş noktası
func main() {
	//1 ile 100 arasındaki sayılardan iki ile bölünebilenlerin toplamı
	toplam := 0
	for i := 0; i < 100; i++ {
		if i%2 == 0 { //Mod operatörü ile kalan hesaplanır
			toplam += i
		}
	}
	fmt.Println("2 ile bölünebilenlerin toplamı=", toplam)

	// switch case kullanımı - sınav notu
	sinav_notu := 70
	switch {
	case sinav_notu >= 0 && sinav_notu < 45:
		fmt.Println("kaldın")
	case sinav_notu >= 45 && sinav_notu < 50:
		fmt.Println("hoca ne derse o olacak")
	case sinav_notu >= 50 && sinav_notu < 75:
		fmt.Println("geçmene yeterli")
	case sinav_notu >= 75 && sinav_notu <= 100:
		fmt.Println("çok iyi")
	default: //Girilen not hiçbir koşula uymuyorsa
		fmt.Println("geçersiz bir not girildi")
	}
}
