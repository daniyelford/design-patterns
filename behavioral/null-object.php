<?php
interface Logger {
    public function log(string $message): void;
}

class FileLogger implements Logger {
    public function log(string $message): void {
        echo "[FILE] $message\n";
    }
}

class NullLogger implements Logger {
    public function log(string $message): void {
        // هیچ کاری نمی‌کند
    }
}

class UserService {
    private Logger $logger;
    public function __construct(Logger $logger){
        $this->logger = $logger;
    }
    public function register(string $name){
        echo "Registering $name...\n";
        $this->logger->log("User $name registered");
    }
}

// استفاده
$service1 = new UserService(new FileLogger());
$service1->register("Ali");

$service2 = new UserService(new NullLogger());
$service2->register("Sara"); // بدون لاگ ولی بدون خطا
