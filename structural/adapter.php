<?php

// کلاس قدیمی
class OldPrinter {
    public function printText(string $text) {
        echo "Printing: $text\n";
    }
}

// رابط مورد انتظار
interface Printer {
    public function print(string $text);
}

// Adapter
class PrinterAdapter implements Printer {
    private OldPrinter $oldPrinter;

    public function __construct(OldPrinter $oldPrinter) {
        $this->oldPrinter = $oldPrinter;
    }

    public function print(string $text) {
        $this->oldPrinter->printText($text);
    }
}

// استفاده
$oldPrinter = new OldPrinter();
$adapter = new PrinterAdapter($oldPrinter);
$adapter->print("Hello Adapter!");


// Printing: Hello Adapter!
