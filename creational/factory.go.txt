package main

import (
    "fmt"
)

// رابط
type MessageSender interface {
    Send(to string, message string)
}

// SMS پیاده‌سازی
type SmsSender struct{}
func (s SmsSender) Send(to string, message string) {
    fmt.Println("Sending SMS to", to+":", message)
}

// Email پیاده‌سازی
type EmailSender struct{}
func (e EmailSender) Send(to string, message string) {
    fmt.Println("Sending Email to", to+":", message)
}

// کارخانه
func CreateSender(t string) (MessageSender, error) {
    switch t {
    case "sms":
        return SmsSender{}, nil
    case "email":
        return EmailSender{}, nil
    default:
        return nil, fmt.Errorf("unknown sender type: %s", t)
    }
}

func main() {
    smsSender, _ := CreateSender("sms")
    smsSender.Send("09120001122", "سلام از SMS!")

    emailSender, _ := CreateSender("email")
    emailSender.Send("test@example.com", "سلام از ایمیل!")
}
