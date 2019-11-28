package ctxerr

//Fields similar to logrus.Fields for entry
type Fields map[string]interface{}

//Entry can be used to add custom additional info to your errors
//when you don't want to create a custom error every time
//it implements standard error interface and can be unwrapped to get original error
type Error struct {
	err    error
	fields Fields
}

//New is the preferrable way to create an Error
func New(err error) *Error {
	return &Error{
		err:    err,
		fields: Fields{},
	}
}

//WithField used to add custom field to an error
func (e Error) WithField(key string, value interface{}) *Error {
	return e.WithFields(Fields{key: value})
}

//WithFields used to add custom fields to an error
func (e Error) WithFields(fields Fields) *Error {
	res := New(e.err)
	for key, value := range e.fields {
		res.fields[key] = value
	}
	for key, value := range fields {
		res.fields[key] = value
	}
	return res
}

//GetField allows you to retrieve the field from error or check if the field exists (by checking for nil)
func (e Error) GetField(key string) interface{} {
	if value, ok := e.fields[key]; ok {
		return value
	}
	return nil
}

//Error to implement standard error interface
//Consider including fields here as well
func (e Error) Error() string {
	return e.err.Error()
}

//Unwrap to be able to get original error
func (e Error) Unwrap() error {
	return e.err
}

var _ error = Error{} //This will fail in case Error doesn't implement the standard error interface
