CREATE TABLE if not exists businesses(
    "id" serial primary key,
    "name" VARCHAR(255) NOT NULL,
    "password" VARCHAR(255) NOT NULL,
    "address" VARCHAR(255),
    "image_url" VARCHAR(255),
    "info" VARCHAR(255),
    "email" VARCHAR(255) NOT NULL,
    "phone_number" VARCHAR(255),
    "web_site" VARCHAR(255),
    "telegram_account" VARCHAR(255),
    "instagram_account" VARCHAR(255),
    "linked_in_account" VARCHAR(255),
    "created_at" TIMESTAMP WITH TIME ZONE default current_timestamp
);