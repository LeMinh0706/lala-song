package util

import (
	"fmt"
	"math/rand"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomX() float64 {
	return 106.256229 + rand.Float64()*(107.583089-106.256229)
}

func RandomY() float64 {
	return 10.354072 + rand.Float64()*(11.257274-10.354072)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomDescription() string {
	return RandomString(int(RandomInt(2, 5))) + " " + RandomString(int(RandomInt(2, 5))) + " " + RandomString(int(RandomInt(2, 5))) + " " + RandomString(int(RandomInt(2, 5)))
}

func RandomEmail() string {
	return RandomString(6) + fmt.Sprint(RandomInt(1000, 9999)) + "@gmail.com"
}

func RandomGender() int32 {
	genders := []int32{0, 1}
	n := len(genders)
	return genders[rand.Intn(n)]
}

func RandomTypeImage() string {
	image := []string{".png", ".jpg"}
	n := len(image)
	return image[rand.Intn(n)]
}

//Dai ac

func RandomImage() string {
	images := []string{"upload/post/1702208792550.jpg", "upload/post/Clerics Do It With Protection Tabletop RPG by pixeptional.jpg", "upload/post/Golang Garbage Collector Wall And Art Print.jpg",
		"upload/post/wallpapersden.com_linux-retro_1920x1080.jpg", "upload/post/f9a95cc71b8e5e7581aa049ed79eaf67.jpg"}
	n := len(images)
	return images[rand.Intn(n)]
}

func RandomAvatar(gender int32) string {
	if gender == 0 {
		women := []string{"upload/avatar/avatar_women_1.jpg", "upload/avatar/avatar_women_2.jpg", "upload/avatar/avatar_women_3.jpg", "upload/avatar/avatar_women_4.jpg"}
		n := len(women)
		return women[rand.Intn(n)]
	} else {
		men := []string{"upload/avatar/avatar_men_1.jpg", "upload/avatar/avatar_men_2.jpg", "upload/avatar/avatar_men_3.jpg", "upload/avatar/avatar_men_4.jpg"}
		n := len(men)
		return men[rand.Intn(n)]
	}
}
