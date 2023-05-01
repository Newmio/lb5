CREATE TABLE users(
    id serial primary key,
    login varchar(250),
    password varchar(250)
);

CREATE TABLE students(
    id serial primary key,
    name varchar(50) not null,
    last_name varchar(50) not null,
    age int not null
);

CREATE TABLE users_students(
    id serial primary key,
    id_user int references users(id) on delete cascade not null,
    id_student int references students(id) on delete cascade not null
);

CREATE TABLE courses(
    id serial primary key,
    name varchar(50) not null,
    total_lessons int
);

CREATE TABLE students_courses(
    id serial primary key,
    id_student int references students(id) on delete cascade not null,
    id_course int references courses(id) on delete cascade not null
);

CREATE TABLE scores(
    id serial primary key,
    score_value int not null,
    id_course int references courses(id) on delete cascade not null,
    id_student int references students(id) on delete cascade not null
);