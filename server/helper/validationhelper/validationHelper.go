package validationhelper

import (
	validator "gopkg.in/go-playground/validator.v9"
)

//Validate method
func Validate(s interface{}) map[string]string {
	var validate *validator.Validate

	validate = validator.New()

	//For custom error message
	// validate.RegisterTranslation("required", trans, func(ut ut.Translator) error {
	// 	return ut.Add("required", "{0} must have a value!", true) // see universal-translator for details

	// }, func(ut ut.Translator, fe validator.FieldError) string {
	// 	t, _ := ut.T("required", fe.Field())

	// 	return t
	// })

	err := validate.Struct(s)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		customErrs := make(map[string]string, len(errs))

		for _, e := range errs {
			// can translate each error one at a time.
			customErrs[e.Namespace()] = e.Translate(trans)
		}
		return customErrs
	}
	return nil
}
