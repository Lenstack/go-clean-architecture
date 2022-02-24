package usecases

type LoggerRepository interface {
	LogError(string, ...interface{})
	LogAccess(string, ...interface{})
}
