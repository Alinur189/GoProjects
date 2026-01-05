INSERT INTO groups (name, faculty_id)
VALUES ('Empty Group', 1);

INSERT INTO students (firstname, lastname, gender, birth_date, group_id)
VALUES ('NoGroup', 'Student', 'male', '2004-01-01', NULL);

SELECT
    s.id,
    s.firstname,
    s.lastname,
    g.name AS group_name
FROM students s
INNER JOIN groups g ON s.group_id = g.id;

SELECT
    s.id,
    s.firstname,
    s.lastname,
    g.name AS group_name
FROM students s
LEFT JOIN groups g ON s.group_id = g.id;

SELECT
    s.id,
    s.firstname,
    s.lastname,
    g.name AS group_name
FROM students s
RIGHT JOIN groups g ON s.group_id = g.id;

SELECT
    s.id,
    s.firstname,
    s.lastname,
    g.name AS group_name
FROM students s
FULL OUTER JOIN groups g ON s.group_id = g.id;
