package catfruits

// Infomation is a repository infomation.
type Infomation struct {
	Frameworks []string `json:"frameworks"`
}

// Scanner is scanner for repository.
type Scanner struct {
	dir    string
	fwFunc map[string]FrameworkFunc
}

// FrameworkFunc is a function that detects framework
type FrameworkFunc func(*Scanner) (ok bool, err error)

// NewScanner returns a Scanner
func NewScanner(dir string) *Scanner {
	return &Scanner{
		dir:    dir,
		fwFunc: DefaultFrameworkFuncs,
	}
}

// Scan scans a repository.
func (sc *Scanner) Scan() (*Infomation, error) {
	info := &Infomation{
		Frameworks: make([]string, 0),
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
