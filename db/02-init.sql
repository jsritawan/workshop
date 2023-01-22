CREATE SEQUENCE IF NOT EXISTS account_id;
CREATE SEQUENCE IF NOT EXISTS cloud_pocket_id;

CREATE TABLE IF NOT EXISTS "cloud_pockets" (
  "id" int4 NOT NULL DEFAULT nextval('cloud_pocket_id'::regclass),
  "name" varchar NOT NULL,
  "budget" float8 NOT NULL DEFAULT 0,
  "balance" float8 NOT NULL DEFAULT 0,
  "is_default" boolean DEFAULT false,
  "description" varchar NOT NULL DEFAULT '',
  "currency" varchar NOT NULL,
  "account_id" int4 NOT NULL,
  PRIMARY KEY ("id")
);

CREATE TABLE IF NOT EXISTS "accounts" (
  "id" int4 NOT NULL DEFAULT nextval('account_id'::regclass),
  "balance" float8 NOT NULL DEFAULT 0,
  PRIMARY KEY ("id")
);

INSERT INTO "cloud_pockets"(name,budget,balance,is_default,description, currency,account_id) VALUES 
('pocket 1',1000.0,2000.0,true,'description','THB',1),
('pocket 2',1000.0,0.0,false,'description','THB',1);