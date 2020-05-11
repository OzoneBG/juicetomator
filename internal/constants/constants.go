package constants

import "errors"

var (
	// ErrFailedToBindJSON is thrown when the server cannot bind the body to the according struct.
	ErrFailedToBindJSON = errors.New("failed to bind json")

	// ErrFailedToExecuteScript is thrown when the script execution fails for some reason.
	ErrFailedToExecuteScript = errors.New("failed to execute script")

	//ErrFailedToFindZipFile is thrown when the zip file was not created or found.
	ErrFailedToFindZipFile = errors.New("failed to find zip file")

	//ErrNotImplemented is thrown when the methods is not implemented.
	ErrNotImplemented = errors.New("method was not implemented")
)
