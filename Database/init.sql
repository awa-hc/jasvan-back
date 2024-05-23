SELECT CASE WHEN EXISTS (
    SELECT 1
    FROM pg_catalog.pg_database
    WHERE datname = "jasvan"
) THEN 0 ELSE 1 END AS db_exists \gset

\if :db_exists
    CREATE DATABASE "jasvan";
\endif