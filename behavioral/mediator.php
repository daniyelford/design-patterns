<?php
interface Mediator {
    public function notify(object $sender, string $event): void;
}

class ChatRoom implements Mediator {
    private array $users = [];
    public function register(User $u): void {
        $this->users[] = $u;
        $u->setMediator($this);
    }
    public function notify(object $sender, string $event): void {
        foreach ($this->users as $u) {
            if ($u !== $sender) {
                $u->receive($event);
            }
        }
    }
}

class User {
    private Mediator $mediator;
    public string $name;
    public function __construct(string $n){ $this->name = $n; }
    public function setMediator(Mediator $m){ $this->mediator = $m; }
    public function send(string $msg){ 
        echo "{$this->name} sends: $msg\n";
        $this->mediator->notify($this, $msg);
    }
    public function receive(string $msg){ echo "{$this->name} got: $msg\n"; }
}

// تست
$chat = new ChatRoom();
$u1 = new User("Ali"); $u2 = new User("Sara");
$chat->register($u1); $chat->register($u2);
$u1->send("سلام!");
