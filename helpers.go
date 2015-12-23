package catfruits

import (
	"io/ioutil"
	"os"
	"path/filepath"
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
