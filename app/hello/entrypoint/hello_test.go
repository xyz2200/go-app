package entrypoint

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewHelloHandler(t *testing.T) {

	tests := []struct {
		name   string
		want   string
		status int
	}{
		{
			name:   "should return hello json",
			want:   `{"message":"Hello","env":"","db_host":"","db_user":"","db_password":"","db_port":""}`,
			status: http.StatusOK,
		},
	}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {

			h := NewHelloHandler()
			h.Register()

			request := httptest.NewRequest(http.MethodGet, "/hello", nil)
			response := httptest.NewRecorder()
			h.handle(response, request)

			got := response.Body.String()

			assert.Equal(t, test.want, got)
			assert.Equal(t, test.status, response.Code)

		})

	}

}
