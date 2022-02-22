// Package ecc (Elliptic Curve Cryptography)
package ecc

import (
	"fmt"
)

type FieldElement struct {
	num   int64
	prime int64
}

func NewFieldElement(num, prime int64) (*FieldElement, error) {
	if num >= prime || num < 0 {
		return nil, fmt.Errorf("'Num %v not in field range 0 to %v'", num, prime-1)
	}
	return &FieldElement{num, prime}, nil
}

func (fe *FieldElement) String() string {
	return fmt.Sprintf("FieldElement_%v(%v)", fe.prime, fe.num)
}

func (fe *FieldElement) Equal(other *FieldElement) bool {
	if other == nil {
		return false
	}
	return fe.num == other.num && fe.prime == other.prime
}

func (fe *FieldElement) NotEqual(other *FieldElement) bool {
	return !fe.Equal(other)
}

func (fe *FieldElement) Add(other *FieldElement) (*FieldElement, error) {
	if fe.prime != other.prime {
		return nil, fmt.Errorf("cannot add two numbers in different Fields")
	}
	num := (fe.num + other.num) % fe.prime
	return NewFieldElement(num, fe.prime)
}

func (fe *FieldElement) Sub(other *FieldElement) (*FieldElement, error) {
	if fe.prime != other.prime {
		return nil, fmt.Errorf("cannot subtract two numbers in different Fields")
	}
	num := (fe.num - other.num) % fe.prime
	if num < 0 {
		num += fe.prime
	}
	return NewFieldElement(num, fe.prime)
}

func (fe *FieldElement) Mul(other *FieldElement) (*FieldElement, error) {
	if fe.prime != other.prime {
		return nil, fmt.Errorf("cannot multiply two numbers in different Fields")
	}
	num := (fe.num % fe.prime) * (other.num % fe.prime) % fe.prime
	return NewFieldElement(num, fe.prime)
}

func (fe *FieldElement) Pow(exponent int64) (*FieldElement, error) {
	if exponent == 0 {
		return NewFieldElement(1, fe.prime)
	}
	// use Fermat's little theorem
	n := exponent % (fe.prime - 1)
	if n < 0 {
		n += fe.prime - 1
	}
	var t *FieldElement
	var err error
	t, err = NewFieldElement(fe.num, fe.prime)
	if err != nil {
		return nil, err
	}
	// won't use big numbers
	for i := 1; i < int(n); i++ {
		t, err = t.Mul(fe)
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}

func (fe *FieldElement) Div(other *FieldElement) (*FieldElement, error) {
	if fe.prime != other.prime {
		return nil, fmt.Errorf("cannot divide two numbers in different Fields")
	}
	if other.num == 0 {
		return nil, fmt.Errorf("cannot divide by zero")
	}
	pow, err := other.Pow(fe.prime - 2)
	if err != nil {
		return nil, err
	}
	return fe.Mul(pow)
}
