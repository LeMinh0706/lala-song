package res

const (
	StatusOk              = 200
	CodeSuccess           = 201
	UpdateFriend          = 20101
	ChangePassword        = 20102
	AddEmail              = 20103
	SeenNoti              = 20104
	DeleteSuccess         = 204
	DeleteNoti            = 20401
	ErrBadRequest         = 40000
	ErrBadRequestPage     = 40001
	ErrBadRequestPageSize = 40002
	ErrBadRequestMime     = 40003
	ErrBadRequestId       = 40004
	ErrAddSinger          = 40005
	ErrAddAlbum           = 40018
	ErrImageSize          = 40006
	ErrGender             = 40007
	ErrAddGenre           = 40017
	ErrUnauthorize        = 40101
	ErrInvalid            = 40102
	ErrYourSelf           = 40103
	ErrUserExist          = 40900
	ErrLogin              = 40104
	ErrNotFoundUser       = 40401
	ErrFindPost           = 40402
	ErrUnlike             = 40403
	ErrLike               = 40404
	ErrFileTooLarge       = 41300
	ErrWrongPassword      = 40105
	ErrEmailExists        = 40901
	ErrUsernameChar       = 40008
	ErrMinPassword        = 40009
	ErrMinFullname        = 40010
	ContentNull           = 40011
	ErrAccountID          = 40012
	ErrPositionField      = 40013
	ErrImageWasDelete     = 40408
	HaveFollow            = 40111
	AcceptForbidden       = 40303
	TheirFriend           = 40411
	WaitngAccept          = 40412
	ErrBadLngLat          = 40020
	ErrSaveImage          = 40021
	ErrEmptyContent       = 40022
	ErrAccountExists      = 40414
	ErrDeleteComment      = 40119
	ErrInputFollow        = 40415
	ErrFullNameNull       = 40014
	ErrTokenInvalid       = 40304
	ErrInputSearch        = 40015
	ErrEmailInvalid       = 40016
	ErrReport             = 40116
	ErrVerify             = 40117
	ErrNotAdmin           = 40118
	ErrEmailNotExists     = 40420
	YouHaveRequest        = 40307
	PasswordHaveChange    = 40308
	ResetPasswordTimeOut  = 40309
	CantDelete            = 40425
	ErrBadRequestQuery    = 40023
	ErrSingerNotfound     = 40423
	ErrGenreNotFound      = 40424
	ErrUsernameSpace      = 40030
	ErrForbidden          = 40301
)

var msg = map[int]string{
	StatusOk:              "Ok",
	CodeSuccess:           "Success",
	DeleteSuccess:         "Delete no error",
	UpdateFriend:          "Success update to friend",
	ChangePassword:        "Success to reset password",
	AddEmail:              "Add email complete",
	SeenNoti:              "Seen notification success",
	DeleteNoti:            "Delete notification success",
	ErrBadRequest:         "Yêu cầu không hợp lệ",
	ErrBadRequestPage:     "Trang phải là số và lớn hơn 1",
	ErrBadRequestPageSize: "Số lượng phải lớn hơn 1",
	ErrBadRequestQuery:    "Trang và số lượng phải lớn hơn 1",
	ErrBadRequestMime:     "Chỉ có thể dùng định dạng hình ảnh với .png, .jpg, .jpeg, .gif",
	ErrAddSinger:          "Images shoud less than 10",
	ErrGender:             "Giới tính chỉ có thể 0 (nữ) hoặc 1 (nam)",
	ErrUnauthorize:        "Unauthorized",
	ErrInvalid:            "Invalid Token",
	ErrYourSelf:           "Không có quyền truy cập",
	ErrUserExist:          "Tài khoảng này đã tồn tại",
	ErrLogin:              "Sai tài khoản hoặc mật khẩu",
	ErrNotFoundUser:       "User not found",
	ErrFindPost:           "Không tìm thấy bài hát",
	ErrUnlike:             "You didn't like this post yet",
	ErrLike:               "You have liked post yet",
	ErrFileTooLarge:       "File quá lớn, chỉ có thể dùng file dưới 10MB",
	ErrWrongPassword:      "Wrong password",
	ErrEmailExists:        "Email exists",
	ErrEmailNotExists:     "Email doesn't exist",
	ErrUsernameSpace:      "Username không được có khoảng trống",
	ErrUsernameChar:       "Username từ 6 đến 16 ký tự",
	ErrMinPassword:        "Password từ 6 ký tự trở lên",
	ErrMinFullname:        "Fullname từ 6 ký tự trở lên",
	ContentNull:           "Description for comment can't null",
	ErrAccountID:          "Account id must be number",
	ErrPositionField:      "LNG or Lat must be both fill or both empty",
	ErrImageWasDelete:     "Image not found or was deleted",
	HaveFollow:            "You have followed this person or they waiting for your acceptance",
	AcceptForbidden:       "Waiting for their reply",
	TheirFriend:           "You're their friend",
	WaitngAccept:          "They're waiting for your acceptance",
	ErrSaveImage:          "Failed to save image",
	ErrEmptyContent:       "Description or images can't be empty",
	ErrAccountExists:      "This account doesn't exist",
	ErrDeleteComment:      "Comment not found",
	ErrInputFollow:        "Error follow status input",
	ErrFullNameNull:       "Fullname can't be empty",
	ErrTokenInvalid:       "Invalid token",
	ErrInputSearch:        "Seaching bar can't be empty",
	ErrEmailInvalid:       "This email is invalid",
	ErrReport:             "You have report this post with this issue",
	YouHaveRequest:        "You requested a few minutes ago, please wait",
	PasswordHaveChange:    "You have changed password before",
	ResetPasswordTimeOut:  "You're to late, try request forgot-password again",
	ErrAddGenre:           "Không thể thêm dòng nhạc này",
	CantDelete:            "Notification not found",
	ErrVerify:             "Please wait for approval",
	ErrNotAdmin:           "Access denied",
	ErrBadRequestId:       "Không tìm thấy nội dung",
	ErrSingerNotfound:     "Không tìm thấy nghệ sĩ",
	ErrGenreNotFound:      "Dòng nhạc này hiện chưa có",
	ErrForbidden:          "Chưa đủ có quyền thực hiện",
	ErrAddAlbum:           "Thêm Album thất bại",
}

// ErrOutOfDate:   "Token is out of date",

// real err
var (
	EmailExists   = "ERROR: duplicate key value violates unique constraint \"users_email_key\" (SQLSTATE 23505)"
	UserExists    = "pq: duplicate key value violates unique constraint \"users_username_key\""
	WrongUsername = "wrong username"
	WrongPassword = "wrong password"
	FollowGhost   = "ERROR: insert or update on table \"follower\" violates foreign key constraint \"follower_to_follow_fkey\" (SQLSTATE 23503)"
)
