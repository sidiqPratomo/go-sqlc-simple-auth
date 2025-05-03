CREATE TABLE user_otps (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    user_id BIGINT NOT NULL,
    otp VARCHAR(10) NOT NULL,
    expired_at DATETIME,
    created_by VARCHAR(100),
    updated_by VARCHAR(100),
    created_time DATETIME,
    updated_time DATETIME,
    status TINYINT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)
);