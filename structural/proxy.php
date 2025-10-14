<?php

// Subject
interface Image {
    public function display(): void;
}

// RealSubject
class RealImage implements Image {
    private string $filename;

    public function __construct(string $filename) {
        $this->filename = $filename;
        $this->loadFromDisk();
    }

    private function loadFromDisk() {
        echo "Loading " . $this->filename . "\n";
    }

    public function display(): void {
        echo "Displaying " . $this->filename . "\n";
    }
}

// Proxy
class ProxyImage implements Image {
    private string $filename;
    private ?RealImage $realImage = null;

    public function __construct(string $filename) {
        $this->filename = $filename;
    }

    public function display(): void {
        if ($this->realImage === null) {
            $this->realImage = new RealImage($this->filename);
        }
        $this->realImage->display();
    }
}

// استفاده
$image = new ProxyImage("photo.jpg");

// تصویر فقط وقتی نمایش داده می‌شود بارگذاری می‌شود
$image->display();
$image->display();


//Loading photo.jpg
//Displaying photo.jpg
//Displaying photo.jpg
