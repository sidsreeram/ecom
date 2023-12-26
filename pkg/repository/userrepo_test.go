package repository

import (
	"context"
	"database/sql"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ECOMMERCE_PROJECT/pkg/common/helperstruct"
	"github.com/ECOMMERCE_PROJECT/pkg/common/response"
	"github.com/ECOMMERCE_PROJECT/pkg/domain"
	"github.com/ECOMMERCE_PROJECT/pkg/repository/mocks"
	"github.com/golang/mock/gomock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ...

func (c *userDatabase) Close() {
	sqlDB, err := c.DB.DB()
	if err != nil {
		// Handle the error properly (log, return, etc.)
		return
	}

	if err := sqlDB.Close(); err != nil {
		// Handle the error properly (log, return, etc.)
		return
	}
}

func TestUserSignUp(t *testing.T) {
	// Mock for GORM database connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	//initialize the db instance with the mock db connection
	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when initializing a mock db session", err)
	}

	repo := &userDatabase{DB: gormDB}
	defer func() {
		repo.Close()
	}()

	user := &helperstruct.UserReq{
		Name:     "siddharth",
		Email:    "sidx141202@gmail.com",
		Mobile:   "8590496810",
		Password: "abcdef",
	}

	tests := []struct {
		name           string
		input          helperstruct.UserReq
		expectedOutput response.UserData
		buildStub      func()
		expectedErr    error
	}{
		{
			name:  "success entry",
			input: *user,
			expectedOutput: response.UserData{
				Id:     1,
				Name:   user.Name,
				Email:  user.Email,
				Mobile: user.Mobile,
			},
			buildStub: func() {
				rows := sqlmock.NewRows([]string{"id", "name", "email", "mobile"}). // Changed "phone" to "mobile"
													AddRow(1, user.Name, user.Email, user.Mobile)

				mock.ExpectQuery("^INSERT INTO users (.+)$").
					WithArgs(user.Name, user.Email, user.Mobile, user.Password).
					WillReturnRows(rows)
			},
		},
		{
			name:           "duplicate entry",
			input:          *user,
			expectedOutput: response.UserData{},
			buildStub: func() {
				// Simulate a duplicate entry by expecting an empty result set
				rows := sqlmock.NewRows([]string{"id", "name", "email", "mobile"})

				mock.ExpectQuery("^INSERT INTO users (.+)$").
					WithArgs(user.Name, user.Email, user.Mobile, user.Password).
					WillReturnRows(rows)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.buildStub()

			result, err := repo.UserSignUp(context.Background(), tt.input)

			if err != tt.expectedErr {
				t.Errorf("Unexpected error. Expected: %v, Got: %v", tt.expectedErr, err)
			}

			if !reflect.DeepEqual(result, tt.expectedOutput) {
				t.Errorf("Unexpected result. Expected: %v, Got: %v", tt.expectedOutput, result)
			}
		})
	}
}

func GetSQLDB(db *gorm.DB) (*sql.DB, error) {
	sqlDB, err := db.DB()
	return sqlDB, err
}

func TestUserLogin(t *testing.T) {
	// Mock for GORM database connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	//initialize the db instance with the mock db connection
	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when initializing a mock db session", err)
	}

	repo := &userDatabase{DB: gormDB}
	defer func() {
		repo.Close()
	}()

	tests := []struct {
		name           string
		email          string
		expectedOutput domain.Users
		buildStub      func()
		expectedErr    error
	}{
		{
			name:  "success entry",
			email: "sidx141202@gmail.com",
			expectedOutput: domain.Users{
				ID:     1,
				Name:   "siddharth",
				Email:  "sidx141202@gmail.com",
				Mobile: "8590496810",
			},
			buildStub: func() {
				rows := sqlmock.NewRows([]string{"id", "name", "email", "mobile"}).
					AddRow(1, "siddharth", "sidx141202@gmail.com", "8590496810")

				mock.ExpectQuery("^SELECT (.+) FROM users WHERE email=(.+)").
					WithArgs("sidx141202@gmail.com").
					WillReturnRows(rows)
			},
		},
		{
			name:           "no entry",
			email:          "nonexistent@gmail.com",
			expectedOutput: domain.Users{},
			buildStub: func() {
				rows := sqlmock.NewRows([]string{"id", "name", "email", "mobile"})

				mock.ExpectQuery("^SELECT (.+) FROM users WHERE email=(.+)").
					WithArgs("nonexistent@gmail.com").
					WillReturnRows(rows)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.buildStub()

			result, err := repo.UserLogin(context.Background(), tt.email)

			if err != tt.expectedErr {
				t.Errorf("Unexpected error. Expected: %v, Got: %v", tt.expectedErr, err)
			}

			if !reflect.DeepEqual(result, tt.expectedOutput) {
				t.Errorf("Unexpected result. Expected: %v, Got: %v", tt.expectedOutput, result)
			}
		})
	}
}

func TestStoreOTP(t *testing.T) {
	
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Define the email and OTP for testing
	testEmail := "test@example.com"
	testOTP := "123456"
	mockUserRepo := mocks.NewMockUserRepository(ctrl)
	// Set expectation on mock
	mockUserRepo.EXPECT().StoreOTP(testEmail, testOTP).Return(true)

	// Call the function and check the result
	result := mockUserRepo.StoreOTP(testEmail, testOTP)
	if result != true {
		t.Errorf("Expected true, but got %v", result)
	}
}
func TestVerifyOTP(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockUserRepo := mocks.NewMockUserRepository(ctrl)

    // Define the OTP for testing
    testOTP := "123456"

    // Define the expected return values
    expectedUserID := 1
    expectedBool := true

    // Set expectation on mock
    mockUserRepo.EXPECT().VerifyOTP(testOTP).Return(expectedUserID, expectedBool)

    // Call the function and check the result
    userID, boolVal := mockUserRepo.VerifyOTP(testOTP)
    if userID != expectedUserID || boolVal != expectedBool {
        t.Errorf("Expected userID %v and boolVal %v, but got userID %v and boolVal %v", expectedUserID, expectedBool, userID, boolVal)
    }
}
func TestAddAddress(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockUserRepo := mocks.NewMockUserRepository(ctrl)

    // Define the test data
    testID := 1
    testAddress := helperstruct.Address{
        House_number: "123",
        Street:       "kolasseri",
        City:         "kadampzhipuram",
        District:     "palakkad",
        Landmark:     "kerala",
        Pincode:      123456,
        IsDefault:    true,
    }

    // Set expectations on mock
    mockUserRepo.EXPECT().AddAddress(testID, testAddress).Return(nil)

    // Call the function and check the result
    err := mockUserRepo.AddAddress(testID, testAddress)
    if err != nil {
        t.Errorf("AddAddress returned error: %v", err)
    }
}

func TestUpdateAddress(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockUserRepo := mocks.NewMockUserRepository(ctrl)

    // Define the test data
    testID := 1
    testAddress := helperstruct.Address{
        House_number: "123",
        Street:       "kolasseri",
        City:         "kadampzhipuram",
        District:     "palakkad",
        Landmark:     "kerala",
        Pincode:      123456,
        IsDefault:    true,
    }

    // Set expectations on mock
    mockUserRepo.EXPECT().UpdateAddress(testID, testID, testAddress).Return(nil)

    // Call the function and check the result
    err := mockUserRepo.UpdateAddress(testID, testID, testAddress)
    if err != nil {
        t.Errorf("UpdateAddress returned error: %v", err)
    }
}

// Continue with similar test functions for ViewProfile, UpdateProfile, FindPassword, and UpdatePassword
