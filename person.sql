-- DROP DATABASE IF EXISTS "person";

CREATE DATABASE "person"
    WITH
    OWNER = postgres
    ENCODING = 'UTF8'
    LC_COLLATE = 'Russian_Russia.1251'
    LC_CTYPE = 'Russian_Russia.1251'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1;

DROP TABLE IF EXISTS public.person;
CREATE TABLE IF NOT EXISTS public.person
(
    id integer NOT NULL DEFAULT nextval('person_id_seq'::regclass),
    email character varying(100) NOT NULL,
    phone character varying(100) NOT NULL,
    first_name character varying(100) NOT NULL,
    last_name character varying(100) NOT NULL,
    CONSTRAINT class_pkey PRIMARY KEY (id)
    )

-- INSERT INTO `person` (email, phone, first_name, last_name) VALUES ('email0@pochta.ru','80123456789','Mary0','Jane');
-- INSERT INTO `person` (email, phone, first_name, last_name) VALUES ('email1@pochta.ru','80123456789','Mary1','Jane');
-- INSERT INTO `person` (email, phone, first_name, last_name) VALUES ('email2@pochta.ru','80123456789','Mary2','Jane');
-- INSERT INTO `person` (email, phone, first_name, last_name) VALUES ('email3@pochta.ru','80123456789','Mary3','Jane');