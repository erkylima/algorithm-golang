package handlers

import "fmt"

var (
	ErrPermissionDenied  = fmt.Errorf("permission denied")
	ErrInvalidParameters = fmt.Errorf("error in payload. verify your payload.")
)
