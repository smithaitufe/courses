CREATE DATABASE courses;


CREATE TABLE roles
(
  id character varying(45) NOT NULL,
  name character varying(250),
  created_at timestamp with time zone NOT NULL DEFAULT now(),
  updated_at timestamp with time zone NOT NULL DEFAULT now(),
  CONSTRAINT roles_pkey PRIMARY KEY (id)
);

CREATE TABLE categories
(
  id character varying(45) NOT NULL,
  name character varying(200),
  created_at timestamp with time zone NOT NULL DEFAULT now(),
  updated_at timestamp with time zone NOT NULL DEFAULT now(),
  CONSTRAINT categories_pkey PRIMARY KEY (id)
);

CREATE TABLE companies
(
  id character varying(45) NOT NULL,
  name character varying(200),
  created_at timestamp with time zone NOT NULL DEFAULT now(),
  updated_at timestamp with time zone NOT NULL DEFAULT now(),
  CONSTRAINT companies_pkey PRIMARY KEY (id)
);

CREATE TABLE users
(
  id character varying(45) NOT NULL,
  last_name character varying(50),
  first_name character varying(50),
  email character varying(250),
  country character varying(150),
  dialing_code character varying(10),
  phone_number character varying(20),
  password character varying(200),
  created_at timestamp with time zone NOT NULL DEFAULT now(),
  updated_at timestamp with time zone NOT NULL DEFAULT now(),
  CONSTRAINT users_pkey PRIMARY KEY (id)
);

CREATE TABLE user_roles
(
  user_id character varying(45),
  role_id character varying(45),
  created_at timestamp with time zone NOT NULL DEFAULT now(),
  updated_at timestamp with time zone NOT NULL DEFAULT now(),
  CONSTRAINT user_roles_pkey PRIMARY KEY (user_id, role_id)
);

create table courses
(
  id character varying(45) not null,
  title character varying(200),
  code character varying(10),
  hours integer default 0,
  amount float default 0.0,
  category_id character varying(45) not null,
  company_id character varying(45) not null,
  created_at timestamp with time zone not null default now(),
  updated_at timestamp with time zone not null default now(),
  constraint courses_pkey primary key (id)
);


ALTER TABLE courses
ADD CONSTRAINT course_company_company_id_fkey
FOREIGN KEY (company_id) REFERENCES companies(id);

ALTER TABLE courses
ADD CONSTRAINT course_category_category_id_fkey
FOREIGN KEY (category_id) REFERENCES categories(id);

CREATE TABLE enrollments
(
  id character varying(45) NOT NULL,
  user_id character varying(45),
  course_id character varying(45),
  created_at timestamp with time zone NOT NULL DEFAULT now(),
  updated_at timestamp with time zone NOT NULL DEFAULT now(),
  CONSTRAINT enrollments_pkey PRIMARY KEY (id)
);

ALTER TABLE enrollments
ADD CONSTRAINT enrollment_user_user_id_fkey
FOREIGN KEY (user_id) REFERENCES users(id);

ALTER TABLE enrollments
ADD CONSTRAINT enrollment_course_course_id_fkey
FOREIGN KEY (course_id) REFERENCES courses(id);
