package config

import (
	"github.com/spf13/viper"
)

var cfg *config

type config struct {
	API APIConfig
	DB  DBConfig
}

type APIConfig struct {
	Port string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Pass     string
	DataBase string
}

func init() {
	//	carregaEnv()
	//Aqui as configurações DEFAULT do bando de dados
	viper.SetDefault("api.port", "9000")
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "5432")
}

func Load() error {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	cfg = new(config)
	cfg.API = APIConfig{
		Port: viper.GetString("api.port"),
	}
	cfg.DB = DBConfig{
		/*             	//POSTGRES
		Host:     viper.GetString("postgres.host"),
		Port:     viper.GetString("postgres.port"),
		User:     viper.GetString("postgres.user"),
		Pass:     viper.GetString("postgres.pass"),
		DataBase: viper.GetString("postgres.name"),
		*/
		//MYSQL
		Host:     viper.GetString("mysql.host"),
		Port:     viper.GetString("mysql.port"),
		User:     viper.GetString("mysql.user"),
		Pass:     viper.GetString("mysql.pass"),
		DataBase: viper.GetString("mysql.name"),
	}

	return nil
}

func GetDB() DBConfig {
	return cfg.DB
}

func GetServerPort() string {
	return cfg.API.Port
}

/*func carregaEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Print("Erro ao carregar arquivo .env: ", err.Error())
	}
}

func retornaEnv(key string) string {
	return os.Getenv(key)
}*/
