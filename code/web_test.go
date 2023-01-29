package owntest

import (
	"context"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

type bookStore struct {
	db *sql.DB
}

func (b *bookStore) InsertBook(ctx context.Context, book *Book) error {
	book.ID = 1
	return nil
}

type Book struct {
	ID       int64
	Title    string
	AuthorID int64
	ISBN     string
	Subject  string
}

// Test_DAO_Insert
func Test_DAO_Insert(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		assert.NoError(t, err, "an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	store := &bookStore{
		db: db,
	}
	book := &Book{
		Title:    "The Go Programming Language",
		AuthorID: 1,
		ISBN:     "978-0134190440",
		Subject:  "computers",
	}

	mock.ExpectExec("INSERT INTO books").WillReturnResult(sqlmock.NewResult(1, 1))
	err = store.InsertBook(context.TODO(), book) // 模拟数据库操作
	require.NoError(t, err)
	assert.NotZero(t, book.ID)
}
