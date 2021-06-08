package src

type ServiceList interface {
	SendNotifcation(p []byte) (int, string, error)
	CheckBody(p []byte) (int, error)
	RequestApiSendNotification() (int, error)
}
