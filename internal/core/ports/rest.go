package ports

type Rest interface {
	Post(url string, body any, out any) error
}
