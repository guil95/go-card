CREATE DATABASE IF NOT EXISTS card_go_db;

USE card_go_db;

DROP TABLE IF EXISTS accounts, operations_types;

CREATE TABLE IF NOT EXISTS accounts (
    id VARCHAR(255) UNIQUE PRIMARY KEY,
    document_number VARCHAR(255) NOT NULL
)  ENGINE=INNODB;

CREATE TABLE IF NOT EXISTS operations_types (
    id INT AUTO_INCREMENT PRIMARY KEY,
    description VARCHAR(255) NOT NULL
)  ENGINE=INNODB;


INSERT INTO operations_types (description) values ("PURCHASE_IN_CASH");
INSERT INTO operations_types (description) values ("PURCHASE_INSTALLMENT");
INSERT INTO operations_types (description) values ("WITHDRAW");
INSERT INTO operations_types (description) values ("PAYMENT");
