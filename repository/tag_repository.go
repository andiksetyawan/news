package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"news/entity"
)

type TagRepository interface {
	CreateTag(ctx context.Context, news entity.Tag) error
	FindTagBySlug(ctx context.Context, slug string) (*entity.Tag, error)
	UpdateTagBySlug(ctx context.Context, payload bson.M, id string) error
	FindTags(ctx context.Context, filter bson.M) (*[]entity.Tag, error)
	DeleteTagBySlug(ctx context.Context, id string) error
}

type tagRepository struct {
	Db *mongo.Collection
}

func NewTagRepo(db *mongo.Database) TagRepository {
	return &tagRepository{
		Db: db.Collection("tag"),
	}
}

func (repository *tagRepository) CreateTag(ctx context.Context, tag entity.Tag) error {
	_, err := repository.Db.InsertOne(ctx, tag)
	if err != nil {
		return err
	}
	return nil
}

func (repository *tagRepository) FindTagBySlug(ctx context.Context, slug string) (*entity.Tag, error) {
	var tag entity.Tag
	err := repository.Db.FindOne(ctx, bson.M{"slug": slug}).Decode(&tag)
	if err != nil {
		return nil, err
	}
	return &tag, nil
}

func (repository *tagRepository) FindTags(ctx context.Context, filter bson.M) (*[]entity.Tag, error) {
	cur, err := repository.Db.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var tags []entity.Tag
	cur.All(ctx, &tags)
	return &tags, nil
}

func (repository *tagRepository) DeleteTagBySlug(ctx context.Context, slug string) error {
	_, err := repository.Db.DeleteOne(ctx, bson.M{"slug": slug})
	if err != nil {
		return err
	}
	return nil
}

func (repository *tagRepository) UpdateTagBySlug(ctx context.Context, payload bson.M, slug string) error {
	_, err := repository.Db.UpdateOne(ctx, bson.M{"slug": slug}, bson.M{"$set": payload})
	if err != nil {
		return err
	}
	return nil
}
