package catfruits

import "testing"

func TestScannerHasPackage(t *testing.T) {
	sc := NewScanner("aaa")
	sc.info.Packages["gem"] = []string{"hoge", "rails", "foo", "bar"}
	if !sc.hasPackage("gem", "rails") {
		t.Errorf("Expected true, but got false")
	}

	sc.info.Packages["gem"] = []string{"poyo"}
	if sc.hasPackage("gem", "rails") {
		t.Errorf("Expected false, but got true")
	}
}
