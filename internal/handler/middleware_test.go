package handler

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bruhlord-s/openboard-go/internal/context"
	"github.com/bruhlord-s/openboard-go/internal/service"
	mock_service "github.com/bruhlord-s/openboard-go/internal/service/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestHandler_userIdentity(t *testing.T) {
	type mockBehavior func(s *mock_service.MockAuthorization, token string)

	testTable := []struct {
		name                 string
		headerName           string
		headerValue          string
		token                string
		mockBehavior         mockBehavior
		exceptedStatusCode   int
		exceptedResponseBody string
	}{
		{
			name:        "OK",
			headerName:  "Authorization",
			headerValue: "Bearer token",
			token:       "token",
			mockBehavior: func(s *mock_service.MockAuthorization, token string) {
				s.EXPECT().ParseToken(token).Return(1, nil)
			},
			exceptedStatusCode:   http.StatusOK,
			exceptedResponseBody: "1",
		},
		{
			name:                 "Missing Header",
			headerName:           "",
			mockBehavior:         func(s *mock_service.MockAuthorization, token string) {},
			exceptedStatusCode:   http.StatusUnauthorized,
			exceptedResponseBody: `{"message":"empty Authorization header"}`,
		},
		{
			name:                 "Invalid Bearer",
			headerName:           "Authorization",
			headerValue:          "Berer token",
			mockBehavior:         func(s *mock_service.MockAuthorization, token string) {},
			exceptedStatusCode:   http.StatusUnauthorized,
			exceptedResponseBody: `{"message":"invalid Authorization header"}`,
		},
		{
			name:                 "Empty Token",
			headerName:           "Authorization",
			headerValue:          "Bearer ",
			mockBehavior:         func(s *mock_service.MockAuthorization, token string) {},
			exceptedStatusCode:   http.StatusUnauthorized,
			exceptedResponseBody: `{"message":"empty Bearer token"}`,
		},
		{
			name:        "Service Failure",
			headerName:  "Authorization",
			headerValue: "Bearer token",
			token:       "token",
			mockBehavior: func(s *mock_service.MockAuthorization, token string) {
				s.EXPECT().ParseToken(token).Return(1, errors.New("error while parsing token"))
			},
			exceptedStatusCode:   http.StatusUnauthorized,
			exceptedResponseBody: `{"message":"error while parsing token"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			auth := mock_service.NewMockAuthorization(c)
			testCase.mockBehavior(auth, testCase.token)

			services := &service.Service{Authorization: auth}
			handler := NewHandler(services)

			r := gin.New()
			r.POST("/test", handler.userIdentity, func(c *gin.Context) {
				id, _ := c.Get(context.UserCtx)
				c.String(http.StatusOK, fmt.Sprintf("%d", id.(int)))
			})

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/test", nil)
			req.Header.Set(testCase.headerName, testCase.headerValue)

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.exceptedStatusCode, w.Code)
			assert.Equal(t, testCase.exceptedResponseBody, w.Body.String())
		})
	}
}
