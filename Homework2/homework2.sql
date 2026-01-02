DROP TABLE IF EXISTS students;
DROP TABLE IF EXISTS schedule;
DROP TABLE IF EXISTS groups;
DROP TABLE IF EXISTS faculties;
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
	-- INSERT FACULTIES
INSERT INTO faculties (name) VALUES
('SEDS'),
('SSH');


INSERT INTO groups (name, faculty_id) VALUES
('Group 1', 1),
('Grop 2', 1),
('Grop 3', 2),
('Group 4', 2);


INSERT INTO students (firstname, lastname, gender, birth_date, group_id) VALUES
('Diana', 'Askarova', 'female', '2004-05-12', 1),
('Maria', 'Petrova', 'female', '2003-11-03', 2),
('Aigerim', 'Kamalova', 'female', '2005-01-22', 3),
('Ualikhan', 'Yertayev', 'male', '2002-07-19', 1),
('Aki', 'Sanatov', 'male', '2003-03-10', 4);


INSERT INTO schedule (subject, day_of_week, lesson_time, group_id, faculty_id) VALUES
('Computer Science', 'Monday', '10:00',3, 2),
('Calculus', 'Monday', '14:00',1, 1),
('History', 'Tuesday', '11:30',4, 2);


ALTER TABLE students ADD COLUMN phone TEXT;
ALTER TABLE students DROP COLUMN phone;

SELECT *
FROM students
WHERE gender = 'female'
ORDER BY birth_date DESC;
SELECT id, firstname, lastname FROM students;
