package cli

import (
	"github.com/JackieLi565/lci/cli/config"
	"github.com/JackieLi565/lci/database"
)

type LCI struct {
	ConfigFile *config.Config
	DB         *database.DB
}

func NewLCI(path string) LCI {
	config, _ := config.Load(path)
	db := database.NewDB(config.Database)

	return LCI{
		ConfigFile: config,
		DB:         db,
	}
}
