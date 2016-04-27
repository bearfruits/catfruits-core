package catfruits

import (
	"encoding/json"
	"regexp"
)

// DefaultPakcageFuncs is default package set
var DefaultPakcageFuncs = map[string]PackageFunc{
	"gem":   pkgGem,
	"npm":   pkgNPM,
	"bower": pkgBower,
}

func pkgGem(sc *Scanner) ([]string, error) {
	if !sc.FileExist("Gemfile") {
		return nil, nil
	}
	b, err := sc.ReadFile("Gemfile")
	if err != nil {
		return nil, err
	}
	re := regexp.MustCompile(`(?m)^\s*gem\s+['"]([^'"]+)['"]`)
	m := re.FindAllSubmatch(b, -1)
	res := make([]string, 0, len(m))
	for _, v := range m {
		res = append(res, string(v[1]))
	}
	return res, nil
}

type packageJSON struct {
	Dependencies    map[string]string `json:"dependencies"`
	DevDependencies map[string]string `json:"devDependencies"`
}

func pkgJSON(sc *Scanner, filename string) ([]string, error) {
	if !sc.FileExist(filename) {
		return nil, nil
	}
	b, err := sc.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	pj := &packageJSON{}

	if err := json.Unmarshal(b, pj); err != nil {
		return nil, err
	}

	res := make([]string, 0, len(pj.Dependencies)+len(pj.DevDependencies))
	for name := range pj.Dependencies {
		res = append(res, name)
	}
	for name := range pj.DevDependencies {
		res = append(res, name)
	}
	return res, nil

}

func pkgNPM(sc *Scanner) ([]string, error) {
	return pkgJSON(sc, "package.json")
}

func pkgBower(sc *Scanner) ([]string, error) {
	return pkgJSON(sc, "bower.json")
}
