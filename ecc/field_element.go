// Package ecc (Elliptic Curve Cryptography)
package ecc

import "fmt"

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

func (fe *FieldElement) Equal(other *FieldElement) bool {
	if other == nil {
		return false
	}
	return fe.num == other.num && fe.prime == other.prime
}

func (fe *FieldElement) NotEqual(other *FieldElement) bool {
	return !fe.Equal(other)
}
