package getresponse

import "testing"

func TestParseEnvironment(t *testing.T) {
	tests := []struct {
		in      string
		want    Environment
		wantErr bool
	}{
		{"SMB", EnvSMB, false},
		{"smb", EnvSMB, false},
		{"PL", EnvMaxEU, false},
		{" pl ", EnvMaxEU, false},
		{"US", EnvMaxUS, false},
		{"", "", true},
		{"EU", "", true},
		{"unknown", "", true},
	}
	for _, tt := range tests {
		got, err := ParseEnvironment(tt.in)
		if tt.wantErr {
			if err == nil {
				t.Errorf("ParseEnvironment(%q): expected error, got nil", tt.in)
			}
			continue
		}
		if err != nil {
			t.Errorf("ParseEnvironment(%q): unexpected error: %v", tt.in, err)
			continue
		}
		if got != tt.want {
			t.Errorf("ParseEnvironment(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}

func TestEnvironmentEndpoint(t *testing.T) {
	cases := map[Environment]string{
		EnvSMB:   endpointSMB,
		EnvMaxEU: endpointMaxEU,
		EnvMaxUS: endpointMaxUS,
	}
	for env, want := range cases {
		if got := env.endpoint(); got != want {
			t.Errorf("%q.endpoint() = %q, want %q", env, got, want)
		}
	}
}

func TestNewValidation(t *testing.T) {
	if _, err := New("", EnvSMB); err == nil {
		t.Error("New with empty key: expected error")
	}
	if _, err := New("key", Environment("bogus")); err == nil {
		t.Error("New with invalid env: expected error")
	}
	c, err := New("key", EnvMaxUS, WithMailingDomain("example.com"))
	if err != nil {
		t.Fatalf("New: unexpected error: %v", err)
	}
	if c.endpoint != endpointMaxUS {
		t.Errorf("endpoint = %q, want %q", c.endpoint, endpointMaxUS)
	}
	if c.mailingDomain != "example.com" {
		t.Errorf("mailingDomain = %q, want example.com", c.mailingDomain)
	}
	if c.httpClient == nil || c.httpClient.Timeout != defaultTimeout {
		t.Errorf("default http client not configured: %+v", c.httpClient)
	}
}
