package settings

import (
	"embed"
	"ever-book/app/global/structs"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

var Config *structs.EnvConfig

func Load(f embed.FS) {
	env := os.Getenv("ENV")

	pathList := []string{
		"env/" + env + "/db.yaml",
	}

	for k := range pathList {
		configFile, err := f.ReadFile(pathList[k])
		if err != nil {
			log.Fatalf("Read File Error: %v", err.Error())
		}

		if err = yaml.Unmarshal(configFile, &Config); err != nil {
			log.Fatalf("Yaml Unmarshal Error: %v", err.Error())
		}
	}
}
