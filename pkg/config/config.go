package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Port       string     `yaml:"port"`
	Database   Database   `yaml:"database"`
	JwtSecret  string     `yaml:"jwt_secret"`
	Redis      Redis      `yaml:"redis"`
	JwtExpTime int        `yaml:"jwt_expire_time"`
	Prothomeus Prometheus `yaml:"prothomeus"`
	Grafana    Grafana    `yaml:"grafana"`
}

type Database struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
	Migrate  bool   `yanl:"migrate"`
	SslMode  string `yaml:"sslmode"`
}

type Redis struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

type Prometheus struct {
	Port string `yaml:"port"`
}
type Grafana struct {
	Port                string `yaml:"port"`
	GfSecurityAdminPass string `yaml:"gf_security_admin_pass"`
}

// type Nats struct {
// 	Host string `yaml:"host"`
// 	Port string `yaml:"port"`
// }

func ReadValue() *Config {
	var configs Config
	filename, _ := filepath.Abs("./config.yaml")
	yamlFile, _ := os.ReadFile(filename)
	yaml.Unmarshal(yamlFile, &configs)
	return &configs
}
