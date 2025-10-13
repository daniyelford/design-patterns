<?php

// Component
interface Text {
    public function render(): string;
}

// Concrete Component
class PlainText implements Text {
    private string $text;
    public function __construct(string $text) { $this->text = $text; }
    public function render(): string { return $this->text; }
}

// Decorator
abstract class TextDecorator implements Text {
    protected Text $wrapped;
    public function __construct(Text $wrapped) { $this->wrapped = $wrapped; }
}

// Concrete Decorators
class BoldText extends TextDecorator {
    public function render(): string {
        return "<b>" . $this->wrapped->render() . "</b>";
    }
}

class ItalicText extends TextDecorator {
    public function render(): string {
        return "<i>" . $this->wrapped->render() . "</i>";
    }
}

// استفاده
$text = new PlainText("Hello Decorator");
$text = new BoldText($text);
$text = new ItalicText($text);

echo $text->render();

//<i><b>Hello Decorator</b></i>
