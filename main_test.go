package catfruits

import (
	"io/ioutil"
	"os"
	"os/exec"
	"testing"
)

func TestMain(m *testing.M) {
	exitCode, err := testMain(m)
	if err != nil {
		panic(err)
	}
	os.Exit(exitCode)
}

var sandboxDir string

func testMain(m *testing.M) (int, error) {
	d, err := ioutil.TempDir("", "catfruits-test")
	if err != nil {
		return 1, err
	}
	sandboxDir = d
	// defer os.RemoveAll(sandboxDir)

	err = exec.Command("git", "clone", "https://github.com/bearfruits/catfruits_sandbox", sandboxDir).Run()
	if err != nil {
		return 1, err
	}

	return m.Run(), nil
}
