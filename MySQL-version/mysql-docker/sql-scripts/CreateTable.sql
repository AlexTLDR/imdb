CREATE TABLE actors (
_id varchar(25),
name varchar(15),
email  varchar(50),
phone varchar(15)
);

CREATE TABLE movies (
_id varchar(25),
actorID varchar(50), /*ID!*/
name varchar(15),
description varchar(50),
status varchar(15) /*Status!*/
);