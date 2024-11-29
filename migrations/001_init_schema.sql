-- Удаление таблиц, если они существуют
DROP TABLE IF EXISTS employees CASCADE;
DROP TABLE IF EXISTS departments CASCADE;
DROP TABLE IF EXISTS offices CASCADE;

-- Таблица офисов
CREATE TABLE offices (
    id SERIAL PRIMARY KEY,                        -- Уникальный идентификатор офиса
    name VARCHAR(255) NOT NULL UNIQUE,           -- Название офиса
    block VARCHAR(50) NOT NULL,                  -- Тип блока (Корпоративный или Розничный)
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Дата создания записи
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP -- Дата обновления записи
);

-- Таблица подразделений
CREATE TABLE departments (
    id SERIAL PRIMARY KEY,                        -- Уникальный идентификатор подразделения
    name VARCHAR(255) NOT NULL,                  -- Название подразделения
    office_id INT NOT NULL REFERENCES offices(id) ON DELETE CASCADE, -- Ссылка на офис
    parent_id INT REFERENCES departments(id) ON DELETE SET NULL, -- Родительское подразделение
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Дата создания записи
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP -- Дата обновления записи
);

-- Таблица сотрудников
CREATE TABLE employees (
    id SERIAL PRIMARY KEY,                         -- Уникальный идентификатор сотрудника
    name VARCHAR(100) NOT NULL,                   -- Имя сотрудника
    email VARCHAR(100) NOT NULL UNIQUE,           -- Почта сотрудника
    phone VARCHAR(15) NOT NULL,                   -- Телефон сотрудника
    address VARCHAR(255),                         -- Адрес сотрудника
    role VARCHAR(50) NOT NULL,                    -- Роль (руководитель или сотрудник)
    job_position VARCHAR(100) NOT NULL,           -- Должность сотрудника
    department_id INT REFERENCES departments(id) ON DELETE CASCADE, -- Ссылка на подразделение
    manager_id INT REFERENCES employees(id) ON DELETE SET NULL,     -- Ссылка на руководителя
    supervisor_id INT,                            -- Поле для идентификатора супервизора
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Дата создания записи
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP -- Дата обновления записи
);

-- Функция для автоматического обновления поля updated_at
CREATE OR REPLACE FUNCTION update_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Триггеры для обновления поля updated_at
CREATE TRIGGER set_updated_at_offices
BEFORE UPDATE ON offices
FOR EACH ROW
EXECUTE FUNCTION update_updated_at();

CREATE TRIGGER set_updated_at_departments
BEFORE UPDATE ON departments
FOR EACH ROW
EXECUTE FUNCTION update_updated_at();

CREATE TRIGGER set_updated_at_employees
BEFORE UPDATE ON employees
FOR EACH ROW
EXECUTE FUNCTION update_updated_at();
