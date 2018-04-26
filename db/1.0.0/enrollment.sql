CREATE TABLE enrollments
(
  id character varying(45) NOT NULL,
  user_id character varying(45),
  course_id character varying(45),
  created_at timestamp with time zone NOT NULL DEFAULT now(),
  updated_at timestamp with time zone NOT NULL DEFAULT now(),
  CONSTRAINT enrollments_pkey PRIMARY KEY (id)
)

ALTER TABLE enrollments
ADD CONSTRAINT enrollment_user_user_id_fkey
FOREIGN KEY (user_id) REFERENCES users(id);

ALTER TABLE enrollments
ADD CONSTRAINT enrollment_course_course_id_fkey
FOREIGN KEY (course_id) REFERENCES courses(id);
