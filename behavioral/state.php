<?php
interface State { public function doAction(Context $c): void; }
class Context { public State $state; public function __construct(State $s){ $this->state = $s; } }
class StartState implements State {
    public function doAction(Context $c): void { echo "Start\n"; $c->state = new StopState(); }
}
class StopState implements State {
    public function doAction(Context $c): void { echo "Stop\n"; }
}

$c = new Context(new StartState());
$c->state->doAction($c); // Start -> switches to Stop
$c->state->doAction($c); // Stop
