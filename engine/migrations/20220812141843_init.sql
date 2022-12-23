-- +goose Up
-- +goose StatementBegin
CREATE TABLE student_group (
    id INT GENERATED ALWAYS AS IDENTITY,
    name VARCHAR(255) NOT NULL,
    number INT NOT NULL,
    year INT NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE schedule (
    id INT GENERATED ALWAYS AS IDENTITY,
    year INT NOT NULL,
    student_group_id INT NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT schedule_fk0 FOREIGN KEY (student_group_id) REFERENCES student_group(id)
);

CREATE TABLE student (
    id INT GENERATED ALWAYS AS IDENTITY,
    name VARCHAR(255) NOT NULL,
    student_group_id INT NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT student_fk0 FOREIGN KEY (student_group_id) REFERENCES student_group(id)
);

CREATE TABLE week (
    id serial NOT NULL,
    type BOOLEAN NOT NULL,
    day_number INT NOT NULL,
    schedule_id INT NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT week_fk0 FOREIGN KEY (schedule_id) REFERENCES schedule(id)
);

CREATE TABLE lesson (
    id INT GENERATED ALWAYS AS IDENTITY,
    name VARCHAR(255) NOT NULL,
    type VARCHAR(255) NOT NULL,
    teacher VARCHAR(255) NOT NULL,
    auditorium VARCHAR(255) NOT NULL,
    sub_group VARCHAR(255) NULL,
    start_time VARCHAR(255) NOT NULL,
    end_time VARCHAR(255) NOT NULL,
    week_id INT NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT lesson_fk0 FOREIGN KEY (week_id) REFERENCES week(id)
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE student_group;

DROP TABLE schedule;

DROP TABLE student;

DROP TABLE week;

DROP TABLE lesson;

-- +goose StatementEnd