CREATE TABLE public.todos (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid() NOT NULL,
    value character varying(150) NOT NULL,
    completed boolean DEFAULT false NOT NULL
);
