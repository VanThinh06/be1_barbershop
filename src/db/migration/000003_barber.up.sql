-- Create Barbers table
CREATE TABLE "Barbers" (
  "barber_id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  "barber_shop_chain_id" UUID,
  "barber_shop_id" UUID,
  "nick_name" VARCHAR(50),
  "barber_code" SERIAL,
  "full_name" VARCHAR(50) NOT NULL,
  "gender_id" INT2,
  "phone" VARCHAR(15) UNIQUE NOT NULL,
  "email" VARCHAR(50) UNIQUE,
  "avatar_url" VARCHAR(120),
  "barber_type_id" INT2 NOT NULL,
  "birth_date" DATE,
  "member_date" DATE,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT now(),
  FOREIGN KEY ("barber_shop_chain_id") REFERENCES "BarberShopChains" ("barber_shop_chain_id"),
  FOREIGN KEY ("barber_shop_id") REFERENCES "BarberShops" ("barber_shop_id"),
  FOREIGN KEY ("gender_id") REFERENCES "Genders" ("id"),
  FOREIGN KEY ("barber_type_id") REFERENCES "BarberTypes" ("id"),
  UNIQUE ("nick_name", "barber_shop_id")
);

-- Create indexes for Barbers table
CREATE INDEX ON "Barbers" ("barber_shop_id");
CREATE INDEX ON "Barbers" ("nick_name");
CREATE INDEX ON "Barbers" ("gender_id");
CREATE INDEX ON "Barbers" ("birth_date");
CREATE INDEX ON "Barbers" ("member_date");
CREATE INDEX ON "Barbers" ("phone");
CREATE INDEX ON "Barbers" ("email");
CREATE INDEX ON "Barbers" ("barber_type_id");

-- Create Users table
CREATE TABLE "Users" (
  "user_id" uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  "phone" varchar(15) UNIQUE NOT NULL,
  "email" varchar(50) UNIQUE,
  "hashed_password" varchar(150),
  "is_active" BOOLEAN NOT NULL DEFAULT TRUE,
  "barber_id" UUID,
  "role_id" int2 NOT NULL,
  "password_changed_at" timestamptz DEFAULT NULL,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  FOREIGN KEY ("role_id") REFERENCES "Roles" ("id"),
  FOREIGN KEY ("barber_id") REFERENCES "Barbers" ("barber_id")
);

-- Create indexes for Users table
CREATE INDEX ON "Users" ("phone");
CREATE INDEX ON "Users" ("email");
CREATE INDEX ON "Users" ("role_id");
CREATE INDEX ON "Users" ("barber_id");

-- Create view_users view
CREATE VIEW "view_users" AS
SELECT
    u."user_id",
    u."phone",
    u."email",
    u."is_active",
    u."created_at",
    u."role_id",
    b."barber_shop_id",
    b."full_name",
    b."nick_name",
    b."avatar_url",
    b."birth_date",
    bt."name" AS "barber_type"
FROM
    "Users" u
LEFT JOIN
    "Barbers" b ON u."barber_id" = b."barber_id"
LEFT JOIN
    "BarberTypes" bt ON b."barber_type_id" = bt."id";

-- Create SessionsUser table
CREATE TABLE "SessionsUser" (
  "id" uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  "user_id" uuid NOT NULL,
  "refresh_token" varchar(255) NOT NULL,
  "user_agent" varchar(255) NOT NULL,
  "client_ip" varchar(50) NOT NULL,
  "fcm_device" varchar(255) NOT NULL,
  "is_blocked" bool NOT NULL DEFAULT false,  
  "expires_at" timestamptz NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  FOREIGN KEY ("user_id") REFERENCES "Users" ("user_id")
);

-- Create index on SessionsUser table
CREATE INDEX ON "SessionsUser" ("user_id");

-- Create OTPRequests table
CREATE TABLE "OTPRequests" (
  "id" uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  "user_id" uuid NOT NULL,
  "otp" varchar(6) NOT NULL,
  "requested_at" timestamptz NOT NULL DEFAULT now(),
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
