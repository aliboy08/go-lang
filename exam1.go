package main

import ( "fmt" )

func main() {

	var inputReps int

	// Get number of repetitions
	// parameters: input_var_container, min_value, max_value, message
	inputReps = getInput(inputReps, 1, 100, `Enter number of repetitions`)

	var results []int // Sum of squares container

	// Loop Repetitions
	i := 0
	loopReps:
	if i < inputReps {

		var inputNum int

		// Get number of inputs
		inputNum = getInput(inputNum, 1, 100, `Enter number of inputs`)

		// Get N numbers
		msg := fmt.Sprintf("Enter %d numbers:", inputNum)
		currentInputs := getMultipleInputs(inputNum, -100, 100, msg)

		// Calculate sum of squared numbers
		results = append(results, calculateSquareSum(currentInputs))

		i++;
		goto loopReps
	}

	// Loop through results to display per line
	i = 0
	loopResults:
	if i < len(results) {
		fmt.Println("Result", (i+1), " = ", results[i])
		i++
		goto loopResults
	}

}

// Get single inputs, loop until input are all valid, return int
func getInput(input int, min int, max int, msg string) (int){

	i := 0
	loop:
	// Loop until input is valid
	if i < 1 {

		fmt.Println(msg); // display input prompt

		fmt.Scan(&input) // get input

		if(input >= min && input <= max) {

			// Valid input
			return input

		} else {

			// Invalid input, ask again
			fmt.Println("Invalid input, allowed input is from", min, "up to", max, ". Try again")

		}

		goto loop
	}

	return input
}

// Get multiple inputs, Loop until input are all valid, return int array
func getMultipleInputs(inputNum int, min int, max int, msg string) ([]int){
	var inputs []int

	i := 0
	loop:
	// Loop until input is valid
	if i < 1 {

		var inputs []int // clear slice, if input is invalid
		var inputValid = true // Input validity flag

		fmt.Println(msg); //display input prompt

		// Loop number of inputs
		j := 0
		loopInputs:
		if j < inputNum {

			var input int

			fmt.Scan(&input) // get input

			inputs = append(inputs, input) // temporarily add to slice

			if(input < min || input > max) {
				// Invalid input, ask again
				fmt.Println("Invalid input, allowed input is from", min, "up to", max, ". Try again")
				inputValid = false
			}

			j++
			goto loopInputs
		}

		if( inputValid ) {
			// All inputs are valid
			return inputs;
		}

		goto loop
	}

	return inputs
}

// Calculate the sum of the squared numbers, exclude negative numbers, return int
func calculateSquareSum(inputs []int) (int){
	var sum int

	i:= 0
	loop:
	// Loop through the set of numbers
	if i < len(inputs) {

		// Only square non-negative numbers
		if( inputs[i] > 0 ) {
			square := inputs[i] * inputs[i]; // square
			sum += square // add to sum
		}

		i++
		goto loop
	}

	return sum
}
