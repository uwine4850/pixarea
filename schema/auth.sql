CREATE TABLE IF NOT EXISTS `pixarea`.`auth` (
    `id` INT NOT NULL AUTO_INCREMENT ,
    `username` VARCHAR(200) NOT NULL ,
    `password` TEXT NOT NULL ,
    PRIMARY KEY (`id`)
);

CREATE TABLE IF NOT EXISTS `pixarea`.`user` (
    `id` INT NOT NULL AUTO_INCREMENT ,
    `auth` INT NOT NULL,
    `name` VARCHAR(200) NULL ,
    `avatar` TEXT NULL,
    `bg_image` TEXT NULL,
    `description` TEXT NULL ,
    PRIMARY KEY (`id`),
    FOREIGN KEY (auth) REFERENCES pixarea.auth(id) ON DELETE CASCADE
);