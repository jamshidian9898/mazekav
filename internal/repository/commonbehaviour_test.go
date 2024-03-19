package repository

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	"mazekav/internal/pkg/testutil"
)

type MyTable struct {
	gorm.Model
	Name string `json:"name"`
}

func (m MyTable) Table() string {
	return "my_table"
}

func TestCommonBehaviour_save(t *testing.T) {

	db, cancel := testutil.CreateTestDB(t)
	defer cancel()

	assert.NoError(t, db.AutoMigrate(&MyTable{}))

	cb := NewCommonBehaviour[MyTable](db)

	// test save first record
	myRow := &MyTable{Name: "mohammad"}
	assert.NoError(t, cb.Save(context.Background(), myRow))
	assert.Equal(t, uint(1), myRow.ID)

	// test update name in first record
	myRow.Name = "ali"
	assert.NoError(t, cb.Save(context.Background(), myRow))
	assert.Equal(t, uint(1), myRow.ID)
}

func TestCommonBehaviour_ByID(t *testing.T) {
	db, cancel := testutil.CreateTestDB(t)
	defer cancel()

	assert.NoError(t, db.AutoMigrate(&MyTable{}))

	db.Create(&MyTable{Name: "mohammad"})
	db.Create(&MyTable{Name: "ali"})

	cb := NewCommonBehaviour[MyTable](db)

	model, err := cb.ByID(context.Background(), 2)

	assert.NoError(t, err)
	assert.Equal(t, "ali", model.Name)
}
