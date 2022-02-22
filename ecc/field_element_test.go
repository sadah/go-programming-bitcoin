package ecc

import "testing"

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
			want:   false,
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
