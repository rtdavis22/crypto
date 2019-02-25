package main

import "fmt"
import "math/big"
import "github.com/ethereum/go-ethereum/crypto/secp256k1"

func main() {
	curve := secp256k1.S256()

	fmt.Println(curve.IsOnCurve(big.NewInt(10), big.NewInt(20)))

	k := []byte("hi")

	x, y := curve.ScalarBaseMult(k)
	fmt.Printf("x = %s, y = %s\n", x.String(), y.String())

	fmt.Println(curve.IsOnCurve(x, y))

	x, y = curve.Double(x, y)
	fmt.Printf("x = %s, y = %s\n", x.String(), y.String())

	fmt.Println(curve.IsOnCurve(x, y))
}
