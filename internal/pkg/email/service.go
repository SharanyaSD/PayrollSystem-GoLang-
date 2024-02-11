package email

//send grid - email service provider
import (
	"fmt"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type Service interface {
	SendEmail(to, subject, content string) error
}

type emailService struct {
	apiKey string //api key for sendGrid/other email service provider
}

func NewEmailService(apiKey string) Service {
	return &emailService{
		apiKey: apiKey,
	}
}

func (es *emailService) SendEmail(to, subject, content string) error {
	from := mail.NewEmail("Payroll ", "sharanyadatrange1@gmail.com")
	toEmail := mail.NewEmail("", to)
	plainTextContent := content
	htmlContent := content
	message := mail.NewSingleEmail(from, subject, toEmail, plainTextContent, htmlContent)

	client := sendgrid.NewSendClient(es.apiKey)
	response, err := client.Send(message)
	if err != nil {
		return err
	}

	if response.StatusCode >= 200 && response.StatusCode < 300 {
		return nil
	} else {
		return fmt.Errorf("Failed to send email, status code : %d", response.StatusCode)
	}

	return nil
}
