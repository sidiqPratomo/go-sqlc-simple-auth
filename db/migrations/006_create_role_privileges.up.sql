CREATE TABLE role_privileges (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    role SMALLINT NOT NULL,
    action VARCHAR(50),
    method VARCHAR(10),
    uri TEXT,
    created_by VARCHAR(100),
    updated_by VARCHAR(100),
    created_time VARCHAR(100),
    updated_time VARCHAR(100),
    status TINYINT NOT NULL
);