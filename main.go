// Extract the BootKey from an offline system hive.
//
// Usage:
//
//	bootkey system
//
// The arguments are:
//
//	system
//	    The system registry hive (required).
package main

import (
	"fmt"
	"os"

	"go.foxforensics.dev/bootkey/bootkey"
)

var Usage = `© 2026 Fox Forensics. Licensed under MIT License.
Usage: bootkey SYSTEM

Report bugs at: foxforensics.dev/issues`

func main() {
	if len(os.Args) == 1 || os.Args[1] == "--help" {
		_, _ = fmt.Fprintln(os.Stderr, Usage)
		os.Exit(2)
	}

	key, err := bootkey.ReadFile(os.Args[1])

	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Printf("%x\n", key)
}
