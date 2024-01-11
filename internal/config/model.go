package config

type Config struct {
	Redis  Redis
	Server Server
	Email  Email
}

type Redis struct {
	Addr string
	Pass string
}

type Server struct {
	Host string
	Port string
}

type Email struct {
	Host string
	Port string
	User string
	Pass string
}
