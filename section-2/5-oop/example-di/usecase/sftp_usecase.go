package usecase

import "fmt"

// Class SmtpSenderUsecase
type SmtpSenderUsecase struct{}

func NewSmtpSenderUsecase() *SmtpSenderUsecase {
	return &SmtpSenderUsecase{}
}

// Interface untuk mengirim email
type ISmtpSenderUsecase interface {
	KirimEmail(penerima, pesan string)
}

func (s SmtpSenderUsecase) KirimEmail(penerima, pesan string) {
	fmt.Printf("[SMTP] Mengirim email ke %s: %s\n", penerima, pesan)
}
