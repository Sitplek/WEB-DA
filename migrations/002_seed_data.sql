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
    -- Обычные сотрудники подразделения Recruitment
    (11, 'Ekaterina', 'Semenova', 'HR Specialist', 'Employee', '777777777', 'ekaterina@example.com', 5, 1, 1, 1),
    (12, 'Nikolay', 'Grigoryev', 'HR Assistant', 'Employee', '888888888', 'nikolay@example.com', 5, 1, 1, 1),
    (23, 'Svetlana', 'Kuzmina', 'HR Coordinator', 'Employee', '222333444', 'svetlana@example.com', 5, 1, 1, 1),
    (24, 'Andrey', 'Volkov', 'HR Specialist', 'Employee', '333444555', 'andrey@example.com', 5, 1, 1, 1),
    (25, 'Natalia', 'Belova', 'HR Intern', 'Employee', '444555666', 'natalia@example.com', 5, 1, 1, 1),
    (37, 'Igor', 'Kozlov', 'HR Specialist', 'Employee', '282828282', 'igor.k@example.com', 5, 1, 1, 1),
    (38, 'Daria', 'Sokolova', 'HR Coordinator', 'Employee', '292929292', 'daria.s@example.com', 5, 1, 1, 1),
    (48, 'Alexandra', 'Petrova', 'HR Specialist', 'Employee', '393939393', 'alexandra.p@example.com', 5, 1, 1, 1),
    (49, 'Vadim', 'Smirnov', 'HR Assistant', 'Employee', '404040404', 'vadim.s@example.com', 5, 1, 1, 1),
    (56, 'Oleg', 'Klimov', 'HR Specialist', 'Employee', '474747474', 'oleg.k@example.com', 5, 1, 1, 1),
    (57, 'Yana', 'Shirokova', 'HR Assistant', 'Employee', '484848484', 'yana.s@example.com', 5, 1, 1, 1),
    (58, 'Ivan', 'Mikhailov', 'HR Specialist', 'Employee', '494949494', 'ivan.m@example.com', 5, 1, 1, 1),
    (73, 'Larisa', 'Vorontsova', 'HR Specialist', 'Employee', '646464646', 'larisa.v@example.com', 5, 1, 1, 1),
    (74, 'Boris', 'Savin', 'HR Assistant', 'Employee', '656565656', 'boris.s@example.com', 5, 1, 1, 1),
    (87, 'Irina', 'Sorokina', 'HR Coordinator', 'Employee', '787878787', 'irina.s@example.com', 5, 1, 1, 1),
    (88, 'Mikhail', 'Fedorov', 'HR Assistant', 'Employee', '797979797', 'mikhail.f@example.com', 5, 1, 1, 1),
    (89, 'Olga', 'Zharkova', 'HR Specialist', 'Employee', '808080808', 'olga.z@example.com', 5, 1, 1, 1),
    (90, 'Sergey', 'Laptev', 'HR Assistant', 'Employee', '818181818', 'sergey.l@example.com', 5, 1, 1, 1),
    (91, 'Natalia', 'Volkova', 'HR Specialist', 'Employee', '828282828', 'natalia.v@example.com', 5, 1, 1, 1),

    -- Обычные сотрудники подразделения Payroll
    (13, 'Diana', 'Pavlova', 'Payroll Clerk', 'Employee', '999999999', 'diana@example.com', 6, 2, 2, 2),
    (14, 'Valentina', 'Kozlova', 'Accounting Specialist', 'Employee', '101010101', 'valentina@example.com', 6, 2, 2, 2),
    (26, 'Konstantin', 'Orlov', 'Payroll Specialist', 'Employee', '555666777', 'konstantin@example.com', 6, 2, 2, 2),
    (27, 'Elizaveta', 'Yakovleva', 'Junior Accountant', 'Employee', '666777888', 'elizaveta@example.com', 6, 2, 2, 2),
    (28, 'Boris', 'Nikitin', 'Payroll Analyst', 'Employee', '777888999', 'boris@example.com', 6, 2, 2, 2),
    (39, 'Semyon', 'Vasiliev', 'Payroll Clerk', 'Employee', '303030303', 'semyon.v@example.com', 6, 2, 2, 2),
    (40, 'Alena', 'Morozova', 'Accounting Clerk', 'Employee', '313131313', 'alena.m@example.com', 6, 2, 2, 2),
    (50, 'Nina', 'Alexeeva', 'Payroll Analyst', 'Employee', '414141414', 'nina.a@example.com', 6, 2, 2, 2),
    (51, 'Evgeny', 'Sorokin', 'Junior Accountant', 'Employee', '424242424', 'evgeny.s@example.com', 6, 2, 2, 2),
    (59, 'Irina', 'Belkina', 'Payroll Clerk', 'Employee', '505050505', 'irina.b@example.com', 6, 2, 2, 2),
    (60, 'Alexey', 'Frolov', 'Accounting Specialist', 'Employee', '515151515', 'alexey.f@example.com', 6, 2, 2, 2),
    (61, 'Marina', 'Egorova', 'Junior Accountant', 'Employee', '525252525', 'marina.e@example.com', 6, 2, 2, 2),
    (62, 'Pavel', 'Petrov', 'Payroll Analyst', 'Employee', '535353535', 'pavel.p@example.com', 6, 2, 2, 2),
    (75, 'Veronika', 'Kuznetsova', 'Accounting Clerk', 'Employee', '666666666', 'veronika.k@example.com', 6, 2, 2, 2),
    (76, 'Igor', 'Vorobyov', 'Payroll Clerk', 'Employee', '676767676', 'igor.v@example.com', 6, 2, 2, 2),
    (77, 'Anna', 'Sokolova', 'Junior Accountant', 'Employee', '686868686', 'anna.s@example.com', 6, 2, 2, 2),
    (78, 'Viktoria', 'Belova', 'Payroll Analyst', 'Employee', '696969696', 'viktoria.b@example.com', 6, 2, 2, 2),
    (92, 'Alexandr', 'Rogov', 'Accounting Specialist', 'Employee', '838383838', 'alexandr.r@example.com', 6, 2, 2, 2),
    (93, 'Tatiana', 'Vlasova', 'Payroll Clerk', 'Employee', '848484848', 'tatiana.v@example.com', 6, 2, 2, 2),
    (94, 'Ivan', 'Sidorov', 'Junior Accountant', 'Employee', '858585858', 'ivan.s@example.com', 6, 2, 2, 2),
    (95, 'Elena', 'Makarova', 'Payroll Analyst', 'Employee', '868686868', 'elena.m@example.com', 6, 2, 2, 2),
    (96, 'Oksana', 'Smirnova', 'Accounting Clerk', 'Employee', '878787878', 'oksana.s@example.com', 6, 2, 2, 2),

    -- Обычные сотрудники подразделения Software Development
    (15, 'Pavel', 'Morozov', 'Frontend Developer', 'Employee', '121212121', 'pavel@example.com', 7, 3, 3, 3),
    (16, 'Anton', 'Smirnov', 'Backend Developer', 'Employee', '131313131', 'anton@example.com', 7, 3, 3, 3),
    (17, 'Yulia', 'Belyaeva', 'QA Engineer', 'Employee', '141414141', 'yulia@example.com', 7, 3, 3, 3),
    (18, 'Igor', 'Popov', 'DevOps Engineer', 'Employee', '151515151', 'igor@example.com', 7, 3, 3, 3),
    (19, 'Viktor', 'Voronov', 'Software Developer', 'Employee', '161616161', 'viktor@example.com', 7, 3, 3, 3),
    (20, 'Artem', 'Sokolov', 'Junior Developer', 'Employee', '171717171', 'artem@example.com', 7, 3, 3, 3),
    (21, 'Olga', 'Mikhailova', 'UI/UX Designer', 'Employee', '181818181', 'olga@example.com', 7, 3, 3, 3),
    (22, 'Roman', 'Lebedev', 'System Administrator', 'Employee', '191919191', 'roman@example.com', 7, 3, 3, 3),
    (29, 'Kirill', 'Semenov', 'Data Scientist', 'Employee', '202020202', 'kirill@example.com', 7, 3, 3, 3),
    (30, 'Marina', 'Frolova', 'Backend Developer', 'Employee', '212121212', 'marina@example.com', 7, 3, 3, 3),
    (31, 'Dmitry', 'Zaitsev', 'QA Engineer', 'Employee', '222222222', 'dmitry@example.com', 7, 3, 3, 3),
    (32, 'Vera', 'Maksimova', 'Junior Developer', 'Employee', '232323232', 'vera@example.com', 7, 3, 3, 3),
    (33, 'Oksana', 'Borodina', 'DevOps Engineer', 'Employee', '242424242', 'oksana@example.com', 7, 3, 3, 3),
    (34, 'Timur', 'Zubkov', 'Software Developer', 'Employee', '252525252', 'timur@example.com', 7, 3, 3, 3)
ON CONFLICT (id) DO NOTHING;
