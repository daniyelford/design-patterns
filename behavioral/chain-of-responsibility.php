<?php
abstract class Handler {
    protected ?Handler $next = null;
    public function setNext(Handler $h){ $this->next = $h; return $h; }
    abstract public function handle(int $amount);
}
class LowHandler extends Handler {
    public function handle(int $amount){ 
        if($amount <= 100) echo "Low handled\n";
        elseif($this->next) $this->next->handle($amount);
    }
}
class HighHandler extends Handler {
    public function handle(int $amount){
        if($amount > 100) echo "High handled\n";
        elseif($this->next) $this->next->handle($amount);
    }
}

$h1 = new LowHandler();
$h2 = new HighHandler();
$h1->setNext($h2);
$h1->handle(150);
