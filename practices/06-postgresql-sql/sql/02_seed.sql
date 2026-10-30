-- Только в ЛИЧНОЙ УЧЕБНОЙ БД tasks_lab!
-- Для повторяемых проверок сбрасываем учебные записи и счётчики.
-- НЕ выполняйте на реальной базе: TRUNCATE удаляет все данные таблиц!
BEGIN;
TRUNCATE TABLE tasks, projects RESTART IDENTITY;

INSERT INTO projects (name) VALUES
    ('Учебный бэкенд'),
    ('Мобильный клиент');

-- Не предполагаем, что ID проекта фиксированы: связываем по UNIQUE name.
INSERT INTO tasks (project_id, title, status, priority)
SELECT p.id, v.title, v.status, v.priority
FROM (
    VALUES
        ('Учебный бэкенд', 'Изучить Go', 'new', 2),
        ('Учебный бэкенд', 'Написать API', 'new', 3),
        ('Учебный бэкенд', 'Временная задача', 'new', 1),
        ('Мобильный клиент', 'Настроить Flutter', 'done', 2)
) AS v(project_name, title, status, priority)
JOIN projects AS p ON p.name = v.project_name;

COMMIT;
