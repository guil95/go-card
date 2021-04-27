CREATE DATABASE IF NOT EXISTS card_go_db;

USE card_go_db;

DROP TABLE IF EXISTS accounts, operations_types, transactions;

CREATE TABLE IF NOT EXISTS accounts (
    id VARCHAR(255) UNIQUE PRIMARY KEY,
    document_number VARCHAR(255) NOT NULL,
    available_credit_limit DECIMAL(12,2) NOT NULL
)  ENGINE=INNODB;

CREATE TABLE IF NOT EXISTS operations_types (
    id INT AUTO_INCREMENT PRIMARY KEY,
    description VARCHAR(255) NOT NULL
)  ENGINE=INNODB;

CREATE TABLE IF NOT EXISTS transactions (
    id INT AUTO_INCREMENT PRIMARY KEY,
    transaction_uuid VARCHAR(255) UNIQUE,
    account_id VARCHAR(255) NOT NULL,
    operation_type_id INT NOT NULL,
    amount DECIMAL(12,2) NOT NULL,
    event_date DATETIME NOT NULL,
    FOREIGN KEY (account_id)
        REFERENCES accounts(id),
    FOREIGN KEY (operation_type_id)
        REFERENCES operations_types(id)
)  ENGINE=INNODB;

INSERT INTO operations_types (description) values ("PURCHASE_IN_CASH");
INSERT INTO operations_types (description) values ("PURCHASE_INSTALLMENT");
INSERT INTO operations_types (description) values ("WITHDRAW");
INSERT INTO operations_types (description) values ("PAYMENT");
