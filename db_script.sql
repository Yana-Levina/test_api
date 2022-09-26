-- Database: person

-- DROP DATABASE IF EXISTS person;

CREATE DATABASE person
    WITH
    OWNER = postgres
    ENCODING = 'UTF8'
    LC_COLLATE = 'Russian_Russia.1251'
    LC_CTYPE = 'Russian_Russia.1251'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1;

-- Table: public.person

-- DROP TABLE IF EXISTS public.person;

CREATE TABLE IF NOT EXISTS public.person
(
    id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 1000 CACHE 1 ),
    email character varying(100) COLLATE pg_catalog."default" NOT NULL,
    phone character varying(100) COLLATE pg_catalog."default" NOT NULL,
    first_name character varying(100) COLLATE pg_catalog."default" NOT NULL,
    last_name character varying(100) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT class_pkey PRIMARY KEY (id)
    )

    TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.person
    OWNER to postgres;

INSERT INTO public.person(email, phone, first_name, last_name)VALUES ('mary@gmail.com', '11111', 'Mary', 'Jane');
INSERT INTO public.person(email, phone, first_name, last_name)VALUES ('jack@gmail.com', '22222', 'Jack', 'London');
INSERT INTO public.person(email, phone, first_name, last_name)VALUES ('tom@gmail.com', '33333', 'Tom', 'Some');
INSERT INTO public.person(email, phone, first_name, last_name)VALUES ('ann@gmail.com', '44444', 'Ann', 'Noname');