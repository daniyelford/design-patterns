<?php
interface Command { public function execute(): void; }
class Light {
    public function on(){ echo "Light on\n"; }
    public function off(){ echo "Light off\n"; }
}
class LightOnCommand implements Command {
    private Light $light;
    public function __construct(Light $l){ $this->light = $l; }
    public function execute(): void { $this->light->on(); }
}
class Remote {
    public function submit(Command $c){ $c->execute(); }
}

// استفاده
$light = new Light();
$cmd = new LightOnCommand($light);
$remote = new Remote();
$remote->submit($cmd);
