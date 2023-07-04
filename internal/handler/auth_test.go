package handler

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	"net/http/httptest"
	"testing"
	"todolist/internal/model"
	"todolist/internal/service"
	mock_service "todolist/internal/service/mocks"
)

func TestHandler_signUp(t *testing.T) {
	type mockBehavior func(s *mock_service.MockAuthorization, user model.User)

	testTable := []struct {
		name               string
		inputBody          string
		inputUser          model.User
		mockBehavior       mockBehavior
		expectedStatusCode int
		expectedStatusBody string
	}{
		{
			name:      "OK",
			inputBody: `{"name": "Test", "email": "testEmail", "username":"test", "password":"qwerty", "age":15}`,
			inputUser: model.User{
				Name:     "Test",
				Email:    "testEmail",
				Username: "test",
				Password: "qwerty",
				Age:      15,
			},
			mockBehavior: func(s *mock_service.MockAuthorization, user model.User) {
				s.EXPECT().CreateUser(user).Return(uint(1), nil)
			},
			expectedStatusCode: 200,
			expectedStatusBody: `{"id":1,"message":"User created"}`,
		},
		{
			name:               "Empty fields",
			inputBody:          `{"name": "Test", "email": "testEmail", "age":15}`,
			mockBehavior:       func(s *mock_service.MockAuthorization, user model.User) {},
			expectedStatusCode: 400,
			expectedStatusBody: `{"message":"invalid input body"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			auth := mock_service.NewMockAuthorization(c)
			testCase.mockBehavior(auth, testCase.inputUser)

			services := &service.Service{Authorization: auth}
			handlers := NewHandler(services)

			r := gin.New()
			r.POST("/sign-up", handlers.signUp)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/sign-up", bytes.NewBufferString(testCase.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, testCase.expectedStatusCode)
			assert.Equal(t, w.Body.String(), testCase.expectedStatusBody)
		})
	}
}

func TestHandler_userIdentity(t *testing.T) {
	type mockBehavior func(s *mock_service.MockAuthorization, token string)

	testTable := []struct {
		name               string
		headerName         string
		headerValue        string
		token              string
		mockBehavior       mockBehavior
		expectedStatusCode int
		expectedStatusBody string
	}{
		{
			name:        "OK",
			headerName:  "Authorization",
			headerValue: "Bearer token",
			token:       "token",
			mockBehavior: func(s *mock_service.MockAuthorization, token string) {
				s.EXPECT().ParseToken(token).Return(uint(1), nil)
			},
			expectedStatusCode: 200,
			expectedStatusBody: "1",
		},
		{
			name:               "Empty auth header",
			headerName:         "",
			token:              "token",
			mockBehavior:       func(s *mock_service.MockAuthorization, token string) {},
			expectedStatusCode: 401,
			expectedStatusBody: `{"message":"auth header is empty. You have to sign-in"}`,
		},
		{
			name:               "Wrong bearer",
			headerName:         "Authorization",
			headerValue:        "Brer token",
			token:              "token",
			mockBehavior:       func(s *mock_service.MockAuthorization, token string) {},
			expectedStatusCode: 401,
			expectedStatusBody: `{"message":"invalid auth header"}`,
		},
		{
			name:               "Invalid auth header",
			headerName:         "Authorization",
			headerValue:        "Bearer token max",
			token:              "token",
			mockBehavior:       func(s *mock_service.MockAuthorization, token string) {},
			expectedStatusCode: 401,
			expectedStatusBody: `{"message":"invalid auth header"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			auth := mock_service.NewMockAuthorization(c)
			testCase.mockBehavior(auth, testCase.token)

			services := &service.Service{Authorization: auth}
			handlers := NewHandler(services)

			r := gin.New()
			r.GET("/recovered", handlers.userIdentity, func(c *gin.Context) {
				id, _ := c.Get(userCtx)
				c.String(200, "%d", id)
			})

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/recovered", nil)
			req.Header.Set(testCase.headerName, testCase.headerValue)

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, testCase.expectedStatusCode)
			assert.Equal(t, w.Body.String(), testCase.expectedStatusBody)
		})
	}
}
