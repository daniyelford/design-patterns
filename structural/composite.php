<?php

// Component
interface FileSystem {
    public function show(): void;
}

// Leaf
class File implements FileSystem {
    private string $name;

    public function __construct(string $name) {
        $this->name = $name;
    }

    public function show(): void {
        echo "File: " . $this->name . PHP_EOL;
    }
}

// Composite
class Folder implements FileSystem {
    private string $name;
    private array $children = [];

    public function __construct(string $name) {
        $this->name = $name;
    }

    public function add(FileSystem $child): void {
        $this->children[] = $child;
    }

    public function show(): void {
        echo "Folder: " . $this->name . PHP_EOL;
        foreach ($this->children as $child) {
            $child->show();
        }
    }
}

// استفاده
$file1 = new File("file1.txt");
$file2 = new File("file2.txt");

$folder1 = new Folder("Documents");
$folder1->add($file1);
$folder1->add($file2);

$folder1->show();

//Folder: Documents
//File: file1.txt
//File: file2.txt
