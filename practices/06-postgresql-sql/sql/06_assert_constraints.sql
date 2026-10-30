-- Автоматические негативные проверки. Выполнить после 05_verify.sql.
-- Каждая ожидаемая SQL-ошибка перехватывается отдельным блоком.
-- Если ограничение НЕ сработало, скрипт завершается с ошибкой.
DO $$
BEGIN
    BEGIN
        INSERT INTO tasks (project_id, title, priority)
        VALUES (999999, 'Несуществующий проект', 2);
        RAISE EXCEPTION 'ОШИБКА ТЕСТА: FK разрешил несуществующий project_id';
    EXCEPTION WHEN foreign_key_violation THEN
        RAISE NOTICE 'PASS: foreign key';
    END;

    BEGIN
        INSERT INTO tasks (project_id, title, priority)
        SELECT id, 'Недопустимый приоритет', 9 FROM projects
        WHERE name = 'Учебный бэкенд';
        RAISE EXCEPTION 'ОШИБКА ТЕСТА: CHECK разрешил приоритет 9';
    EXCEPTION WHEN check_violation THEN
        RAISE NOTICE 'PASS: priority check';
    END;

    BEGIN
        INSERT INTO projects (name) VALUES ('   ');
        RAISE EXCEPTION 'ОШИБКА ТЕСТА: пустое имя разрешено';
    EXCEPTION WHEN check_violation THEN
        RAISE NOTICE 'PASS: project name check';
    END;

    BEGIN
        DELETE FROM projects WHERE name = 'Учебный бэкенд';
        RAISE EXCEPTION 'ОШИБКА ТЕСТА: удаление связанного проекта разрешено';
    EXCEPTION WHEN restrict_violation THEN
        RAISE NOTICE 'PASS: ON DELETE RESTRICT';
    END;

    BEGIN
        INSERT INTO projects(name) VALUES ('Учебный бэкенд');
        RAISE EXCEPTION 'ОШИБКА ТЕСТА: дубликат имени разрешён';
    EXCEPTION WHEN unique_violation THEN
        RAISE NOTICE 'PASS: unique project name';
    END;

    IF (SELECT COUNT(*) FROM projects) <> 2 OR
       (SELECT COUNT(*) FROM tasks) <> 4 THEN
        RAISE EXCEPTION 'После тестов количество строк изменилось';
    END IF;
    RAISE NOTICE 'OK: 5 ограничений действуют, данные остались неизменными';
END $$;
