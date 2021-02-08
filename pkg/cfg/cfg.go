package cfg

import (
	"github.com/i582/php-parser/pkg/errors"
	"github.com/i582/php-parser/pkg/version"
)

type Config struct {
	Version          *version.Version
	ErrorHandlerFunc func(e *errors.Error)
}
