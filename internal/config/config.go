package config

import (
	"github.com/spf13/viper"
)

type DBConfig struct {
	Host     string
	Port     string
	Database string
	User     string
	Password string
}

func GetConfig(path string) (DBConfig, error) {
	v := viper.New()
	c := DBConfig{}

	v.SetConfigName("config")
	v.AddConfigPath(path)

	err := v.ReadInConfig()
	if err != nil {
		return c, err
	}

	c.User = v.GetString("db.user")
	c.Password = v.GetString("db.password")
	c.Host = v.GetString("db.host")
	c.Database = v.GetString("db.database")
	c.Port = v.GetString("db.port")

	return c, nil
}
