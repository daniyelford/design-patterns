<?php

interface Button {
    public function render(): string;
}

interface Checkbox {
    public function render(): string;
}

// پیاده‌سازی‌های Windows
class WindowsButton implements Button {
    public function render(): string {
        return "Render Windows Button";
    }
}

class WindowsCheckbox implements Checkbox {
    public function render(): string {
        return "Render Windows Checkbox";
    }
}

// پیاده‌سازی‌های Mac
class MacButton implements Button {
    public function render(): string {
        return "Render Mac Button";
    }
}

class MacCheckbox implements Checkbox {
    public function render(): string {
        return "Render Mac Checkbox";
    }
}

// کارخانه‌ی کلی (Abstract Factory)
interface GUIFactory {
    public function createButton(): Button;
    public function createCheckbox(): Checkbox;
}

// کارخانه‌ی Windows
class WindowsFactory implements GUIFactory {
    public function createButton(): Button {
        return new WindowsButton();
    }
    public function createCheckbox(): Checkbox {
        return new WindowsCheckbox();
    }
}

// کارخانه‌ی Mac
class MacFactory implements GUIFactory {
    public function createButton(): Button {
        return new MacButton();
    }
    public function createCheckbox(): Checkbox {
        return new MacCheckbox();
    }
}

// برنامه
function renderUI(GUIFactory $factory) {
    $button = $factory->createButton();
    $checkbox = $factory->createCheckbox();

    echo $button->render() . PHP_EOL;
    echo $checkbox->render() . PHP_EOL;
}

// تست
$os = 'mac'; // تغییر بده به 'windows'

if ($os === 'windows') {
    $factory = new WindowsFactory();
} else {
    $factory = new MacFactory();
}

renderUI($factory);
