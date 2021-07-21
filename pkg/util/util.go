package util

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"math/rand"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/hieuphq/question-be/pkg/model/errors"
)

var letters = []rune("012345789_ABCDEFGHIJKLMNOPQRSTUVWXYZ")

const randomUserSaltLength = 4

// GenRandomInRange return a random int in a range
func GenRandomInRange(min, max int) int {
	return rand.Intn(max-min) + min
}

// SplitAndTrimSpaceString is a helper for split and strim space from the results
func SplitAndTrimSpaceString(s string, sep string) []string {
	if s == "" {
		return nil
	}

	if sep == "" {
		return []string{strings.TrimSpace(s)}
	}

	l := strings.Split(s, sep)
	rs := []string{}
	for i := range l {
		tmp := strings.TrimSpace(l[i])
		if tmp != "" {
			rs = append(rs, tmp)
		}
	}
	return rs
}

// RandomString generate random string with lenght
func RandomString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// ValidateEmail validate a string is email by regular expression
func ValidateEmail(email string) bool {
	Re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return Re.MatchString(email)
}

// ValidatePhone validate a string is phone by regular expression
func ValidatePhone(phone string) bool {
	re := regexp.MustCompile(`^[+]*[(]{0,1}[0-9]{1,4}[)]{0,1}[-\s\./0-9]*$`)
	return re.MatchString(phone)
}

// GenerateSaltedPassword generate salted password from string and salt
// Return salted string
func GenerateSaltedPassword(password, salt string, loops int) (string, error) {
	salted := salt
	passwd := password

	r := regexp.MustCompile(`^\$([0-9]+)\$(.*)`)
	subStrs := r.FindStringSubmatch(salted)
	if len(subStrs) == 3 {
		i, err := strconv.Atoi(subStrs[1])
		if err != nil {
			return "", err
		}
		loops = i
		salted = subStrs[2]
	}

	for i := 0; i < loops; i++ {
		h := sha1.New()
		h.Write([]byte(salted + passwd))
		passwd = fmt.Sprintf("%x", h.Sum(nil))
	}

	return fmt.Sprintf("$%d$%s", loops, salted+passwd), nil
}

// HashNumber hash a number to string by sha1 algolithm
// Return hashed string
func HashNumber(val int64) string {
	b := []byte(strconv.FormatInt(val, 10))
	h := sha1.New()
	h.Write(b)
	return hex.EncodeToString(h.Sum(nil))
}

// CopyMap make a copy from map
// Return map struct data
func CopyMap(src map[string]interface{}) map[string]interface{} {
	rs := map[string]interface{}{}
	for k, v := range src {
		rs[k] = v
	}

	return rs
}

// ParseErrorCode parse error code from errors.Error
func ParseErrorCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	switch arg := err.(type) {
	case *errors.Error:
		return int(arg.Code)

	case error:
		return http.StatusInternalServerError

	default:
		return http.StatusOK
	}
}

// HandleError handler error from errors.Error
func HandleError(c *gin.Context, err error) {
	if err == nil {
		return
	}

	switch arg := err.(type) {
	// case validator.ValidatorError:
	// 	ds := []errors.Error{}
	// 	for k, v := range arg {
	// 		itm := errors.E(
	// 			errors.Target(k),
	// 			errors.Message(v),
	// 		)
	// 		ds = append(ds, *itm)
	// 	}

	// 	newErr := errors.Error{
	// 		Message: "bad request",
	// 		Errors:  ds,
	// 		Code:    errors.Code(http.StatusBadRequest),
	// 	}
	// 	c.JSON(http.StatusBadRequest, newErr)

	case *errors.Error:
		c.JSON(int(arg.Code), arg)

	case error:
		c.JSON(http.StatusInternalServerError, errors.Error{
			Code:    http.StatusInternalServerError,
			Message: arg.Error(),
		})
	}
}
