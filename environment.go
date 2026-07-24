package getresponse

import (
	"fmt"
	"strings"
)

// Environment identifies which GetResponse API host a client talks to.
//
// The underlying string values match the short codes historically used by
// callers ("SMB", "PL", "US"), so values persisted elsewhere map cleanly via
// ParseEnvironment.
type Environment string

const (
	// EnvSMB is the standard GetResponse (self-service) API.
	EnvSMB Environment = "SMB"
	// EnvMaxEU is the GetResponse MAX (360) EU API.
	EnvMaxEU Environment = "PL"
	// EnvMaxUS is the GetResponse MAX (360) US API.
	EnvMaxUS Environment = "US"
)

const (
	endpointSMB   = "https://api.getresponse.com/v3/"
	endpointMaxEU = "https://api3.getresponse360.pl/v3/"
	endpointMaxUS = "https://api3.getresponse360.com/v3/"
)

// endpoint returns the base API URL (with trailing slash) for the environment.
func (e Environment) endpoint() string {
	switch e {
	case EnvMaxEU:
		return endpointMaxEU
	case EnvMaxUS:
		return endpointMaxUS
	default:
		return endpointSMB
	}
}

// valid reports whether e is a known environment.
func (e Environment) valid() bool {
	switch e {
	case EnvSMB, EnvMaxEU, EnvMaxUS:
		return true
	default:
		return false
	}
}

// ParseEnvironment converts a short code ("SMB", "PL", "US", case-insensitive)
// into an Environment. Unlike the legacy behaviour it does not silently fall
// back to SMB: an unknown code returns an error.
func ParseEnvironment(s string) (Environment, error) {
	switch Environment(strings.ToUpper(strings.TrimSpace(s))) {
	case EnvSMB:
		return EnvSMB, nil
	case EnvMaxEU:
		return EnvMaxEU, nil
	case EnvMaxUS:
		return EnvMaxUS, nil
	default:
		return "", fmt.Errorf("getresponse: unknown environment %q (want SMB, PL or US)", s)
	}
}
