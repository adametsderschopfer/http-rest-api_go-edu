package teststore_test

import (
	"github.com/adametsderschopfer/http-rest-api_go-edu/internal/store"
	"github.com/adametsderschopfer/http-rest-api_go-edu/internal/store/teststore"
	"github.com/adametsderschopfer/http-rest-api_go-edu/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	s := teststore.New()
	u := model.TestUser(t)
	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	email := "user@example.org"

	s := teststore.New()
	_, err := s.User().FindByEmail(email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u := model.TestUser(t)
	u.Email = email

	s.User().Create(u)

	u, err = s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
