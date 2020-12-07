/*
	dosya işlemleri
*/

package main

// kullanacağımız paketleri import ediyoruz
import (
	"fmt"
	"os"
)

// uygulamanın giriş noktası
func main() {

	/*
		O_RDWR : Dosyayı hem okuma hem de yazma için (read-write) açar. En çok kullanılan budur.
		O_CREATE : Dosyayı açarke eğer belirtilen isimde bir dosya yoksa oluşturur.
	*/

	_, err := os.OpenFile("read.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}

	/*
		Dosya bilgilerine ulaşma
	*/
	fi, err := os.Stat("read.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("File Name: " + fi.Name())

	/*
		Dosya/Klasör türlerini öğrenme
	*/
	switch mode := fi.Mode(); {
	case mode.IsRegular():
		fmt.Println("regular file")

	case mode.IsDir():
		fmt.Println("directory")

	case mode&os.ModeSymlink != 0:
		fmt.Println("symbolic link")

	case mode&os.ModeNamedPipe != 0:
		fmt.Println("named pipe")
	}
}
