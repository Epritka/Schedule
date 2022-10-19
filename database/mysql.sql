CREATE TABLE `schedule` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `year` INT NOT NULL,
    `group_id` INT NOT NULL,
    PRIMARY KEY (`id`)
);

CREATE TABLE `person` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL,
    `group_id` INT NOT NULL,
    PRIMARY KEY (`id`)
);

CREATE TABLE `group` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL,
    `number` INT NOT NULL,
    `year` INT NOT NULL,
    PRIMARY KEY (`id`)
);

CREATE TABLE `lesson` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL,
    `type` VARCHAR(255) NOT NULL,
    `teacher` VARCHAR(255) NOT NULL,
    `auditorium` VARCHAR(255) NOT NULL,
    `sub_group` INT NOT NULL,
    `start_time` TIME NOT NULL,
    `end_time` TIME NOT NULL,
    `week_id` INT NOT NULL,
    PRIMARY KEY (`id`)
);

CREATE TABLE `week` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `type` BOOLEAN NOT NULL,
    `day_number` INT NOT NULL,
    `schedule_id` INT NOT NULL,
    PRIMARY KEY (`id`)
);

ALTER TABLE
    `schedule`
ADD
    CONSTRAINT `schedule_fk0` FOREIGN KEY (`group_id`) REFERENCES `group`(`id`);

ALTER TABLE
    `person`
ADD
    CONSTRAINT `person_fk0` FOREIGN KEY (`group_id`) REFERENCES `group`(`id`);

ALTER TABLE
    `lesson`
ADD
    CONSTRAINT `lesson_fk0` FOREIGN KEY (`week_id`) REFERENCES `week`(`id`);

ALTER TABLE
    `week`
ADD
    CONSTRAINT `week_fk0` FOREIGN KEY (`schedule_id`) REFERENCES `schedule`(`id`);