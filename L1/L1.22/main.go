package main

import (
	"fmt"
	"math/big"
)

func add(a, b *big.Int) *big.Int {
	return new(big.Int).Add(a, b)
}

func sub(a, b *big.Int) *big.Int {
	return new(big.Int).Sub(a, b)
}

func mul(a, b *big.Int) *big.Int {
	return new(big.Int).Mul(a, b)
}

func div(a, b *big.Int) (*big.Int, error) {
	if b.Sign() == 0 {
		return nil, fmt.Errorf("divide by zero")
	}
	return new(big.Int).Quo(a, b), nil
}

func main() {
	a := new(big.Int).Lsh(big.NewInt(1), 128)
	b, ok := new(big.Int).SetString("12329437591874581745917495871948571983457981374598713498573988745878478574857842024", 10)
	if !ok {
		panic("can't create a bigInt froms tr")
	}

	fmt.Println("a =", a)
	fmt.Println("b =", b)

	fmt.Println("a + b =", add(a, b))
	fmt.Println("a - b =", sub(a, b))
	fmt.Println("a * b =", mul(a, b))

	ab, _ := div(a, b)
	fmt.Println("a / b =", ab)
	ba, _ := div(b, a)
	fmt.Println("b / a =", ba)
}