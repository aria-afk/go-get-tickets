CREATE TABLE IF NOT EXISTS "vendors" (
  "uuid" UUID NOT NULL DEFAULT gen_random_uuid(),
  "name" TEXT NOT NULL UNIQUE,
  "owner_uuid" UUID NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT NOW(),
  "updated_at" TIMESTAMP NOT NULL DEFAULT NOW(),
  PRIMARY KEY ("uuid")
);

-- to enable crypt() function for passwords
CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE IF NOT EXISTS "vendor_users" (
  "uuid" UUID NOT NULL DEFAULT gen_random_uuid(),
  "name" TEXT NOT NULL,
  "vendor_uuid" UUID,
  "permissions" TEXT,
  "email" TEXT NOT NULL UNIQUE,
  "password" TEXT NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT NOW(),
  "updated_at" TIMESTAMP NOT NULL DEFAULT NOW(),
  PRIMARY KEY ("uuid"),
  FOREIGN KEY ("vendor_uuid") REFERENCES "vendors"("uuid")
);

-- drop before adding, since there is no builtin IF NOT EXISTS support for constraint
ALTER TABLE "vendors" DROP CONSTRAINT IF EXISTS "vendors_owner_uuid_fkey";
ALTER TABLE "vendors" ADD CONSTRAINT "vendors_owner_uuid_fkey" FOREIGN KEY ("owner_uuid") REFERENCES "vendor_users"("uuid");

CREATE TABLE IF NOT EXISTS "events" (
  "uuid" UUID NOT NULL DEFAULT gen_random_uuid(),
  "name" TEXT NOT NULL,
  "vendor_uuid" UUID NOT NULL,
  "address" TEXT,
  "image_url" TEXT,
  "start_date" DATE,
  "time" TIME,
  "capacity" INTEGER NOT NULL,
  "status" TEXT NOT NULL,
  "ticket_price" DECIMAL(12, 2)[],
  "currency" TEXT,
  "created_at" TIMESTAMP NOT NULL DEFAULT NOW(),
  "updated_at" TIMESTAMP NOT NULL DEFAULT NOW(),
  PRIMARY KEY ("uuid"),
  FOREIGN KEY ("vendor_uuid") REFERENCES "vendors"("uuid")
);

CREATE TABLE IF NOT EXISTS "receipts" (
  "uuid" UUID NOT NULL DEFAULT gen_random_uuid(),
  "payment_method" text NOT NULL,
  "amount" DECIMAL(12, 2),
  "currency" TEXT,
  "number_of_tickets" INTEGER,
  "customer_email" TEXT,
  "customer_name" TEXT,
  "vendor_uuid" UUID,
  "event_uuid" UUID,
  "stripe_payment_ref" TEXT,
  "paypal_payment_ref" TEXT,
  "created_at" TIMESTAMP NOT NULL DEFAULT NOW(),
  "updated_at" TIMESTAMP NOT NULL DEFAULT NOW(),
  PRIMARY KEY ("uuid"),
  FOREIGN KEY ("vendor_uuid") REFERENCES "vendors"("uuid"),
  FOREIGN KEY ("event_uuid") REFERENCES "events"("uuid")
);

CREATE TABLE IF NOT EXISTS "tickets" (
  "uuid" UUID NOT NULL DEFAULT gen_random_uuid(),
  "name" TEXT,
  "purchaser_name" TEXT,
  "purchaser_email" TEXT,
  "status" TEXT,
  "event_uuid" UUID NOT NULL,
  "receipt_uuid" UUID,
  "scanned_at" TIMESTAMP,
  "marked_at" TIMESTAMP,
  "created_at" TIMESTAMP NOT NULL DEFAULT NOW(),
  "updated_at" TIMESTAMP NOT NULL DEFAULT NOW(),
  PRIMARY KEY ("uuid"),
  FOREIGN KEY ("event_uuid") REFERENCES "events"("uuid"),
  FOREIGN KEY ("receipt_uuid") REFERENCES "receipts"("uuid")
);

