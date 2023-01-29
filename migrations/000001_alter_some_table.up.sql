CREATE TABLE if not exists "courses"(
    "id" serial primary key,
    "name" VARCHAR(255) NOT NULL,
    "course_price" integer NOT NULL,
    "category_id" integer DEFAULT 2,
    "info" VARCHAR(255),
    "business_id" integer not NULL,
    "image_url" VARCHAR(255),
    "sale_of" integer DEFAULT 0,
    "created_at" TIMESTAMP WITH TIME ZONE default current_timestamp
);
