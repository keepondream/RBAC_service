package utils

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/shopspring/decimal"
	"github.com/spf13/cast"
)

// LoadingEnv 将配置文件加载至系统环境变量
func LoadingEnv() {
	APP_ENV := os.Getenv("APP_ENV")

	envFileName := fmt.Sprintf("%s.env", APP_ENV)
	secretFileName := fmt.Sprintf("%s.secret.env", APP_ENV)
	err := godotenv.Load(envFileName, secretFileName)
	if err != nil {
		log.Panicf("Loading env file error : %v ", err)
	}
}

// ParsePageAndPageSize 统一解析page,page_size
func ParsePageAndPageSize(pageValue, pageSizeValue string) (page, pageSize string) {
	page = "1"
	pageSize = "20"
	if cast.ToInt(pageValue) > 0 {
		page = pageValue
	}

	if cast.ToInt(pageSizeValue) > 0 {
		pageSize = pageSizeValue
	}

	return
}

// GetPageAndPageSize 从http.request解析分页参数 page, per_page
func GetPageAndPageSize(r *http.Request) (string, string) {
	pageValue := r.URL.Query().Get("page")
	pageSizeValue := r.URL.Query().Get("per_page")

	return ParsePageAndPageSize(pageValue, pageSizeValue)
}

// ParsePagination 根据page,page_size解析偏移量
func ParsePagination(pageValue, pageSizeValue string) (offset, limit int) {
	page, pageSize := ParsePageAndPageSize(pageValue, pageSizeValue)
	limit = cast.ToInt(pageSize)
	offset = limit * (cast.ToInt(page) - 1)
	return
}

// GetPagination 从http.request解析分页偏移量
func GetPagination(r *http.Request) (offset, limit int) {
	pageValue := r.URL.Query().Get("page")
	pageSizeValue := r.URL.Query().Get("per_page")

	return ParsePagination(pageValue, pageSizeValue)
}

// ParseQuery 解析筛选条件
// 请求头需要设置 "application/x-www-form-urlencoded" , 否则空格无法解析
// 规则: 查询关键字以空格分割,多个维度用冒号连接关键字,每个关键字需要用encode
// 示例: 查询默认关键字为 "xiaoming!! xiaofei" 用户名是 xiaoming!! 或者 xiaofei 的数据
// 完整请求示例: http://host.com?query=${encodeURIComponent('LEGO 10272 Creator: Old Trafford - Manchester United')} username:${encodeURIComponent('xiaoming!!')} username:${encodeURIComponent('xiaofei')}
func ParseQuery(query string) map[string][]string {
	conditions := map[string][]string{}
	queryArr := strings.Split(query, " ")

	for _, encodeValue := range queryArr {
		if encodeValue == "" {
			continue
		}
		conditionArr := strings.Split(encodeValue, ":")
		if len(conditionArr) == 2 {
			// 多维度条件筛选
			conditionValue, _ := url.QueryUnescape(conditionArr[1])
			conditions[conditionArr[0]] = append(conditions[conditionArr[0]], conditionValue)
		} else {
			// 默认筛选关键字 "default"
			conditionValue, _ := url.QueryUnescape(encodeValue)
			conditions["default"] = append(conditions["default"], conditionValue)
		}
	}

	return conditions
}

// GetQuery 解析http.Request中query的搜索条件 query
func GetQuery(r *http.Request) map[string][]string {
	query := r.URL.Query().Get("query")
	return ParseQuery(query)
}

// ParseOrder 解析排序字段
func ParseOrder(order string) []string {
	fields := []string{}
	orderArr := strings.Split(order, ",")
	for _, v := range orderArr {
		if v != "" {
			fields = append(fields, v)
		}
	}

	return fields
}

// GetOrder 解析http.Request中 order 排序字段
func GetOrder(r *http.Request) []string {
	order := r.URL.Query().Get("order")
	return ParseOrder(order)
}

// ParseSort 解析排序方式 支持 asc , desc , 默认:desc
func ParseSort(sort string) string {
	sortVal := "desc"

	if sort == "asc" || sort == "desc" {
		sortVal = sort
	}

	return sortVal
}

// GetSort 解析http.Request中 sort 排序方式, 支持 asc , desc , 默认:desc
func GetSort(r *http.Request) string {
	sort := r.URL.Query().Get("sort")
	return ParseSort(sort)
}

// MakePassword 生成随机密码
func MakePassword(length int) string {
	rand.Seed(time.Now().UnixNano())
	digits := "0123456789"
	specials := "~=+%^*/()[]{}/!@#$?|"
	all := "ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		digits + specials
	buf := make([]byte, length)
	buf[0] = digits[rand.Intn(len(digits))]
	buf[1] = specials[rand.Intn(len(specials))]
	for i := 2; i < length; i++ {
		buf[i] = all[rand.Intn(len(all))]
	}
	rand.Shuffle(len(buf), func(i, j int) {
		buf[i], buf[j] = buf[j], buf[i]
	})
	str := string(buf)

	return str
}

// MakeVerifyCode 生成验证码
func MakeVerifyCode(length int) string {
	rand.Seed(time.Now().UnixNano())
	digits := "0123456789"
	buf := make([]byte, length)
	for i := 0; i < length; i++ {
		buf[i] = digits[rand.Intn(len(digits))]
	}
	rand.Shuffle(len(buf), func(i, j int) {
		buf[i], buf[j] = buf[j], buf[i]
	})
	str := string(buf)

	return str
}

// String2Decimal 字符串转decimal,默认为零值
func String2Decimal(str string) decimal.Decimal {
	d, err := decimal.NewFromString(str)
	if err != nil {
		return decimal.Decimal{}
	}

	return d
}
