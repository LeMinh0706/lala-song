package util

import (
	"path/filepath"
	"regexp"
	"strings"
)

var AllowType = map[string]bool{
	".png":  true,
	".jpg":  true,
	".jpeg": true,
	".gif":  true,
}

var AllowMp3 = map[string]bool{
	".mp3": true,
}

var AllowTxt = map[string]bool{
	".txt": true,
}

func Mp3Check(mp3 string) bool {
	ext := strings.ToLower(filepath.Ext(mp3))
	return AllowMp3[ext]
}

func LyricCheck(lyric string) bool {
	ext := strings.ToLower(filepath.Ext(lyric))
	return AllowTxt[ext]
}

func FileExtCheck(image string) bool {
	ext := strings.ToLower(filepath.Ext(image))

	return AllowType[ext]
}

func EmailCheck(email string) bool {
	const emailCheck = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	regex := regexp.MustCompile(emailCheck)
	return regex.MatchString(email)
}

func UsernameNotSpace(username string) bool {
	const usernameCheck = `^\S+$`
	regex := regexp.MustCompile(usernameCheck)
	return regex.MatchString(username)
}
