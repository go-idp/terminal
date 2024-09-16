package core

// Config is the configuration for the server
type Config struct {
	// Port is the port the server listens on
	Port int `config:"port"`

	// Path is the path the server listens on
	Path string

	// Username is the username for basic auth
	Username string
	// Password is the password for basic auth
	Password string

	// Driver is the Driver runtime, options: host, docker, kubernetes, ssh, default: host
	Driver      string
	DriverImage string
}
