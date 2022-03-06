package request

type Response struct {
	Body []byte
	Code int
	Err  error
}

func (r *Response) Write(b []byte) (int, error) {
	r.Body = b
	return len(b), nil
}
