-- Active: 1709309448667@@127.0.0.1@3306
create schema if not exists puuclocks;

use puuclocks;

DROP TABLE IF EXISTS `user`;

CREATE TABLE `user` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;