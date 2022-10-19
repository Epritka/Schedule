CREATE TABLE "public.schedule" (
    "id" serial NOT NULL,
    "year" integer NOT NULL,
    "group_id" integer NOT NULL,
    CONSTRAINT "schedule_pk" PRIMARY KEY ("id")
) WITH (OIDS = FALSE);

CREATE TABLE "public.person" (
    "id" serial NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "group_id" integer NOT NULL,
    CONSTRAINT "person_pk" PRIMARY KEY ("id")
) WITH (OIDS = FALSE);

CREATE TABLE "public.group" (
    "id" serial NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "number" integer NOT NULL,
    "year" integer NOT NULL,
    CONSTRAINT "group_pk" PRIMARY KEY ("id")
) WITH (OIDS = FALSE);

CREATE TABLE "public.lesson" (
    "id" serial NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "type" VARCHAR(255) NOT NULL,
    "teacher" VARCHAR(255) NOT NULL,
    "auditorium" VARCHAR(255) NOT NULL,
    "sub_group" integer NOT NULL,
    "start_time" TIME NOT NULL,
    "end_time" TIME NOT NULL,
    "week_id" integer NOT NULL,
    CONSTRAINT "lesson_pk" PRIMARY KEY ("id")
) WITH (OIDS = FALSE);

CREATE TABLE "public.week" (
    "id" serial NOT NULL,
    "type" BOOLEAN NOT NULL,
    "day_number" integer NOT NULL,
    "schedule_id" integer NOT NULL,
    CONSTRAINT "week_pk" PRIMARY KEY ("id")
) WITH (OIDS = FALSE);

ALTER TABLE
    "schedule"
ADD
    CONSTRAINT "schedule_fk0" FOREIGN KEY ("group_id") REFERENCES "group"("id");

ALTER TABLE
    "person"
ADD
    CONSTRAINT "person_fk0" FOREIGN KEY ("group_id") REFERENCES "group"("id");

ALTER TABLE
    "lesson"
ADD
    CONSTRAINT "lesson_fk0" FOREIGN KEY ("week_id") REFERENCES "week"("id");

ALTER TABLE
    "week"
ADD
    CONSTRAINT "week_fk0" FOREIGN KEY ("schedule_id") REFERENCES "schedule"("id");