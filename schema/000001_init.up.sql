CREATE TABLE users
(
    id            serial       not null unique,
    name          varchar(255) ,
    username      varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE courses (
    id SERIAL PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    description TEXT
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