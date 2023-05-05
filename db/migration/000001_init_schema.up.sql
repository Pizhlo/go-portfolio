CREATE TABLE "invitations" (
  "id" serial PRIMARY KEY,
  "hr_name" varchar NOT NULL,
  "phone" varchar NOT NULL,
  "email" varchar NOT NULL,
  "company" varchar NOT NULL,
  "message" varchar NOT NULL,
  "date" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "projects" (
  "id" serial PRIMARY KEY,
  "name" varchar NOT NULL,
  "link" varchar NOT NULL,
  "description" varchar NOT NULL,
  "date" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "education" (
  "id" serial PRIMARY KEY,
  "name" varchar NOT NULL,
  "link" varchar NOT NULL,
  "description" varchar NOT NULL,
  "date" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "admin" (
  "login" varchar NOT NULL,
  "password" varchar NOT NULL
);
