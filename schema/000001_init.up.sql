CREATE TABLE users
(
    id SERIAL PRIMARY KEY,
    name          varchar(255) ,
    username      varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE courses (
    id SERIAL PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    description TEXT,
    difficulty_level INT NOT NULL,
    field_of_activity VARCHAR(50) NOT NULL,
    duration_days INT NOT NULL,
    lang VARCHAR(50) NOT NULL,
    rating NUMERIC(3,2) DEFAULT 0.00,
    author VARCHAR(100) NOT NULL
);

CREATE TABLE user_courses (
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    course_id INT REFERENCES courses(id) ON DELETE CASCADE,
    PRIMARY KEY (user_id, course_id)
);

CREATE TABLE topics (
    id SERIAL PRIMARY KEY,
    course_id INT REFERENCES courses(id) ON DELETE CASCADE,
    title VARCHAR(100) NOT NULL,
    content TEXT,
    materials TEXT,
    assignments TEXT
);


CREATE TABLE form (
    id SERIAL PRIMARY KEY ,
    name text
);

CREATE TABLE questions (
    id SERIAL PRIMARY KEY ,
    question TEXT,
    form_id int references form(id)
);

CREATE TABLE answers (
    id SERIAL PRIMARY KEY ,
    user_id INT references users(id),
    form_id INT REFERENCES  form(id),
    difficulty_level INT NOT NULL,
    field_of_activity VARCHAR(50) NOT NULL,
    duration_days INT NOT NULL,
    lang VARCHAR(50) NOT NULL,
    rating NUMERIC(3,2) DEFAULT 0.00,
    author VARCHAR(100) NOT NULL
);

