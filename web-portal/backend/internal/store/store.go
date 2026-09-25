package store

import (
	"database/sql"
	"time"

	_ "modernc.org/sqlite"
)

type Release struct {
	ID          int64
	Repo        string
	TagName     string
	Name        string
	Author      string
	HTMLURL     string
	PublishedAt time.Time
	ReceivedAt  time.Time
}

type Store struct {
	db *sql.DB
}

func New(path string) (*Store, error) {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}

	if err := migrate(db); err != nil {
		return nil, err
	}

	return &Store{db: db}, nil
}

func migrate(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS releases (
			id          INTEGER PRIMARY KEY AUTOINCREMENT,
			repo        TEXT NOT NULL,
			tag_name    TEXT NOT NULL,
			name        TEXT NOT NULL,
			author      TEXT NOT NULL,
			html_url    TEXT NOT NULL,
			published_at DATETIME NOT NULL,
			received_at  DATETIME NOT NULL
		)
	`)
	return err
}

func (s *Store) InsertRelease(r Release) error {
	_, err := s.db.Exec(`
		INSERT INTO releases (repo, tag_name, name, author, html_url, published_at, received_at)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`,
		r.Repo, r.TagName, r.Name, r.Author, r.HTMLURL, r.PublishedAt, r.ReceivedAt,
	)
	return err
}

func (s *Store) ListReleases(limit int) ([]Release, error) {
	rows, err := s.db.Query(`
		SELECT id, repo, tag_name, name, author, html_url, published_at, received_at
		FROM releases
		ORDER BY received_at DESC
		LIMIT ?
	`, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var releases []Release
	for rows.Next() {
		var r Release
		if err := rows.Scan(&r.ID, &r.Repo, &r.TagName, &r.Name, &r.Author, &r.HTMLURL, &r.PublishedAt, &r.ReceivedAt); err != nil {
			return nil, err
		}
		releases = append(releases, r)
	}
	return releases, nil
}
