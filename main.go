package main

import (
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Tag struct {
	ID        int64     `db:"id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
}

type Article struct {
	ID        int64     `db:"id"`
	Title     string    `db:"title"`
	Body      string    `db:"body"`
	UserId    int64     `db:"user_id"`
	CreatedAt time.Time `db:"created_at"`
}

type ArticleTag struct {
	ID        int64 `db:"id"`
	ArticleId int64 `db:"article_id"`
	TagId     int64 `db:"tag_id"`
}

type ArticleResponse struct {
	ID        int64     `db:"id"`
	Title     string    `db:"title"`
	Body      string    `db:"body"`
	UserId    int64     `db:"user_id"`
	CreatedAt time.Time `db:"created_at"`
	Tags      []Tag
}

func main() {
	db, err := sqlx.Connect("postgres", "host=pg user=admin password=password dbname=sqlx_db sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	tagSlice := []string{"Curl", "HTTP", "HTTP2"}
	tagIdSlice := []int64{}
	for _, v := range tagSlice {
		var tagId int64
		err := db.QueryRowx(`INSERT INTO tags (name) VALUES ($1) on conflict (name) do UPDATE SET name=$1 RETURNING id`, v).Scan(&tagId)
		if err != nil {
			log.Panic(err)
		}
		tagIdSlice = append(tagIdSlice, tagId)
	}

	fmt.Println(tagIdSlice)

	var createId int64
	err = db.QueryRowx(`INSERT INTO articles (title, body, user_id) VALUES ($1, $2, $3) RETURNING id`, "Title1", "Body1", 1).Scan(&createId)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(createId)

	for _, t := range tagIdSlice {
		var articletagId int64
		err = db.QueryRowx(`INSERT INTO articles_tags (article_id, tag_id) VALUES ($1, $2) RETURNING id`, createId, t).Scan(&articletagId)
		if err != nil {
			log.Panic(err)
		}
		fmt.Println(articletagId)
	}

	var articletag []ArticleTag
	err = db.Select(&articletag, "SELECT * FROM articles_tags")
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(articletag)

	var articleDetail Article
	err = db.Get(&articleDetail, "SELECT id, title, body, user_id, created_at FROM articles WHERE id = $1", createId)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(articleDetail)

	articleTags := []Tag{}
	err = db.Select(&articleTags, "SELECT tags.id, tags.name, tags.created_at FROM tags LEFT JOIN articles_tags ON tags.id = articles_tags.tag_id WHERE articles_tags.article_id = $1", createId)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(articleTags)

	articleShow := &ArticleResponse{
		ID:        articleDetail.ID,
		Title:     articleDetail.Title,
		Body:      articleDetail.Body,
		UserId:    articleDetail.UserId,
		CreatedAt: articleDetail.CreatedAt,
		Tags:      articleTags,
	}

	fmt.Println(articleShow)
}
