package zhenai

import (
	"log"
	"regexp"
	"strconv"
	"zhenaiSpider/simpleSpider/engine"
	"zhenaiSpider/simpleSpider/model"
)

var (
	ageRe       = regexp.MustCompile(`<div class="m-btn purple"[^>]*>([\d]+)岁</div>`)
	marriageRe  = regexp.MustCompile(`<div class="m-btn purple"[^>]*>(未婚|离异)</div>`)
	heightRe    = regexp.MustCompile(`<div class="m-btn purple"[^>]*>([\d]+)cm</div>`)
	weightRe    = regexp.MustCompile(`<div class="m-btn purple"[^>]*>([\d]+)kg</div>`)
	genderRe    = regexp.MustCompile(`"genderString":"(男|女)士"`)
	educationRe = regexp.MustCompile(`"educationString":"([^"]+)"`)
	incomeRe    = regexp.MustCompile(`<div class="m-btn purple"[^>]*>月收入:([^<]+)</div>`)
	nationRe    = regexp.MustCompile(`<div class="m-btn pink"[^>]*>([^>]+族)</div>`)
	carRe       = regexp.MustCompile(`<div class="m-btn pink"[^>]*>([已|未]买车)</div>`)
	houseRe     = regexp.MustCompile(`<div class="m-btn pink"[^>]*>([已|未]购房)</div>`)
	useCityRe   = regexp.MustCompile(`<div class="des f-cl"[^>]*>([^ ]+) [^<]*</div>`)
	idRe        = regexp.MustCompile(`http://album.zhenai.com/u/([0-9]+)`)
	photoRe     = regexp.MustCompile(`background-image:url\(([^?]+)?`)
	nameRe      = regexp.MustCompile(`<h1 class="nickName"[^>]*>([^<]+)</h1>`)
)

func ParserProfile(contents []byte, name string) engine.ParserResult {
	profile := model.Profile{}
	profile.Age = convertStringToInt(extractString(ageRe, contents))
	profile.Marriage = extractString(marriageRe, contents)
	profile.Height = convertStringToInt(extractString(heightRe, contents))
	profile.Weight = convertStringToInt(extractString(weightRe, contents))
	profile.Gender = extractString(genderRe, contents)
	profile.Education = extractString(educationRe, contents)
	profile.Income = extractString(incomeRe, contents)
	profile.Car = extractString(carRe, contents)
	profile.House = extractString(houseRe, contents)
	profile.Nation = extractString(nationRe, contents)
	profile.City = extractString(useCityRe, contents)
	profile.Photo = extractString(photoRe, contents)
	//extractString(nameRe, contents)
	profile.Name = name

	result := engine.ParserResult{}
	result.Items = append(result.Items, profile)

	return result
}
func convertStringToInt(num string) int {
	if num == "" {
		return 0
	}

	i, err := strconv.Atoi(num)
	if err != nil {
		log.Panicf("num is %s, %v\n", num, err)
		return 0
	}

	return i
}

func extractString(re *regexp.Regexp, contents []byte) string {
	matches := re.FindSubmatch(contents)
	if len(matches) > 0 {
		return string(matches[1])
	}

	return ""
}
