package service

type MailService interface {
	SendEmailConfirmation(userID, callbackURL string) error
	SendResetPassword(userID, callbackURL string) error
	SendNotification(userID, notificationContent string) error
	Subscribe(email, userID string) error
	Unsubscribe(email string) error
}