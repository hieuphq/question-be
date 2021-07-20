package config

import "github.com/spf13/viper"

// Config environment configuration when run application
type Config struct {
	Env       string
	Port      string
	DBHost    string
	DBPort    string
	DBUser    string
	DBName    string
	DBPass    string
	DBSSLMode string
}

// DefaultConfigLoaders is default loader list
func defaultConfigLoaders() []Loader {
	loaders := []Loader{}
	fileLoader := NewFileLoader(".env", ".")
	loaders = append(loaders, fileLoader)
	loaders = append(loaders, NewENVLoader())

	return loaders
}

func LoadConfig() Config {
	v := viper.New()
	v.SetDefault("PORT", "3000")
	v.SetDefault("ENV", "dev")

	for _, loader := range defaultConfigLoaders() {
		newV, err := loader.Load(*v)

		if err == nil {
			v = newV
		}
	}
	return generateConfigFromViper(v)
}

// generateConfigFromViper generate config from viper data
func generateConfigFromViper(v *viper.Viper) Config {

	return Config{
		Port: v.GetString("PORT"),
		Env:  v.GetString("ENV"),

		DBHost:    v.GetString("DB_HOST"),
		DBPort:    v.GetString("DB_PORT"),
		DBUser:    v.GetString("DB_USER"),
		DBName:    v.GetString("DB_NAME"),
		DBPass:    v.GetString("DB_PASS"),
		DBSSLMode: v.GetString("DB_SSL_MODE"),
	}
}
