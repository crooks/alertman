package main

import (
	"log"
)

func main() {
	var err error
	f := parseFlags()
	cfg, err = newConfig(f.configFile)
	if err != nil {
		log.Fatalf("Unable to parse config file: %v", err)
	}
	a := App{}
	a.Initialize(
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.DBName,
	)
	a.Run(":8010")
}
