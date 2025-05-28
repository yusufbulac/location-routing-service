package integration

import (
	"github.com/yusufbulac/location-routing-service/test/integration/testutils"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	testutils.SetupTestDB()
	testutils.CleanDatabase()
	testutils.SeedLocations()

	code := m.Run()
	os.Exit(code)
}
