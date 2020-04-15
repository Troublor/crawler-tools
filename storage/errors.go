package storage

import "errors"

var (
	IsDirErr        = errors.New("file path is a directory")
	IsNotDirErr     = errors.New("file path is not a directory")
	NotActivatedErr = errors.New("storage is not activated")
	NonPointerErr   = errors.New("payload is not a pointer")
)
