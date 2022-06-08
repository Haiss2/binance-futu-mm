create table "public"."logs" (
    "id" serial primary key not null,
    "key" text not null,
    "value" text not null,
    "timestamp" bigint not null
)
