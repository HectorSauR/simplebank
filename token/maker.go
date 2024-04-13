package token

import "time"

// MMaker is an interfacce for managing tokens
type Maker interface {
	// CreateToken creates a new token for a specific username and duration
	CreateToken(username string, duration time.Duration) (string, error)

	// Verifytoken checks if the token is valid or not
	VerifyToken(token string) (*Payload, error)
}
