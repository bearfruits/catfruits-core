package catfruits

// DefaultFlameworkFuncs is default framework set
var DefaultFlameworkFuncs = map[string]FlameworkFunc{
	"Ruby on Rails": FWRails,
}

// FWRails detects Ruby on Rails
func FWRails(sc *Scanner) (bool, error) {
	return sc.FileExist("script/rails") || sc.FileExist("bin/rails"), nil
}
