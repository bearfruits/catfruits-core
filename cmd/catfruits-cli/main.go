package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/bearfruits/catfruits-core"
)

func main() {
	for _, dir := range os.Args[1:] {
		sc := catfruits.NewScanner(dir)
		info, err := sc.Scan()
		if err != nil {
			panic(err)
		}

		fullpath, err := filepath.Abs(dir)
		if err != nil {
			panic(err)
		}
		fmt.Println(fullpath)
		b, err := json.MarshalIndent(info, "", "  ")
		if err != nil {
			panic(err)
		}
		fmt.Println(string(b))
	}
}
