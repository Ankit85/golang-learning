package main

import "fmt"

func main() {
	fmt.Println("Learning Loops")

	// days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	/* for d := 0; d < len(days); d++ {
		fmt.Printf("Days are %v \n", days[d])
	}
	*/

	/* for i := range days {
		fmt.Printf("Days are %v \n", days[i])
	} */
	/* for index, day := range days {
		fmt.Printf("value at index %v Days are %v \n", index, day)
	} */

	number := 2
	for number < 10 {
		if number == 4 {
			number++
			continue
		}
		/* if number == 5 {
			goto sika
		} */
		if number == 6 {
			break
		}
		fmt.Printf("Days are %v \n", number)
		number++
	}
	/*sika:
	fmt.Println("Came here from number 5") */
}
