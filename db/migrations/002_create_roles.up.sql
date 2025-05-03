CREATE TABLE roles (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(100) NOT NULL,
    code VARCHAR(100) NOT NULL,
    created_by VARCHAR(100),
    updated_by VARCHAR(100),
    created_time VARCHAR(100),
    updated_time VARCHAR(100),
    status TINYINT NOT NULL
);