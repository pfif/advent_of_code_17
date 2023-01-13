package input

import (
	"embed"
	"fmt"
	"os"
)

//go:embed *.txt
var fs embed.FS

// GetInput opens input file, finding its input in the command line
// and returns each lines splitted
func GetInput() ([]byte, error) {
	if len(os.Args) < 2 {
		return nil, fmt.Errorf("No file name provided as command line argument")
	}

	file, err := fs.ReadFile(os.Args[1])
	if err != nil {
		return nil, fmt.Errorf("Could not read input file: %w", err)
	}

	return file, nil
}
