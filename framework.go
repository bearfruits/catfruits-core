package catfruits

// DefaultFrameworkFuncs is default framework set
var DefaultFrameworkFuncs = map[string]FrameworkFunc{
	"Ruby on Rails": FWRails,
}

// FWRails detects Ruby on Rails
func FWRails(sc *Scanner) (bool, error) {
	return sc.FileExist("script/rails") || sc.FileExist("bin/rails"), nil
}
