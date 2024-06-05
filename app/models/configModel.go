package models

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
)

type app struct {
	Name string `yaml:"name"`
	Addr string `yaml:"addr"`
	Port int    `yaml:"port"`
	Ssl  bool   `yaml:"ssl"`
}
type jwt struct {
	Realm            string `yaml:"realm"`
	Key              string `yaml:"key"`
	Timeout          int64  `yaml:"timeout"`
	MaxRefresh       int64  `yaml:"maxRefresh"`
	SendCookie       bool   `yaml:"sendCookie"`
	TokenLookup      string `yaml:"tokenLookup"`
	SigningAlgorithm string `yaml:"signingAlgorithm"`
	TokenHeadName    string `yaml:"tokenHeadName"`
	CookieName       string `yaml:"cookieName"`
}
type db struct {
	DbDriver                  string `yaml:"dbDriver"`
	Dsn                       string `yaml:"dsn"`
	DefaultStringSize         uint   `yaml:"defaultStringSize"`
	DisableDatetimePrecision  bool   `yaml:"disableDatetimePrecision"`
	DontSupportRenameIndex    bool   `yaml:"dontSupportRenameIndex"`
	DontSupportRenameColumn   bool   `yaml:"dontSupportRenameColumn"`
	SkipInitializeWithVersion bool   `yaml:"skipInitializeWithVersion"`
}

type Config struct {
	App app `yaml:"app"`
	Jwt jwt `yaml:"jwt"`
	Db  db  `yaml:"db"`
}

func (c Config) LoadConfigForYaml() (Config, error) {
	f, err := os.Open("config/config.yml")
	if err != nil {
		fmt.Printf("loadConfigForYaml os.Open err: %+v\n", err)
		return Config{}, err
	}
	defer f.Close()

	err = yaml.NewDecoder(f).Decode(&c)
	return c, err
}

func (c Config) SetDbConfig() mysql.Config {
	{
		rc := mysql.Config{
			DSN:                       c.Db.Dsn,
			DefaultStringSize:         c.Db.DefaultStringSize,
			DisableDatetimePrecision:  c.Db.DisableDatetimePrecision,
			DontSupportRenameIndex:    c.Db.DontSupportRenameIndex,
			DontSupportRenameColumn:   c.Db.DontSupportRenameColumn,
			SkipInitializeWithVersion: c.Db.SkipInitializeWithVersion,
		}
		return rc
	}
}
