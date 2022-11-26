CREATE TABLE "user" (
    "id" UUid DEFAULT gen_random_uuid() NOT NULL,
    "first_name" VARCHAR(255) NOT NULL,
    "last_name" VARCHAR(255),
    "email" VARCHAR(320) NOT NULL,
    "password" VARCHAR(255) NOT NULL,
    "created_at" Timestamp Without Time Zone NOT NULL DEFAULT NOW(),
    PRIMARY KEY ("id"),
    CONSTRAINT "unique_user_id" UNIQUE("id"),
    CONSTRAINT "unique_user_email" UNIQUE("email")
);