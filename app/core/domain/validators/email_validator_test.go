package validators

import (
	"strings"
	"testing"
)

func TestValidatorEmail_Validate(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "valid email",
			args: args{
				email: "test@example.com",
			},
			wantErr: false,
		},
		{
			name: "empty email",
			args: args{
				email: "",
			},
			wantErr: true,
		},
		{
			name: "long email",
			args: args{
				email: "a@b." + strings.Repeat("c", 256),
			},
			wantErr: true,
		},
		{
			name: "invalid email missing @",
			args: args{
				email: "testexample.com",
			},
			wantErr: true,
		},
		{
			name: "invalid email with multiple @",
			args: args{
				email: "test@ex@mple.com",
			},
			wantErr: true,
		},
		{
			name: "Invalid email has only one char before @",
			args: args{
				email: "t@ex@mple.com",
			},
			wantErr: true,
		},

		{
			name: "Invalid email has more than 64 char before @",
			args: args{
				email: strings.Repeat("a", 100) + "@ex@mple.com",
			},
			wantErr: true,
		},

		{
			name: "Invalid Email domain does not cotain a dot",
			args: args{
				email: "pmello@examplecom",
			},
			wantErr: true,
		},

		{
			name: "invalid email with invalid characters",
			args: args{
				email: "test@example.c#m",
			},
			wantErr: true,
		},
		{
			name: "invalid email with subdomain longer than 63 characters",
			args: args{
				email: "test@subdomain." + strings.Repeat("a", 64),
			},
			wantErr: true,
		},
		{
			name: "invalid email with tld longer than 6 characters",
			args: args{
				email: "test@example.longtld",
			},
			wantErr: true,
		},
		{
			name: "invalid email with subdomain starting with hyphen",
			args: args{
				email: "test@-subdomain.example.com",
			},
			wantErr: true,
		},
		{
			name: "invalid email with subdomain ending with hyphen",
			args: args{
				email: "test@subdomain-.example.com",
			},
			wantErr: true,
		},
		{
			name: "invalid email with domain starting with hyphen",
			args: args{
				email: "test@example.-com",
			},
			wantErr: true,
		},
		{
			name: "invalid email with domain ending with hyphen",
			args: args{
				email: "test@example.com-",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := ValidatorEmail{}
			if err := v.Validate(tt.args.email); (err != nil) != tt.wantErr {
				t.Errorf("ValidatorEmail.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
