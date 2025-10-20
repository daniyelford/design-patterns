<?php
abstract class Game {
    final public function play() {
        $this->start(); $this->playTurn(); $this->end();
    }
    abstract protected function start();
    abstract protected function playTurn();
    protected function end(){ echo "Game over\n"; }
}
class Chess extends Game {
    protected function start(){ echo "Chess starts\n"; }
    protected function playTurn(){ echo "Move\n"; }
}

$g = new Chess();
$g->play();
