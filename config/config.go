package config

import "github.com/spf13/viper"

type confInfo struct {
	Name string
	Type string
	Path string
}
type config struct {
	viper *viper.Viper
}

var (
	Conf   *config
	Secret *config
)

func init() {
	si := confInfo{
		Name: "secrets",
		Type: "yaml",
		Path: "config/secret",
	}
	ci := confInfo{
		Name: "confs",
		Type: "yaml",
		Path: "config/conf",
	}
	Conf = &config{getConf(ci)}
	Secret = &config{getConf(si)}
}

// viper用于解析yaml
func getConf(si confInfo) *viper.Viper {
	v := viper.New()
	v.SetConfigName(si.Name) // 与yaml文件名一致
	v.SetConfigType(si.Type)
	v.AddConfigPath(si.Path)
	v.ReadInConfig()
	return v
}

func (c *config) GetString(key string) string {
	return c.viper.GetString(key)
}
