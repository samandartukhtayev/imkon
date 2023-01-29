CREATE TABLE if not exists "users"(
    "id" serial primary key,
    "first_name" VARCHAR(255) NOT NULL,
    "last_name" VARCHAR(255),
    "email" VARCHAR(255) NOT NULL,
    "phone_number" VARCHAR(255),
    "password" VARCHAR(255) NOT NULL,
    "image_url" VARCHAR(255),
    "portfolia_url" VARCHAR(255),
    "created_at" TIMESTAMP WITH TIME ZONE default current_timestamp
);