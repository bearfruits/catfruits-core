package repo

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

type Infomation struct {
	Flameworks []string `json:"flameworks"`
}

type RepoScanner struct {
	dir    string
	fwFunc map[string]FlameworkFunc
}

type FlameworkFunc func(*RepoScanner) (ok bool, err error)

func NewScanner(dir string) *RepoScanner {
	return &RepoScanner{
		dir:    dir,
		fwFunc: DefaultFlameworkFuncs,
	}
}

func (sc *RepoScanner) Scan() (*Infomation, error) {
	info := &Infomation{
		Flameworks: make([]string, 0),
	}

	for name, f := range sc.fwFunc {
		ok, err := f(sc)
		if err != nil {
			return nil, err
		}
		if !ok {
			continue
		}

		info.Flameworks = append(info.Flameworks, name)
	}

	return info, nil
}

// ----------- Helper functions

func (sc *RepoScanner) FilePath(path string) string {
	return filepath.Join(sc.dir, path)
}

func (sc *RepoScanner) FileExist(path string) bool {
	fname := sc.FilePath(path)
	_, err := os.Stat(fname)
	return err == nil
}

func (sc *RepoScanner) ReadFile(path string) ([]byte, error) {
	return ioutil.ReadFile(sc.FilePath(path))
}
