DROP DATABASE IF EXISTS app;
CREATE DATABASE app;
USE app;
CREATE TABLE `user`
(
    `id`         int(11) NOT NULL AUTO_INCREMENT,
    `email`      varchar(255) NOT NULL,
    `username`   varchar(255) NOT NULL,
    `password`   varchar(255) NOT NULL,
    `created_on` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_on` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `email` (`email`),
    UNIQUE KEY `username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
CREATE TABLE `token`
(
    `id`                int          NOT NULL AUTO_INCREMENT,
    `token_id`          varchar(255) NOT NULL,
    `name`              varchar(255) NOT NULL,
    `ticker`            varchar(255) NOT NULL,
    `symbol`            varchar(255) NOT NULL,
    `price`             float        NOT NULL,
    `change_percentage` float                 DEFAULT NULL,
    `created_on`        timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_on`        timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `token_id` (`token_id`),
    KEY                 `name` (`name`),
    KEY                 `ticker` (`ticker`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
CREATE TABLE `favorite`
(
    `id`         int      NOT NULL AUTO_INCREMENT,
    `user_id`    int      NOT NULL,
    `token_id`   int      NOT NULL,
    `created_on` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_on` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY          `user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;