package main

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestValidatePhone(t *testing.T) {
	type args struct {
		phoneNumber string
		countryCode string
	}
	tests := []struct {
		name string
		args 	args
		want    PhoneInformation
		wantErr bool
	}{
		{
			name: "exp1",
			args: args{
				phoneNumber: "631 311 8150",
				countryCode: "",
			},
		want: PhoneInformation{
			PhoneNumber: "631 311 8150",
			Error: ErrorBody{
				CountryCode: "required value is missing",
			},
		}, wantErr: false,
	},
		{
			name: "exp2",
			args: args{
				phoneNumber: "+12125690123",
				countryCode: "",
			},
			want: PhoneInformation{
					PhoneNumber: "+12125690123",
					CountryCode: "US",
					AreaCode: "212",
					LocalPhoneNumber: "5690123",
			}, wantErr: false,
		},
		{
			name: "exp2",
			args: args{
				phoneNumber: "12125690123",
				countryCode: "US",
			},
			want: PhoneInformation{
				PhoneNumber: "12125690123",
				CountryCode: "US",
				AreaCode: "212",
				LocalPhoneNumber: "5690123",
			}, wantErr: false,
		},
		{
			name: "exp3",
			args: args{
				phoneNumber: "+52 631 3118150",
				countryCode: "",
			},
			want: PhoneInformation{
				PhoneNumber: "+52 631 3118150",
				CountryCode: "MX",
				AreaCode: "631",
				LocalPhoneNumber: "3118150",
			}, wantErr: false,
		},
		{
			name: "exp4",
			args: args{
				phoneNumber: "34 915 872200",
				countryCode: "",
			},
			want: PhoneInformation{
				PhoneNumber: "34 915 872200",
				CountryCode: "ES",
				AreaCode: "915",
				LocalPhoneNumber: "872200",
			}, wantErr: false,
		},
		{
			name: "exp5",
			args: args{
				phoneNumber: "25690123",
				countryCode: "",
			},
			want: PhoneInformation{
				PhoneNumber: "25690123",
				Error: ErrorBody{
					CountryCode: "required value is missing",
				},
			}, wantErr: false,
		},
	}


	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg, err := validatePhone(tt.args.phoneNumber,tt.args.countryCode)
			if (err != nil) != tt.wantErr {
				t.Errorf("validatePhone error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(msg, tt.want) {
				t.Errorf("validatePhone = %v", cmp.Diff(msg, tt.want))
			}

		})
	}
}


