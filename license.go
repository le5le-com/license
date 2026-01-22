package license

import (
	"time"
)

type LicenseInfo struct {
	IsLe5le          bool
	AppExpiration    time.Time
	UserLimit        int64
	DBDiskLimitG     int64
	DBDiskLimit      int64
	CollectionLimit  int64
	DocumentLimit    int64
	Access           uint16
	UpdateExpiration time.Time
}

func GetSecretKey() string {
	return "your-secret#le5le.com"
}

func DecryptLicense(license, secret string) (info LicenseInfo) {
	info.AppExpiration = time.Now().AddDate(100, 0, 0)
	info.UserLimit = 100000000
	info.DBDiskLimitG = 100000000
	info.CollectionLimit = 100000000
	info.DocumentLimit = 100000000
	info.Access = 255

	info.UpdateExpiration = info.AppExpiration

	return
}

// Mk 生成key
func Mk(nums string) string {
	return nums
}

func Encode(key, text string) string {
	return text
}

// Dc 解码
func Dc(key, text string) string {
	return text
}
