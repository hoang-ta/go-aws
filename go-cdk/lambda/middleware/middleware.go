package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/golang-jwt/jwt/v5"
)

func ValidateJWTMiddleware(next func(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)) func(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return func(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		tokenString := extractTokenFromHeader(request.Headers)
		if tokenString == "" {
			return events.APIGatewayProxyResponse{
				Body : "Missing auth token",
				StatusCode: http.StatusBadRequest,
			}, nil
		}
		claims, err := parseToken(tokenString)
		if err != nil {
			return events.APIGatewayProxyResponse{
				Body : "User unauthorized",
				StatusCode: http.StatusUnauthorized,
			}, err
		}
		expires := int64(claims["expires"].(float64))
		if expires < time.Now().Unix() {
			return events.APIGatewayProxyResponse{
				Body : "Token expired",
				StatusCode: http.StatusUnauthorized,
			}, nil
		}
		return next(request)
	}
}

func extractTokenFromHeader(headers map[string]string) string {
	authHeader, ok := headers["Authorization"]
	if !ok {
		return ""
	}
	splitToken := strings.Split(authHeader, "Bearer")
	if len(splitToken) != 2 { 
		return ""
	}
	return splitToken[1]
}

func parseToken(tokenString string) (jwt.MapClaims, error) {
	fmt.Println(tokenString)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		secret := "secret"
		return []byte(secret), nil	
	})
	if err != nil {
		return nil, fmt.Errorf("unauthorized")
	}
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid claims")
	}
	return claims, nil
 }