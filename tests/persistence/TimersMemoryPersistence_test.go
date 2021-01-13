package test_persistence

import (
	persist "github.com/expproletariy/pip-timers-service/persistence"
	cconf "github.com/pip-services3-go/pip-services3-commons-go/config"
	"testing"
)

type TimersMemoryPersistenceTest struct {
	persistence *persist.TimersMemoryPersistence
	fixture     *TimersPersistenceFixture
}

func newTimersMemoryPersistenceTest() *TimersMemoryPersistenceTest {
	persistence := persist.NewTimersMemoryPersistence()
	persistence.Configure(cconf.NewEmptyConfigParams())

	fixture := NewTimersPersistenceFixture(persistence)

	return &TimersMemoryPersistenceTest{
		persistence: persistence,
		fixture:     fixture,
	}
}

func (c *TimersMemoryPersistenceTest) setup(t *testing.T) {
	err := c.persistence.Open("")
	if err != nil {
		t.Error("Failed to open persistence", err)
	}

	err = c.persistence.Clear("")
	if err != nil {
		t.Error("Failed to clear persistence", err)
	}
}

func (c *TimersMemoryPersistenceTest) teardown(t *testing.T) {
	err := c.persistence.Close("")
	if err != nil {
		t.Error("Failed to close persistence", err)
	}
}

func TestTimersMemoryPersistence(t *testing.T) {
	c := newTimersMemoryPersistenceTest()
	if c == nil {
		return
	}

	c.setup(t)
	t.Run("CRUD Operations", c.fixture.TestCrudOperations)
	c.teardown(t)

	c.setup(t)
	t.Run("Get with filters", c.fixture.TestGetWithFilters)
	c.teardown(t)
}
