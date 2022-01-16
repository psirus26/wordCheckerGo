package folder1

import "fmt"

func Function1() {
	var (
		firstNum  int = 1
		secondNum     = 20
	)

	const (
		Monday = iota
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		Sunday
	)

	fmt.Println("called func in pkg", firstNum, secondNum)
	fmt.Println("called func in pkg", Monday, Tuesday)
	fmt.Println("called func in pkg", Wednesday, Thursday)
	fmt.Println("called func in pkg", Friday, Saturday, Sunday)

	// fmt.Scan(&firstNum, &secondNum)
	// fmt.Println(firstNum, secondNum)

	green := 'A'
	red := `SUS`
	var purple string = "AMOGUS"
	var blue = 75
	fmt.Println(green, red, purple, blue)

	fmt.Println(fmt.Sprintf("%[3]*.[2]*[1]f", 12.0, 2, 6))

	fmt.Println(fmt.Sprintf("%d", 500))
	// 'first' variable now has the value of 500

	fmt.Println(fmt.Sprintf("%b", 500))
	// 'binaryVariable' variable now has the value of 111110100

}

func Function2() {
	var val int
	var res int = 1
	fmt.Scan(&val)

	for i := val; i > 0; i-- {
		res *= i
	}
	fmt.Print(res)
}

func Function3() {
	var input string
	var output string
	fmt.Scan(&input)
	var length = len(input) - 1

	for i := length; i >= 0; i-- {
		output = output + string(input[i])
	}
	fmt.Print(output)
}

func Function5() {
	var p = new(string)
	fmt.Println(p) // Will print some memory address, for eg. 0xc000040240
	// This address can be different in your case
	*p = "this is my value"
	fmt.Println(*p) // Will print the actual value — an empty string
	fmt.Println(p)  // Will print the actual value — an empty string

	var x *string
	x = new(string)
	*x = "test"
	fmt.Println(*x)

	var s = "yes, it is possible"
	var p1 = &s
	var p2 = &p1

	fmt.Println(*p2)  // Will print true
	fmt.Println(p1)   // Will print true
	fmt.Println(*p1)  // Will print true
	fmt.Println(**p2) // Will print true
	fmt.Println(s)    // Will print true
	fmt.Println(&s)   // Will print true

}
