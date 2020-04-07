package http

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"product-service/internals/controller"
	"testing"
)

func TestServer(t *testing.T) {
	ctr := controller.GatewayCtrl{
		PingCtrl: controller.NewPingController(),
		ProductCtrl: controller.NewProductController(),
	}
	svc := NewHTTPServer(ctr)
	svc.Configure()
	t.Run("GET /ping ",func( t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/ping", nil)
		if err != nil {
			t.Error(err)
		}
		w := httptest.NewRecorder()
		svc.route.ServeHTTP(w, req)
		assert.Equal(t, 200, w.Code)
		assert.Equal(t,"pong",w.Body.String())
	})

}