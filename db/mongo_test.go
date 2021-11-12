package db

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestInitMongoSuccess(t *testing.T) {
	os.Setenv("MONGO_URI","mongodb://207.148.76.65:27017")
	assert.NotNil(t, InitMongo())
	assert.NotPanics(t, func() { InitMongo() })
}

func TestInitMongoError(t *testing.T) {
	os.Setenv("MONGO_URI","mongodb://NOTVALID.HOST:27017")//Wrong Mongouri
	assert.Panics(t, func() { InitMongo() })
	os.Setenv("MONGO_URI","mongodb://207.148.76.65:27017")
}
