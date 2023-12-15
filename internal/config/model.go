package config

type Config struct {
	Server Server
	DB     Database
}

type Database struct {
	Host string
	Port string
	User string
	Pass string
	Name string
}

type Server struct {
	Host string
	Port string
}
