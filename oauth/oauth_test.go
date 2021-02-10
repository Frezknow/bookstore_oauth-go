package oauth

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	fmt.Println("about to start oauth tests")
	rest.StartMockupServer()
	os.Exit(m.Run())
}
func TestOauthConstants(t *testing.T) {
	assert.EqualValues(t, "X-Public", headerXPublic)
	assert.EqualValues(t, "X-Client-ID", headerXClientID)
	assert.EqualValues(t, "X-Caller-ID", headerXCallerID)
	assert.EqualValues(t, "access_token", paramAccessToken)
}

func TestIsPublicNilRequest(t *testing.T) {
	assert.True(t, IsPublic(nil))
}
func TestIsPublicNoError(t *testing.T) {
	request := http.Request{
		Header: make(http.Header),
	}
	assert.False(t, IsPublic(&request))

	request.Header.Add("X-Public", "true")
	assert.True(t, IsPublic(&request))
}

func TestGetCallerIDNilRequest(t *testing.T) {
	assert.EqualValues(t, 0, GetCallerID(nil))
}

func TestGetCallerInvalidCallerFormat(t *testing.T) {
	request := http.Request{
		Header: make(http.Header),
	}
	assert.EqualValues(t, 0, GetCallerID(&request))
}
func TestGetCallerNoError(t *testing.T) {
	request := http.Request{
		Header: make(http.Header),
	}
	request.Header.Add("X-Caller-ID", "121")
	assert.EqualValues(t, 0, GetCallerID(&request))
}

func TestGetAccessTokenInvalidRestclientResponse(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodGet,
		URL:          "http://localhost:8090/oauth/access_token/abc123",
		ReqBody:      ``,
		RespHTTPCode: -1,
		RespBody:     `{}`,
	})
	accessToken, err := getAccessToken("abc123")
	assert.Nil(t, accessToken)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid restclient response when trying to get access token", err.Message())
}

//TODO: Add Complete coverage for the getaccesstoken function
