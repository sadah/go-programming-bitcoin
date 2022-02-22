package ecc

import (
	"reflect"
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
			want:    &FieldElement{0, 31},
			wantErr: false,
		},
		{
			name:    "OK",
			args:    args{1, 31},
			want:    &FieldElement{1, 31},
			wantErr: false,
		},
		{
			name:    "OK",
			args:    args{32, 31},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewFieldElement(tt.args.num, tt.args.prime)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewFieldElement() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFieldElement() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFieldElement_Equal(t *testing.T) {
	type fields struct {
		num   int64
		prime int64
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
			fields: fields{0, 31},
			args:   args{other: &FieldElement{0, 31}},
			want:   true,
		},
		{
			name:   "OK",
			fields: fields{2, 31},
			args:   args{other: &FieldElement{2, 31}},
			want:   true,
		},
		{
			name:   "NG",
			fields: fields{2, 31},
			args:   args{other: &FieldElement{15, 31}},
			want:   false,
		},
		{
			name:   "NG",
			fields: fields{2, 31},
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
		num   int64
		prime int64
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
			fields: fields{2, 31},
			args:   args{other: &FieldElement{2, 31}},
			want:   false,
		},
		{
			name:   "NG",
			fields: fields{2, 31},
			args:   args{other: &FieldElement{15, 31}},
			want:   true,
		},
		{
			name:   "NG",
			fields: fields{2, 31},
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
		num   int64
		prime int64
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
			fields:  fields{2, 31},
			args:    args{other: &FieldElement{15, 31}},
			want:    &FieldElement{17, 31},
			wantErr: false,
		},
		{
			name:    "OK",
			fields:  fields{17, 31},
			args:    args{other: &FieldElement{21, 31}},
			want:    &FieldElement{7, 31},
			wantErr: false,
		},
		{
			name:    "OK",
			fields:  fields{0, 31},
			args:    args{other: &FieldElement{21, 31}},
			want:    &FieldElement{21, 31},
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
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Add() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFieldElement_Sub(t *testing.T) {
	type fields struct {
		num   int64
		prime int64
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
			fields:  fields{29, 31},
			args:    args{other: &FieldElement{4, 31}},
			want:    &FieldElement{25, 31},
			wantErr: false,
		},
		{
			name:    "OK",
			fields:  fields{15, 31},
			args:    args{other: &FieldElement{30, 31}},
			want:    &FieldElement{16, 31},
			wantErr: false,
		},
		{
			name:    "OK",
			fields:  fields{0, 31},
			args:    args{other: &FieldElement{30, 31}},
			want:    &FieldElement{1, 31},
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
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sub() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFieldElement_Mul(t *testing.T) {
	type fields struct {
		num   int64
		prime int64
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
			fields:  fields{24, 31},
			args:    args{other: &FieldElement{19, 31}},
			want:    &FieldElement{22, 31},
			wantErr: false,
		},
		{
			name:    "OK",
			fields:  fields{0, 31},
			args:    args{other: &FieldElement{1, 31}},
			want:    &FieldElement{0, 31},
			wantErr: false,
		},
		{
			name:    "OK",
			fields:  fields{1, 31},
			args:    args{other: &FieldElement{0, 31}},
			want:    &FieldElement{0, 31},
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
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Mul() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFieldElement_Pow(t *testing.T) {
	type fields struct {
		num   int64
		prime int64
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
			fields:  fields{17, 31},
			args:    args{exponent: 3},
			want:    &FieldElement{15, 31},
			wantErr: false,
		},
		{
			name:    "OK",
			fields:  fields{7, 31},
			args:    args{exponent: -3},
			want:    &FieldElement{16, 31},
			wantErr: false,
		},
		{
			name:    "OK",
			fields:  fields{7, 31},
			args:    args{exponent: 1},
			want:    &FieldElement{7, 31},
			wantErr: false,
		},
		{
			name:    "OK",
			fields:  fields{7, 31},
			args:    args{exponent: 0},
			want:    &FieldElement{1, 31},
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
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pow() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFieldElement_Div(t *testing.T) {
	type fields struct {
		num   int64
		prime int64
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
			fields:  fields{3, 31},
			args:    args{other: &FieldElement{24, 31}},
			want:    &FieldElement{4, 31},
			wantErr: false,
		},
		{
			name:    "OK",
			fields:  fields{0, 31},
			args:    args{other: &FieldElement{24, 31}},
			want:    &FieldElement{0, 31},
			wantErr: false,
		},
		{
			name:    "NG",
			fields:  fields{3, 31},
			args:    args{other: &FieldElement{0, 31}},
			want:    nil,
			wantErr: true,
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
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Div() got = %v, want %v", got, tt.want)
			}
		})
	}
}
