package middlewares

import "github.com/rs/cors"

func SettingsCors() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})
}
