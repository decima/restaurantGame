package mongoAdapter

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"restaurantAPI/lib/database"
)

type Collection[T database.Entity] struct {
	name            database.CollectionName
	mongoCollection *mongo.Collection
}

func (c Collection[T]) Insert(t T) error {
	insert, err := c.mongoCollection.InsertOne(context.Background(), t)

	if err != nil {
		return err
	}
	t.SetID(database.ID(insert.InsertedID.(string)))
	return nil
}

func (c Collection[T]) Update(t T) error {
	_, err := c.mongoCollection.ReplaceOne(context.Background(), map[string]interface{}{"_id": t.GetID()}, t)
	return err
}

func (c Collection[T]) Delete(id database.ID) error {
	_, err := c.mongoCollection.DeleteOne(context.Background(), map[string]interface{}{"_id": id})
	return err
}

func (c Collection[T]) FindBy(criteria []database.Criterion, sorts []database.Sort, limit database.Limit) ([]*T, error) {
	var res []*T

	filters := make(map[string]interface{})
	for _, criterion := range criteria {
		filters[criterion.Field] = criterion.Value
	}

	limitCount := int64(limit.Count)
	limitOffset := int64(limit.Offset)

	findOpts := options.FindOptions{}
	if limitCount > 0 {
		findOpts.SetLimit(limitCount)
	}

	if limitOffset > 0 {
		findOpts.SetSkip(limitOffset)
	}

	if sorts != nil {
		sort := make(map[string]int)
		for _, s := range sorts {
			sort[s.Field] = s.Order
		}
		findOpts.SetSort(sort)

	}

	cursor, err := c.mongoCollection.Find(context.Background(), filters, &findOpts)
	if err != nil {
		return nil, err
	}

	for cursor.Next(context.Background()) {
		var item T
		err := cursor.Decode(&item)
		if err != nil {
			return nil, err
		}

		res = append(res, &item)
	}
	return res, nil
}

func (c Collection[T]) FindOneBy(criteria []database.Criterion) (*T, error) {
	res, err := c.FindBy(criteria, nil, database.Limit{Count: 1, Offset: 0})
	if err != nil {
		return nil, err
	}
	if len(res) == 0 {
		return nil, errors.New("No results found")
	}

	return res[0], nil
}

func (c Collection[T]) Find(id database.ID) (*T, error) {
	return c.FindOneBy([]database.Criterion{{Field: "_id", Value: id}})
}

func (c Collection[T]) FindAll() ([]*T, error) {
	return c.FindBy(nil, nil, database.Limit{Count: 0, Offset: 0})
}

func (c Collection[T]) Truncate() error {
	_, err := c.mongoCollection.DeleteMany(context.Background(), map[string]interface{}{})
	return err
}
