package util

import (
	"strings"
	"testing"
)

func TestEncrypt(t *testing.T) {
	type args struct {
		password string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "successful case 1",
			args:    args{"xinchao"},
			want:    "$2a$10$1CmM8UESBamKuFhGDEFwVetkWKQ2rPK3KNd75vPe3XYiL8mPj1eNC",
			wantErr: false,
		},
		{
			name:    "successful case 2",
			args:    args{"xinchao 2"},
			want:    "$2a$10$hlk55037PCFraw1iLk2Yz.wYwebbtYHcOu8egL/g5sQTvCgkQKRXK",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Encrypt(tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("Encrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !strings.HasPrefix(got, "$2a$10$") {
				t.Errorf("Encrypt() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidatePassword(t *testing.T) {
	type args struct {
		plain     string
		encrypted string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "successful validation",
			args: args{
				plain:     "xinchao",
				encrypted: "$2a$10$hlk55037PCFraw1iLk2Yz.wYwebbtYHcOu8egL/g5sQTvCgkQKRXK",
			},
			want: true,
		},
		{
			name: "successful validation",
			args: args{
				plain:     "xinchao",
				encrypted: "$2a$10$hlk55037PCFraw1iLk2Yz.wYwebbtYHcOu8egL/g5sQTvCgkQKRXP",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidatePassword(tt.args.plain, tt.args.encrypted); got != tt.want {
				t.Errorf("ValidatePassword() = %v, want %v", got, tt.want)
			}
		})
	}
}
