package ValueObject

type Notification struct {
	Title   string
	Message string
}

func NewNotification(title string, message string) *Notification {
	return &Notification{
		Title:   title,
		Message: message,
	}
}
