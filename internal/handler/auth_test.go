package handler

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bruhlord-s/openboard-go/internal/model"
	"github.com/bruhlord-s/openboard-go/internal/service"
	mock_service "github.com/bruhlord-s/openboard-go/internal/service/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestHandler_signUp(t *testing.T) {
	type mockBehavior func(s *mock_service.MockAuthorization, user model.User)

	testTable := []struct {
		name                string
		inputBody           string
		inputUser           model.User
		mockBehavior        mockBehavior
		exceptedStatusCode  int
		exceptedRequestBody string
	}{
		{
			name: "OK",
			inputBody: `{
				"email": "test@test.com",
				"username": "admin",
				"name": "admin",
				"password": "12345678"
				}`,
			inputUser: model.User{
				Email:    "test@test.com",
				Username: "admin",
				Name:     "admin",
				Password: "12345678",
			},
			mockBehavior: func(s *mock_service.MockAuthorization, user model.User) {
				s.EXPECT().CreateUser(user).Return(1, nil)
			},
			exceptedStatusCode:  http.StatusOK,
			exceptedRequestBody: `{"id":1}`,
		},
		{
			name: "Empty Fields",
			inputBody: `{
				"email": "test@test.com",
				"name": "admin",
				}`,
			mockBehavior:        func(s *mock_service.MockAuthorization, user model.User) {},
			exceptedStatusCode:  http.StatusUnprocessableEntity,
			exceptedRequestBody: `{"message":"invalid input body"}`,
		},
		{
			name: "Service Failure",
			inputBody: `{
				"email": "test@test.com",
				"username": "admin",
				"name": "admin",
				"password": "12345678"
				}`,
			inputUser: model.User{
				Email:    "test@test.com",
				Username: "admin",
				Name:     "admin",
				Password: "12345678",
			},
			mockBehavior: func(s *mock_service.MockAuthorization, user model.User) {
				s.EXPECT().CreateUser(user).Return(1, errors.New("service failure"))
			},
			exceptedStatusCode:  http.StatusInternalServerError,
			exceptedRequestBody: `{"message":"service failure"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			auth := mock_service.NewMockAuthorization(c)
			testCase.mockBehavior(auth, testCase.inputUser)

			services := &service.Service{Authorization: auth}
			handler := NewHandler(services)

			r := gin.New()
			r.POST("/sign-up", handler.signUp)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/sign-up",
				bytes.NewBufferString(testCase.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.exceptedStatusCode, w.Code)
			assert.Equal(t, testCase.exceptedRequestBody, w.Body.String())
		})
	}
}

func TestHandler_signIn(t *testing.T) {
	type mockBehavior func(s *mock_service.MockAuthorization, input signInInput)

	testTable := []struct {
		name                string
		inputBody           string
		inputUser           signInInput
		mockBehavior        mockBehavior
		exceptedStatusCode  int
		exceptedRequestBody string
	}{
		{
			name: "OK",
			inputBody: `{
				"username": "admin",
				"password": "12345678"
				}`,
			inputUser: signInInput{
				Username: "admin",
				Password: "12345678",
			},
			mockBehavior: func(s *mock_service.MockAuthorization, input signInInput) {
				s.EXPECT().GenerateToken(input.Username, input.Password).Return("token", nil)
			},
			exceptedStatusCode:  http.StatusOK,
			exceptedRequestBody: `{"token":"token"}`,
		},
		{
			name: "Empty Fields",
			inputBody: `{
				"email": "test@test.com",
				}`,
			mockBehavior:        func(s *mock_service.MockAuthorization, input signInInput) {},
			exceptedStatusCode:  http.StatusUnprocessableEntity,
			exceptedRequestBody: `{"message":"invalid input body"}`,
		},
		{
			name: "Service Failure",
			inputBody: `{
				"username": "admin",
				"password": "12345678"
				}`,
			inputUser: signInInput{
				Username: "admin",
				Password: "12345678",
			},
			mockBehavior: func(s *mock_service.MockAuthorization, input signInInput) {
				s.
					EXPECT().
					GenerateToken(input.Username, input.Password).
					Return("token", errors.New("service failure"))
			},
			exceptedStatusCode:  http.StatusInternalServerError,
			exceptedRequestBody: `{"message":"service failure"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			auth := mock_service.NewMockAuthorization(c)
			testCase.mockBehavior(auth, testCase.inputUser)

			services := &service.Service{Authorization: auth}
			handler := NewHandler(services)

			r := gin.New()
			r.POST("/sign-in", handler.signIn)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/sign-in",
				bytes.NewBufferString(testCase.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.exceptedStatusCode, w.Code)
			assert.Equal(t, testCase.exceptedRequestBody, w.Body.String())
		})
	}
}
