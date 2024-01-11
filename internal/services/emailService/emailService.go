package emailService

import (
	"async-queue/domain"
	"async-queue/dto"
	"async-queue/internal/config"
	"context"
	"encoding/json"
	"log"
	"net/smtp"

	"github.com/hibiken/asynq"
)

type emailService struct {
	cnf *config.Config
}

func NewEmailService(cnf *config.Config) domain.EmailService {
	return &emailService{
		cnf: cnf,
	}
}

// Send implements domain.EmailService.
func (svc *emailService) Send(to string, subject string, body string) error {
	auth := smtp.PlainAuth("", svc.cnf.Email.User, svc.cnf.Email.Pass, svc.cnf.Email.Host)
	msg := []byte("From: fiangumilar <" + svc.cnf.Email.User + ">\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n" +
		body)

	if err := smtp.SendMail(svc.cnf.Email.Host+":"+svc.cnf.Email.Port, auth, svc.cnf.Email.User, []string{to}, msg); err != nil {
		log.Printf("[ERR]: error sending =email: %v", err.Error())
		return err
	}

	return nil
}

// SendEmailQueue implements domain.EmailService.
func (svc *emailService) SendEmailQueue() (string, func(context.Context, *asynq.Task) error) {
	return "send:email", func(ctx context.Context, task *asynq.Task) error {
		var data dto.EmailSendReq
		_ = json.Unmarshal(task.Payload(), &data)

		log.Printf("Execute send email: %v", data.To)
		return svc.Send(data.To, data.Subject, data.Body)
	}

}
