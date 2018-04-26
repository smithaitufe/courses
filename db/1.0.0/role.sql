CREATE TABLE roles
(
  id character varying(45) NOT NULL,
  name character varying(250),
  created_at timestamp with time zone NOT NULL DEFAULT now(),
  updated_at timestamp with time zone NOT NULL DEFAULT now(),
  CONSTRAINT roles_pkey PRIMARY KEY (id)
)
