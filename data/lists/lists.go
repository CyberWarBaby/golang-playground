package lists

import "fmt"

// func main() {
// 	prices := []float64{10.99, 11.0}

// 	updatedPrices := append(prices, 5.99)
// 	fmt.Println(updatedPrices)
// }

// func main() {
// 	var productNames [4]string = [4]string{"A book", "A layer"}
// 	prices := [4]float64{1.99, 9.99, 45.99, 20.00}

// 	productNames[2] = "A carpet"
// 	fmt.Println(prices)
// 	fmt.Println(productNames)
// 	fmt.Println(prices[1])

// 	featuredPrices := prices[1:3] // or [:3] or [1:]
// 	highlightedPrices := featuredPrices[:1]
// 	fmt.Println(featuredPrices)
// 	fmt.Println(highlightedPrices)
// 	fmt.Println(len(highlightedPrices), cap(highlightedPrices))

// }

type Product struct {
	title string
	id    int
	price float64
}

func main() {
	hobbies := [3]string{"coding,", "speaking,", "talking to my babe"}

	//	Output (print) that array in the command line.
	fmt.Println(hobbies)

	//	The first element (standalone)
	fmt.Println(hobbies[0])

	//	The second and third element combined as a new list
	newHobbies := hobbies[1:]
	fmt.Println(newHobbies)

	//		Create that slice in two different ways (i.e. create two slices in the end)
	// newHobbies2 := hobbies[:1]
	newHobbies3 := hobbies[0:]

	// 4) Re-slice the slice from (3) and change it to contain the second
	//		and last element of the original array.
	reslicedHobbies := newHobbies3[1:]
	fmt.Println(reslicedHobbies)

	// newHobbies3 := newHobbies[1]
	// myBaby := []string{"abiola", "iyawo alabi", "faith"}
	// fmt.Println("i love", myBaby)

	courseGoals := []string{"Mad ass open source advocate", "Next level cncf contributor"}
	courseGoals[1] = "A next level maintainer of cncf projects"
	courseGoals = append(courseGoals, "Become a gde Go and global talent")
	fmt.Println(courseGoals)

	products := []Product{
		{
			"first-product",
			123,
			1.78},
		{
			"second-product",
			123,
			1.88},
	}

	newProduct := Product{
		"third-product",
		123,
		1.88,
	}

	fmt.Println(products)

	products = append(products, newProduct)
	fmt.Println(products)
}
