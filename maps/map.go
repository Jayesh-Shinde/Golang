package main

import (
	"fmt"
)

func main() {
	websites := map[string]string{
		"Google":              "www.google.com",
		"Amazon web services": "www.aws.com",
	}
	fmt.Println("Map data:", websites)

	fmt.Println("Google key value:", websites["Google"])
	fmt.Println("Foogle key value:", websites["Foogle"])

	websites["LinkedIn"] = "www.linkedin.com"

	delete(websites, "Amazon web services")
	fmt.Println("Map data:", websites)

	//make function

	userNames := make([]string, 2, 5)

	userNames[0] = "jay"
	//userNames[2] = "shinde"// will throw error as length is only 2
	/*
			This is one of the most common points of confusion when starting with Go. The error happens because you are confusing **Length** with **Capacity**.

		In Go, a slice is a window into an underlying array. When you use `make([]string, 2, 5)`, you are creating a slice with two distinct boundaries:

		1. **Length (2):** The number of elements the slice currently "owns." You can only access indices from `0` to `len - 1`.
		2. **Capacity (5):** The total space available in the underlying memory before Go has to create a new, larger array.

		### Why `userNames[2]` fails

		Even though there is "room" for 5 items in memory, your slice is currently defined as having a size of **2**.

		In Go, indexing (using `[]`) is strictly bound by the **Length**, not the Capacity. If you try to access index `2`, Go stops you immediately with an "out of bounds" error to prevent you from accidentally accessing memory that hasn't been "activated" yet.

		---

		### How to use that extra capacity

		To use the 5 slots you reserved, you have two options:

		#### Option 1: Use `append` (Recommended)

		This is the idiomatic way to grow a slice. `append` checks the length, sees there is extra capacity, and safely puts the new value in index `2` while increasing the length for you.

		```go
		userNames := make([]string, 2, 5)
		userNames[0] = "jay"
		userNames = append(userNames, "shinde") // This works! Length becomes 3.

		```

		#### Option 2: Reslice it

		You can manually "stretch" the window of the slice to reveal the hidden capacity.

		```go
		userNames := make([]string, 2, 5)
		userNames = userNames[0:3] // Manually expand length to 3
		userNames[2] = "shinde"    // Now this works!

		```

		---

		### Summary Table

		| Feature | Length (`len`) | Capacity (`cap`) |
		| --- | --- | --- |
		| **Definition** | What is currently "in" the slice. | How much it *can* hold without reallocating. |
		| **Access** | You can use `slice[i]` if `i < len`. | You **cannot** use `slice[i]` if `i >= len`. |
		| **Purpose** | Used for loops and indexing. | Used for performance (pre-allocating memory). |

		**Think of it like a hotel:**

		* **Capacity (5)** is the number of rooms the hotel built.
		* **Length (2)** is the number of guests currently checked in.
		* If you try to walk into Room 3 without checking a guest in first (`append`), the manager (Go) will stop you!

		**Would you like to see what happens to the underlying array when you `append` more items than the capacity (e.g., adding a 6th item to a capacity-5 slice)?**
	*/

	// below append will start appending after initial length only , at index 1 we still have blank
	userNames = append(userNames, "shinde")
	userNames = append(userNames, "madhura")

	fmt.Println("username:", userNames)

	courses := make(map[string]float64, 3) //only one int parameter instead of 2
	//to specify the size of pre-allocated memory
	courses["Java"] = 9.0
	courses["Go"] = 8.0
	courses["Spring boot"] = 7.0 // here it will allocate new memory
	fmt.Println("courses:", courses)

	grocery := make(grocery, 2)

	grocery["apple"] = 5.0
	grocery["milk"] = 2.0

	fmt.Println("grocery:", grocery)

	for index, value := range userNames {
		fmt.Println("index:", index)
		fmt.Println("value:", value)
	}

	for i := 0; i < 3; i++ {
		fmt.Println("index:", i)
	}

	for i := range 3 {
		fmt.Println(i)
	}

	for key, value := range courses {
		fmt.Println("Key", key)
		fmt.Println("Value", value)
	}

}

type grocery map[string]float64
