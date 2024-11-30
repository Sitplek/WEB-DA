-- Добавление данных в таблицы

-- Добавляем департаменты
INSERT INTO departments (id, name)
VALUES 
    (1, 'Human Resources'), 
    (2, 'Finance'), 
    (3, 'Engineering')
ON CONFLICT (id) DO NOTHING;

-- Добавляем подразделения
INSERT INTO divisions (id, name, department_id)
VALUES
    (1, 'Recruitment', 1),
    (2, 'Payroll', 2),
    (3, 'Software Development', 3)
ON CONFLICT (id) DO NOTHING;

-- Добавляем отделы
INSERT INTO units (id, name, division_id)
VALUES
    (1, 'Candidate Screening', 1),
    (2, 'Salary Processing', 2),
    (3, 'Backend Development', 3)
ON CONFLICT (id) DO NOTHING;

-- Добавляем сотрудников
INSERT INTO employees (id, first_name, last_name, position, role, phone, email, manager_id, department_id, division_id, unit_id)
VALUES
    -- Генеральный директор
    (1, 'Dmitry', 'Ivanov', 'CEO', 'General Director', '123456789', 'ceo@example.com', NULL, NULL, NULL, NULL),
    -- Руководители департаментов
    (2, 'Maria', 'Petrova', 'HR Director', 'Head of HR', '987654321', 'hr@example.com', 1, 1, NULL, NULL),
    (3, 'Alexey', 'Sidorov', 'Finance Director', 'Head of Finance', '555555555', 'finance@example.com', 1, 2, NULL, NULL),
    (4, 'Ivan', 'Kuznetsov', 'Engineering Director', 'Head of Engineering', '999999999', 'engineering@example.com', 1, 3, NULL, NULL),
    -- Руководители подразделений
    (5, 'Anna', 'Smirnova', 'Recruitment Manager', 'Head of Recruitment', '333333333', 'recruitment@example.com', 2, 1, 1, NULL),
    (6, 'Sergey', 'Vasilyev', 'Payroll Manager', 'Head of Payroll', '444444444', 'payroll@example.com', 3, 2, 2, NULL),
    (7, 'Elena', 'Fedorova', 'Software Manager', 'Head of Software', '555666777', 'software@example.com', 4, 3, 3, NULL),
    -- Сотрудники
    (8, 'Oleg', 'Novikov', 'Recruiter', 'Employee', '111111111', 'oleg@example.com', 5, 1, 1, 1),
    (9, 'Irina', 'Zhukova', 'Payroll Specialist', 'Employee', '222222222', 'irina@example.com', 6, 2, 2, 2),
    (10, 'Maxim', 'Karpov', 'Backend Developer', 'Employee', '888888888', 'maxim@example.com', 7, 3, 3, 3)
ON CONFLICT (id) DO NOTHING;
