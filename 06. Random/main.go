package main;

import (
	"math/big"
	// "math/rand"
	"crypto/rand"
	"fmt"
	// "time"
)

func main()  {
	/* ========================================================== */
	// /* Generating Random Numbers Using the Math random package */
	/* ========================================================== */

	// rand.Seed(time.Now().UnixNano());
	// rand := rand.Intn(5);
	// fmt.Println("seed is", rand);

	
	/* ======================================================== */
	/* Generating Random Number Using the Crypty random package */
	/* ======================================================== */

	randNum, _ := rand.Int(rand.Reader, big.NewInt(9));
	fmt.Println("Random crypto number is: ", randNum);
}