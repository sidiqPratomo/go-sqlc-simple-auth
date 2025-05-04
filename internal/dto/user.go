package dto

import (
	"time"

	"github.com/sidiqPratomo/Go-simple-auth/internal/db"
	"github.com/sidiqPratomo/Go-simple-auth/internal/util"
)

// LoginRequest digunakan saat user mencoba login
type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// RoleDetail represents detailed information about a role
type RoleDetail struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	Code      string `json:"code"`
	CreatedBy string `json:"created_by"`
	UpdatedBy string `json:"updated_by"`
	Status    int    `json:"status"`
}

// LoginResponse adalah response utama ketika user berhasil login
type LoginResponse struct {
	User        User  `json:"user"`
	Role        Role  `json:"role"`
	AccessToken string `json:"access_token"`
	ExpiresAt   string `json:"expires_at"`
}

// User hanya field yang diperlukan untuk frontend
type User struct {
	Id              int64      `json:"id"`
	StatusOTP       *int16      `json:"status_otp,omitempty"`
	Nik             *string    `json:"nik,omitempty"`
	Photo           *string    `json:"photo,omitempty"`
	FirstName       string     `json:"first_name"`
	LastName        string     `json:"last_name"`
	Username        string     `json:"username"`
	Email           string     `json:"email"`
	Gender          *string    `json:"gender,omitempty"`
	Address         *string    `json:"address,omitempty"`
	PhoneNumber     *string    `json:"phone_number,omitempty"`
	EmailVerifiedAt *time.Time `json:"email_verified_at,omitempty"`
	Status          int8       `json:"status"`
}

// Role berisi daftar roles & privileges milik user
type Role struct {
	Role       []RoleData       `json:"role"`
	Privileges []PrivilegeData  `json:"privileges"`
}

// RoleData mewakili satu peran user
type RoleData struct {
	RoleID   int64  `json:"role_id"`
	RoleName string `json:"role_name"`
	RoleCode string `json:"role_code"`
	Status   int8   `json:"status"`
}

// PrivilegeData mewakili hak akses user
type PrivilegeData struct {
	Id        int64     `json:"id"`
	Module    string    `json:"module"`
	Submodule *string   `json:"submodule,omitempty"`
	Ordering  *string   `json:"ordering,omitempty"`
	Action    string    `json:"action"`
	Method    string    `json:"method"`
	Uri       string    `json:"uri"`
	Status    int8      `json:"status"`
}

// Privilege represents a privilege with additional fields
type Privilege struct {
	Id        int64  `json:"id"`
	Role      int64  `json:"role"`
	Action    string `json:"action"`
	Uri       string `json:"uri"`
	Method    string `json:"method"`
	CreatedBy string `json:"created_by"`
	UpdatedBy string `json:"updated_by"`
	Status    int    `json:"status"`
}

// UserRole represents the mapping of user roles
type UserRole struct {
	Id          int64      `json:"id"`
	UsersId     int64      `json:"users_id"`
	RolesId     RoleDetail `json:"roles_id"`
	CreatedBy   string     `json:"created_by"`
	UpdatedBy   string     `json:"updated_by"`
	CreatedTime time.Time  `json:"created_time"`
	UpdatedTime time.Time  `json:"updated_time"`
	Status      int        `json:"status"`
}

// func MapRolesToDTOs(roles []db.GetUserRolesRow) []UserRole {
// 	parseTime := func(timeStr string) time.Time {
// 		parsedTime, _ := time.Parse(time.RFC3339, timeStr)
// 		return parsedTime
// 	}
// 	var roleDTOs []UserRole
// 	for _, role := range roles {
// 		roleDTOs = append(roleDTOs, UserRole{
// 			Id:      role.ID,
// 			UsersId: role.UserID,
// 			RolesId: RoleDetail{
// 				Id:        role.RoleID,
// 				Name:      role.RoleName,
// 				Code:      role.RoleCode,
// 				CreatedBy: role.RoleCreatedBy.String,
// 				UpdatedBy: role.RoleUpdatedBy.String,
// 				Status:    int(role.RoleStatus),
// 			},
// 			CreatedBy:   role.RuCreatedBy.String,
// 			UpdatedBy:   role.RuUpdatedBy.String,
// 			CreatedTime: parseTime(role.RuCreatedTime.String),
// 			UpdatedTime: parseTime(role.RuUpdatedTime.String),
// 			Status:      int(role.RuStatus),
// 		})
// 	}
// 	return roleDTOs
// }

// func MapPrivilegesToDTOs(privileges []db.GetUserPrivilegesRow) []Privilege {
// 	var dto []Privilege
// 	for _, p := range privileges {
// 		dto = append(dto, Privilege{
// 			Id:        int64(p.ID),
// 			Role:      int64(p.Role),
// 			Action:    p.Action.String,
// 			Uri:       p.Uri.String,
// 			Method:    p.Method.String,
// 			CreatedBy: p.CreatedBy.String,
// 			UpdatedBy: p.UpdatedBy.String,
// 			Status:    int(p.Status),
// 		})
// 	}
// 	return dto
// }

func MapRolesToDTOs(roles []db.GetUserRolesRow) []RoleData {
	var roleDTOs []RoleData
	for _, role := range roles {
		roleDTOs = append(roleDTOs, RoleData{
			RoleID:   role.RoleID,
			RoleName: role.RoleName,
			RoleCode: role.RoleCode,
			Status:   int8(role.RoleStatus),
		})
	}
	return roleDTOs
}

func MapPrivilegesToDTOs(privileges []db.GetUserPrivilegesRow) []PrivilegeData {
	var dto []PrivilegeData
	for _, p := range privileges {
		dto = append(dto, PrivilegeData{
			Id:        p.ID,
			Module:    p.Module.String,
			Submodule: util.NullStringToPtr(p.Submodule),
			Ordering:  util.NullStringToPtr(p.Ordering),
			Action:    p.Action.String,
			Method:    p.Method.String,
			Uri:       p.Uri.String,
			Status:    int8(p.Status),
		})
	}
	return dto
}