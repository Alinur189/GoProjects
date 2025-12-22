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


INSERT INTO schedule (subject, day_of_week, lesson_time, faculty_id) VALUES
('Computer Science', 'Monday', '10:00', 2),
('Calculus', 'Monday', '14:00', 1),
('History', 'Tuesday', '11:30', 2);


ALTER TABLE students ADD COLUMN phone TEXT;
ALTER TABLE students DROP COLUMN phone;

SELECT *
FROM students
WHERE gender = 'female'
ORDER BY birth_date DESC;
