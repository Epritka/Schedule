-- +goose Up
-- +goose StatementBegin
CREATE TYPE WEEK_TYPE AS ENUM ('even', 'odd');

CREATE TABLE educational_institution (
    id INT GENERATED ALWAYS AS IDENTITY,
    name VARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE faculty (
    id INT GENERATED ALWAYS AS IDENTITY,
    name VARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE year (
    id INT GENERATED ALWAYS AS IDENTITY,
    name VARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE student_group (
    id INT GENERATED ALWAYS AS IDENTITY,
    name VARCHAR(255) NOT NULL,
    educational_institution_id INT NOT NULL,
    faculty_id INT NOT NULL,
    year_id INT NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT student_group_fk0 FOREIGN KEY (educational_institution_id) REFERENCES educational_institution(id),
    CONSTRAINT student_group_fk1 FOREIGN KEY (faculty_id) REFERENCES faculty(id),
    CONSTRAINT student_group_fk3 FOREIGN KEY (year_id) REFERENCES year(id)
);

CREATE TABLE student (
    id INT GENERATED ALWAYS AS IDENTITY,
    user_id INT NOT NULL,
    group_id INT NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT student_fk0 FOREIGN KEY (group_id) REFERENCES student_group(id)
);

CREATE TABLE day (
    id INT GENERATED ALWAYS AS IDENTITY,
    group_id INT NOT NULL,
    number INT NOT NULL,
    week_type WEEK_TYPE NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT day_fk0 FOREIGN KEY (group_id) REFERENCES student_group(id)
);

CREATE TABLE lesson (
    id INT GENERATED ALWAYS AS IDENTITY,
    day_id INT NOT NULL,
    name VARCHAR(255) NOT NULL,
    type VARCHAR(255) NOT NULL,
    teacher VARCHAR(255) NOT NULL,
    auditorium VARCHAR(255) NOT NULL,
    sub_group VARCHAR(255) NULL,
    start_time VARCHAR(255) NOT NULL,
    end_time VARCHAR(255) NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT lesson_fk0 FOREIGN KEY (day_id) REFERENCES day(id)
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE educational_institution;

DROP TABLE faculty;

DROP TABLE year;

DROP TABLE student_group;

DROP TABLE schedule;

DROP TABLE student;

DROP TABLE week;

DROP TABLE lesson;

DROP TYPE WEEK_TYPE;

-- +goose StatementEnd