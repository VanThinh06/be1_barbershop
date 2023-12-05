ALTER TABLE "Customers"
DROP COLUMN is_social_auth;

ALTER TABLE "SessionsCustomer"
DROP COLUMN timezone;

DROP INDEX IF EXISTS idx_customers_email;