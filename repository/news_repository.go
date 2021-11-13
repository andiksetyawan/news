package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"news/entity"
	"news/model"
	"time"
)

type NewsRepository interface {
	CreateNews(ctx context.Context, news entity.News) error
	FindNewsByID(ctx context.Context, id string) (*entity.News, error)
	UpdateNewsByID(ctx context.Context, payload model.NewsUpdate, id string) error
	FindNews(ctx context.Context, filter bson.M) (*[]entity.News, error)
	DeleteNewsByID(ctx context.Context, id string) error
	FindNewsBySlug(ctx context.Context, slug string) (*entity.News, error)
}

type newsRepository struct {
	Db *mongo.Collection
}

func NewNewsRepo(db *mongo.Database) NewsRepository {
	return &newsRepository{
		Db: db.Collection("news"),
	}
}

func (repository *newsRepository) CreateNews(ctx context.Context, news entity.News) error {
	_, err := repository.Db.InsertOne(ctx, news)
	if err != nil {
		return err
	}
	return nil
}

func (repository *newsRepository) FindNewsByID(ctx context.Context, id string) (*entity.News, error) {
	filter := bson.M{"id": id}
	filter["$or"] = []bson.M{
		{"deleted_at": 0},
		{"deleted_at": bson.M{"$exists": false}},
	}
	var news entity.News
	err := repository.Db.FindOne(ctx, filter).Decode(&news)
	if err != nil {
		return nil, err
	}
	return &news, nil
}

func (repository *newsRepository) FindNews(ctx context.Context, filter bson.M) (*[]entity.News, error) {
	filter["$or"] = []bson.M{
		{"deleted_at": 0},
		{"deleted_at": bson.M{"$exists": false}},
	}
	cur, err := repository.Db.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var news []entity.News
	cur.All(ctx, &news)
	return &news, nil
}

func (repository *newsRepository) FindNewsBySlug(ctx context.Context, slug string) (*entity.News, error) {
	filter := bson.M{"slug": slug}
	filter["$or"] = []bson.M{
		{"deleted_at": 0},
		{"deleted_at": bson.M{"$exists": false}},
	}
	var news entity.News
	err := repository.Db.FindOne(ctx, filter).Decode(&news)
	if err != nil {
		return nil, err
	}
	return &news, nil
}

func (repository *newsRepository) DeleteNewsByID(ctx context.Context, id string) error {
	_, err := repository.Db.UpdateOne(ctx, bson.M{"id": id}, bson.M{"$set": bson.M{"deleted_at": time.Now().Unix()}})
	if err != nil {
		return err
	}
	return nil
}

func (repository *newsRepository) UpdateNewsByID(ctx context.Context, payload model.NewsUpdate, id string) error {
	_, err := repository.Db.UpdateOne(ctx, bson.M{"id": id}, bson.M{"$set": payload})
	if err != nil {
		return err
	}
	return nil
}
