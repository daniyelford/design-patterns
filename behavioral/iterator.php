<?php
class MyCollection implements Iterator {
    private array $items; private int $i = 0;
    public function __construct(array $items){ $this->items = $items; }
    public function current(){ return $this->items[$this->i]; }
    public function key(){ return $this->i; }
    public function next(){ $this->i++; }
    public function rewind(){ $this->i = 0; }
    public function valid(){ return isset($this->items[$this->i]); }
}
$col = new MyCollection([1,2,3]);
foreach($col as $k=>$v) echo "$k:$v\n";
