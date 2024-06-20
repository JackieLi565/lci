package database

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/JackieLi565/lci/cli/config"
	"github.com/jackc/pgx/v5"
)

type DB struct {
	conn *pgx.Conn
}

type Problem struct {
	ID            int
	Difficulty    config.DatabaseConfig
	URL           string
	CreatedAt     time.Time
	Topics        []int
	Hint          string
	LastAttempted *time.Time
	Status        *int
}

type Tag struct {
	ID   int
	Name string
}

type Pattern struct {
	Tag
	Count int
}

func NewDB(config config.DatabaseConfig) *DB {
	ctx := context.Background()
	url := generateDatabaseURL(config.User, config.Password, config.DBName)
	conn, err := pgx.Connect(ctx, url)
	if err != nil {
		log.Fatal("lci: database connection failed\nSee 'lci config --help'")
	}

	if err := conn.Ping(ctx); err != nil {
		log.Fatal("lci: database ping failed\nSee 'lci config --help'")
	}

	return &DB{conn: conn}
}

func (d *DB) Init() error {
	sql := `
		create type difficulty as enum ('easy', 'medium', 'hard');
		
		create table patterns (
			pattern_id serial primary key,
			name text not null unique
		);

		create table tags (
			tag_id serial primary key,
			name text not null unique
		);
	`
	_, err := d.conn.Exec(context.Background(), sql)
	return err
}

func (d *DB) AddPattern(pattern string) error {
	patternTable := patternPrefix(pattern)
	tx, err := d.conn.Begin(context.Background())
	if err != nil {
		return err
	}

	defer tx.Rollback(context.Background())

	createTableSQL := fmt.Sprintf(`
        create table %s (
            id serial primary key, 
            difficulty difficulty not null,
            url text not null,
            created_at date not null,
            topics int[] not null,
            hint text not null,
            status int references status(id),
            last_attempted date
        );
    `, patternTable)

	if _, err := tx.Exec(context.Background(), createTableSQL); err != nil {
		return err
	}

	insertSQL := `INSERT INTO patterns (name) VALUES ($1);`
	if _, err := tx.Exec(context.Background(), insertSQL, patternTable); err != nil {
		return err
	}

	return tx.Commit(context.Background())
}

func (d *DB) ListPatterns() ([]*Pattern, error) {
	rows, err := d.conn.Query(context.Background(), "select * from patterns")
	if err != nil {
		return nil, err
	}

	patterns, err := pgx.CollectRows(rows, pgx.RowToStructByName[Tag])
	if err != nil {
		return nil, err
	}

	var patternList = make([]*Pattern, 0, len(patterns))

	for _, pattern := range patterns {
		var count int
		sql := fmt.Sprintf("select count(*) from %s;", pattern.Name)
		err := d.conn.QueryRow(context.Background(), sql).Scan(&count)
		if err != nil {
			return patternList, err
		}

		patternList = append(patternList, &Pattern{
			Tag: Tag{
				ID:   pattern.ID,
				Name: strings.TrimPrefix(pattern.Name, "pattern_"),
			},
			Count: count,
		})
	}

	return patternList, nil
}

func (d *DB) RemovePattern(name string) error {
	patternTable := patternPrefix(name)
	tx, err := d.conn.Begin(context.Background())
	if err != nil {
		return err
	}

	defer tx.Rollback(context.Background())

	if _, err := tx.Exec(context.Background(), "delete from patterns where name = $1", patternTable); err != nil {
		return err
	}

	sql := fmt.Sprintf("drop table %s", patternTable)
	if _, err := tx.Exec(context.Background(), sql); err != nil {
		return err
	}

	return tx.Commit(context.Background())
}

func (d *DB) AddTag(tag string) error {
	tagTable := tagPrefix(tag)
	tx, err := d.conn.Begin(context.Background())
	if err != nil {
		return err
	}

	defer tx.Rollback(context.Background())

	createTableSQL := fmt.Sprintf(`
		create table %s (
			id serial primary key,
			name text not null
		);
	`, tagTable)

	if _, err := tx.Exec(context.Background(), createTableSQL); err != nil {
		return err
	}

	insertSQL := `insert into tags (name) values ($1);`
	if _, err := tx.Exec(context.Background(), insertSQL, tagTable); err != nil {
		return err
	}

	return tx.Commit(context.Background())
}

func (d *DB) ListTags() ([]*Tag, error) {
	rows, err := d.conn.Query(context.Background(), "select * from tags")
	if err != nil {
		return nil, err
	}

	tags, err := pgx.CollectRows(rows, pgx.RowToStructByName[*Tag])
	if err != nil {
		return nil, err
	}

	for _, tag := range tags {
		tag.Name = strings.TrimPrefix("tags_", tag.Name)
	}

	return tags, nil
}

func (d *DB) RemoveTag(name string) error {
	tagTable := tagPrefix(name)
	tx, err := d.conn.Begin(context.Background())
	if err != nil {
		return err
	}

	defer tx.Rollback(context.Background())

	if _, err := tx.Exec(context.Background(), "delete from tags where name = $1", tagTable); err != nil {
		return err
	}

	sql := fmt.Sprintf("drop table %s", tagTable)
	if _, err := tx.Exec(context.Background(), sql); err != nil {
		return err
	}

	return tx.Commit(context.Background())
}

func (d *DB) AddProblem(pattern string, p Problem) (int, error) {
	sql := fmt.Sprintf("insert into %s (difficulty, url, status, created_at, topics, last_attempted, hint) values ($1, $2, $3, $4, $5, $6, $7) returning id", pattern)
	var id int
	err := d.conn.QueryRow(context.Background(), sql, p.Difficulty, p.URL, p.Status, p.CreatedAt, p.Topics, p.LastAttempted, p.Hint).Scan(&id)
	if err != nil {
		return id, err
	}
	return id, nil
}

func (d *DB) AddTagEntry(tag string, t Tag) (int, error) {
	sql := fmt.Sprintf("insert into %s (name) values ($1) returning id", tag)
	var id int
	err := d.conn.QueryRow(context.Background(), sql, t.Name).Scan(&id)
	if err != nil {
		return id, err
	}
	return id, nil
}

func (d *DB) GetTagEntry(tag string, name string) *Tag {
	var t Tag
	sql := fmt.Sprintf("select id, name from %s where name=$1", tag)
	err := d.conn.QueryRow(context.Background(), sql, name).Scan(t.ID, t.Name)
	if err != nil {
		return nil
	}
	return &t
}

func generateDatabaseURL(username string, password string, dbname string) string {
	addr := "localhost:5432"
	return fmt.Sprintf("postgres://%s:%s@%s/%s", username, password, addr, dbname)
}

func patternPrefix(name string) string {
	return fmt.Sprintf("pattern_%s", name)
}

func tagPrefix(name string) string {
	return fmt.Sprintf("tag_%s", name)
}
