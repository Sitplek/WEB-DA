-- Удаление таблиц, если они существуют (для чистого запуска)
DROP TABLE IF EXISTS employees CASCADE;
DROP TABLE IF EXISTS departments CASCADE;

-- Создание таблицы департаментов
CREATE TABLE departments (
    id SERIAL PRIMARY KEY,                        -- Уникальный идентификатор департамента
    name VARCHAR(100) NOT NULL UNIQUE,           -- Название департамента
    parent_id INT REFERENCES departments(id) ON DELETE SET NULL, -- Родительский департамент
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Дата создания записи
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP -- Дата обновления записи
);

-- Создание таблицы сотрудников
CREATE TABLE employees (
    id SERIAL PRIMARY KEY,                         -- Уникальный идентификатор сотрудника
    name VARCHAR(100) NOT NULL,                   -- Имя сотрудника
    job_position VARCHAR(50) NOT NULL,            -- Должность сотрудника
    department_id INT NOT NULL REFERENCES departments(id) ON DELETE CASCADE, -- Ссылка на департамент
    skills JSONB DEFAULT '{}'::jsonb,             -- Навыки сотрудника в формате JSON
    region VARCHAR(50) NOT NULL,                  -- Регион сотрудника
    role VARCHAR(50) NOT NULL,                    -- Роль сотрудника
    manager_id INT REFERENCES employees(id) ON DELETE SET NULL, -- Ссылка на руководителя
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Дата создания записи
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP -- Дата обновления записи
);

-- Создание индексов
CREATE INDEX idx_employees_job_position ON employees(job_position); -- Индекс на должность
CREATE INDEX idx_employees_department_id ON employees(department_id); -- Индекс на департамент
CREATE INDEX idx_employees_region ON employees(region); -- Индекс на регион
CREATE INDEX idx_employees_fulltext ON employees USING GIN ( -- Полнотекстовый индекс
    to_tsvector('english', name || ' ' || coalesce(skills::text, ''))
);

-- Индекс на имя департамента
CREATE INDEX idx_departments_name ON departments(name);

-- Функция для автоматического обновления поля updated_at
CREATE OR REPLACE FUNCTION update_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Триггеры для обновления поля updated_at
CREATE TRIGGER set_updated_at_departments
BEFORE UPDATE ON departments
FOR EACH ROW
EXECUTE FUNCTION update_updated_at();

CREATE TRIGGER set_updated_at_employees
BEFORE UPDATE ON employees
FOR EACH ROW
EXECUTE FUNCTION update_updated_at();

-- Функция для рекурсивной выборки иерархии департаментов
CREATE OR REPLACE FUNCTION department_hierarchy(root_id INT)
RETURNS TABLE(id INT, name VARCHAR, parent_id INT, level INT) AS $$
BEGIN
    RETURN QUERY
    WITH RECURSIVE hierarchy AS (
        SELECT id, name, parent_id, 1 AS level
        FROM departments
        WHERE id = root_id
        UNION ALL
        SELECT d.id, d.name, d.parent_id, h.level + 1
        FROM departments d
        INNER JOIN hierarchy h ON d.parent_id = h.id
    )
    SELECT * FROM hierarchy;
END;
$$ LANGUAGE plpgsql;

-- Функция для рекурсивной выборки иерархии сотрудников
CREATE OR REPLACE FUNCTION employee_hierarchy(root_id INT)
RETURNS TABLE(id INT, name VARCHAR, job_position VARCHAR, manager_id INT, level INT) AS $$
BEGIN
    RETURN QUERY
    WITH RECURSIVE hierarchy AS (
        SELECT id, name, job_position, manager_id, 1 AS level
        FROM employees
        WHERE id = root_id
        UNION ALL
        SELECT e.id, e.name, e.job_position, e.manager_id, h.level + 1
        FROM employees e
        INNER JOIN hierarchy h ON e.manager_id = h.id
    )
    SELECT * FROM hierarchy;
END;
$$ LANGUAGE plpgsql;
