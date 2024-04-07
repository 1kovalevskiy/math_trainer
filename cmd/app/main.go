package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"github.com/1kovalevskiy/math_trainer/config"
	"github.com/1kovalevskiy/math_trainer/internal/app"
)

func main() {

	var configPath string
	flag.StringVar(&configPath, "config-path", "", "path to config")
	flag.Parse()

	if !filepath.IsAbs(configPath) {
		dir, err := os.Getwd()
		if err != nil {
			log.Println(err)
			return
		}
		configPath = filepath.Join(dir, configPath)
	}

	if filepath.Ext(configPath) != ".yaml" {
		configPath = filepath.Join(configPath, "config.yaml")
	}

	if _, err := os.Stat(configPath); err != nil {
		log.Println(err)
		return
	}

	log.Println(configPath)
	cfg, err := config.NewConfig(configPath)
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	app.Run(cfg)
}
