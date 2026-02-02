package config

type EolConfig struct {
	Config []ProductConfig `yaml:"config"`
}

type ProductConfig struct {
	Id      string `yaml:"id"`
	Product string `yaml:"product"`
	Path    string `yaml:"path"`
	Method  string `yaml:"method"`
}
