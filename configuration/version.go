package configuration

import "fmt"

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func Info() string {
	return fmt.Sprintf(`Version: %v
Commit: %v
Date: %v`, version, commit, date)
}
