// hello.go

package main

import (
	"fmt"
	"github.com/go-gota/gota/dataframe"
)

func main() {
	// Veri çerçevesi oluşturma
	data := [][]string{
		{"John", "Doe", "30"},
		{"Jane", "Doe", "25"},
		{"Bob", "Smith", "45"},
	}

	df := dataframe.LoadRecords(data)

	// Veri çerçevesini ekrana yazdırma
	fmt.Println("DataFrame:")
	fmt.Println(df)
}
