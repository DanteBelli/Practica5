package main

import "fmt"

func Revez(s string) string {
	b := []byte(s)
	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}
func main() {
	input := "Boca perdio en la clasificacion y no entro a la Libertadores"
	rev := Revez(input)
	dosrev := Revez(rev)
	fmt.Printf("La frase original es: %q\n", input)
	fmt.Printf("La frase revez es: %q\n", rev)
	fmt.Printf("La Doble revez es: %q\n", dosrev)
}
