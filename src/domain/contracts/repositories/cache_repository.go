package repositories

type CacheRepository interface {
	SetByKey(key string, object string) error
	GetByKey(key string) (string, error)
}
