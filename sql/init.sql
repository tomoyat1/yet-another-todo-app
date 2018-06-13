CREATE TABLE todos
(
    id CHAR(36) NOT NULL,
    incrementalId serial PRIMARY KEY,
    title varchar(256),
    details varchar(512),
    done boolean DEFAULT false  NOT NULL
);
CREATE UNIQUE INDEX todos_id_uindex ON todos (id);
CREATE UNIQUE INDEX todos_incrementalId_uindex ON todos (incrementalId);
