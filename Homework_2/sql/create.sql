CREATE TABLE faculties (
    id SERIAL PRIMARY KEY,
    name TEXT);
CREATE TABLE groups (
    id SERIAL PRIMARY KEY,
    name TEXT,
    faculty_id INT REFERENCES faculties(id));
CREATE TABLE students (
    id SERIAL PRIMARY KEY,
    firstname TEXT,
    lastname TEXT,
    gender TEXT,
    birth_date DATE,
    group_id INT REFERENCES groups(id));
CREATE TABLE schedule (
    id SERIAL PRIMARY KEY,
    subject TEXT,
    day_of_week TEXT,
    lesson_time TEXT,
    group_id INT,
    faculty_id INT);
