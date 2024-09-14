create database url_collections;

use url_collections;

CREATE TABLE urls (
    id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    original_url VARCHAR(200) NOT NULL,
    short_code VARCHAR(50) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT(NOW()),
    updated_at TIMESTAMP NOT NULL DEFAULT(NOW())
);


insert into urls 
(original_url, short_code) 
values 
("https://www.goguru.dev", "goguru");




SELECT * FROM urls;
