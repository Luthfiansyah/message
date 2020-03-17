package handlers

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	_ "encoding/json"
	"errors"
	"fmt"
	"github.com/Luthfiansyah/warpin-message/database"
	"github.com/spf13/viper"
	"github.com/Luthfiansyah/warpin-message/config"
	"reflect"
	"regexp"
	"time"
	"unicode"

	"strconv"
	"github.com/gin-gonic/gin"
)

func GetRunMode() string {
	serverMode := viper.GetString("RUN_MODE")
	return serverMode
}

type GeneralResponseType struct {
	ResponseStatus    bool   `json:"response_status"`
	ResponseCode      int64  `json:"response_code"`
	ResponseMessage   string `json:"response_message"`
	ResponseTimestamp string `json:"response_timestamp"`
}

func GeneralResponseErrorBuild(ResponseStatus bool, ResponseCode int64, ResponseMessage string) *GeneralResponseType {
	var generalResponseType GeneralResponseType
	generalResponseType.ResponseTimestamp = GetCurrentTimestampTimeZone(config.TIME_ZONE).String()
	generalResponseType.ResponseStatus = ResponseStatus
	generalResponseType.ResponseCode = ResponseCode
	generalResponseType.ResponseMessage = ResponseMessage
	return &generalResponseType
}

func GeneralResponseSuccessBuild(ResponseStatus bool, ResponseCode int64, ResponseMessage string) *GeneralResponseType {
	var generalResponseType GeneralResponseType
	generalResponseType.ResponseTimestamp = GetCurrentTimestampTimeZone(config.TIME_ZONE).String()
	generalResponseType.ResponseStatus = ResponseStatus
	generalResponseType.ResponseCode = ResponseCode
	generalResponseType.ResponseMessage = ResponseMessage
	return &generalResponseType
}

func ShowResponseError(httpStatusCode int, c *gin.Context, errorCode int64, errorMessage string) {
	gr := GeneralResponseErrorBuild(false, errorCode, errorMessage)
	c.JSON(httpStatusCode, gin.H{
		"general_response": gr,
	})
}

type GenericResponse struct {
	Success   bool        `json:"success"`
	ErrorID   int         `json:"errorid"`
	MessageEN string      `json:"error_en"`
	MessageIN string      `json:"error_id"`
	Data      interface{} `json:"data"`
}

func GetCurrentTimestampTimeZone(timeZone string) time.Time {
	theTime := time.Now()
	loc, _ := time.LoadLocation(timeZone)
	theTime = theTime.In(loc)
	return theTime
}

func GetCurrentTimeTimeZone(timeZone string) string {
	theTime := time.Now()
	loc, _ := time.LoadLocation(timeZone)
	theTime = theTime.In(loc)
	time := theTime.Format("2006-01-02 15:04:05")
	return time
}

func GetCurrentTimeTimeZoneUnix(timeZone string) int64 {
	theTime := time.Now()
	loc, _ := time.LoadLocation(timeZone)
	theTime = theTime.In(loc)
	time := theTime.Unix()
	return time
}

func convertToInt(value string) int {
	i, err2 := strconv.ParseInt(value, 10, 32)
	if err2 != nil {
		panic(err2)
	}
	return int(i)
}

func convertToInt64(value string) int64 {
	i, err2 := strconv.ParseInt(value, 10, 32)
	if err2 != nil {
		panic(err2)
	}
	return int64(i)
}

func IsLetter(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func IsNumber(s string) bool {
	for _, r := range s {
		if !unicode.IsNumber(r) {
			return false
		}
	}
	return true
}

func RemoveOuterQuotes(string string) string {
	return regexp.MustCompile(`^"(.*)"$`).ReplaceAllString(string, `$1`)
}

func TrimQuotes(s string) string {
	if len(s) >= 2 {
		if c := s[len(s)-1]; s[0] == c && (c == '"' || c == '\'') {
			return s[1 : len(s)-1]
		}
	}
	return s
}

func EncodeMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func IsItemExists(arrayType interface{}, item interface{}) bool {
	arr := reflect.ValueOf(arrayType)

	/*if arr.Kind() != reflect.Array {
		panic("Invalid data-type")
	}*/

	for i := 0; i < arr.Len(); i++ {
		if arr.Index(i).Interface() == item {
			return true
		}
	}

	return false
}

func RemoveSpecialCharacters(string string) string {
	var re = regexp.MustCompile(`[!@#$%^&*(),.?";:{}|<>'=]`)
	s := re.ReplaceAllString(string, `$1$2`)
	return s
}

func CheckSpecialCharacter(string string) bool {
	matched, _ := regexp.MatchString(`[!@#$%^&*(),.?";:{}|<>'=_]`, string)
	return matched
}

func EscapeSpecialCharacter(string string) string {
	reg, err := regexp.Compile(`[!?;{}+|<>%'=]`)
	if err != nil {
		fmt.Println(err.Error())
	}
	newString := reg.ReplaceAllString(string, "")
	return newString
}

func CheckAlphaNumeric(string string) bool {
	matched, _ := regexp.MatchString(`[^-_A-Za-z0-9 ]+`, string)
	return matched
}

func generateUniqueCodeJob(companyID int64, jobNo string, orderNo string, jobType int64) string {
	//uniqueCode := strconv.Itoa(int(companyID)) + jobNo + orderNo + strconv.Itoa(int(jobType))
	return jobNo
}

func generateUniqueCodeJobDetail(productID int64, productPackingID int64, fromWarehouseLocationID int64, ToWarehouseLocationID int64) string {
	//uniqueCode := strconv.Itoa(int(productID)) + "." + strconv.Itoa(int(productPackingID)) + "." + strconv.Itoa(int(fromWarehouseLocationID)) + "." + strconv.Itoa(int(ToWarehouseLocationID)) + "." + strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
	uniqueCode := strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
	return uniqueCode
}

func ValidateFormatDate(data string) error {
	var layout = "2006-01-02 15:04:05"

	_, err := time.Parse(layout, data)

	if err != nil {
		return err
	}

	return nil
}

func ExtractorInterface(param interface{}) ([]interface{}, error) {

	s := reflect.ValueOf(param)
	if s.Kind() != reflect.Slice {
		return nil, errors.New("InterfaceSlice() given a non-slice type")
	}

	k := make([]interface{}, s.Len())
	for i := 0; i < s.Len(); i++ {
		k[i] = s.Index(i).Interface()
	}

	return k, nil
}

func CacheKeyMessages() string {
	return EncodeMD5Hash("messages" )
	//return "CITY" + strconv.Itoa(int(id))
}

func SetCache(key string, param interface{}) error {

	var appKey = viper.GetString("APP_KEY")

	byteData, _ := json.Marshal(param)

	err := database.SetRedis(appKey+key, string(byteData))
	return err
}

func GetCache(key string) (string, error) {

	appKey := viper.GetString("APP_KEY")
	value := database.GetRedis(appKey + key)

	return value, nil
}
