package entity

import "errors"

// Global Error
var (
	ErrGlobalNotFound               = errors.New("not_found")
	ErrGlobalServerErr              = errors.New("internal_server_error")
	ErrGlobalFileSizeExceedLimit    = errors.New("file_size_exceed_limit")
	ErrGlobalInvalidFileContentType = errors.New("invalid_file_content_type")
)

// Blog Error
var (
	ErrBlogTagMustBeUnique = errors.New("tag_must_be_unique")
)

// Auth Error
var (
	ErrAuthThisEmailOrUsernameIsAlreadyUsed = errors.New("this_email_or_username_is_already_used")
	ErrAuthWrongEmailOrPassword             = errors.New("wrong_email_or_password")
	ErrAuthTokenExpired                     = errors.New("token_expired")
	ErrAuthTokenInvalid                     = errors.New("token_invalid")
	ErrAuthTokenNotProvided                 = errors.New("token_not_provided")
)
