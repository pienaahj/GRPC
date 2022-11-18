package main

import (
	"log"

	"github.com/spf13/viper"
)

// Get the config and return it as a connection string
func config() (string, string) {

	//  Log the function call
	log.Println("Config was invoked")

	// Set the file name of the configurations file
	viper.SetConfigName("config")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yaml")

	// Set the path to look for the configurations file
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME/config")

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Error reading config file, %s", err)
	}

	// Set undefined variables
	viper.SetDefault("DB.HOST", "127.0.0.1")

	// getting env variables DB.PORT
	// viper.Get() returns an empty interface{}
	// so we have to do the type assertion, to get the value
	DBPort, ok := viper.Get("DB.PORT").(string)

	// if type assert is not valid it will throw an error
	if !ok {
		log.Fatalf("Invalid type assertion")
	}

	log.Printf("viper : %s = %s \n", "Database Port", DBPort)

	//  build the mongodb config URI "mongodb://user:password@localhost:27017/"
	buildConfigString := func() string {
		buildString := "mongodb://"
		//  get the port from the config
		DBPort, ok := viper.Get("DB.PORT").(string)

		// if type assert is not valid it will throw an error
		if !ok {
			log.Fatalf("Invalid type assertion : port")
		}

		//  get the host from the config
		DBHost, ok := viper.Get("DB.HOST").(string)

		// if type assert is not valid it will throw an error
		if !ok {
			log.Fatalf("Invalid type assertion : host")
		}

		// Get the user
		DBUser, ok := viper.Get("DB.USERNAME").(string)

		// if type assert is not valid it will throw an error
		if !ok {
			log.Fatalf("Invalid type assertion : user")
		}
		// Get the password
		DBPass, ok := viper.Get("DB.PASWORD").(string)

		// if type assert is not valid it will throw an error
		if !ok {
			log.Fatalf("Invalid type assertion : password")
		}

		buildString += DBUser + ":" + DBPass + "@" + DBHost + ":" + DBPort + "/"
		return buildString
	}

	buildGRPCString := func() string {
		var buildString string
		// Get the user
		buildString, ok := viper.Get("GRPC.ADDRESS").(string)

		// if type assert is not valid it will throw an error
		if !ok {
			log.Fatalf("Invalid type assertion : address")
		}
		return buildString
	}
	return buildConfigString(), buildGRPCString()
}
