package db

import (
	"github.com/M1ralai/me-portfolio/cmd/api/types"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var logger = types.NewLogger("database")

func InitDb() {
	dbp, err := sqlx.Open("sqlite3", "db/posts.db")
	if err != nil {
		logger.Fatal(err)
	}
	defer dbp.Close()
	_, err = dbp.Exec(createPostTable)
	if err != nil {
		logger.Println(err)
	}

	dbe, err := sqlx.Open("sqlite3", "db/mails.db")
	if err != nil {
		logger.Fatal(err)
	}
	defer dbe.Close()

	_, err = dbe.Exec(createMailTable)
	if err != nil {
		logger.Println(err)
	}
	logger.Println("db successfully initialized")
}

func GetPostsForIndex(from int, count int) ([]types.Post, error) {
	db, err := sqlx.Open("sqlite3", "db/posts.db")
	if err != nil {
		logger.Println(err)
		return nil, err
	}

	var Posts []types.Post

	err = db.Select(&Posts, getPostForIndex, from, count)
	if err != nil {
		logger.Println(err)
		return nil, err
	}
	return Posts, nil
}

func CreatePost(post types.Post) error {
	db, err := sqlx.Open("sqlite3", "db/posts.db")
	if err != nil {
		logger.Println(err)
		return err
	}

	_, err = db.Exec(createPost, post.Title, post.Content, post.Excerpt, post.Date)
	if err != nil {
		logger.Println(err)
		return err
	}

	return nil
}

func DeletePost(id int) error {
	db, err := sqlx.Open("sqlite3", "db/posts.db")
	if err != nil {
		logger.Println(err)
		return err
	}

	_, err = db.Exec(deletePost, id)
	if err != nil {
		logger.Println(err)
		return err
	}

	return nil
}

func GetPostById(id int) (types.Post, error) {
	db, err := sqlx.Open("sqlite3", "db/posts.db")
	if err != nil {
		logger.Println(err)
		return types.Post{}, err
	}
	var p types.Post
	err = db.Get(&p, getPostById, id)
	if err != nil {
		logger.Println(err)
		return types.Post{}, err
	}
	return p, nil
}
