package config

import (
	"reflect"
	"testing"
)

func TestAddress(t *testing.T) {
	type args struct {
		addr string
	}
	tests := []struct {
		name string
		args args
		want string // Changed to string since we can't compare functions directly
	}{
		{
			name: "sets address",
			args: args{addr: "localhost:8080"},
			want: "localhost:8080",
		},
		{
			name: "sets empty address",
			args: args{addr: ""},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Address(tt.args.addr)
			// Test the function by applying it to an empty config
			result := got(Config{})
			if result.Addr != tt.want {
				t.Errorf("Address() result.Addr = %v, want %v", result.Addr, tt.want)
			}
		})
	}
}

func TestConfig_IsValid(t *testing.T) {
	type fields struct {
		Addr string
	}
	tests := []struct {
		name    string
		fields  fields
		want    bool
		wantErr bool
	}{
		{
			name:    "valid config with address",
			fields:  fields{Addr: "localhost:8080"},
			want:    true,
			wantErr: false,
		},
		{
			name:    "invalid config with empty address",
			fields:  fields{Addr: ""},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Config{
				Addr: tt.fields.Addr,
			}
			got, err := c.IsValid()
			if (err != nil) != tt.wantErr {
				t.Errorf("IsValid() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsValid() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewConfig(t *testing.T) {
	type args struct {
		opts []Option
	}
	tests := []struct {
		name string
		args args
		want Config
	}{
		{
			name: "no options returns empty config",
			args: args{opts: []Option{}},
			want: Config{},
		},
		{
			name: "single address option",
			args: args{opts: []Option{Address("localhost:8080")}},
			want: Config{Addr: "localhost:8080"},
		},
		{
			name: "multiple options - last one wins",
			args: args{opts: []Option{
				Address("first:8080"),
				Address("second:9090"),
			}},
			want: Config{Addr: "second:9090"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewConfig(tt.args.opts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_missingConfigField(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		errMsg  string
	}{
		{
			name:    "returns error with field name",
			args:    args{name: "Address"},
			wantErr: true,
			errMsg:  "configuration field 'Address' is required",
		},
		{
			name:    "returns error with different field name",
			args:    args{name: "Port"},
			wantErr: true,
			errMsg:  "configuration field 'Port' is required",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := missingConfigField(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("missingConfigField() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && err.Error() != tt.errMsg {
				t.Errorf("missingConfigField() error message = %v, want %v", err.Error(), tt.errMsg)
			}
		})
	}
}
