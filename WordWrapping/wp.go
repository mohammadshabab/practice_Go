package main

import (
	"fmt"
	"unicode"
)

func main() {
	const text = `Galaksinin Batı Sarmal Kolu'nun bir ucunda, haritası bile çıkarılmamış ücra bir köşede, gözlerden uzak, küçük ve sarı bir güneş vardır.
  Bu güneşin yörüngesinde, kabaca yüz kırksekiz milyon kilometre uzağında, tamamıyla önemsiz ve mavi-yeşil renkli, küçük bir gezegen döner.
  Gezegenin maymun soyundan gelen canlıları öyle ilkeldir ki dijital kol saatinin hâlâ çok etkileyici bir buluş olduğunu düşünürler.`

	const maxWidth = 40
	var lw int
	for _, w := range text {
		fmt.Printf("%c", w)

		switch lw++; {
		case lw > maxWidth && w != '\n' && unicode.IsSpace(w):
			fmt.Println()
			fallthrough
		case w == '\n':
			lw = 0
		}
	}
	fmt.Println()

}
