package ping_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/cmd/server/http/ping"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/pkg/testutil"
)

func TestHandlePing(t *testing.T) {
	endpointPing := "/ping"
	router := testutil.SetUpRouter()

	controller := ping.NewController()
	router.GET(endpointPing, controller.HandlePing)

	response := testutil.ExecuteTestRequest(router, http.MethodGet, endpointPing, nil)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, "pong", response.Body.String())
}
