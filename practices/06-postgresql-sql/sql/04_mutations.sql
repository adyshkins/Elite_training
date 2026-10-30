-- CREATE / UPDATE / DELETE в одной транзакции.
-- Выполняйте один раз после 02_seed.sql; для повторения сбросьте данные через 02_seed.sql.
BEGIN;

INSERT INTO tasks (project_id, title, priority)
SELECT id, 'Протестировать API', 1
FROM projects WHERE name = 'Мобильный клиент'
RETURNING id, project_id, title, status, priority;

UPDATE tasks
SET status = 'done'
WHERE title = 'Изучить Go'
RETURNING id, title, status;

UPDATE tasks
SET priority = 2
WHERE title = 'Написать API'
RETURNING id, title, priority;

DELETE FROM tasks
WHERE title = 'Временная задача'
RETURNING id, title;

COMMIT;

-- После операций ожидаются 2 проекта и 4 задачи.
SELECT t.id, t.title, p.name AS project_name, t.status, t.priority
FROM tasks AS t
JOIN projects AS p ON p.id = t.project_id
ORDER BY t.id;
