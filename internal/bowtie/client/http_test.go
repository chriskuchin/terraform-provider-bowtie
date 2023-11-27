package client

import (
	"net/http"
	"testing"
)

func TestClient_getHostURL(t *testing.T) {
	type fields struct {
		HTTPClient *http.Client
		hostURL    string
		auth       AuthPayload
	}
	type args struct {
		path string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "login",
			fields: fields{
				HTTPClient: nil,
				hostURL:    "https://example.com",
				auth: AuthPayload{
					Username: "test@example.com",
					Password: "passw0rd123",
				},
			},
			args: args{
				path: "/user/login",
			},
			want: "https://example.com/-net/api/v0/user/login",
		},
		{
			name: "local login",
			fields: fields{
				HTTPClient: nil,
				hostURL:    "http://localhost:3000",
				auth: AuthPayload{
					Username: "test@example.com",
					Password: "passw0rd123",
				},
			},
			args: args{
				path: "/user/login",
			},
			want: "http://localhost:3000/-net/api/v0/user/login",
		},
		{
			name: "me",
			fields: fields{
				HTTPClient: nil,
				hostURL:    "https://example.com",
				auth: AuthPayload{
					Username: "test@example.com",
					Password: "passw0rd123",
				},
			},
			args: args{
				path: "/user/me",
			},
			want: "https://example.com/-net/api/v0/user/me",
		},
		{
			name: "error",
			fields: fields{
				HTTPClient: nil,
				hostURL:    "https://example.com",
				auth: AuthPayload{
					Username: "test@example.com",
					Password: "passw0rd123",
				},
			},
			args: args{
				path: "user/me",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				HTTPClient: tt.fields.HTTPClient,
				hostURL:    tt.fields.hostURL,
				auth:       tt.fields.auth,
			}
			if got := c.getHostURL(tt.args.path); got != tt.want {
				t.Errorf("Client.getHostURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
