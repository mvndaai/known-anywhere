package config

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
)

var c Config

func init() {
	c = Config{}
	processStruct(reflect.ValueOf(&c).Elem())
}

type (
	Config struct {
		Env          string `key:"ENVIRONMENT" default:"dev"`
		Postgres     postgres
		JWTSecret    string `key:"JWT_SECRET" required:"true"`
		DebugErrors  bool   `key:"DEBUG_ERRORS" default:"false"`
		Formated     formated
		FrontendPath string `key:"FRONTEND_PATH" default:"./bin/frontend"`
	}

	postgres struct {
		Host     string `key:"POSTGRES_HOST" default:"localhost"`
		Port     int    `key:"POSTGRES_PORT" default:"5432"`
		User     string `key:"POSTGRES_USER"`
		Password string `key:"POSTGRES_PASSWORD"`
		DBName   string `key:"POSTGRES_DB"`
		SSLMode  string `key:"POSTGRES_SSLMODE" default:"disable"`
	}

	formated struct {
		Port int `key:"PORT" default:"8080"`
	}
)

func Get() Config             { return c }
func (c Config) Port() string { return fmt.Sprintf(":%d", c.Formated.Port) }

func (v postgres) DataSourceName() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", v.Host, v.Port, v.User, v.Password, v.DBName, v.SSLMode)
}

func processStruct(v reflect.Value) {
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Anonymous {
			continue
		}

		// If it's a struct, process it recursively
		if field.Type.Kind() == reflect.Struct {
			processStruct(v.Field(i))
			continue
		}

		key := field.Tag.Get("key")
		defaultVal := field.Tag.Get("default")
		required, _ := strconv.ParseBool(field.Tag.Get("required"))

		// Get environment variable
		val := os.Getenv(key)
		if val == "" {
			val = defaultVal
		}
		if val == "" {
			if required {
				panic(fmt.Sprintf("Required environment variable '%s' not set", key))
			}
		}

		//log.Printf("key: %s, default: %s, required: %s, value: %s\n", key, defaultVal, required, val)
		// Set the field value
		if val != "" {
			switch field.Type.Kind() {
			case reflect.String:
				v.Field(i).SetString(val)
			case reflect.Int:
				iVal, err := strconv.Atoi(val)
				if err != nil {
					panic(fmt.Sprintf("Error converting '%s' for '%s'  to int: %s", val, key, err))
				}
				v.Field(i).SetInt(int64(iVal))
			case reflect.Bool:
				bVal, _ := strconv.ParseBool(val)
				v.Field(i).SetBool(bVal)
			}
		}
	}
}
