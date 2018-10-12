package database

type Config struct {
	Driver   string
	Debug    bool
	Username string
	Password string
	Host     string
	Port     uint
	Database string
	Params   string
}
