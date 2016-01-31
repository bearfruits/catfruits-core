package catfruits

// Infomation is a repository infomation.
type Infomation struct {
	Frameworks []string            `json:"frameworks"`
	Packages   map[string][]string `json:"packages"`
}

// Scanner is scanner for repository.
type Scanner struct {
	dir    string
	fwFunc map[string]FrameworkFunc
	info   *Infomation
}

// FrameworkFunc is a function that detects framework
type FrameworkFunc func(*Scanner) (ok bool, err error)

// PackageFunc detects any packages
type PackageFunc func(*Scanner) (pkgs []string, err error)

// NewScanner returns a Scanner
func NewScanner(dir string) *Scanner {
	return &Scanner{
		dir:    dir,
		fwFunc: DefaultFrameworkFuncs,
		info: &Infomation{
			Frameworks: make([]string, 0),
			Packages:   make(map[string][]string),
		},
	}
}

// Scan scans a repository.
func (sc *Scanner) Scan() (*Infomation, error) {
	info := sc.info
	for name, f := range DefaultPakcageFuncs {
		pkgs, err := f(sc)
		if err != nil {
			return nil, err
		}
		info.Packages[name] = pkgs
	}

	for name, f := range sc.fwFunc {
		ok, err := f(sc)
		if err != nil {
			return nil, err
		}
		if !ok {
			continue
		}

		info.Frameworks = append(info.Frameworks, name)
	}

	return info, nil
}
