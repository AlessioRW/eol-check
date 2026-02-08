package config

type EolConfig struct {
	Config []ProductConfig `yaml:"config"`
}

type ProductConfig struct {
	Id      string `yaml:"id"`
	Product string `yaml:"product"`
	Method  string `yaml:"method"`
	Args    []any  `yaml:"args"`
}
