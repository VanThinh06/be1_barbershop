CREATE EXTENSION IF NOT EXISTS "uuid-ossp";


CREATE TABLE "employee" ("id" uuid UNIQUE PRIMARY KEY DEFAULT (uuid_generate_v4()), "username" varchar NOT NULL,
                                                                                                       "role" varchar NOT NULL,
                                                                                                                      "image" varchar, "store_id" uuid,
                                                                                                                                       "created_at" timestamp NOT NULL DEFAULT (now()), "update_at" timestamp);


CREATE TABLE "manager" ("id" uuid UNIQUE PRIMARY KEY DEFAULT (uuid_generate_v4()), "username" varchar NOT NULL,
                                                                                                      "role" varchar NOT NULL,
                                                                                                                     "image" varchar, "store_id" uuid[], "created_at" timestamp NOT NULL DEFAULT (now()), "update_at" timestamp);
CREATE TABLE "store" (
    "id" uuid UNIQUE PRIMARY KEY DEFAULT (uuid_generate_v4()),
    "name_store" varchar, 
    "location" integer,
    "manager_id" uuid,
    "employee_id" uuid,
    "status" varchar, 
    "created_at" timestamp NOT NULL DEFAULT (now()), 
    "update_at" timestamp);


CREATE INDEX ON "employee" ("store_id");


CREATE INDEX ON "employee" ("username");


CREATE INDEX ON "manager" ("username");


CREATE INDEX ON "store" ("id");


CREATE INDEX ON "store" ("name_store");


CREATE INDEX ON "store" ("manager_id");


CREATE INDEX ON "store" ("employee_id");


ALTER TABLE "store" ADD
FOREIGN KEY ("manager_id") REFERENCES "manager" ("id");


ALTER TABLE "store" ADD
FOREIGN KEY ("employee_id") REFERENCES "employee" ("id");

