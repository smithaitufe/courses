CREATE TABLE user_roles
(
  user_id character varying(45),
  course_id character varying(45),
  created_at timestamp with time zone NOT NULL DEFAULT now(),
  updated_at timestamp with time zone NOT NULL DEFAULT now(),
  CONSTRAINT user_roles_pkey PRIMARY KEY (user_id, role_id)
)
