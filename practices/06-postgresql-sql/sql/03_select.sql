-- Чтение (Read) после 02_seed.sql.
SELECT id, name, created_at FROM projects ORDER BY id;
SELECT id, project_id, title, status, priority FROM tasks ORDER BY id;

-- WHERE фильтрует, ORDER BY сортирует: высокий приоритет равен 3.
SELECT id, title, priority
FROM tasks
WHERE status = 'new' AND priority >= 2
ORDER BY priority DESC, id ASC;

-- JOIN: название связанного проекта.
SELECT t.id, t.title, p.name AS project_name, t.status, t.priority
FROM tasks AS t
JOIN projects AS p ON p.id = t.project_id
ORDER BY p.name, t.id;

-- LEFT JOIN + COUNT: показывает также проекты без задач.
SELECT p.name, COUNT(t.id) AS tasks_total
FROM projects AS p
LEFT JOIN tasks AS t ON t.project_id = p.id
GROUP BY p.id, p.name
ORDER BY p.name;
