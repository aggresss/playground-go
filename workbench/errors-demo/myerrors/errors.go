package myerrors

import "errors"

var (
	ERR_EOF            = errors.New("EOF")
	ERR_CLOSED_PIPE    = errors.New("io: read/write on closed pipe")
	ERR_NO_PROGRESS    = errors.New("multiple Read calls return no data or error")
	ERR_SHORT_BUFFER   = errors.New("short buffer")
	ERR_SHORT_WRITE    = errors.New("short write")
	ERR_UNEXPECTED_EOF = errors.New("unexpected EOF")
)
