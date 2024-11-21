-- Создание таблицы сотрудников
CREATE TABLE employees (
    id SERIAL PRIMARY KEY,                         -- Уникальный идентификатор сотрудника
    name VARCHAR(100) NOT NULL,                   -- Имя сотрудника
    position VARCHAR(50) NOT NULL,                -- Должность сотрудника
    department INT NOT NULL REFERENCES departments(id) ON DELETE CASCADE, -- Ссылка на департамент
    skills JSONB,                                 -- Навыки сотрудника в формате JSON
    region VARCHAR(50) NOT NULL,                  -- Регион сотрудника
    role VARCHAR(50) NOT NULL,                    -- Роль сотрудника
    manager_id INT REFERENCES employees(id) ON DELETE SET NULL, -- Ссылка на руководителя
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Дата создания записи
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP -- Дата обновления записи
);

-- Создание таблицы департаментов
CREATE TABLE departments (
    id SERIAL PRIMARY KEY,                        -- Уникальный идентификатор департамента
    name VARCHAR(100) NOT NULL UNIQUE,           -- Название департамента
    parent_id INT REFERENCES departments(id) ON DELETE SET NULL, -- Родительский департамент
    head_id INT REFERENCES employees(id) ON DELETE SET NULL, -- Руководитель департамента
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Дата создания записи
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP -- Дата обновления записи
);

-- Создание индексов для оптимизации
-- Индекс на поле position в таблице employees
CREATE INDEX idx_employees_position ON employees(position);

-- Индекс на поле department в таблице employees
CREATE INDEX idx_employees_department ON employees(department);

-- Индекс на поле region в таблице employees
CREATE INDEX idx_employees_region ON employees(region);

-- Индекс для полнотекстового поиска по полям name и skills в таблице employees
CREATE INDEX idx_employees_fulltext ON employees USING GIN (
    to_tsvector('english', name || ' ' || coalesce(skills::text, ''))
);

-- Индекс на поле name в таблице departments
CREATE INDEX idx_departments_name ON departments(name);

-- Настройка обновления полей created_at и updated_at
CREATE OR REPLACE FUNCTION update_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_updated_at_employees
BEFORE UPDATE ON employees
FOR EACH ROW
EXECUTE FUNCTION update_updated_at();

CREATE TRIGGER set_updated_at_departments
BEFORE UPDATE ON departments
FOR EACH ROW
EXECUTE FUNCTION update_updated_at();

-- Рекурсивный CTE для выборки иерархии департаментов
-- Пример: SELECT * FROM department_hierarchy(1);
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

-- Рекурсивный CTE для выборки иерархии сотрудников
-- Пример: SELECT * FROM employee_hierarchy(1);
CREATE OR REPLACE FUNCTION employee_hierarchy(root_id INT)
RETURNS TABLE(id INT, name VARCHAR, position VARCHAR, manager_id INT, level INT) AS $$
BEGIN
    RETURN QUERY
    WITH RECURSIVE hierarchy AS (
        SELECT id, name, position, manager_id, 1 AS level
        FROM employees
        WHERE id = root_id
        UNION ALL
        SELECT e.id, e.name, e.position, e.manager_id, h.level + 1
        FROM employees e
        INNER JOIN hierarchy h ON e.manager_id = h.id
    )
    SELECT * FROM hierarchy;
END;
$$ LANGUAGE plpgsql;
