package adapters

type BucketAdapter interface {
	CreatePutObjectUrl(key string) (string, error)
	GetObjectUrl(key string) string
	DeleteObject(key string) error
}
