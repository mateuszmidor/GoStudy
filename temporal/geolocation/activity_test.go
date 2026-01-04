package geolocation_test

import (
	"io"
	"net/http"
	"strings"
	"testing"

	geolocation "github.com/mateuszmidor/GoStudy/temporal/geolocation"
	"github.com/stretchr/testify/assert"
	"go.temporal.io/sdk/testsuite"
)

type MockHTTPClient struct {
	Response *http.Response
	Err      error
}

func (m *MockHTTPClient) Get(url string) (*http.Response, error) {
	return m.Response, m.Err
}

// TestGetIP tests the GetIP activity with a mock server.
func TestGetIP(t *testing.T) {
	// set up test environment
	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestActivityEnvironment()

	// Create a mock response that returns the fake IP address
	mockResponse := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(strings.NewReader("127.0.0.1\n")),
	}

	// load Activities and inject mock response
	ipActivities := &geolocation.IPActivities{
		HTTPClient: &MockHTTPClient{Response: mockResponse},
	}
	env.RegisterActivity(ipActivities)

	// Call the GetIP function
	val, err := env.ExecuteActivity(ipActivities.GetIP)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// get the Activity result
	var ip string
	val.Get(&ip)

	// Validate the returned IP
	expectedIP := "127.0.0.1"
	assert.Equal(t, ip, expectedIP)
}

// TestGetLocationInfo tests the GetLocationInfo activity with a mock server.
func TestGetLocationInfo(t *testing.T) {
	// set up test environment
	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestActivityEnvironment()

	mockResponse := &http.Response{
		StatusCode: 200,
		Body: io.NopCloser(strings.NewReader(`{
            "city": "San Francisco",
            "regionName": "California",
            "country": "United States"
        }`)),
	}

	ipActivities := &geolocation.IPActivities{
		HTTPClient: &MockHTTPClient{Response: mockResponse},
	}

	env.RegisterActivity(ipActivities)

	ip := "127.0.0.1"
	val, err := env.ExecuteActivity(ipActivities.GetLocationInfo, ip)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	var location string
	val.Get(&location)

	expectedLocation := "San Francisco, California, United States"
	assert.Equal(t, location, expectedLocation)
}
