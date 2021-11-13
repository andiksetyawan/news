package service

import (
	"context"
	"fmt"
	"github.com/gosimple/slug"
	"github.com/stretchr/testify/assert"
	"news/db"
	"news/model"
	"news/repository"
	"testing"
	"time"
)

var tagServTest = NewTagService(repository.NewTagRepo(db.InitMongo()))

func TestTagService_CreateTag(t *testing.T) {
	name := "Test Tag Name" + fmt.Sprint(time.Now().Unix())
	s := slug.Make(name)
	newTag := model.TagCreateRequest{Name: name}
	err := tagServTest.CreateTag(context.Background(), newTag)
	assert.Nil(t, err)
	tagServTest.DeleteTagBySlug(context.Background(), s)
}

func TestTagService_FindTagBySlug(t *testing.T) {
	name := "Test Tag Name" + fmt.Sprint(time.Now().Unix())
	s := slug.Make(name)
	newTag := model.TagCreateRequest{Name: name}
	err := tagServTest.CreateTag(context.Background(), newTag)
	assert.Nil(t, err)

	tag, err := tagServTest.FindTagBySlug(context.Background(), s)
	assert.Nil(t, err)
	assert.Equal(t, s, tag.Slug)
	assert.Equal(t, name, tag.Name)
	assert.Greater(t, tag.CreatedAt, int64(0))

	tagServTest.DeleteTagBySlug(context.Background(), s)
}

func TestTagService_UpdateTagBySlug(t *testing.T) {
	name := "Test Tag Name" + fmt.Sprint(time.Now().Unix())
	s := slug.Make(name)
	newTag := model.TagCreateRequest{Name: name}
	err := tagServTest.CreateTag(context.Background(), newTag)
	assert.Nil(t, err)

	err = tagServTest.UpdateTagBySlug(context.Background(), model.TagCreateRequest{Name: "Changed Name"}, s)
	assert.Nil(t, err)

	tag, _ := tagServTest.FindTagBySlug(context.Background(), s)

	assert.Equal(t, s, tag.Slug)
	assert.Equal(t, "Changed Name", tag.Name)
	assert.Greater(t, tag.UpdatedAt, int64(0))

	tagServTest.DeleteTagBySlug(context.Background(), s)
}
