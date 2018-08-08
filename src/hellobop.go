package hellobop

import "fmt"

func Hello(name string) string {
	return fmt.Sprintf("Hello %s from binary-only package!", name)
}
