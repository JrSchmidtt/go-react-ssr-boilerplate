package config

import (
	"os"
	gossr "github.com/natewong1313/go-react-ssr"
)

var SsrConfig = gossr.Config{
	AppEnv:             os.Getenv("APP_ENV"),
	AssetRoute:         "/assets",
	FrontendDir:        "./src/views",
	GeneratedTypesPath: "./src/views/generated.d.ts",
	PropsStructsPath:   "./src/models/props.go",
}

var RenderConfig = gossr.RenderConfig{
	Title: "Boilerplate",
	MetaTags: map[string]string{
		"og:title":    "Gin example app",
		"description": "Hello world!",
	},
}