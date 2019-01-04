package plugin

import (
	"strings"

	"github.com/Masterminds/semver"
	"github.com/pkg/errors"
)

// Flag is the unparsed flag value from <plugin-flag>=<impl>[@version]
// For example, --plugin-metrics=prometheus@1.0.0
type Flag struct {
	Val string
}

const (
	// NoImplementation is a placeholder for a plugin with no
	// implementation defined
	NoImplementation = "none"
)

// Type represents a type of plugin
type Type int

const (
	// TypeCNI is the CNI plugin type
	TypeCNI Type = iota
	// TypeCSI is the CSI plugin type
	TypeCSI
	// TypeCCM is the CCM plugin type
	TypeCCM
	// TypeClusterManagement is the ClusterManagement plugin type
	TypeClusterManagement
	// TypeAutoscaler is the Autoscaler plugin type
	TypeAutoscaler
	// TypeMetrics is the Metrics plugin type
	TypeMetrics
	// TypeLogs is the Logs plugin type
	TypeLogs
)

func (p Type) String() string {
	names := [...]string{
		"Container Network Interface (CNI)",
		"Container Storage Interface (CSI)",
		"Cloud Controller Manager (CCM)",
		"Cluster Management",
		"Autoscaler",
		"Metrics",
		"Logs",
	}

	if p < TypeCNI || p > TypeLogs {
		return "Unknown"
	}

	return names[p]
}

// Parse parses the Flag into an implementation and version or returns an error.
// If the flag is well-formed but no implementation is provided, empty strings
// are returned with no error so the caller may default values if desired.
// If the flag is well-formed and an implementation is provided but no version,
// then the implementation is returned along with an empty string for version.
// If the flag is malformed, an error is returned.
func (f Flag) Parse() (string, string, error) {
	fields := strings.SplitN(f.Val, "@", 2)
	impl := fields[0]

	if impl == "" {
		// No implementation, that's ok (caller should default)
		return "", "", nil
	}

	if len(fields) == 1 {
		// No version provided, that's ok (caller should default)
		return impl, "", nil
	}

	// Must have two fields if we got to here
	version, err := semver.NewVersion(fields[1])
	if err != nil {
		return "", "", errors.Wrap(err, "parsing semver")
	}

	return impl, version.String(), nil
}
