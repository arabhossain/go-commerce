

CREATE TABLE IF NOT EXISTS users
(
    `id`            INT             NOT NULL AUTO_INCREMENT,
    `first_name`    TEXT            NULL,
    `last_name`     TEXT            NULL,
     username       VARCHAR(100)    NOT NULL UNIQUE,
     email          VARCHAR(100)    NOT NULL UNIQUE,
    `password`      TEXT            NOT NULL,
    `created`       DATETIME        NOT NULL,
    `modified`      DATETIME        NOT NULL,
    PRIMARY KEY (`id`)
);
