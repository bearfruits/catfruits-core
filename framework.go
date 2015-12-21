package repo

var DefaultFlameworkFuncs = map[string]FlameworkFunc{
	"Ruby on Rails": FWRails,
}

func FWRails(sc *RepoScanner) (bool, error) {
	return sc.FileExist("script/rails") || sc.FileExist("bin/rails"), nil
}
