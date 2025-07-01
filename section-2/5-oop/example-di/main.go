package main

import (
	"github.com/likuisa13/learn-golang/section-2/5-oop/example-di/usecase"
)

// Contoh Dependency Injection
func main() {
	senderUsecase := usecase.NewSmtpSenderUsecase()
	userUsecase := usecase.NewUserUsecase(senderUsecase)

	userUsecase.RegistrasiUser("dwiki@mail.com")
}
