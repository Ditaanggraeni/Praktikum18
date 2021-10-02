-- Table: public.mahasiswa

-- DROP TABLE public.mahasiswa;

CREATE TABLE IF NOT EXISTS public.mahasiswa
(
    id integer NOT NULL DEFAULT nextval('mahasiswa_id_seq'::regclass),
    nama character varying COLLATE pg_catalog."default",
    nim integer,
    jurusan character varying COLLATE pg_catalog."default"
)

TABLESPACE pg_default;

ALTER TABLE public.mahasiswa
    OWNER to postgres;