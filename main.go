package main

import "fmt"
import "math/big"
import "github.com/ethereum/go-ethereum/crypto"
import "github.com/ethereum/go-ethereum/crypto/secp256k1"

func main() {
	curve := secp256k1.S256()

	fmt.Println(curve.IsOnCurve(big.NewInt(10), big.NewInt(20)))

	k := []byte("hi")

	x, y := curve.ScalarBaseMult(k)
	fmt.Printf("x = %s, y = %s\n", x, y)

	fmt.Println(curve.IsOnCurve(x, y))

	x, y = curve.Double(x, y)
	fmt.Printf("x = %s, y = %s\n", x, y)

	fmt.Println(curve.IsOnCurve(x, y))

	K := x.String() + y.String()

	hash := crypto.Keccak256([]byte(K))

	fmt.Printf("hash = %v\n", hash)
}
