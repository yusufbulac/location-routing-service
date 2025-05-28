package integration

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/yusufbulac/location-routing-service/internal/model"
	"github.com/yusufbulac/location-routing-service/test/integration/testutils"
)

func readAndLogBody(t *testing.T, resp *http.Response) []byte {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Logf("Failed to read response body: %v", err)
		return nil
	}
	t.Logf("Response Body: %s", string(body))
	return body
}

func TestCreateLocation(t *testing.T) {
	reqBody := map[string]interface{}{
		"name":      "Test Location",
		"latitude":  40.7128,
		"longitude": 29.0060,
		"color":     "#abcdef",
	}
	resp := testutils.Post(t, "/api/v1/locations", reqBody)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	body := readAndLogBody(t, resp)

	var loc model.Location
	err := json.Unmarshal(body, &loc)
	require.NoError(t, err, "Failed to decode created location JSON")
	assert.Equal(t, "Test Location", loc.Name)
	assert.Equal(t, "#abcdef", loc.Color)
}

func TestGetAllLocations(t *testing.T) {
	resp := testutils.Get(t, "/api/v1/locations")
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	body := readAndLogBody(t, resp)

	var locations []model.Location
	err := json.Unmarshal(body, &locations)
	require.NoError(t, err, "Failed to decode locations list JSON")
	assert.GreaterOrEqual(t, len(locations), 3, "Expected at least 3 seeded locations")
}

func TestUpdateLocation(t *testing.T) {
	initial := model.Location{Name: "Initial", Latitude: 41, Longitude: 29, Color: "#123456"}
	if err := testutils.TestDB.Create(&initial).Error; err != nil {
		t.Fatalf("Failed to create location: %v", err)
	}

	t.Logf("Created location with ID: %d", initial.ID)
	require.NotZero(t, initial.ID, "Initial test location must have an ID")

	payload := map[string]interface{}{
		"name":      "Updated",
		"latitude":  42,
		"longitude": 30,
		"color":     "#654321",
	}
	url := "/api/v1/locations/" + strconv.Itoa(int(initial.ID))

	resp := testutils.Put(t, url, payload)
	defer resp.Body.Close()

	body := readAndLogBody(t, resp)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var updated model.Location
	err := json.Unmarshal(body, &updated)
	require.NoError(t, err, "Failed to decode updated location JSON")
	assert.Equal(t, "Updated", updated.Name, "Updated name mismatch")
	assert.Equal(t, "#654321", updated.Color, "Updated color mismatch")
}

func TestCreateLocation_InvalidPayload(t *testing.T) {
	reqBody := map[string]interface{}{
		"latitude":  40.7128,
		"longitude": 29.0060,
	}
	resp := testutils.Post(t, "/api/v1/locations", reqBody)
	defer resp.Body.Close()

	_ = readAndLogBody(t, resp)

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	t.Log("Expected bad request due to missing required fields")
}

func TestUpdateLocation_NotFound(t *testing.T) {
	nonExistentID := 999999
	payload := map[string]interface{}{
		"name":      "Ghost",
		"latitude":  10,
		"longitude": 20,
		"color":     "#ffffff",
	}
	url := "/api/v1/locations/" + strconv.Itoa(nonExistentID)
	resp := testutils.Put(t, url, payload)
	defer resp.Body.Close()

	_ = readAndLogBody(t, resp)

	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	t.Log("Expected not found when updating non-existent location")
}
