package middleware

import "fmt"

func RoleValid(role string, token *string) bool {
	tokenStr := *token
	verify := verifyToken(tokenStr)
	if !verify {
		fmt.Println("Verify failed")
		return false
	}

	payload, err := extractToken(tokenStr)
	if err != nil {
		fmt.Println(err)
	}

	if payload.Subject != role {
		fmt.Println("role not equal")
		return false
	}

	return true
}
