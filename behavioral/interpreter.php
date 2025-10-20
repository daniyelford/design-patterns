<?php
interface Expression {
    public function interpret(array $context): int;
}

class Number implements Expression {
    public function __construct(private int $number){}
    public function interpret(array $context): int { return $this->number; }
}

class Plus implements Expression {
    public function __construct(private Expression $left, private Expression $right){}
    public function interpret(array $context): int {
        return $this->left->interpret($context) + $this->right->interpret($context);
    }
}

// (5 + 10)
$expr = new Plus(new Number(5), new Number(10));
echo $expr->interpret([]); // 15
