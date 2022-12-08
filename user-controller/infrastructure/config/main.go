package config

import (
	"github.com/spf13/viper"
	"github.com/xo/dburl"
)

type DatabaseConfig struct {
	Driver   string // Драйвер БД
	Host     string // Хост БД
	Port     string // Порт БД
	Name     string // Имя БД
	Username string // Логин БД
	Password string // Пароль БД
}

// Конфигурация приложения
type Config struct {
	IsDebug bool   `mapstructure:"DEBUG"`
	DBUrl   string `mapstructure:"DATABASE_URL"` // Строка подключения к БД
	Port    int    `mapstructure:"PORT"`
	DB      DatabaseConfig
}

// Загрузка конфигурации из .env файла
func New(path string) (config Config) {
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err.Error())
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		panic(err.Error())
	}
	config.ParseDbUrl()
	viper.WatchConfig()
	return
}

// Разбиение строки подключения к БД на параметры
func (config *Config) ParseDbUrl() {
	url, err := dburl.Parse(config.DBUrl)
	if err != nil {
		panic(err.Error())
	}
	config.DB.Driver = url.Driver
	config.DB.Host = url.Hostname()
	config.DB.Port = url.Port()
	config.DB.Name = url.Path
	config.DB.Username = url.User.Username()
	config.DB.Password, _ = url.User.Password()
}
