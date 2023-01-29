CREATE TABLE if not exists "accepted_vacancies"(
    "id" serial primary key,
    "user_id" integer not null,
    "vacancy_id" integer not null,
    "created_at" TIMESTAMP WITH TIME ZONE default current_timestamp
);