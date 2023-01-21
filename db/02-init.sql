CREATE SEQUENCE IF NOT EXISTS account_id;
CREATE SEQUENCE IF NOT EXISTS cloud_pocket_id;

CREATE TABLE IF NOT EXISTS "cloud_pockets" (
  "id" int4 NOT NULL DEFAULT nextval('cloud_pocket_id'::regclass),
  "name" varchar NOT NULL,
  "budget" float8 NOT NULL DEFAULT 0,
  "balance" float8 NOT NULL DEFAULT 0,
  "is_default" boolean,
  "description" varchar,
  "currency" varchar NOT NULL,
  "account_id" int4,
  PRIMARY KEY ("id")
);

CREATE TABLE IF NOT EXISTS "accounts" (
  "id" int4 NOT NULL DEFAULT nextval('account_id'::regclass),
  "balance" float8 NOT NULL DEFAULT 0,
  PRIMARY KEY ("id")
);

ALTER TABLE "cloud_pockets" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");