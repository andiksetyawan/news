package service

import (
	"context"
	"errors"
	"github.com/gosimple/slug"
	"go.mongodb.org/mongo-driver/bson"
	"news/entity"
	"news/model"
	"news/repository"
	"strings"
	"time"
)

type TagService interface {
	CreateTag(ctx context.Context, tag model.TagCreateRequest) error
	FindTagBySlug(ctx context.Context, slug string) (*entity.Tag, error)
	UpdateTagBySlug(ctx context.Context, tag model.TagCreateRequest, id string) error
	FindTags(ctx context.Context, filter bson.M) (*[]entity.Tag, error)
	DeleteTagBySlug(ctx context.Context, id string) error
}

func NewTagService(repository repository.TagRepository) TagService {
	return &tagService{
		repository: repository,
	}
}

type tagService struct {
	repository repository.TagRepository
}

func (service *tagService) CreateTag(ctx context.Context, tag model.TagCreateRequest) error {
	s := slug.Make(strings.TrimSpace(tag.Name))
	//find exist
	bySlug, err := service.repository.FindTagBySlug(ctx, s)
	if err == nil && bySlug.Slug != "" {
		return errors.New("exist tag")
	}
	newTag := entity.Tag{
		Slug:      s,
		Name:      tag.Name,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: 0,
	}
	err = service.repository.CreateTag(ctx, newTag)
	if err != nil {
		return err
	}

	return nil
}

func (service *tagService) FindTagBySlug(ctx context.Context, slug string) (*entity.Tag, error) {
	if slug == "" {
		return nil, errors.New("tag not exist")
	}

	t, err := service.repository.FindTagBySlug(ctx, slug)

	if err != nil {
		return nil, err
	}

	return t, nil
}

func (service *tagService) UpdateTagBySlug(ctx context.Context, tag model.TagCreateRequest, slug string) error {
	bySlug, err := service.repository.FindTagBySlug(ctx, slug)
	if err != nil || bySlug.Slug == "" {
		return errors.New("tag not exist")
	}

	payload := bson.M{"name": tag.Name, "updated_at": time.Now().Unix()}
	err = service.repository.UpdateTagBySlug(ctx, payload, slug)

	if err != nil {
		return err
	}

	return nil
}

func (service *tagService) FindTags(ctx context.Context, filter bson.M) (*[]entity.Tag, error) {
	news, err := service.repository.FindTags(ctx, filter)

	if err != nil {
		return nil, err
	}

	return news, nil
}

func (service *tagService) DeleteTagBySlug(ctx context.Context, slug string) error {
	err := service.repository.DeleteTagBySlug(ctx, slug)
	if err != nil {
		return err
	}

	return nil
}
