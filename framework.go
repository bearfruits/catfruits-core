package catfruits

// DefaultFrameworkFuncs is default framework set
var DefaultFrameworkFuncs = map[string]FrameworkFunc{
	// Ruby
	"Ruby on Rails": fwRails,
	"Padrino":       fwPadrino,

	// JavaScript
	"Vue":      fwJSGen("vue"),
	"React":    fwJSGen("react"),
	"Angular":  fwJSGen("angular"),
	"backbone": fwJSGen("backbone"),

	// PHP
	"CakePHP": fwCakePHP,

	// JVM
	"Play Framework": fwPlayFramework,
}

// FWRails detects Ruby on Rails
func fwRails(sc *Scanner) (bool, error) {
	return sc.hasPackage("gem", "rails") ||
		sc.FileExist("script/rails") ||
		sc.FileExist("bin/rails"), nil
}

func fwPadrino(sc *Scanner) (bool, error) {
	return sc.FileContains("config/apps.rb", "Padrino") || sc.FileContains("config/boot.rb", "Padrino"), nil
}

// bower と npm から探索する関数をかえす
func fwJSGen(name string) FrameworkFunc {
	return func(sc *Scanner) (bool, error) {
		return sc.hasPackage("bower", name) ||
			sc.hasPackage("npm", name), nil
	}
}

func fwCakePHP(sc *Scanner) (bool, error) {
	return sc.FileExist("cake/bootstrap.php") ||
		sc.FileExist("lib/Cake/bootstrap.php"), nil
}

func fwPlayFramework(sc *Scanner) (bool, error) {
	return sc.FileExist("activator") ||
		sc.FileExist("activator.bat"), nil
}
