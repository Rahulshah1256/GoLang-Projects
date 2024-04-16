CREATE TABLE public.tb_login (
	username varchar NOT NULL,
	"password" varchar NOT NULL,
	CONSTRAINT login_pk PRIMARY KEY (username)
);
CREATE INDEX login_username_idx ON public.tb_login (username);

CREATE TABLE public.tb_registration (
	id int4 NOT NULL,
	username varchar NOT NULL,
	full_name varchar NOT NULL,
	email varchar NULL,
	"password" varchar NOT NULL,
	create_at timestamp with time zone NULL DEFAULT now(),
	updated_at timestamp with time zone NULL,
	CONSTRAINT registration_pk PRIMARY KEY (id)
);
CREATE INDEX registration_username_idx ON public.tb_registration (username);

CREATE TABLE public.tb_admission (
	full_name varchar NOT NULL,
	id int4 NOT NULL,
	gender varchar NOT NULL,
	dob date NOT NULL,
	address varchar NOT NULL,
	phone_no int4 NOT NULL,
	previous_company varchar NULL,
	designation varchar NULL,
	skills varchar NULL
);
CREATE INDEX tb_admission_id_idx ON public.tb_admission (id);
