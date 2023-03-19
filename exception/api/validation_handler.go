package api

import (
	"database/sql"
	"golang-jwt/helper"
	"reflect"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

type Validation struct {
	DB *sql.DB
}

func NewValidation(DB *sql.DB) *Validation {
	return &Validation{
		DB: DB,
	}
}

func (v *Validation) Init() (*validator.Validate, ut.Translator) {
	translator := en.New()
	uni := ut.New(translator, translator)

	trans, _ := uni.GetTranslator("en")
	validate := validator.New()
	en_translations.RegisterDefaultTranslations(validate, trans)

	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := field.Tag.Get("label")
		return name
	})

	validate.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} harus diisi", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())
		return t

	})

	validate.RegisterTranslation("isunique", trans, func(ut ut.Translator) error {
		return ut.Add("isunique", "{0} tidak boleh sama", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("isunique", fe.Field())
		return t

	})
	validate.RegisterValidation("isunique", func(fl validator.FieldLevel) bool {
		params := fl.Param()
		split_params := strings.Split(params, "-")

		tableName := split_params[0]
		fieldName := split_params[1]
		fieldValue := fl.Field().String()

		return v.checkIsUnique(tableName, fieldName, fieldValue)
	})

	return validate, trans
}

func (v *Validation) Struct(s interface{}) interface{} {
	validate, trans := v.Init()
	vErrors := make(map[string]interface{})
	err := validate.Struct(s)

	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			vErrors[e.StructField()] = e.Translate(trans)
		}
	}

	if len(vErrors) > 0 {
		return vErrors
	}

	return nil
}

func (v *Validation) checkIsUnique(table_name string, field_name string, field_value string) bool {
	tx, err := v.DB.Begin()
	helper.SetPanicError(err)

	SQL := "SELECT " + field_name + " FROM " + table_name + " WHERE " + field_name + " = ?"
	rows, err := tx.Query(
		SQL,
		field_value,
	)

	helper.SetPanicError(err)
	defer rows.Close()

	result := field_name
	if rows.Next() {
		err := rows.Scan(&result)
		helper.SetPanicError(err)
	}

	return result != field_value
}
