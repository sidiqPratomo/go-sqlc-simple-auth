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
