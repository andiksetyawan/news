package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/gosimple/slug"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"news/entity"
	"news/model"
	"news/repository"
	"time"
)

type NewsService interface {
	CreateNews(ctx context.Context, req model.NewsCreateRequest) error
	FindNewsByID(ctx context.Context, id string) (*model.NewsDetailResponse, error)
	UpdateNewsByID(ctx context.Context, payload model.NewsUpdateRequest, id string) error
	FindNews(ctx context.Context, topic, status string) (*[]entity.News, error)
	DeleteNewsByID(ctx context.Context, id string) error
}

func NewNewsService(newsRepository repository.NewsRepository, tagRepository repository.TagRepository) NewsService {
	return &newsService{
		newsRepository: newsRepository,
		tagRepository:  tagRepository,
	}
}

type newsService struct {
	newsRepository repository.NewsRepository
	tagRepository  repository.TagRepository
}

func (service *newsService) CreateNews(ctx context.Context, req model.NewsCreateRequest) error {
	newNews := entity.News{
		ID:        uuid.New().String(),
		Slug:      slug.Make(req.Title),
		Title:     req.Title,
		Text:      req.Text,
		Tags:      req.Tags,
		Status:    "draft",
		CreatedAt: time.Now().Unix(),
		UpdatedAt: 0,
		DeletedAt: 0,
	}
	err := service.newsRepository.CreateNews(ctx, newNews)

	if err != nil {
		return err
	}

	return nil
}

func (service *newsService) FindNewsByID(ctx context.Context, id string) (*model.NewsDetailResponse, error) {
	news, err := service.newsRepository.FindNewsByID(ctx, id)

	if err != nil {
		return nil, err
	}

	newsRes := model.NewsDetailResponse{
		ID:        news.ID,
		Title:     news.Title,
		Slug:      news.Slug,
		Text:      news.Text,
		Tags:      make([]entity.Tag, 0),
		Status:    news.Status,
		CreatedAt: news.CreatedAt,
		UpdatedAt: news.UpdatedAt,
	}

	for _, tag := range news.Tags {
		bySlug, err := service.tagRepository.FindTagBySlug(ctx, tag)
		if err != nil {
			log.Println(err)
			continue
		}
		newsRes.Tags = append(newsRes.Tags, *bySlug)
	}

	return &newsRes, nil
}

func (service *newsService) UpdateNewsByID(ctx context.Context, payload model.NewsUpdateRequest, id string) error {
	news := model.NewsUpdate{
		Title:     payload.Title,
		Text:      payload.Text,
		Tags:      payload.Tags,
		Status:    payload.Status,
		UpdatedAt: time.Now().Unix(),
	}
	err := service.newsRepository.UpdateNewsByID(ctx, news, id)

	if err != nil {
		return err
	}

	return nil
}

func (service *newsService) FindNews(ctx context.Context, topic, status string) (*[]entity.News, error) {
	filter := bson.M{}
	if topic != "" {
		filter["tags"] = topic
	}
	if status != "" {
		filter["status"] = status
	}
	news, err := service.newsRepository.FindNews(ctx, filter)

	if err != nil {
		return nil, err
	}

	return news, nil
}

func (service *newsService) DeleteNewsByID(ctx context.Context, id string) error {
	err := service.newsRepository.DeleteNewsByID(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
