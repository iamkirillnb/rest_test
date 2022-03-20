create table data (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    name varchar not null,
    email character varying not null,
    question text not null
);
