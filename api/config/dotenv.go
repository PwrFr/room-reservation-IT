package config

import "github.com/joho/godotenv"

func InitEnv() map[string]string {
	_ = godotenv.Load()
	env, _ := godotenv.Read()

	return env
}

//calling env
func Env(val string) string {
	return InitEnv()[val]
}
