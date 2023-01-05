package structs

type EnvConfig struct {
	DBConfig DBConfig `yaml:"db_config"`
}

type DBConfig struct {
	Host     string `yaml:"host"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}
