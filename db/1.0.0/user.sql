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
)
