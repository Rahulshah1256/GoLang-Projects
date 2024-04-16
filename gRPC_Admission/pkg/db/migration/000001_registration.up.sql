CREATE TABLE registration (
  "id" int4 NOT NULL,
  "username" varchar PRIMARY KEY,
  "hashed_password" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "updated_at" timestamptz NOT NULL DEFAULT('0001-01-01 00:00:00Z'),  
  "created_at" timestamptz NOT NULL DEFAULT (now())
);
CREATE INDEX registration_idx ON registration (username);

CREATE TABLE login (
	"username" varchar NOT NULL,
	"hashed_password" varchar NOT NULL,
	CONSTRAINT login_pk PRIMARY KEY (username)
);
CREATE INDEX login_username_idx ON login (username);

CREATE TABLE admission (
	"full_name" varchar NOT NULL,
	"id" int4 NOT NULL,
	"gender" varchar NOT NULL,
	"dob" date NOT NULL,
	"address" varchar NOT NULL,
	"phone_no" int4 NOT NULL,
	"previous_company" varchar NULL,
	"designation" varchar NULL,
	"skills" varchar NULL,
    CONSTRAINT id_pk PRIMARY KEY (id)
);
CREATE INDEX admission_id_idx ON admission (id);