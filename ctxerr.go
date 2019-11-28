package ctxerr

type Fields map[string]interface{}

type Error struct {
	err error
	fields Fields
}

func New(err error) *Error {
	return &Error {
		err: err,
		fields: Fields{},
	}
}

func (e Error) WithField(key string, value interface{}) *Error {
	return e.WithFields(Fields{key:value})
}

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

func (e Error) GetField(key string) interface{} {
	if value, ok := e.fields[key]; ok {
		return value
	}
	return nil
}

func (e Error) Error() string {
	return e.err.Error()
}

func (e Error) Unwrap() error {
	return e.err
}
