<?php
interface Observer { public function update(string $event, $data): void; }
class EmailNotifier implements Observer {
    public function update(string $event, $data): void { echo "Email: $event\n"; }
}
class Logger implements Observer {
    public function update(string $event, $data): void { echo "Log: $event\n"; }
}

class Subject {
    private array $observers = [];
    public function attach(Observer $o){ $this->observers[] = $o; }
    public function notify(string $event, $data = null){
        foreach($this->observers as $o) $o->update($event, $data);
    }
}

// استفاده
$s = new Subject();
$s->attach(new EmailNotifier());
$s->attach(new Logger());
$s->notify("order.created", ["id"=>123]);
