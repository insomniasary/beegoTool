package services

import (
	"bytes"
	"compress/gzip"
	"crypto/md5"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"html/template"
	"io/ioutil"
	"math"
	"math/rand"
	"net"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"time"
)

var CstZone = time.FixedZone("CST", 8*3600)

// 取值范围 [0, max)
func RandNum(max int) int {
	if max <= 0 {
		return max
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	num := r.Intn(max)
	return num
}

func InStringSlice(slice []string, element string) bool {
	element = strings.TrimSpace(element)
	for _, v := range slice {
		if strings.TrimSpace(v) == element {
			return true
		}
	}
	return false
}

func RandInt64(min, max int64) int64 {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return rand.Int63n(max-min) + min
}
func RandInt(min, max int) int {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return rand.Intn(max-min) + min
}

func InIntSlice(slice []int, element int) bool {
	if slice == nil {
		return false
	}
	for _, v := range slice {
		if v == element {
			return true
		}
	}

	return false
}

func Round(f float64, n int) float64 {
	n10 := math.Pow10(n)
	return math.Trunc((f+0.5/n10)*n10) / n10
}

/**
 * int64转换string
 * @method parseInt
 * @param  {[type]} b string        [description]
 * @return {[type]}   [description]
 */
func ParseStringInt64(b int64) string {
	id := strconv.FormatInt(b, 10)
	return id
}

func ParseInt64String(b string) int64 {
	id, _ := strconv.ParseInt(b, 10, 64)
	return id
}

/**
 * 转换浮点数为string
 * @method func
 * @param  {[type]} t *             Tools [description]
 * @return {[type]}   [description]
 */
func ParseFlostToString(f float64, prec int) string {
	return strconv.FormatFloat(f, 'f', prec, 64)
}

func GetRandomStringByLetter(n int, letterBytes string) string {
	lettersArr := strings.Split(letterBytes, "")
	retStr := ""
	b := make([]byte, n)
	for range b {
		randInt := rand.Int63() % int64(len(lettersArr))
		retStr += lettersArr[randInt]
	}
	return retStr
}

//生成随机字符串2
func GetRandomStringNew(n int) string {
	const letterBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}

//生成随机字符串2
func GetRandomInt(n int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < n; i++ {
		fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()
}

/**
 * @method func
 * @param token(string)
 * @return string
 */

func SortByKey(sortArrays map[string]string) string {
	var keys []string
	for k := range sortArrays {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	//To perform the opertion you want
	items := make([]string, 0)
	for _, k := range keys {
		if sortArrays[k] != "" {
			items = append(items, k+"="+sortArrays[k])
		}
	}
	return strings.Join(items, "&")
}

func Substring(source string, start int, end int) string {
	var r = []rune(source)
	length := len(r)

	if start < 0 || end > length || start > end {
		return ""
	}

	if start == 0 && end == length {
		return source
	}

	return string(r[start:end])
}

func Split2IntSlice(source string, sep string) []int {
	rpRe := strings.Split(source, sep)
	res := make([]int, 0)
	for _, v := range rpRe {
		t, _ := strconv.Atoi(v)
		res = append(res, t)
	}
	return res
}

func UnmarshalGzip(data []byte) ([]byte, error) {
	b := new(bytes.Buffer)
	binary.Write(b, binary.LittleEndian, data)
	r, err := gzip.NewReader(b)
	if err != nil {
		return nil, err
	} else {
		defer r.Close()
		undatas, err := ioutil.ReadAll(r)
		if err != nil {
			return nil, err
		}
		return undatas, nil
	}
}

func ReplaceStrings(s string, old []string, replace []string) string {
	if s == "" {
		return s
	}
	if len(old) != len(replace) {
		return s
	}

	for i, v := range old {
		s = strings.Replace(s, v, replace[i], 1000)
	}

	return s
}

func Ip2Long(ipStr string) uint32 {
	var long uint32
	binary.Read(bytes.NewBuffer(net.ParseIP(ipStr).To4()), binary.BigEndian, &long)
	return long
}

func VerifyMobileFormat(mobileNum string) bool {
	regular := `^1[0-9]\d{9}$`

	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobileNum)
}

func FormatAppVersion(appVersion, GitCommit, BuildDate string) (string, error) {
	content := `
   Version: {{.Version}}
Go Version: {{.GoVersion}}
Git Commit: {{.GitCommit}}
     Built: {{.BuildDate}}
   OS/ARCH: {{.GOOS}}/{{.GOARCH}}
`
	tpl, err := template.New("version").Parse(content)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	err = tpl.Execute(&buf, map[string]string{
		"Version":   appVersion,
		"GoVersion": runtime.Version(),
		"GitCommit": GitCommit,
		"BuildDate": BuildDate,
		"GOOS":      runtime.GOOS,
		"GOARCH":    runtime.GOARCH,
	})
	if err != nil {
		return "", err
	}

	return buf.String(), err
}

func FileLock(path string) bool {
	if FileExist(path) {
		return false
	}
	err := ioutil.WriteFile(path, []byte(time.Now().String()), 0644)
	if err != nil {
		return false
	}
	return true
}

func DelFileLock(fileName string) {
	os.Remove(fileName)
}

// 判断文件是否存在及是否有权限访问
func FileExist(file string) bool {
	_, err := os.Stat(file)
	if os.IsNotExist(err) {
		return false
	}
	if os.IsPermission(err) {
		return false
	}

	return true
}

func UnicodeIndex(str, substr string) int {
	// 子串在字符串的字节位置
	result := strings.Index(str, substr)
	if result >= 0 {
		// 获得子串之前的字符串并转换成[]byte
		prefix := []byte(str)[0:result]
		// 将子串之前的字符串转换成[]rune
		rs := []rune(string(prefix))
		// 获得子串之前的字符串的长度，便是子串在字符串的字符位置
		result = len(rs)
	}

	return result
}

func Md5_encode(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func Wechat_makeSign(params map[string]string, key string,toUpper bool) string {
	buildStr := SortByKey(params)
	signStr := buildStr + "&key=" + key
	fmt.Println(signStr)
	sign := Md5_encode(signStr)
	if toUpper {
		sign = strings.ToUpper(sign)
	}
	return sign
}

func MapStringToInterface(src map[string]string) map[string]interface{} {
	mapString := make(map[string]interface{})
	for key, value := range src {
		strKey := fmt.Sprintf("%v", key)
		strValue := fmt.Sprintf("%v", value)
		mapString[strKey] = strValue
	}

	return mapString
}

func FileLoad(filepath string) []byte {

	privatefile, err := os.Open(filepath)

	defer privatefile.Close()

	if err != nil {

	}

	privateKey := make([]byte, 2048)

	num, err := privatefile.Read(privateKey)

	return privateKey[:num]

}

func GetFirstDateOfMonth(d time.Time) time.Time {
	d = d.AddDate(0, 0, -d.Day()+1)
	return GetZeroTime(d)
}

//获取某一天的0点时间
func GetZeroTime(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
}

func InterfaceToString(v interface{}) string {
	var ret string
	//fmt.Println(v.(type))
	switch v.(type) {
	case []byte:
		ret = string(v.([]byte))
	case string:
		ret = v.(string)
	case int:
		ret = strconv.Itoa(v.(int))
	case float64:
		//ret = ParseString((int(v.(float64))))
		ret = ParseStringInt64((int64(v.(float64))))
	case []interface{}:
		ret = ""
	case error:
		return ""
	default:
		ret = ""
	}
	return ret
}
func FilePutContents(filename string, data []byte) error {
	if dir := filepath.Dir(filename); dir != "" {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}
	return ioutil.WriteFile(filename, data, 0644)
}

//func GetFirstDateOfMonth(d time.Time) time.Time {
//	d = d.AddDate(0, 0, -d.Day() + 1)
//	return GetZeroTime(d)
//}
////获取某一天的0点时间
//func GetZeroTime(d time.Time) time.Time {
//	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
//}

func Long2ip(ip uint32) net.IP {
	a := byte((ip >> 24) & 0xFF)
	b := byte((ip >> 16) & 0xFF)
	c := byte((ip >> 8) & 0xFF)
	d := byte(ip & 0xFF)
	return net.IPv4(a, b, c, d)
}

// Strstr strstr()
func Strstr(haystack string, needle string) string {
	if needle == "" {
		return ""
	}
	idx := strings.Index(haystack, needle)
	if idx == -1 {
		return ""
	}
	return haystack[idx+len([]byte(needle)):]
}

func TodayExitSecond() int {
	todayLast := time.Now().Format("2006-01-02") + " 23:59:59"

	todayLastTime, _ := time.ParseInLocation("2006-01-02 15:04:05", todayLast, time.Local)

	remainSecond := time.Duration(todayLastTime.Unix() - time.Now().Local().Unix())

	return int(remainSecond)
}

func MonthDay(year int, month int) (days int) {
	if month != 2 {
		if month == 4 || month == 6 || month == 9 || month == 11 {
			days = 30

		} else {
			days = 31
			//fmt.Fprintln(os.Stdout, "The month has 31 days");
		}
	} else {
		if ((year%4) == 0 && (year%100) != 0) || (year%400) == 0 {
			days = 29
		} else {
			days = 28
		}
	}
	//fmt.Fprintf(os.Stdout, "The %d-%d has %d days.\n", year, month, days)
	return
}
