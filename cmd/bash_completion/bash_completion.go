// usage: ./bash_completion <FILE>
package main

import (
	"fmt"
	"os"

	"github.com/hpcng/warewulf/internal/wwctl"
)

func main() {
	fh, err := os.Create(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	defer fh.Close()

	if err := root.GenBashCompletion(fh); err != nil {
		fmt.Println(err)
		return
	}
}
