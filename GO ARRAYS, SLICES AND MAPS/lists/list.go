package lists

import "fmt"

func List() {
	//Array creation
	productPrices := [4]float64{1.99, 9.89, 9.06, 2.45}

	//Dynamic array using slices
	names := []string{}

	namesForNames := [4]string{"Benard", "Allan", "Meshack", "Victor"}

	//Adding a value to a dynamic array or rather slice
	for _, name := range namesForNames {
		names = append(names, name)
	}

	names = append(names, "Brian")

	//Print the new slice
	fmt.Println(names)

	//slicing an array
	featuredPrices := productPrices[1:]
	//You can also create a slice from anather slice
	highlightedPrices := featuredPrices[1:2]
	fmt.Println(productPrices)
	fmt.Println(featuredPrices)
	fmt.Println(highlightedPrices)
}
