package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/ECOMMERCE_PROJECT/pkg/common/helperstruct"
	"github.com/ECOMMERCE_PROJECT/pkg/common/response"
	mock "github.com/ECOMMERCE_PROJECT/pkg/usecase/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUserSignUp(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userusecase := mock.NewMockUserUseCase(ctrl)
	cartusecase := mock.NewMockCartUseCase(ctrl)
	UserHandler := NewUserHandelr(userusecase, cartusecase)

	user := &helperstruct.UserReq{
		Name:     "siddharth",
		Email:    "sidx141202@gmail.com",
		Mobile:   "8590496810",
		Password: "abcdef",
	}

	testCases := []struct {
		name           string
		input          helperstruct.UserReq
		expectedOutput response.UserData
		buildStub      func()
		expectedErr    error
	}{
		{
			name:  "SuccessfulSignUp",
			input: *user,
			buildStub: func() {
				userusecase.EXPECT().UserSignUp(gomock.Any(), gomock.Any()).Return(response.UserData{
					Id:     1,
					Name:   "siddharth",
					Email:  "sidx141202@gmail.com",
					Mobile: "8590496810",
				}, nil)
				cartusecase.EXPECT().CreateCart(gomock.Any()).Return(nil)
			},
			expectedOutput: response.UserData{
				Id:     1,
				Name:   "siddharth",
				Email:  "sidx141202@gmail.com",
				Mobile: "8590496810",
			},
			expectedErr: nil,
		},
		{
			name:  "Duplicate Entry",
			input: *user,
			buildStub: func() {
				userusecase.EXPECT().UserSignUp(gomock.Any(), gomock.Any()).Return(response.UserData{}, errors.New("duplicate user"))

			},
			expectedOutput: response.UserData{},
			expectedErr:    errors.New("duplicate user"),
		},
	}

	_, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.buildStub()
			router := gin.Default()
			router.POST("/signup", UserHandler.UserSignUp)

			reqBody := fmt.Sprintf(`{"name":"%s","email":"%s","mobile":"%s","password":"%s"}`, tc.input.Name, tc.input.Email, tc.input.Mobile, tc.input.Password)
			req, _ := http.NewRequest(http.MethodPost, "/signup", strings.NewReader(reqBody))
			req.Header.Set("Content-Type", "application/json")

			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)

			if tc.expectedErr != nil {
				assert.Equal(t, http.StatusBadRequest, resp.Code)
			} else {
				assert.Equal(t, http.StatusCreated, resp.Code)
			}
		})
	}
}
func TestUserLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	userusecase := mock.NewMockUserUseCase(ctrl)
	cartusecase := mock.NewMockCartUseCase(ctrl)
	UserHandler := NewUserHandelr(userusecase, cartusecase)
	testUser := helperstruct.LoginReq{
		Email:   "sidx141202@gmail.com",
		Password: "abcdef",
	}
	jsonTestUser, err := json.Marshal(testUser)
if err != nil {
	t.Fatal(err)
}

// Create a mock HTTP request and response writer
req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonTestUser))
if err != nil {
	t.Fatal(err)
}
	rr := httptest.NewRecorder()

	// Create a mock gin context
	c, _ := gin.CreateTestContext(rr)
	c.Request = req

	// Set expectations on mock
	userusecase.EXPECT().UserLogin(c.Request.Context(), testUser).Return(nil)

	// Call the function and check the result
	UserHandler.UserLogin(c)

	// Check the status code and response body
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	expected := `{"StatusCode":200,"Message":"otp send successfully,  Please enter otp","Data":null,"Errors":null}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
