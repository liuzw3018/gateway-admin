package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/universal-translator"
	"github.com/liuzw3018/gateway_admin/public"
	"gopkg.in/go-playground/validator.v9"
	en_translations "gopkg.in/go-playground/validator.v9/translations/en"
	zh_translations "gopkg.in/go-playground/validator.v9/translations/zh"
	"reflect"
	"regexp"
	"strings"
)

// TranslationMiddleware 设置Translation
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
			val.RegisterValidation("validUsername", func(fl validator.FieldLevel) bool {
				return fl.Field().String() == "admin"
			})

			val.RegisterValidation("validServiceName", func(fl validator.FieldLevel) bool {
				match, _ := regexp.Match(`^[a-zA-Z0-9_]{6,128}$`, []byte(fl.Field().String()))
				return match
			})
			val.RegisterValidation("validRule", func(fl validator.FieldLevel) bool {
				//return fl.Field().String() == "admin"
				match, _ := regexp.Match(`^\S+$`, []byte(fl.Field().String()))
				return match
			})

			val.RegisterValidation("validURLRewrite", func(fl validator.FieldLevel) bool {
				if fl.Field().String() == "" {
					return true
				}
				for _, ms := range strings.Split(fl.Field().String(), ",") {
					if len(strings.Split(ms, " ")) != 2 {
						return false
					}
				}
				return true
			})
			val.RegisterValidation("validHeaderTransfor", func(fl validator.FieldLevel) bool {
				if fl.Field().String() == "" {
					return true
				}
				for _, ms := range strings.Split(fl.Field().String(), ",") {
					if len(strings.Split(ms, " ")) != 3 {
						return false
					}
				}
				return true
			})

			val.RegisterValidation("validIpList", func(fl validator.FieldLevel) bool {
				if fl.Field().String() == "" {
					return true
				}
				for _, item := range strings.Split(fl.Field().String(), ",") {
					matched, _ := regexp.Match(`\S+`, []byte(item)) //ip_addr
					if !matched {
						return false
					}
				}
				return true
			})
			val.RegisterValidation("validWeightList", func(fl validator.FieldLevel) bool {
				for _, ms := range strings.Split(fl.Field().String(), ",") {
					if match, _ := regexp.Match(`^\d+$`, []byte(ms)); !match {
						return false
					}
				}
				return true
			})
			val.RegisterValidation("validIpPortList", func(fl validator.FieldLevel) bool {
				for _, ms := range strings.Split(fl.Field().String(), ",") {
					if matched, _ := regexp.Match(`^\S+:\d+$`, []byte(ms)); !matched {
						return false
					}
				}
				return true
			})

			//自定义翻译器
			//https://github.com/go-playground/validator/blob/v9/_examples/translations/main.go
			val.RegisterTranslation("validUsername", trans, func(ut ut.Translator) error {
				return ut.Add("validUsername", "{0} 填写不正确", true)
			}, func(ut ut.Translator, fe validator.FieldError) string {
				t, _ := ut.T("validUsername", fe.Field())
				return t
			})

			val.RegisterTranslation("validServiceName", trans, func(ut ut.Translator) error {
				return ut.Add("validServiceName", "{0} 不符合输入格式", true)
			}, func(ut ut.Translator, fe validator.FieldError) string {
				t, _ := ut.T("validServiceName", fe.Field())
				return t
			})
			val.RegisterTranslation("validRule", trans, func(ut ut.Translator) error {
				return ut.Add("validRule", "{0} 必须是一个非空字符", true)
			}, func(ut ut.Translator, fe validator.FieldError) string {
				t, _ := ut.T("validRule", fe.Field())
				return t
			})

			val.RegisterTranslation("validURLRewrite", trans, func(ut ut.Translator) error {
				return ut.Add("validURLRewrite", "{0} 不符合输入格式", true)
			}, func(ut ut.Translator, fe validator.FieldError) string {
				t, _ := ut.T("validURLRewrite", fe.Field())
				return t
			})
			val.RegisterTranslation("validHeaderTransfor", trans, func(ut ut.Translator) error {
				return ut.Add("validHeaderTransfor", "{0} 不符合输入格式", true)
			}, func(ut ut.Translator, fe validator.FieldError) string {
				t, _ := ut.T("validHeaderTransfor", fe.Field())
				return t
			})

			val.RegisterTranslation("validIpList", trans, func(ut ut.Translator) error {
				return ut.Add("validIpList", "{0} 不符合输入格式", true)
			}, func(ut ut.Translator, fe validator.FieldError) string {
				t, _ := ut.T("validIpList", fe.Field())
				return t
			})
			val.RegisterTranslation("validWeightList", trans, func(ut ut.Translator) error {
				return ut.Add("validWeightList", "{0} 不符合输入格式", true)
			}, func(ut ut.Translator, fe validator.FieldError) string {
				t, _ := ut.T("validWeightList", fe.Field())
				return t
			})
			val.RegisterTranslation("validIpPortList", trans, func(ut ut.Translator) error {
				return ut.Add("validIpPortList", "{0} 不符合输入格式", true)
			}, func(ut ut.Translator, fe validator.FieldError) string {
				t, _ := ut.T("validIpPortList", fe.Field())
				return t
			})

			break
		}
		c.Set(public.TranslatorKey, trans)
		c.Set(public.ValidatorKey, val)
		c.Next()
	}
}
