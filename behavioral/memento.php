<?php
class Editor {
    private string $text = "";
    public function type(string $words){ $this->text .= " " . $words; }
    public function getText(): string { return $this->text; }
    public function save(): Memento { return new Memento($this->text); }
    public function restore(Memento $m){ $this->text = $m->getText(); }
}

class Memento {
    private string $text;
    public function __construct(string $t){ $this->text = $t; }
    public function getText(): string { return $this->text; }
}

// تست
$e = new Editor();
$e->type("salam");
$saved = $e->save();
$e->type("chetori?");
echo $e->getText()."\n";   // salam chetori?
$e->restore($saved);
echo $e->getText()."\n";   // salam
