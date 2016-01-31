package catfruits

import "testing"

func TestScannerHasPackage(t *testing.T) {
	sc := NewScanner("aaa")
	sc.info.Packages["Gem"] = []string{"hoge", "rails", "foo", "bar"}
	if !sc.hasPackage("Gem", "rails") {
		t.Errorf("Expected true, but got false")
	}

	sc.info.Packages["Gem"] = []string{"poyo"}
	if sc.hasPackage("Gem", "rails") {
		t.Errorf("Expected false, but got true")
	}
}
