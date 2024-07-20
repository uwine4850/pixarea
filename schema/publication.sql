CREATE TABLE IF NOT EXISTS `pixarea`.`caledories_list` (
    `id` INT NOT NULL AUTO_INCREMENT ,
    `name` VARCHAR(200) NOT NULL ,
    PRIMARY KEY (`id`)
);

INSERT INTO pixarea.caledories_list (`id`, `name`) VALUES (NULL, 'Pixelart'), (NULL, 'Gamedev'), (NULL, "Paint");

CREATE TABLE IF NOT EXISTS `pixarea`.`publication` (
    `id` INT NOT NULL AUTO_INCREMENT ,
    `author` INT NOT NULL,
    `name` VARCHAR(200) NOT NULL ,
    `description` VARCHAR(600) NOT NULL ,
    `category1` INT NULL,
    `category2` INT NULL,
    `date` DATETIME NOT NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY (author) REFERENCES pixarea.auth(id) ON DELETE CASCADE,
    FOREIGN KEY (category1) REFERENCES pixarea.caledories_list(id) ON DELETE CASCADE,
    FOREIGN KEY (category2) REFERENCES pixarea.caledories_list(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS `pixarea`.`publication_images` (
    `id` INT NOT NULL AUTO_INCREMENT ,
    `publication` INT NOT NULL,
    `image_path` VARCHAR(600) NOT NULL ,
    PRIMARY KEY (`id`),
    FOREIGN KEY (publication) REFERENCES pixarea.publication(id) ON DELETE CASCADE
);