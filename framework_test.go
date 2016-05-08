package catfruits

import (
	"fmt"
	"os/exec"
	"testing"
)

func TestFWRails(t *testing.T) {
	fwTestHelper("Ruby on Rails", "rails/Gemfile", t)
}

func TestFWVue(t *testing.T) {
	fwTestHelper("Vue", "vue/npm", t)
}

func TestFWPadrino(t *testing.T) {
	fwTestHelper("Padrino", "padrino", t)
}

func fwTestHelper(name, branch string, t *testing.T) {
	exec.Command("bash", "-c", fmt.Sprintf("cd %s; git checkout %s", sandboxDir, branch)).Run()
	sc := NewScanner(sandboxDir)

	i, err := sc.Scan()
	if err != nil {
		t.Error(err)
	}
	if !include(i.Frameworks, name) {
		t.Errorf("Branch %s should be %s, but isn't", branch, name)
	}
}

func include(s []string, v string) bool {
	for _, a := range s {
		if a == v {
			return true
		}
	}
	return false
}
