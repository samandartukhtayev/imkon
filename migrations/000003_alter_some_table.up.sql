CREATE TABLE if not exists vacancies(
    "id" serial primary key,
    "name" VARCHAR(255) NOT NULL,
    "category_id" integer default 2,
    "image_url" VARCHAR(255),
    "address" VARCHAR(255),
    "job_type" VARCHAR(255) NOT NULL,
    "min_salary" integer,
    "max_salary" integer,
    "info" VARCHAR(255),
    "views_count" integer,
    "business_id" integer,
    "created_at"  TIMESTAMP WITH TIME ZONE default current_timestamp
);