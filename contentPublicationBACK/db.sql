CREATE TABLE IF NOT EXISTS mangas
(
    id                  INT AUTO_INCREMENT PRIMARY KEY,
    titleName           VARCHAR(200) NOT NULL,
    originalAuthor      VARCHAR(100) NOT NULL,
    creationDate        DATETIME,
    descriptionManga    TEXT,
    titleImage          LONGTEXT
);