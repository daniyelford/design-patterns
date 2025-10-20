<?php
interface ShippingStrategy {
    public function calculate(float $weight): float;
}

class AirShipping implements ShippingStrategy {
    public function calculate(float $weight): float { return 10 + $weight * 2; }
}
class SeaShipping implements ShippingStrategy {
    public function calculate(float $weight): float { return 5 + $weight * 0.5; }
}

class Order {
    private ShippingStrategy $strategy;
    public function __construct(ShippingStrategy $s) { $this->strategy = $s; }
    public function setStrategy(ShippingStrategy $s) { $this->strategy = $s; }
    public function shippingCost(float $weight): float { return $this->strategy->calculate($weight); }
}
$order = new Order(new AirShipping());
echo $order->shippingCost(3.5);
$order->setStrategy(new SeaShipping());
echo $order->shippingCost(3.5);
