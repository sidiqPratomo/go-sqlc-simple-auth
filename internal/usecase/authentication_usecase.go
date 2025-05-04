package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/sidiqPratomo/DJKI-Pengaduan/apperror"
	"github.com/sidiqPratomo/Go-simple-auth/internal/db"
	"github.com/sidiqPratomo/Go-simple-auth/internal/dto"
	"github.com/sidiqPratomo/Go-simple-auth/internal/util"
)

type AuthenticationUsecase interface {
	LoginUser(ctx context.Context, req dto.LoginRequest) (*dto.LoginResponse, error)
}

type authenticationUsecaseImpl struct {
	db          db.Querier
	hashHelper  util.HashHelperIntf
	jwtHelper   util.JwtAuthentication
	emailHelper util.EmailHelper
}

func NewAuthenticationUsecase(q db.Querier, hash util.HashHelperIntf, jwt util.JwtAuthentication, email util.EmailHelper) AuthenticationUsecase {
	return &authenticationUsecaseImpl{
		db:          q,
		hashHelper:  hash,
		jwtHelper:   jwt,
		emailHelper: email,
	}
}

func (u *authenticationUsecaseImpl) LoginUser(ctx context.Context, req dto.LoginRequest) (*dto.LoginResponse, error) {
	user, err := u.db.GetUserByUsername(ctx, req.Username)
	if err != nil {
		return nil, apperror.BadRequestError(errors.New("username not found"))
	}

	ok, err := u.hashHelper.CheckPassword(req.Password, []byte(user.Password))
	if err != nil || !ok {
		return nil, apperror.WrongPasswordError(err)
	}

	if !user.EmailVerifiedAt.Valid {
		return nil, apperror.NewAppError(400, err, "User Not verified")
	}

	claims := util.JwtCustomClaims{UserId: user.ID, Email: user.Email, Role: ""} // role akan diisi setelah fetch

	if user.StatusOtp.Valid && user.StatusOtp.Int16 == 0 {
		token, expiresAt, err := u.jwtHelper.CreateAndSign(claims, u.jwtHelper.Config.AccessSecret)
		if err != nil {
			return nil, apperror.InternalServerError(err)
		}

		roles, err := u.db.GetUserRoles(ctx, user.ID)
		if err != nil {
			return nil, apperror.InternalServerError(err)
		}

		var roleIds []int16
		for _, r := range roles {
			roleIds = append(roleIds, int16(r.RoleID))
		}

		privileges, err := u.db.GetUserPrivileges(ctx, user.ID)
		if err != nil {
			return nil, apperror.InternalServerError(err)
		}

		roleDTOs := dto.MapRolesToDTOs(roles)
		privilegeDTOs := dto.MapPrivilegesToDTOs(privileges)

		return &dto.LoginResponse{
			User: dto.User{
				Id:              user.ID,
				StatusOTP:       util.NullInt16ToPtr(user.StatusOtp),
				Nik:             util.NullStringToPtr(user.Nik),
				Photo:           util.NullStringToPtr(user.Photo),
				FirstName:       user.FirstName,
				LastName:        user.LastName,
				Email:           user.Email,
				Username:        user.Username,
				Gender:          util.NullStringToPtr(user.Gender),
				Address:         util.NullStringToPtr(user.Address),
				PhoneNumber:     util.NullStringToPtr(user.PhoneNumber),
				EmailVerifiedAt: util.NullTimeToPtr(user.EmailVerifiedAt),
				Status:          user.Status,
			},
			Role: dto.Role{
				Role:       roleDTOs,
				Privileges: privilegeDTOs,
			},
			AccessToken: *token,
			ExpiresAt:   *expiresAt,
		}, nil
	}

	// Kirim OTP jika status OTP = 1
	err = u.createAndSendOTP(ctx, int(user.ID), user.Email)
	if err != nil {
		return nil, apperror.InternalServerError(err)
	}

	return nil, nil
}

func (u *authenticationUsecaseImpl) createAndSendOTP(ctx context.Context, userId int, email string) error {
	otp := util.GenerateOTP(6)
	now := time.Now()
	expired := now.Add(10 * time.Minute)

	_, err := u.db.CreateUserOtp(ctx, db.CreateUserOtpParams{
		UserID:     int64(userId),
		Otp:        otp,
		ExpiredAt:  util.TimeToNullTime(expired),
		CreatedBy:  util.IntToNullString(userId),
		UpdatedBy:  util.IntToNullString(userId),
		CreatedTime: util.TimeToNullTime(now),
		UpdatedTime: util.TimeToNullTime(now),
		Status:     1,
	})
	if err != nil {
		return err
	}

	template := `<p>Your OTP code is <strong>{{.OTP}}</strong>. It will expire in 10 minutes.</p>`
	u.emailHelper.AddRequest([]string{email}, "Pengaduan DJKI-OTP")
	if err := u.emailHelper.CreateBody(template, map[string]interface{}{"OTP": otp}); err != nil {
		return err
	}
	return u.emailHelper.SendEmail()
}
