package Database

type Config struct {
	Driver   string `yaml:"driver"`
	Debug    bool   `yaml:"debug"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     uint   `yaml:"port"`
	Database string `yaml:"database"`
	Params   string `yaml:"params"`
	Selected bool   `yaml:"selected"`
}
