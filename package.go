package catfruits

import "regexp"

var DefaultPakcageFuncs = map[string]PackageFunc{
	"Gem": pkgGem,
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
