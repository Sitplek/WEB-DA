-- Удаление существующих данных для чистого заполнения
TRUNCATE employees RESTART IDENTITY CASCADE;
TRUNCATE departments RESTART IDENTITY CASCADE;

-- Добавление тестовых департаментов
INSERT INTO departments (name, parent_id, created_at, updated_at)
VALUES 
    ('Head Office', NULL, DEFAULT, DEFAULT), -- Главный офис
    ('IT Department', 1, DEFAULT, DEFAULT), -- IT Департамент подчинён Главному офису
    ('HR Department', 1, DEFAULT, DEFAULT), -- HR Департамент подчинён Главному офису
    ('Engineering', 2, DEFAULT, DEFAULT),  -- Подразделение IT
    ('Support', 2, DEFAULT, DEFAULT);       -- Подразделение IT

-- Добавление тестовых сотрудников
INSERT INTO employees (name, job_position, department_id, skills, region, role, manager_id, created_at, updated_at)
VALUES 
    -- Сотрудники Главного офиса
    ('Alice Johnson', 'CEO', 1, '{"leadership": 5, "strategy": 4}', 'North America', 'Executive', NULL, DEFAULT, DEFAULT),
    ('Bob Smith', 'Head of IT', 2, '{"management": 4, "development": 5}', 'Europe', 'Manager', 1, DEFAULT, DEFAULT),
    ('Carol Lee', 'Head of HR', 3, '{"communication": 5, "recruitment": 5}', 'Asia', 'Manager', 1, DEFAULT, DEFAULT),

    -- Сотрудники IT Департамента
    ('David Kim', 'Software Engineer', 4, '{"python": 5, "databases": 4}', 'North America', 'Employee', 2, DEFAULT, DEFAULT),
    ('Eve Torres', 'DevOps Engineer', 4, '{"linux": 5, "docker": 4}', 'Europe', 'Employee', 2, DEFAULT, DEFAULT),
    ('Frank Wright', 'IT Support Specialist', 5, '{"troubleshooting": 5, "networking": 4}', 'Asia', 'Employee', 2, DEFAULT, DEFAULT),

    -- Сотрудники HR Департамента
    ('Grace Hall', 'Recruiter', 3, '{"interviewing": 5, "hiring": 4}', 'Europe', 'Employee', 3, DEFAULT, DEFAULT),
    ('Hank Green', 'HR Specialist', 3, '{"policy_making": 4, "onboarding": 5}', 'North America', 'Employee', 3, DEFAULT, DEFAULT),

    -- Другие сотрудники
    ('Ivy Clarke', 'Junior Engineer', 4, '{"java": 3, "spring": 4}', 'Asia', 'Employee', 4, DEFAULT, DEFAULT),
    ('Jake Black', 'Support Technician', 5, '{"hardware": 4, "customer_service": 5}', 'Europe', 'Employee', 5, DEFAULT, DEFAULT);
