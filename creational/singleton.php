<?php
// جلوگیری از ساخت شیء جدید
class Database
{
    private static ?Database $instance = null;
    private string $connection;
    private function __construct()
    {
        $this->connection = "Connected to database";
    }

    // جلوگیری از clone
    private function __clone() {}

    // جلوگیری از unserialize
    public function __wakeup()
    {
        throw new \Exception("Cannot unserialize a singleton.");
    }

    public static function getInstance(): Database
    {
        if (self::$instance === null) {
            self::$instance = new self();
        }

        return self::$instance;
    }

    public function getConnection(): string
    {
        return $this->connection;
    }
}


$db1 = Database::getInstance();
$db2 = Database::getInstance();

echo $db1->getConnection() . PHP_EOL;


