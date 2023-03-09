CREATE TABLE
IF NOT EXISTS tracks
(
    track_id int NOT NULL,
    track_name character varying (40) COLLATE pg_catalog."default",
    artist character varying (40) COLLATE pg_catalog."default",
    track_length int NULL,

    CONSTRAINT tracks_pkey PRIMARY KEY (track_id)
)
WITH
(
    OIDS = FALSE
)
TABLESPACE pg_default;