package database

type CollectionName string

const (
	Recipes     CollectionName = "recipes"
	Employees   CollectionName = "employees"
	Restaurants CollectionName = "restaurants"
)

type Entity interface {
	GetID() ID
	SetID(id ID)
}

type Collection[T Entity] interface {
	Insert(t T) error
	Update(t T) error
	Delete(id ID) error
	FindBy(criteria []Criterion, sorts []Sort, limit Limit) ([]*T, error)
	FindOneBy(criteria []Criterion) (*T, error)
	Find(id ID) (*T, error)
	FindAll() ([]*T, error)
	Truncate() error
}
