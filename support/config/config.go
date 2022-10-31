package config

const (
	LibsPath         = "support/libs/"
	AssetsStylePath  = "assets/css/"
	AssetsScriptPath = "assets/js/"
	AssetsMapPath    = "assets/json/"
)

func AvailableStyles() []string {
	return []string{"sample", "bootstrap", "bulma"}
}
