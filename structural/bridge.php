<?php

// Implementor
interface Color {
    public function applyColor(): string;
}

// Concrete Implementor
class RedColor implements Color {
    public function applyColor(): string { return "Red"; }
}

class BlueColor implements Color {
    public function applyColor(): string { return "Blue"; }
}

// Abstraction
abstract class Shape {
    protected Color $color;

    public function __construct(Color $color) {
        $this->color = $color;
    }

    abstract public function draw(): string;
}

// Refined Abstraction
class Circle extends Shape {
    public function draw(): string {
        return "Circle with color " . $this->color->applyColor();
    }
}

class Rectangle extends Shape {
    public function draw(): string {
        return "Rectangle with color " . $this->color->applyColor();
    }
}

// استفاده
$redCircle = new Circle(new RedColor());
$blueRect = new Rectangle(new BlueColor());

echo $redCircle->draw() . PHP_EOL;
echo $blueRect->draw() . PHP_EOL;


//Circle with color Red
//Rectangle with color Blue
