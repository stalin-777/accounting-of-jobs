CREATE TABLE IF NOT EXISTS workplace(
   id serial PRIMARY KEY,
   hostname VARCHAR (100) NOT NULL,
   ip inet NOT NULL,
   username VARCHAR (100) UNIQUE NOT NULL
);