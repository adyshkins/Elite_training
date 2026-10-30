-- Запуск после: 01_schema -> 02_seed -> 03_select -> 04_mutations.
-- Нарушения вызывают ошибку; использовать psql -v ON_ERROR_STOP=1.
DO $$
DECLARE
    project_count INTEGER;
    task_count INTEGER;
BEGIN
    SELECT COUNT(*) INTO project_count FROM projects;
    SELECT COUNT(*) INTO task_count FROM tasks;
    IF project_count <> 2 OR task_count <> 4 THEN
        RAISE EXCEPTION 'Ожидалось 2 проекта и 4 задачи, получено % и %', project_count, task_count;
    END IF;
    IF NOT EXISTS (
        SELECT 1 FROM tasks WHERE title = 'Изучить Go' AND status = 'done'
    ) THEN
        RAISE EXCEPTION 'UPDATE статуса не выполнен';
    END IF;
    IF NOT EXISTS (
        SELECT 1 FROM tasks t JOIN projects p ON p.id = t.project_id
        WHERE t.title = 'Протестировать API' AND p.name = 'Мобильный клиент' AND t.status = 'new'
    ) THEN
        RAISE EXCEPTION 'INSERT задачи или JOIN не работает';
    END IF;
    IF EXISTS (SELECT 1 FROM tasks WHERE title = 'Временная задача') THEN
        RAISE EXCEPTION 'DELETE не выполнен';
    END IF;
    IF EXISTS (
        SELECT 1 FROM tasks t LEFT JOIN projects p ON p.id = t.project_id WHERE p.id IS NULL
    ) THEN
        RAISE EXCEPTION 'Обнаружена задача без проекта';
    END IF;
    RAISE NOTICE 'OK: 2 проекта, 4 задачи, JOIN, INSERT, UPDATE, DELETE';
END $$;
