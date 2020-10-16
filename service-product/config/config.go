package config

type Config struct {
	Port        string
	AuthService AuthService `mapstructure:"auth_service"`
	Database    Database
}

type Database struct {
	Driver string `mapstructure:"driver"`
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
	User string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName string `mapstructure:"db_name"`
	Config string `mapstructure:"config"`
}
