package err

type Err string

func (er Err) Error() string {
	return string(er)
}
