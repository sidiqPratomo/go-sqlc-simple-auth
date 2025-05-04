-- name: CreateUser :execresult
INSERT INTO users (
    status_otp, nik, photo, first_name, last_name, username, email,
    gender, address, phone_number, password, email_verified_at,
    remember_token, created_by, updated_by, created_time, updated_time, status
) VALUES (
    ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
);

-- name: GetUserByID :one
SELECT * FROM users WHERE id = ? LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
WHERE status = COALESCE(?, status)
ORDER BY created_time DESC
LIMIT ? OFFSET ?;

-- name: UpdateUserStatus :exec
UPDATE users SET status = ?, updated_by = ?, updated_time = ? WHERE id = ?;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = ?;

-- name: CreateRole :execresult
INSERT INTO roles (name, code, created_by, updated_by, created_time, updated_time, status)
VALUES (?, ?, ?, ?, ?, ?, ?);

-- name: GetRoleByID :one
SELECT * FROM roles WHERE id = ?;

-- name: ListRoles :many
SELECT * FROM roles WHERE status = COALESCE(?, status);

-- name: CreateUserOtp :execresult
INSERT INTO user_otps (user_id, otp, expired_at, created_by, updated_by, created_time, updated_time, status)
VALUES (?, ?, ?, ?, ?, ?, ?, ?);

-- name: GetUserOtpByUserID :one
SELECT * FROM user_otps WHERE user_id = ? ORDER BY created_time DESC LIMIT 1;

-- name: CreateRoleUser :execresult
INSERT INTO role_users (user_id, roles_id, created_by, updated_by, created_time, updated_time, status)
VALUES (?, ?, ?, ?, ?, ?, ?);

-- name: GetRoleUserByUserID :one
SELECT * FROM role_users WHERE user_id = ?;

-- name: CreatePrivelege :execresult
INSERT INTO priveleges (module, submodule, ordering, action, method, uri, created_by, updated_by, created_time, updated_time, status)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: ListPriveleges :many
SELECT * FROM priveleges WHERE status = COALESCE(?, status);

-- name: CreateRolePrivelege :execresult
INSERT INTO role_privileges (role, action, method, uri, created_by, updated_by, created_time, updated_time, status)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: GetRolePrivilegesByRole :many
SELECT * FROM role_privileges WHERE role = ? AND status = 1;

-- name: GetUserByUsername :one
SELECT * FROM users WHERE username = ? LIMIT 1;

-- name: GetUserRoles :many
SELECT 
    ru.id AS id,
    ru.user_id AS user_id,
    r.id AS role_id,
    r.name AS role_name,
    r.code AS role_code,
    r.created_by AS role_created_by,
    r.updated_by AS role_updated_by,
    r.created_time AS role_created_time,
    r.updated_time AS role_updated_time,
    r.status AS role_status,
    ru.created_by AS ru_created_by,
    ru.updated_by AS ru_updated_by,
    ru.created_time AS ru_created_time,
    ru.updated_time AS ru_updated_time,
    ru.status AS ru_status
FROM role_users ru
JOIN roles r ON ru.roles_id = r.id
WHERE ru.user_id = ?;

-- name: GetUserPrivileges :many
SELECT 
    rp.id AS id,
    rp.role AS role,
    p.module AS module,
    p.submodule AS submodule,
    p.ordering AS ordering,
    rp.action AS action,
    rp.uri AS uri,
    rp.method AS method,
    rp.created_by AS created_by,
    rp.updated_by AS updated_by,
    rp.created_time AS created_time,
    rp.updated_time AS updated_time,
    rp.status AS status
FROM role_privileges rp
JOIN priveleges p ON 
    rp.action = p.action AND 
    rp.uri = p.uri AND 
    rp.method = p.method
WHERE rp.role IN (
    SELECT roles_id FROM role_users WHERE user_id = ?
) AND rp.status = 1;