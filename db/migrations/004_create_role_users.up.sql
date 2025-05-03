CREATE TABLE role_users (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    user_id BIGINT NOT NULL,
    roles_id BIGINT NOT NULL,
    created_by VARCHAR(100),
    updated_by VARCHAR(100),
    created_time VARCHAR(100),
    updated_time VARCHAR(100),
    status TINYINT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (roles_id) REFERENCES roles(id)
);
