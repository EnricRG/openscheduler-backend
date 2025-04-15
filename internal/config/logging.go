package config

import (
	"reflect"
	"strings"

	"github.com/EnricRG/openscheduler-backend/internal/logging"
)

// LoggingConfig contains all config values related to logging system configuration
type LoggingConfig struct {
	Level  string            // Global logging level
	Levels map[string]string // Map of logging levels. Keys are formatted as <package>.<type-name>
}

// resolveLevel determines the logging level based on the package and type names.
func (c LoggingConfig) ResolveLevel(typeValue any) (level logging.Level) {
	if typeValue == nil {
		return parseLevel(c.Level)
	}

	valueType := reflect.TypeOf(typeValue)
	packageName := valueType.PkgPath()
	typeName := valueType.Name()

	levelStr, packageLevelDefined := c.Levels[packageName+"."+typeName]
	if packageLevelDefined {
		level = parseLevel(levelStr)
	} else {
		level = parseLevel(c.Level)
	}
	return
}

func getPackageName(pkgPath string) string {
	segments := strings.Split(pkgPath, `/`)
	return segments[len(segments)-1]
}

const (
	traceStr = "TRACE"
	debugStr = "DEBUG"
	infoStr  = "INFO"
	warnStr  = "WARN"
	errorStr = "ERROR"
)

func parseLevel(levelStr string) (level logging.Level) {
	cleanLevel := strings.TrimSpace(levelStr)

	if strings.EqualFold(traceStr, cleanLevel) {
		level = logging.Trace
	} else if strings.EqualFold(debugStr, cleanLevel) {
		level = logging.Debug
	} else if strings.EqualFold(infoStr, cleanLevel) {
		level = logging.Info
	} else if strings.EqualFold(warnStr, cleanLevel) {
		level = logging.Warn
	} else if strings.EqualFold(errorStr, cleanLevel) {
		level = logging.Error
	} else {
		panic("Unexpected logging value '" + cleanLevel + "'")
	}
	return
}
