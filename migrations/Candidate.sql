CREATE DATABASE IF NOT EXISTS `BeerData` CHARACTER SET utf8 COLLATE utf8_unicode_ci;

USE `BeerData`;

CREATE TABLE `BeerData`.`beer`
(
    `id`   int NOT NULL,
    `beer_name` VARCHAR(50) NOT NULL,
    `beer_type` VARCHAR(50) NOT NULL,
    `beer_desc` VARCHAR(50) NOT NULL,
    `bear_image` VARCHAR(2000) NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARACTER SET = utf8
  COLLATE = utf8_unicode_ci;

ALTER TABLE `beer`
    MODIFY `id` int NOT NULL AUTO_INCREMENT;

CREATE TABLE `BeerData`.`log`
(
    `id`   int NOT NULL,
    `log_method` VARCHAR(6) NOT NULL,
    `log_desc` VARCHAR(2000) NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARACTER SET = utf8
  COLLATE = utf8_unicode_ci;

ALTER TABLE `log`
    MODIFY `id` int NOT NULL AUTO_INCREMENT;
  


