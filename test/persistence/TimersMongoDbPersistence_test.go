package test_persistence

import (
	persist "github.com/expproletariy/pip-timers-service/persistence"
	cconf "github.com/pip-services3-go/pip-services3-commons-go/config"
	"os"
	"testing"
)

func TestTimersMongoDbPersistence(t *testing.T) {

	var persistence *persist.TimersMongoDBPersistence
	var fixture *TimersPersistenceFixture

	mongoUri := os.Getenv("MONGO_SERVICE_URI")
	mongoHost := os.Getenv("MONGO_SERVICE_HOST")

	if mongoHost == "" {
		mongoHost = "localhost"
	}
	mongoPort := os.Getenv("MONGO_SERVICE_PORT")
	if mongoPort == "" {
		mongoPort = "27017"
	}

	mongoDatabase := os.Getenv("MONGO_DB")
	if mongoDatabase == "" {
		mongoDatabase = "test"
	}

	// Exit if mongo connection is not set
	if mongoUri == "" && mongoHost == "" {
		return
	}

	persistence = persist.NewTimersMongoDBPersistence()
	persistence.Configure(cconf.NewConfigParamsFromTuples(
		"connection.uri", mongoUri,
		"connection.host", mongoHost,
		"connection.port", mongoPort,
		"connection.database", mongoDatabase,
	))

	fixture = NewTimersPersistenceFixture(persistence)

	opnErr := persistence.Open("")
	if opnErr == nil {
		persistence.Clear("")
	}

	defer persistence.Close("")

	t.Run("TimersMongoDbPersistence:CRUD Operations", fixture.TestCrudOperations)
	persistence.Clear("")
	t.Run("TimersMongoDbPersistence:Get with Filters", fixture.TestGetWithFilters)
	persistence.Clear("")
}
