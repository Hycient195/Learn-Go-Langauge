package main

import (
	"strings"
	"strconv"
	"fmt"
	"os"
	"bufio"
)

func strToNum()  {
	fmt.Println("Give us a rating from 1 to 10");
	
	reader := bufio.NewReader(os.Stdin);
	input, _ := reader.ReadString('\n');

	convertedInput, convErr := strconv.ParseFloat(strings.TrimSpace(input), 64);

	if (convErr != nil) {
		fmt.Println(convErr);
	} else {
		fmt.Println("Rating is ", convertedInput)
		fmt.Println("Rating plus 1 is ", (convertedInput + 1))
	}
}

func numToStr() {
	num := 89;
	convertedNum := strconv.Itoa(num);
	fmt.Println("This is the converted number in string ", convertedNum)
	fmt.Printf("Type of converted number is %T \n\n", convertedNum);
}

func main()  {
	numToStr();
	strToNum();
}