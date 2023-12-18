package repository

type Repository interface {
	Create(obj interface{}) (int64, error)
	Exists(obj interface{}) (bool, error)
	Get(args, obj interface{}) (interface{}, error)
	GetAll() ([]interface{}, error)
	Update(obj interface{}) (int64, error)
	Delete(id int64) (int64, error)
}
