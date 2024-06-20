package cli

import (
	"github.com/JackieLi565/lci/cli/config"
	"github.com/JackieLi565/lci/database"
)

type LCI struct {
	ConfigFile *config.Config
	DB         *database.DB
}

func NewLCI() LCI {
	config := config.LoadConfig("./config.json")
	db := database.NewDB(config.Database)

	return LCI{
		ConfigFile: config,
		DB:         db,
	}
}
