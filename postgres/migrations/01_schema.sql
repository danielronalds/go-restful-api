CREATE SCHEMA api;

CREATE TABLE IF NOT EXISTS api.Events (
  Id            SERIAL PRIMARY KEY,
  Name          varchar(255) NOT NULL,
  Description   varchar(255) NOT NULL,
  Startdate     varchar(255) NOT NULL,
  Enddate       varchar(255) NOT NULL
) ;
