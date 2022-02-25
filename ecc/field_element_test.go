package ecc

import (
	"github.com/google/go-cmp/cmp"
	"math/big"
	"testing"
)

func TestNewFieldElement(t *testing.T) {
	type args struct {
		num   int64
		prime int64
	}
	tests := []struct {
		name    string
		args    args
		want    *FieldElement
		wantErr bool
	}{
		{
			name:    "OK",
			args:    args{0, 31},
			want:    &FieldElement{big.NewInt(0), big.NewInt(31)},
			wantErr: false,
		},
		{
			name:    "OK",
			args:    args{1, 31},
			want:    &FieldElement{big.NewInt(1), big.NewInt(31)},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewFieldElement(tt.args.num, tt.args.prime)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewFieldElement() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("NewFieldElement() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFieldElement_Equal(t *testing.T) {
	type fields struct {
		num   *big.Int
		prime *big.Int
	}
	type args struct {
		other *FieldElement
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name:   "OK",
			fields: fields{big.NewInt(0), big.NewInt(31)},
			args:   args{other: &FieldElement{big.NewInt(0), big.NewInt(31)}},
			want:   true,
		},
		{
			name:   "OK",
			fields: fields{big.NewInt(2), big.NewInt(31)},
			args:   args{other: &FieldElement{big.NewInt(2), big.NewInt(31)}},
			want:   true,
		},
		{
			name:   "NG",
			fields: fields{big.NewInt(2), big.NewInt(31)},
			args:   args{other: &FieldElement{big.NewInt(1), big.NewInt(31)}},

			want: false,
		},
		{
			name:   "NG",
			fields: fields{big.NewInt(2), big.NewInt(31)},
			args:   args{other: nil},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fe := &FieldElement{
				num:   tt.fields.num,
				prime: tt.fields.prime,
			}
			if got := fe.Equal(tt.args.other); got != tt.want {
				t.Errorf("Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFieldElement_NotEqual(t *testing.T) {
	type fields struct {
		num   *big.Int
		prime *big.Int
	}
	type args struct {
		other *FieldElement
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name:   "NG",
			fields: fields{big.NewInt(2), big.NewInt(31)},
			args:   args{other: &FieldElement{big.NewInt(2), big.NewInt(31)}},
			want:   false,
		},
		{
			name:   "OK",
			fields: fields{big.NewInt(2), big.NewInt(31)},
			args:   args{other: &FieldElement{big.NewInt(15), big.NewInt(31)}},
			want:   true,
		},
		{
			name:   "OK",
			fields: fields{big.NewInt(2), big.NewInt(31)},
			args:   args{other: nil},
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fe := &FieldElement{
				num:   tt.fields.num,
				prime: tt.fields.prime,
			}
			if got := fe.NotEqual(tt.args.other); got != tt.want {
				t.Errorf("NotEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFieldElement_Add(t *testing.T) {
	type fields struct {
		num   *big.Int
		prime *big.Int
	}
	type args struct {
		other *FieldElement
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *FieldElement
		wantErr bool
	}{
		{
			name:    "OK",
			fields:  fields{big.NewInt(2), big.NewInt(31)},
			args:    args{other: &FieldElement{big.NewInt(15), big.NewInt(31)}},
			want:    &FieldElement{big.NewInt(17), big.NewInt(31)},
			wantErr: false,
		},
		{
			name:    "OK",
			fields:  fields{big.NewInt(17), big.NewInt(31)},
			args:    args{other: &FieldElement{big.NewInt(21), big.NewInt(31)}},
			want:    &FieldElement{big.NewInt(7), big.NewInt(31)},
			wantErr: false,
		},
		{
			name:    "OK",
			fields:  fields{big.NewInt(0), big.NewInt(31)},
			args:    args{other: &FieldElement{big.NewInt(21), big.NewInt(31)}},
			want:    &FieldElement{big.NewInt(21), big.NewInt(31)},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fe := &FieldElement{
				num:   tt.fields.num,
				prime: tt.fields.prime,
			}
			got, err := fe.Add(tt.args.other)
			if (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("Add() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFieldElement_Sub(t *testing.T) {
	type fields struct {
		num   *big.Int
		prime *big.Int
	}
	type args struct {
		other *FieldElement
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *FieldElement
		wantErr bool
	}{
		{
			name:    "OK",
			fields:  fields{big.NewInt(29), big.NewInt(31)},
			args:    args{other: &FieldElement{big.NewInt(4), big.NewInt(31)}},
			want:    &FieldElement{big.NewInt(25), big.NewInt(31)},
			wantErr: false,
		},
		{
			name:    "OK",
			fields:  fields{big.NewInt(15), big.NewInt(31)},
			args:    args{other: &FieldElement{big.NewInt(30), big.NewInt(31)}},
			want:    &FieldElement{big.NewInt(16), big.NewInt(31)},
			wantErr: false,
		},
		{
			name:    "OK",
			fields:  fields{big.NewInt(0), big.NewInt(31)},
			args:    args{other: &FieldElement{big.NewInt(30), big.NewInt(31)}},
			want:    &FieldElement{big.NewInt(1), big.NewInt(31)},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fe := &FieldElement{
				num:   tt.fields.num,
				prime: tt.fields.prime,
			}
			got, err := fe.Sub(tt.args.other)
			if (err != nil) != tt.wantErr {
				t.Errorf("Sub() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("Sub() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFieldElement_Mul(t *testing.T) {
	type fields struct {
		num   *big.Int
		prime *big.Int
	}
	type args struct {
		other *FieldElement
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *FieldElement
		wantErr bool
	}{
		{
			name:   "OK",
			fields: fields{big.NewInt(24), big.NewInt(31)},
			args:   args{other: &FieldElement{big.NewInt(19), big.NewInt(31)}},
			want:   &FieldElement{big.NewInt(22), big.NewInt(31)},
		},
		{
			name:    "OK",
			fields:  fields{big.NewInt(0), big.NewInt(31)},
			args:    args{other: &FieldElement{big.NewInt(1), big.NewInt(31)}},
			want:    &FieldElement{big.NewInt(0), big.NewInt(31)},
			wantErr: false,
		},
		{
			name:    "OK",
			fields:  fields{big.NewInt(1), big.NewInt(31)},
			args:    args{other: &FieldElement{big.NewInt(0), big.NewInt(31)}},
			want:    &FieldElement{big.NewInt(0), big.NewInt(31)},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fe := &FieldElement{
				num:   tt.fields.num,
				prime: tt.fields.prime,
			}
			got, err := fe.Mul(tt.args.other)
			if (err != nil) != tt.wantErr {
				t.Errorf("Mul() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("Mul() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFieldElement_Pow(t *testing.T) {
	type fields struct {
		num   *big.Int
		prime *big.Int
	}
	type args struct {
		exponent int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *FieldElement
		wantErr bool
	}{
		{
			name:    "OK",
			fields:  fields{big.NewInt(17), big.NewInt(31)},
			args:    args{exponent: 3},
			want:    &FieldElement{big.NewInt(15), big.NewInt(31)},
			wantErr: false,
		},
		{
			name:    "OK",
			fields:  fields{big.NewInt(7), big.NewInt(31)},
			args:    args{exponent: -3},
			want:    &FieldElement{big.NewInt(16), big.NewInt(31)},
			wantErr: false,
		},
		{
			name:    "OK",
			fields:  fields{big.NewInt(7), big.NewInt(31)},
			args:    args{exponent: 1},
			want:    &FieldElement{big.NewInt(7), big.NewInt(31)},
			wantErr: false,
		},
		{
			name:    "OK",
			fields:  fields{big.NewInt(7), big.NewInt(31)},
			args:    args{exponent: 0},
			want:    &FieldElement{big.NewInt(1), big.NewInt(31)},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fe := &FieldElement{
				num:   tt.fields.num,
				prime: tt.fields.prime,
			}
			got, err := fe.Pow(tt.args.exponent)
			if (err != nil) != tt.wantErr {
				t.Errorf("Pow() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("Pow() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFieldElement_Div(t *testing.T) {
	type fields struct {
		num   *big.Int
		prime *big.Int
	}
	type args struct {
		other *FieldElement
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *FieldElement
		wantErr bool
	}{
		{
			name:    "OK",
			fields:  fields{big.NewInt(3), big.NewInt(31)},
			args:    args{other: &FieldElement{big.NewInt(24), big.NewInt(31)}},
			want:    &FieldElement{big.NewInt(4), big.NewInt(31)},
			wantErr: false,
		},
		{
			name:    "OK",
			fields:  fields{big.NewInt(0), big.NewInt(31)},
			args:    args{other: &FieldElement{big.NewInt(24), big.NewInt(31)}},
			want:    &FieldElement{big.NewInt(0), big.NewInt(31)},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fe := &FieldElement{
				num:   tt.fields.num,
				prime: tt.fields.prime,
			}
			got, err := fe.Div(tt.args.other)
			if (err != nil) != tt.wantErr {
				t.Errorf("Div() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("Div() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
