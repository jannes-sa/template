CREATE TYPE service_type AS ENUM ('grpc', 'http', 'mq');

CREATE TABLE service_log
(
    type service_type,
    req json,
    res json,
    errcode json,
    job_id character varying(100),
    PRIMARY KEY (job_id)
)