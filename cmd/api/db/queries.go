package db

var createPostTable = `
CREATE TABLE post (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    excerpt TEXT,
    date TEXT NOT NULL
);
`

var createMailTable = `
CREATE TABLE mail (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    date TEXT NOT NULL
);
`

var getPostForIndex = `SELECT excerpt, title, id FROM post ORDER BY date DESC LIMIT ? OFFSET ?;`
var createPost = `
INSERT INTO post (title, content, excerpt, date)
VALUES (?, ?, ?, ?);`

var deletePost = `
DELETE FROM post WHERE id = ?;`

var getPostById = `SELECT id, title, content, date FROM post WHERE id = ? ;`
