package catfruits

import (
	"encoding/json"
	"regexp"
)

// DefaultPakcageFuncs is default package set
var DefaultPakcageFuncs = map[string]PackageFunc{
	"gem": pkgGem,
	"npm": pkgNPM,
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

func pkgNPM(sc *Scanner) ([]string, error) {
	if !sc.FileExist("package.json") {
		return nil, nil
	}
	b, err := sc.ReadFile("package.json")
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
