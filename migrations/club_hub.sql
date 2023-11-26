-- database script

CREATE TABLE IF NOT EXISTS countries (
    ID SERIAL PRIMARY KEY,
    name VARCHAR(60) NOT NULL
);

CREATE TABLE IF NOT EXISTS locations (
    ID SERIAL  PRIMARY KEY,
    city VARCHAR(60) NOT NULL,
    address VARCHAR(100),
    zip_code VARCHAR(10),
    country_ID INT NOT NULL,
    FOREIGN KEY (country_ID) REFERENCES country(ID)
);

CREATE TABLE IF NOT EXISTS companies (
    ID SERIAL PRIMARY KEY,
    name VARCHAR(60) NOT NULL,
    tax_number VARCHAR(20) NOT NULL,
    location_ID INT NOT NULL,
    FOREIGN KEY (location_ID) REFERENCES location(ID)
);

CREATE TABLE IF NOT EXISTS owners (
    ID SERIAL  PRIMARY KEY,
    first_name VARCHAR(30) NOT NULL,
    last_name VARCHAR(30) NOT NULL,
    email VARCHAR(255) NOT NULL,
    phone VARCHAR(20),
    location_ID INT NOT NULL,
    company_ID INT NOT NULL,
    FOREIGN KEY (location_ID) REFERENCES location(ID),
    FOREIGN KEY (company_ID) REFERENCES company(ID)
);

CREATE TABLE IF NOT EXISTS franchises (
    ID SERIAL PRIMARY KEY,
    name VARCHAR(60) NOT NULL,
    url VARCHAR(150),
    location_ID INT NOT NULL,
    company_ID INT NOT NULL,
    FOREIGN KEY (location_ID) REFERENCES location(ID),
    FOREIGN KEY (company_ID) REFERENCES company(ID)
);

