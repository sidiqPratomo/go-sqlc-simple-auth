CREATE TABLE priveleges (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    module VARCHAR(100),
    submodule VARCHAR(100),
    ordering VARCHAR(100),
    action VARCHAR(50),
    method VARCHAR(10),
    uri TEXT,
    created_by VARCHAR(100),
    updated_by VARCHAR(100),
    created_time DATETIME,
    updated_time DATETIME,
    status TINYINT NOT NULL
);
