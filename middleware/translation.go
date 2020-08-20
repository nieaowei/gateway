package middleware

import (
	"gateway/public"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	en_translations "gopkg.in/go-playground/validator.v9/translations/en"
	zh_translations "gopkg.in/go-playground/validator.v9/translations/zh"
	"reflect"
	"regexp"
	"strings"
)

//设置Translation
func TranslationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//参照：https://github.com/go-playground/validator/blob/v9/_examples/translations/main.go

		//设置支持语言
		en := en.New()
		zh := zh.New()

		//设置国际化翻译器
		uni := ut.New(zh, zh, en)
		val := validator.New()

		//根据参数取翻译器实例
		locale := c.DefaultQuery("locale", "zh")
		trans, _ := uni.GetTranslator(locale)

		//翻译器注册到validator
		switch locale {
		case "en":
			en_translations.RegisterDefaultTranslations(val, trans)
			val.RegisterTagNameFunc(func(fld reflect.StructField) string {
				return fld.Tag.Get("en_comment")
			})
			break
		default:
			zh_translations.RegisterDefaultTranslations(val, trans)
			val.RegisterTagNameFunc(func(fld reflect.StructField) string {
				return fld.Tag.Get("comment")
			})

			//自定义验证方法
			//https://github.com/go-playground/validator/blob/v9/_examples/custom-validation/main.go
			val.RegisterValidation("is_valid_username", func(fl validator.FieldLevel) bool {
				return true
			})

			//自定义翻译器
			//https://github.com/go-playground/validator/blob/v9/_examples/translations/main.go
			val.RegisterTranslation("is_valid_username", trans, func(ut ut.Translator) error {
				return ut.Add("is_valid_username", "{0} 填写不正确哦", true)
			}, func(ut ut.Translator, fe validator.FieldError) string {
				t, _ := ut.T("is_valid_username", fe.Field())
				return t
			})

			//自定义验证方法
			//https://github.com/go-playground/validator/blob/v9/_examples/custom-validation/main.go
			val.RegisterValidation("is_valid_service_name", func(fl validator.FieldLevel) bool {
				matched, _ := regexp.MatchString("[a-zA-Z0-9_]{6,128}", fl.Field().String())
				return matched
			})

			//自定义翻译器
			//https://github.com/go-playground/validator/blob/v9/_examples/translations/main.go
			val.RegisterTranslation("is_valid_service_name", trans, func(ut ut.Translator) error {
				return ut.Add("is_valid_service_name", "{0} 不符合输入格式", true)
			}, func(ut ut.Translator, fe validator.FieldError) string {
				t, _ := ut.T("is_valid_service_name", fe.Field())
				return t
			})

			val.RegisterValidation("valid_url_rewrite", func(fl validator.FieldLevel) bool {
				if fl.Field().String() == "" {
					return true
				}
				data := strings.Split(fl.Field().String(), "\n")
				for _, datum := range data {
					if datum == "" {
						continue
					}
					if len(strings.Split(datum, " ")) != 2 {
						return false
					}
				}
				return true
			})

			//自定义翻译器
			//https://github.com/go-playground/validator/blob/v9/_examples/translations/main.go
			val.RegisterTranslation("valid_url_rewrite", trans, func(ut ut.Translator) error {
				return ut.Add("valid_url_rewrite", "{0} 输入格式错误", true)
			}, func(ut ut.Translator, fe validator.FieldError) string {
				t, _ := ut.T("valid_url_rewrite", fe.Field())
				return t
			})

			val.RegisterValidation("valid_header_transform", func(fl validator.FieldLevel) bool {
				if fl.Field().String() == "" {
					return true
				}
				data := strings.Split(fl.Field().String(), "\n")
				for _, datam := range data {
					if datam == "" {
						continue
					}
					if len(strings.Split(datam, " ")) != 3 {
						return false
					}
				}
				return true
			})

			//自定义翻译器
			//https://github.com/go-playground/validator/blob/v9/_examples/translations/main.go
			val.RegisterTranslation("valid_header_transform", trans, func(ut ut.Translator) error {
				return ut.Add("valid_header_transform", "{0} 输入格式错误", true)
			}, func(ut ut.Translator, fe validator.FieldError) string {
				t, _ := ut.T("valid_header_transform", fe.Field())
				return t
			})

			val.RegisterValidation("valid_ip_list", func(fl validator.FieldLevel) bool {
				if fl.Field().String() == "" {
					return true
				}
				data := strings.Split(fl.Field().String(), "\n")
				for _, datum := range data {
					if datum == "" {
						continue
					}
					matched, _ := regexp.Match("^\\S+:\\d+$", []byte(datum))
					if !matched {
						return false
					}
				}
				return true
			})

			//自定义翻译器
			//https://github.com/go-playground/validator/blob/v9/_examples/translations/main.go
			val.RegisterTranslation("valid_ip_list", trans, func(ut ut.Translator) error {
				return ut.Add("valid_ip_list", "{0} 输入格式错误", true)
			}, func(ut ut.Translator, fe validator.FieldError) string {
				t, _ := ut.T("valid_ip_list", fe.Field())
				return t
			})

			val.RegisterValidation("valid_weight_list", func(fl validator.FieldLevel) bool {
				if fl.Field().String() == "" {
					return true
				}
				data := strings.Split(fl.Field().String(), "\n")
				for _, datum := range data {
					if datum == "" {
						continue
					}
					matched, _ := regexp.Match("^\\d+$", []byte(datum))
					if !matched {
						return false
					}
				}
				return true
			})

			//自定义翻译器
			//https://github.com/go-playground/validator/blob/v9/_examples/translations/main.go
			val.RegisterTranslation("valid_weight_list", trans, func(ut ut.Translator) error {
				return ut.Add("valid_weight_list", "{0} 输入格式错误", true)
			}, func(ut ut.Translator, fe validator.FieldError) string {
				t, _ := ut.T("valid_weight_list", fe.Field())
				return t
			})
			break
		}
		c.Set(public.TranslatorKey, trans)
		c.Set(public.ValidatorKey, val)
		c.Next()
	}
}
