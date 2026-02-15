package main

import "fmt"

func main() {
	var productArray [4]string = [4]string{"apple"}
	prices := [4]int{1, 2, 3, 4}
	fmt.Println("Prices", prices)
	fmt.Println(productArray)
	productArray[2] = "Banana"
	fmt.Println(productArray)
	fmt.Println("Price", prices[2])
	fmt.Println("Slice of prices", prices[1:3])
	fmt.Println("Prices", prices)

	fmt.Println("length Prices", len(prices))

	fmt.Println("Slice of prices", prices[:3])
	fmt.Println("Slice of prices", prices[1:])

	slicedPrice := prices[:1]
	slicedPrice[0] = 99 //changes orignal array as well as slice doe not create new array just kind of pointer to same array
	fmt.Println("Slied array ", slicedPrice)
	fmt.Println("Orignal array: ", prices)

	fmt.Println("Slied array length ", len(slicedPrice)) // length is 1 but below capacity is 4 as pointing to same orignal array
	fmt.Println("Slied array capacity", cap(slicedPrice))

	tempratures := []int{0} // this creates a dynalic array

	tempratures[0] = 1
	fmt.Println("temp:", tempratures)
	//tempratures[1] = 2 // this will cause error as already initial capacity in array was 1,
	// use append to allocate new array of with new element and size
	//fmt.Println("temp:", tempratures)
	newTemps := append(tempratures, 2) // this will add a value and return the new array keeping old one as it is
	fmt.Println("temp orignal:", tempratures)
	fmt.Println("new temps:", newTemps)

	tempratures = append(tempratures, 2, 3)
	fmt.Println("temp overwritten:", tempratures) // here it is overwritten

	list1 := []int{1, 2}
	fmt.Println("lsit1 : ", list1)

	list2 := []int{3, 4}
	fmt.Println("list2 :", list2)

	list1 = append(list1, list2...) // this merge works only on dynamic array where array are like
	//list1 := []int{1, 2} where size is specified but not list1 := [2]int{1, 2} where size is specified

	fmt.Println("merged list: ", list1)
}
