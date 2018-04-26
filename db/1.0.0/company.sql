CREATE TABLE companies
(
  id character varying(45) NOT NULL,
  name character varying(200),
  created_at timestamp with time zone NOT NULL DEFAULT now(),
  updated_at timestamp with time zone NOT NULL DEFAULT now(),
  CONSTRAINT companies_pkey PRIMARY KEY (id)
)
