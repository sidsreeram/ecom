package usecase

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/ECOMMERCE_PROJECT/pkg/common/helperstruct"
	"github.com/ECOMMERCE_PROJECT/pkg/common/response"
	"github.com/ECOMMERCE_PROJECT/pkg/repository/mocks"
	mock "github.com/ECOMMERCE_PROJECT/pkg/usecase/mocks"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	"golang.org/x/crypto/bcrypt"
)

type eqCreateUserParamsMatcher struct {
	arg      helperstruct.UserReq
	password string
}

func (e eqCreateUserParamsMatcher) Matches(x interface{}) bool {
	arg, ok := x.(helperstruct.UserReq)
	if !ok {
		return false
	}
	if err := bcrypt.CompareHashAndPassword([]byte(arg.Password), []byte(e.password)); err != nil {
		return false
	}
	e.arg.Password = arg.Password
	return reflect.DeepEqual(e.arg, arg)
}

func (e eqCreateUserParamsMatcher) String() string {
	return fmt.Sprintf("matches arg %v and password %v", e.arg, e.password)
}

func EqCreateUserParams(arg helperstruct.UserReq, password string) gomock.Matcher {
	return eqCreateUserParamsMatcher{arg, password}
}

func TestUserSignup(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUserRepo := mocks.NewMockUserRepository(ctrl)
	userUseCase := NewUserUseCase(mockUserRepo)

	user := &helperstruct.UserReq{
		Name:     "siddharth",
		Email:    "sidx141202@gmail.com",
		Mobile:   "8590496810",
		Password: "abcdef",
	}

	testCases := []struct {
		name           string
		input          helperstruct.UserReq
		setupMock      func(mockUserRepo *mocks.MockUserRepository)
		expectedOutput response.UserData
		expectedError  error
	}{
		{
			name:  "SuccessfulSignUp",
			input: *user,
			setupMock: func(mockUserRepo *mocks.MockUserRepository) {
				expectedOutput := response.UserData{
					Id:     1,
					Name:   "siddharth",
					Email:  "sidx141202@gmail.com",
					Mobile: "8590496810",
				}
				mockUserRepo.EXPECT().UserSignUp(gomock.Any(), gomock.Any()).Return(expectedOutput, nil)
			},
			expectedOutput: response.UserData{
				Id:     1,
				Name:   "siddharth",
				Email:  "sidx141202@gmail.com",
				Mobile: "8590496810",
			},
			expectedError: nil,
		},
		{
			name:  "Duplicate Entry",
			input: *user,
			setupMock: func(mockUserRepo *mocks.MockUserRepository) {
				mockUserRepo.EXPECT().UserSignUp(gomock.Any(), gomock.Any()).Return(response.UserData{}, errors.New("duplicate user"))
			},
			expectedOutput: response.UserData{},
			expectedError:  errors.New("duplicate user"),
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.setupMock(mockUserRepo)
			actualOutput, actualError := userUseCase.UserSignUp(ctx, tc.input)
			assert.Equal(t, tc.expectedOutput, actualOutput)
			assert.Equal(t, tc.expectedError, actualError)
		})
	}
}
func TestUserLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserUseCase := mock.NewMockUserUseCase(ctrl)

	// Define the test data
	testUser := helperstruct.LoginReq{
		Email:    "sidx141202@gmail.com",
		Password: "abcdef",
	}

	// Set expectations on mock
	mockUserUseCase.EXPECT().UserLogin(gomock.Any(), testUser).Return(nil).AnyTimes()

	// Call the function and check the result
	err := mockUserUseCase.UserLogin(context.Background(), testUser)
	if err != nil {
		t.Errorf("UserLogin returned unexpected error: %v", err)
	}
}

func TestVerifyOTP(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mocks.NewMockUserRepository(ctrl)

	// Define the test data
	testOTP := "123456"

	// Set expectations on mock
	mockUserRepo.EXPECT().VerifyOTP(gomock.Eq(testOTP)).Return(42, true)

	// Call the function and check the result
	userID, isValid := mockUserRepo.VerifyOTP(testOTP)

	// Now you can use userID and isValid in your test logic if needed
	fmt.Println(userID, isValid)
}
func TestAddAddress(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUserRepo := mocks.NewMockUserRepository(ctrl)
	userUseCase := NewUserUseCase(mockUserRepo)

	testAddress := helperstruct.Address{
		// Fill in the fields of the address
	}

	// Set expectations on mock
	mockUserRepo.EXPECT().AddAddress(1, testAddress).Return(nil)

	// Call the function and check the result
	err := userUseCase.AddAddress(1, testAddress)
	if err != nil {
		t.Errorf("AddAddress returned error: %v", err)
	}
}

func TestUpdateAddress(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mocks.NewMockUserRepository(ctrl)
	userUseCase := NewUserUseCase(mockUserRepo)

	testAddress := helperstruct.Address{
		// Fill in the fields of the address
	}

	// Set expectations on mock
	mockUserRepo.EXPECT().UpdateAddress(1, 1, testAddress).Return(nil)

	// Call the function and check the result
	err := userUseCase.UpdateAddress(1, 1, testAddress)
	if err != nil {
		t.Errorf("UpdateAddress returned error: %v", err)
	}
}
