-- Создание таблиц
CREATE TABLE departments (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    head_id INT DEFAULT NULL
);

CREATE TABLE divisions (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    head_id INT DEFAULT NULL,
    department_id INT REFERENCES departments(id) ON DELETE CASCADE
);

CREATE TABLE units (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    head_id INT DEFAULT NULL,
    division_id INT REFERENCES divisions(id) ON DELETE CASCADE
);

CREATE TABLE employees (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    position VARCHAR(100) NOT NULL,
    photo_path VARCHAR(255),
    role VARCHAR(50),
    phone VARCHAR(20),
    email VARCHAR(100),
    manager_id INT REFERENCES employees(id) ON DELETE SET NULL,
    department_id INT REFERENCES departments(id) ON DELETE SET NULL,
    division_id INT REFERENCES divisions(id) ON DELETE SET NULL,
    unit_id INT REFERENCES units(id) ON DELETE SET NULL
);
