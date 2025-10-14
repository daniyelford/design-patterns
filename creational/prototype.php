<?php

// Prototype interface
interface Prototype
{
    public function clone(): Prototype;
}

// کلاس اصلی
class User implements Prototype
{
    public string $name;
    public string $role;
    public array $permissions;

    public function __construct(string $name, string $role, array $permissions)
    {
        $this->name = $name;
        $this->role = $role;
        $this->permissions = $permissions;
    }

    public function clone(): Prototype
    {
        // Clone عمیق (deep copy)
        return clone $this;
    }

    public function show()
    {
        echo "👤 {$this->name} ({$this->role})\n";
        echo "🔑 Permissions: " . implode(', ', $this->permissions) . "\n\n";
    }
}

// نمونه‌ی اولیه
$admin = new User("Admin", "Administrator", ["create", "edit", "delete"]);
$admin->show();

// ساخت کاربر جدید با کپی
$user1 = $admin->clone();
$user1->name = "Daniyal";
$user1->role = "Editor";
$user1->permissions[] = "publish";

$user1->show();
$admin->show(); // هنوز دست‌نخورده‌ست


//👤 Admin (Administrator)
//🔑 Permissions: create, edit, delete

//👤 Daniyal (Editor)
//🔑 Permissions: create, edit, delete, publish

//👤 Admin (Administrator)
//🔑 Permissions: create, edit, delete
