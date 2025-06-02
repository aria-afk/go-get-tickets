DROP TABLE IF EXISTS "tickets";
DROP TABLE IF EXISTS "receipts";
DROP TABLE IF EXISTS "events";
ALTER TABLE "vendors" DROP CONSTRAINT IF EXISTS "vendors_owner_uuid_fkey";
DROP TABLE IF EXISTS "vendor_users";
DROP TABLE IF EXISTS "vendors";

