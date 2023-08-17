package test

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/h2non/gock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSimple(t *testing.T) {
	// given
	gock.New("http://foo.com").
		Get("/bar").
		Reply(http.StatusOK).
		JSON(map[string]string{"foo": "bar"})
	defer gock.Off()

	// when
	res, err := http.Get("http://foo.com/bar")

	// then
	assert.NoError(t, err)
	require.NotNil(t, res)
	assert.Equal(t, http.StatusOK, res.StatusCode)

	body, _ := ioutil.ReadAll(res.Body)
	assert.Equal(t, `{"foo":"bar"}`, strings.TrimSpace(string(body)))
	assert.True(t, gock.IsDone()) // Verify that we don't have pending mocks
}
