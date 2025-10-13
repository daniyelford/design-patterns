<?php

// رابط مشترک برای همه‌ی ارسال‌کننده‌ها
interface MessageSender {
    public function send(string $to, string $message): void;
}

// پیاده‌سازی SMS
class SmsSender implements MessageSender {
    public function send(string $to, string $message): void {
        echo "Sending SMS to $to: $message\n";
    }
}

// پیاده‌سازی Email
class EmailSender implements MessageSender {
    public function send(string $to, string $message): void {
        echo "Sending Email to $to: $message\n";
    }
}

// کلاس Factory
class MessageFactory {
    public static function create(string $type): MessageSender {
        return match ($type) {
            'sms' => new SmsSender(),
            'email' => new EmailSender(),
            default => throw new Exception("Unknown sender type: $type"),
        };
    }
}

// تست
$sender1 = MessageFactory::create('sms');
$sender1->send('09123456789', 'سلام از SMS!');

$sender2 = MessageFactory::create('email');
$sender2->send('test@example.com', 'سلام از ایمیل!');


// Sending SMS to 09120001122: سلام از SMS!
// Sending Email to test@example.com: سلام از ایمیل!
