package catfruits

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// ----------- Helper functions

// FilePath returns a filepath
func (sc *Scanner) FilePath(path string) string {
	return filepath.Join(sc.dir, path)
}

// FileExist returns a bool
func (sc *Scanner) FileExist(path string) bool {
	fname := sc.FilePath(path)
	_, err := os.Stat(fname)
	return err == nil
}

// ReadFile returns a file value
func (sc *Scanner) ReadFile(path string) ([]byte, error) {
	return ioutil.ReadFile(sc.FilePath(path))
}

func (sc *Scanner) hasPackage(pkg, name string) bool {
	pkgs, ok := sc.info.Packages[pkg]
	if !ok {
		return false
	}

	for _, v := range pkgs {
		if v == name {
			return true
		}
	}
	return false
}

func (sc *Scanner) FileContains(path, content string) bool {
	b, err := sc.ReadFile(path)
	if err != nil {
		return false
	}
	return strings.Contains(string(b), content)
}
