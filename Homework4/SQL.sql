
DROP TABLE IF EXISTS attendance;
DROP TABLE IF EXISTS schedule;
DROP TABLE IF EXISTS students;
DROP TABLE IF EXISTS groups;
DROP TABLE IF EXISTS faculties;
DROP TABLE IF EXISTS subjects;


CREATE TABLE faculties (
    id SERIAL PRIMARY KEY,
    name TEXT
);

CREATE TABLE groups (
    id SERIAL PRIMARY KEY,
    name TEXT,
    faculty_id INT REFERENCES faculties(id)
);

CREATE TABLE students (
    id SERIAL PRIMARY KEY,
    firstname TEXT,
    lastname TEXT,
    gender TEXT,
    birth_date DATE,
    group_id INT REFERENCES groups(id)
);

CREATE TABLE schedule (
    id SERIAL PRIMARY KEY,
    subject TEXT,
    day_of_week TEXT,
    lesson_time TEXT,
    group_id INT,
    faculty_id INT
);


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
('Computer Science', 'Monday', '10:00', 3, 2),
('Calculus', 'Monday', '14:00', 1, 1),
('History', 'Tuesday', '11:30', 4, 2);


CREATE TABLE subjects (
    id SERIAL PRIMARY KEY,
    name TEXT UNIQUE NOT NULL
);

ALTER TABLE schedule
ADD COLUMN subject_id INT REFERENCES subjects(id);

INSERT INTO subjects (name)
SELECT DISTINCT subject
FROM schedule
ON CONFLICT (name) DO NOTHING;

UPDATE schedule s
SET subject_id = sub.id
FROM subjects sub
WHERE s.subject = sub.name;


CREATE TABLE attendance (
    id SERIAL PRIMARY KEY,
    subject_id INT NOT NULL REFERENCES subjects(id),
    student_id INT NOT NULL REFERENCES students(id),
    visit_day DATE NOT NULL,
    visited BOOLEAN NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    UNIQUE (subject_id, student_id, visit_day)
);


SELECT * FROM subjects ORDER BY id;


INSERT INTO attendance (subject_id, student_id, visit_day, visited) VALUES
(1, 1, '2026-01-07', true),
(1, 2, '2026-01-07', false),
(2, 3, '2026-01-08', true);


INSERT INTO attendance (subject_id, student_id, visit_day, visited) VALUES

(1, 1, '2026-01-14', true),


(2, 1, '2026-01-08', true),


(3, 4, '2026-01-09', true),
(3, 5, '2026-01-09', false);


SELECT id, firstname, lastname FROM students ORDER BY id;

SELECT
    a.id,
    s.firstname,
    s.lastname,
    sub.name AS subject,
    a.visit_day,
    a.visited
FROM attendance a
JOIN students s ON s.id = a.student_id
JOIN subjects sub ON sub.id = a.subject_id
ORDER BY a.visit_day, a.id;


SELECT * FROM subjects ORDER BY id; 
SELECT * FROM attendance ORDER BY id;
