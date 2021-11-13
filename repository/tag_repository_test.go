package repository

import (
	"context"
	"fmt"
	"github.com/gosimple/slug"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"news/db"
	"news/entity"
	"testing"
	"time"
)

var tagRepoTest = NewTagRepo(db.InitMongo())

func TestTagRepository_CreateTag_FindTagBySlug(t *testing.T) {
	name := "Test Tag Name" + fmt.Sprint(time.Now().Unix())
	s := slug.Make(name)
	createdAt := time.Now().Unix()

	tagRepoTest.CreateTag(context.Background(), entity.Tag{
		Slug:      s,
		Name:      name,
		CreatedAt: createdAt,
		UpdatedAt: 0,
	})

	tag, _ := tagRepoTest.FindTagBySlug(context.Background(), s)
	assert.Equal(t, name, tag.Name)
	assert.Equal(t, s, tag.Slug)
	assert.Equal(t, createdAt, tag.CreatedAt)
	assert.Equal(t, int64(0), tag.UpdatedAt)

	tagRepoTest.DeleteTagBySlug(context.Background(), s)
}

func TestTagRepository_UpdateTagBySlug(t *testing.T) {
	name := "Test Tag Name" + fmt.Sprint(time.Now().Unix())
	s := slug.Make(name)
	createdAt := time.Now().Unix()

	tagRepoTest.CreateTag(context.Background(), entity.Tag{
		Slug:      s,
		Name:      name,
		CreatedAt: createdAt,
		UpdatedAt: 0,
	})
	tagRepoTest.UpdateTagBySlug(context.Background(), bson.M{"Name": "ubah nama tag"}, s)

	tag, _ := tagRepoTest.FindTagBySlug(context.Background(), s)
	assert.Equal(t, "ubah nama tag", tag.Name)

	tagRepoTest.DeleteTagBySlug(context.Background(), s)
}

func TestTagRepository_DeleteTagBySlug(t *testing.T) {
	name := "Test Tag Name" + fmt.Sprint(time.Now().Unix())
	s := slug.Make(name)
	createdAt := time.Now().Unix()

	tagRepoTest.CreateTag(context.Background(), entity.Tag{
		Slug:      s,
		Name:      name,
		CreatedAt: createdAt,
		UpdatedAt: 0,
	})

	tagRepoTest.DeleteTagBySlug(context.Background(), s)

	tag, err := tagRepoTest.FindTagBySlug(context.Background(), s)
	assert.Error(t, err)
	assert.Nil(t, tag)
}
