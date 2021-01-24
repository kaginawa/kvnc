package kvnc

import (
	"fmt"
	"io"
	"os"
)

// SafeClose closes closer safely.
func SafeClose(closer io.Closer, name string) {
	if err := closer.Close(); err != nil {
		if err.Error() == "EOF" {
			return
		}
		if _, err := fmt.Fprintf(os.Stderr, "failed to close %s: %v\n", err, name); err != nil {
			fmt.Printf("failed to close %s: %v\n", err, name)
		}
	}
}
