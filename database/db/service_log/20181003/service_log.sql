CREATE TYPE service_type AS ENUM ('grpc', 'http', 'mq');

CREATE TABLE service_log
(
    type service_type,
    req text,
    res text,
    errcode text,
    job_id character varying(100),
    PRIMARY KEY (job_id)
)