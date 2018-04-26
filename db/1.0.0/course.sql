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
FOREIGN KEY (category_id) REFERENCES categories(id)
