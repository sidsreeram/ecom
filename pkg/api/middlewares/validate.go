package middlewares

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)
func ValidateToken(token string) (int, error) {
    Tokenvalue, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
        if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
        }
        return []byte("secret"), nil
    })

    if err != nil {
        return 0, err
    }

    if Tokenvalue == nil || !Tokenvalue.Valid {
        return 0, fmt.Errorf("invalid token")
    }

    var userID int
    if claims, ok := Tokenvalue.Claims.(jwt.MapClaims); ok && Tokenvalue.Valid {
        userID, ok = claims["userId"].(int)
        if !ok {
            return 0, fmt.Errorf("expected a string value for userId, but got %T", claims["usersId"])
        }

        // Check the token expiration
        if float64(time.Now().Unix()) > claims["exp"].(float64) {
            return 0, fmt.Errorf("token expired")
        }
    }

    return userID, err
}


// func ValidateToken(token string) (int, error) {
// 	Tokenvalue, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {

// 		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
// 		}

// 		return []byte("secret"), nil
// 	})

// 	if err != nil {
// 		return 0, err
// 	}

// 	if Tokenvalue == nil || !Tokenvalue.Valid {
// 		return 0, fmt.Errorf("invalid token")
// 	}

// 	var parsedID interface{}
// 	if claims, ok := Tokenvalue.Claims.(jwt.MapClaims); ok && Tokenvalue.Valid {
// 		parsedID = claims["id"]
// 		//Check the expir
// 		if float64(time.Now().Unix()) > claims["exp"].(float64) {
// 			return 0, fmt.Errorf("token expired")
// 		}
// 		// fmt.Println(claims["exp"])
// 	}
// 	value, ok := parsedID.(float64)
// 	if !ok {
// 		return 0, fmt.Errorf("expected an int value, but got %T", parsedID)
// 	}

// 	id := int(value)
// 	return id, err
// }

