# Проверка ограничений PostgreSQL: ожидаемые ошибки

[← ПР № 6](README.md) · [Схема](sql/01_schema.sql)

Выполняйте примеры **по одному** в интерактивном `psql` после создания схемы и загрузки данных. Каждый запрос должен завершиться ошибкой — это подтверждает корректное ограничение. Не запускайте в одной транзакции `BEGIN`: после ошибки она будет отменена.

**1. Несуществующий проект → SQLSTATE 23503 (foreign key):**

~~~sql
INSERT INTO tasks (project_id, title, priority)
VALUES (999999, 'Несуществующий проект', 2);
~~~

**2. Недопустимый приоритет → SQLSTATE 23514 (CHECK):**

~~~sql
INSERT INTO tasks (project_id, title, priority)
SELECT id, 'Слишком высокий приоритет', 9
FROM projects WHERE name = 'Учебный бэкенд';
~~~

**3. Пустое название проекта → SQLSTATE 23514 (CHECK):**

~~~sql
INSERT INTO projects (name) VALUES ('   ');
~~~

**4. Удаление проекта, на который ссылаются задачи → SQLSTATE 23001 (RESTRICT):**

~~~sql
DELETE FROM projects WHERE name = 'Учебный бэкенд';
~~~

`ON DELETE RESTRICT` защищает связанные задачи от удаления проекта. Не подменяйте `RESTRICT` на `CASCADE`.

**5. Дублирование имени проекта → SQLSTATE 23505 (UNIQUE):**

~~~sql
INSERT INTO projects (name) VALUES ('Учебный бэкенд');
~~~

Эти же случаи автоматически проверяет [`sql/06_assert_constraints.sql`](sql/06_assert_constraints.sql) — запускайте его **после** `05_verify.sql`. После ручных тестов проверьте: `SELECT COUNT(*) FROM projects;` и `SELECT COUNT(*) FROM tasks;` — количество не изменилось. Ожидаемо ошибочные команды **не включены** в обычный сценарий с `ON_ERROR_STOP=1`, иначе он завершится на первой ошибке.
