CREATE TABLE "Barbers" (
  "barber_id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  "barber_shop_chain_id" UUID,
  "barber_shop_id" UUID,
  "nick_name" VARCHAR(50),
  "barber_code" SERIAL UNIQUE NOT NULL,
  "full_name" VARCHAR(50) NOT NULL,
  "gender_id" INT2,
  "phone" varchar(15) UNIQUE NOT NULL,
  "email" varchar(50) UNIQUE,
  "avatar_url" VARCHAR(120),
  "barber_role_id" INT2 NOT NULL,
  "birth_date" DATE,
  "member_date" DATE,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT now(),
  FOREIGN KEY ("barber_shop_chain_id") REFERENCES "BarberShopChains" ("barber_shop_chain_id"),
  FOREIGN KEY ("barber_shop_id") REFERENCES "BarberShops" ("barber_shop_id"),
  FOREIGN KEY ("gender_id") REFERENCES "Genders" ("id"),
  FOREIGN KEY ("barber_role_id") REFERENCES "BarberRoles" ("barber_role_id"),
  UNIQUE ("nick_name", "barber_shop_id")
);

CREATE INDEX ON "Barbers" ("barber_shop_id");
CREATE INDEX ON "Barbers" ("nick_name");
CREATE INDEX ON "Barbers" ("gender_id");
CREATE INDEX ON "Barbers" ("birth_date");
CREATE INDEX ON "Barbers" ("member_date");
CREATE INDEX ON "Barbers" ("phone");
CREATE INDEX ON "Barbers" ("email");
CREATE INDEX ON "Barbers" ("barber_role_id");


-- Create Users table
CREATE TABLE "Users" (
  "user_id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "phone" varchar(15) UNIQUE NOT NULL,
  "email" varchar(50) UNIQUE,
  "hashed_password" varchar(150),
  "is_active" BOOLEAN NOT NULL DEFAULT TRUE,
  "barber_id" UUID,
  "role_id" int2 NOT NULL,
  "password_changed_at" timestamptz DEFAULT NULL,
  "create_at" timestamptz NOT NULL DEFAULT (now()),
  FOREIGN KEY ("role_id") REFERENCES "Roles" ("id"),
  FOREIGN KEY ("barber_id") REFERENCES "Barbers" ("barber_id")
);

-- Create indexes on Users
CREATE INDEX ON "Users" ("phone");
CREATE INDEX ON "Users" ("email");
CREATE INDEX ON "Users" ("role_id");
CREATE INDEX ON "Users" ("barber_id");

CREATE VIEW "view_users" AS
SELECT
    u."user_id",
    u."phone",
    u."email",
    u."is_active",
    u."create_at",
    b."barber_shop_chain_id" AS "client_id",
    b."full_name" AS "barber_full_name",
    b."nick_name" AS "barber_nick_name",
    b."avatar_url" AS "barber_avatar_url", 
    b."birth_date" AS "barber_birth_date", 
    br."name" AS "barber_type",
    r."name" AS "role_name"
FROM
    "Users" u
LEFT JOIN
    "Barbers" b ON u."barber_id" = b."barber_id"
LEFT JOIN
    "Barber_Roles" br ON b."barber_role_id" = br."barber_role_id"
LEFT JOIN
    "Roles" r ON u."role_id" = r."id";

-- Create SessionsUser table
CREATE TABLE "SessionsUser" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "user_id" uuid NOT NULL,
  "refresh_token" varchar NOT NULL,
  "user_agent" varchar NOT NULL,
  "client_ip" varchar NOT NULL,
  "fcm_device" varchar NOT NULL,
  "is_blocked" bool NOT NULL DEFAULT false,  
  "expires_at" timestamptz NOT NULL,
  "create_at" timestamptz NOT NULL DEFAULT (now()),
  FOREIGN KEY ("user_id") REFERENCES "Users" ("user_id")
);

-- Create index on SessionsBarber
CREATE INDEX ON "SessionsBarber" ("user_id");

-- Create OTPRequests table
CREATE TABLE "OTPRequests" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "user_id" uuid NOT NULL,
  "otp" varchar(6) NOT NULL,
  "requested_at" timestamptz NOT NULL DEFAULT (now()),
  "is_confirm" bool NOT NULL DEFAULT false,
  FOREIGN KEY ("user_id") REFERENCES "Users" ("user_id")
);

-- Create or replace function to clean old OTP requests
CREATE OR REPLACE FUNCTION clean_old_otp_requests()
RETURNS TRIGGER AS $$
BEGIN
    DELETE FROM "OTPRequests"
    WHERE "user_id" = NEW."user_id" 
      AND "requested_at"::date <> CURRENT_DATE;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Create trigger to invoke the function before insert
CREATE TRIGGER before_insert_otp_request
BEFORE INSERT ON "OTPRequests"
FOR EACH ROW
EXECUTE FUNCTION clean_old_otp_requests();