<?php
interface Visitor {
    public function visitCircle(Circle $c);
    public function visitSquare(Square $s);
}

interface Shape {
    public function accept(Visitor $v);
}

class Circle implements Shape {
    public function accept(Visitor $v){ $v->visitCircle($this); }
}
class Square implements Shape {
    public function accept(Visitor $v){ $v->visitSquare($this); }
}

class AreaVisitor implements Visitor {
    public function visitCircle(Circle $c){ echo "Circle area calculated\n"; }
    public function visitSquare(Square $s){ echo "Square area calculated\n"; }
}

// استفاده
$shapes = [new Circle(), new Square()];
$v = new AreaVisitor();
foreach($shapes as $s){ $s->accept($v); }
