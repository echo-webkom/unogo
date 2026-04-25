package unogo

import "fmt"

func unoError(code int, format string, args ...any) error {
	return fmt.Errorf("uno responded with non-success code (%d): %s", code, fmt.Sprintf(format, args...))
}
