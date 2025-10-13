<?php

// Flyweight
class TreeType {
    private string $name;
    private string $color;

    public function __construct(string $name, string $color) {
        $this->name = $name;
        $this->color = $color;
    }

    public function draw(int $x, int $y) {
        echo "Drawing {$this->name} ({$this->color}) at ($x, $y)\n";
    }
}

// Factory برای اشتراک
class TreeFactory {
    private array $treeTypes = [];

    public function getTreeType(string $name, string $color): TreeType {
        $key = $name . $color;
        if (!isset($this->treeTypes[$key])) {
            $this->treeTypes[$key] = new TreeType($name, $color);
        }
        return $this->treeTypes[$key];
    }
}

// استفاده
$factory = new TreeFactory();

$tree1 = $factory->getTreeType("Oak", "Green");
$tree2 = $factory->getTreeType("Oak", "Green"); // همان شیء قبلی

$tree1->draw(10, 20);
$tree2->draw(15, 25);

//Drawing Oak (Green) at (10, 20)
//Drawing Oak (Green) at (15, 25)
