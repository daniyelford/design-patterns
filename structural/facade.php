<?php

// Subsystems
class Light {
    public function on() { echo "Lights ON\n"; }
    public function off() { echo "Lights OFF\n"; }
}

class AC {
    public function on() { echo "AC ON\n"; }
    public function off() { echo "AC OFF\n"; }
}

class Security {
    public function arm() { echo "Security Armed\n"; }
    public function disarm() { echo "Security Disarmed\n"; }
}

// Facade
class SmartHomeFacade {
    private Light $light;
    private AC $ac;
    private Security $security;

    public function __construct() {
        $this->light = new Light();
        $this->ac = new AC();
        $this->security = new Security();
    }

    public function leaveHome() {
        $this->light->off();
        $this->ac->off();
        $this->security->arm();
    }

    public function arriveHome() {
        $this->security->disarm();
        $this->light->on();
        $this->ac->on();
    }
}

// استفاده
$home = new SmartHomeFacade();
$home->leaveHome();
$home->arriveHome();


//Lights OFF
//AC OFF
//Security Armed
//Security Disarmed
//Lights ON
//AC ON
