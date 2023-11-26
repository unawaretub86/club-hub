package apperrors

import "errors"

var (
	FranchiseNotFound = errors.New("franchise not found")
	RecordNotFound = errors.New("record not found")
)
