CREATE TABLE countries (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) UNIQUE
);


INSERT INTO countries (name) VALUES
    ('Argentina'),
    ('Australia'),
    ('Brasil'),
    ('Canadá'),
    ('China'),
    ('España'),
    ('Estados Unidos'),
    ('Francia'),
    ('India'),
    ('Italia'),
    ('Japón'),
    ('México'),
    ('Reino Unido'),
    ('Rusia'),
    ('Sudáfrica'),
   	('Colombia');


CREATE TABLE locations (
    id SERIAL PRIMARY KEY,
    city VARCHAR(100),
    country_id INT REFERENCES countries(id),
    address VARCHAR(100),
    zip_code VARCHAR(20)
);

CREATE TABLE contacts (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    phone VARCHAR(20) UNIQUE NOT NULL,
    location_id INTEGER REFERENCES locations(id)
);

CREATE TABLE owners (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    contact_id INTEGER REFERENCES contacts(id)
);

CREATE TABLE information (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    tax_number VARCHAR(20) UNIQUE NOT NULL,
    location_id INTEGER REFERENCES locations(id)
);

CREATE TABLE companies (
    id  SERIAL PRIMARY KEY,
    owner_id INTEGER REFERENCES owners(id),
    information_id INTEGER REFERENCES information(id)
);

CREATE TABLE franchises (
    id SERIAL PRIMARY KEY,
    company_id INT REFERENCES companies(id),
    name VARCHAR (100) UNIQUE,
    url VARCHAR (255) UNIQUE,
    location_id INT REFERENCES locations(id)
);