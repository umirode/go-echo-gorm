package Connection

import (
	"io/ioutil"
	"path/filepath"
	"sync"

	"github.com/Sirupsen/logrus"
	"github.com/umirode/go-rest/Database"
	"gopkg.in/yaml.v2"
)

var connections struct {
	Connections map[string]map[string]Database.Config `yaml:"connections"`
}

var allConnectionsOnce sync.Once

func GetConnections() map[string]map[string]Database.Config {
	allConnectionsOnce.Do(func() {
		filename, _ := filepath.Abs("database.yaml")
		yamlFile, err := ioutil.ReadFile(filename)
		if err != nil {
			logrus.Fatal(err)
		}

		err = yaml.Unmarshal(yamlFile, &connections)
		if err != nil {
			logrus.Fatal(err)
		}
	})

	return connections.Connections
}

func GetConfig(connectionType string) *Database.Config {
	connections := GetConnections()[connectionType]

	for _, connection := range connections {
		if connection.Selected {
			return &connection
		}
	}

	logrus.Fatal("No database connection selected")

	return nil
}
