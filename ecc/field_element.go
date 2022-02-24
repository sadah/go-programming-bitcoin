// Package ecc (Elliptic Curve Cryptography)
package ecc

import (
	"fmt"
	"math/big"
)

type FieldElement struct {
	num   *big.Int
	prime *big.Int
}

func NewFieldElement(num, prime int64) (*FieldElement, error) {
	if num >= prime || num < 0 {
		return nil, fmt.Errorf("'Num %v not in field range 0 to %v'", num, prime-1)
	}
	return &FieldElement{big.NewInt(num), big.NewInt(prime)}, nil
}

func (fe *FieldElement) String() string {
	return fmt.Sprintf("FieldElement_%v(%v)", fe.prime, fe.num)
}

func (fe *FieldElement) Num() *big.Int {
	return fe.num
}

func (fe *FieldElement) Prime() *big.Int {
	return fe.prime
}

func (fe *FieldElement) Equal(other *FieldElement) bool {
	if other == nil {
		return false
	}
	return fe.num.Cmp(other.num) == 0 && fe.prime.Cmp(other.prime) == 0
}

func (fe *FieldElement) NotEqual(other *FieldElement) bool {
	return !fe.Equal(other)
}

func (fe *FieldElement) Add(other *FieldElement) (*FieldElement, error) {
	if fe.prime.Cmp(other.prime) != 0 {
		return nil, fmt.Errorf("cannot add two numbers in different Fields")
	}
	ret := big.NewInt(0)
	ret.Add(fe.num, other.num).Mod(ret, fe.prime)
	return &FieldElement{ret, fe.prime}, nil
}

func (fe *FieldElement) Sub(other *FieldElement) (*FieldElement, error) {
	if fe.prime.Cmp(other.prime) != 0 {
		return nil, fmt.Errorf("cannot subtract two numbers in different Fields")
	}
	ret := big.NewInt(0)
	ret.Sub(fe.num, other.num).Mod(ret, fe.prime)
	return &FieldElement{ret, fe.prime}, nil
}

func (fe *FieldElement) Mul(other *FieldElement) (*FieldElement, error) {
	if fe.prime.Cmp(other.prime) != 0 {
		return nil, fmt.Errorf("cannot multiply two numbers in different Fields")
	}
	// num := (fe.num % fe.prime) * (other.num % fe.prime) % fe.prime
	ret := big.NewInt(0)
	ret.Mul(fe.num, other.num).Mod(ret, fe.prime)
	return &FieldElement{ret, fe.prime}, nil
}

func (fe *FieldElement) Pow(exponent int64) (*FieldElement, error) {
	if exponent == 0 {
		return &FieldElement{big.NewInt(1), fe.prime}, nil
	}
	// use Fermat's little theorem
	// n = exponent % (prime - 1)
	n := big.NewInt(0)
	prime := big.NewInt(0).Set(fe.prime)
	n.Mod(big.NewInt(exponent), prime.Sub(prime, big.NewInt(1)))

	ret := big.NewInt(0)
	ret.Exp(fe.num, n, fe.prime)
	return &FieldElement{ret, fe.prime}, nil
}

func (fe *FieldElement) Div(other *FieldElement) (*FieldElement, error) {
	if fe.prime.Cmp(other.prime) != 0 {
		return nil, fmt.Errorf("cannot divide two numbers in different Fields")
	}
	if big.NewInt(0).Cmp(other.num) == 0 {
		return nil, fmt.Errorf("cannot divide by zero")
	}

	prime := big.NewInt(0).Set(fe.prime)
	n := prime.Sub(prime, big.NewInt(2))
	ret := big.NewInt(0)
	// ret = num * pow(other.num, prime - 2, prime) % prime
	ret.Exp(other.num, n, fe.prime)
	ret.Mul(ret, fe.num).Mod(ret, fe.prime)
	return &FieldElement{ret, fe.prime}, nil
}
