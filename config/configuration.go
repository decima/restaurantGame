package config

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"net/url"
	"os"
	"restaurantAPI/lib/database"
	"restaurantAPI/lib/database/mongoAdapter"
	"strings"
)

const (
	serverPort = "server.port"
	serverHost = "server.host"
	dbEngine   = "db.engine"
	dbHost     = "db.host"
	dbPort     = "db.port"
	dbUser     = "db.user"
	dbPassword = "db.password"
	dbName     = "db.name"
)

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	viper.SetDefault(serverPort, 8080)
	viper.SetDefault(serverHost, "localhost")
	viper.SetDefault(dbEngine, "mongodb")
	viper.SetDefault(dbHost, "localhost")
	viper.SetDefault(dbPort, 27017)
	viper.SetDefault(dbUser, "root")
	viper.SetDefault(dbPassword, "root")
	viper.SetDefault(dbName, "restaurant")

	viper.SetConfigName("restaurant")
	viper.AddConfigPath(".")
	viper.AddConfigPath("/etc/restaurant")
	viper.AddConfigPath("$HOME/.restaurant")
	viper.SetConfigType("yaml")

	viper.SetEnvPrefix("restaurant")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

}

func Load() {

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			err := viper.WriteConfigAs("./restaurant.yaml")
			if err != nil {
				panic("Config file was not found and could not be created; aborting")
			}
			log.Info().Msg("Config file was not found and was created; please fill in the necessary information and restart the application.")
			os.Exit(0)
		} else {
			panic("Config file was found but another error was produced; aborting")
		}
	}
}

func HostAndPort() string {
	return fmt.Sprintf("%s:%d", viper.GetString(serverHost), viper.GetInt(serverPort))
}
func DatabaseURI() url.URL {
	dbURL := url.URL{
		Scheme: viper.GetString(dbEngine),
		User:   url.UserPassword(viper.GetString(dbUser), viper.GetString(dbPassword)),
		Host:   fmt.Sprintf("%s:%d", viper.GetString(dbHost), viper.GetInt(dbPort)),
		Path:   viper.GetString(dbName),
	}
	return dbURL
}

var databaseInstance database.Client

func Database() (database.Client, error) {

	var err error

	if databaseInstance != nil {
		return databaseInstance, nil
	}

	switch viper.GetString(dbEngine) {
	case "mongodb":
		databaseInstance = mongoAdapter.NewClient(DatabaseURI())

	default:
		return nil, fmt.Errorf("unsupported database engine: %s", viper.GetString(dbEngine))
	}

	return databaseInstance, err

}

func Collection[T database.Entity](name database.CollectionName) (database.Collection[T], error) {
	db, err := Database()
	if err != nil {
		return nil, err
	}

	switch viper.GetString(dbEngine) {
	case "mongodb":
		return mongoAdapter.GetCollection[T](db.(*mongoAdapter.Client), name), nil
	}
	return nil, fmt.Errorf("unsupported database engine: %s", viper.GetString(dbEngine))
}
