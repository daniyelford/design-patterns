class AutoSaver {
    private bool $saving = false;

    public function save(): void {
        if ($this->saving) {
            echo "⏳ در حال ذخیره، درخواست نادیده گرفته شد\n";
            return; // Balk!
        }

        $this->saving = true;
        echo "💾 در حال ذخیره...\n";
        sleep(2);
        $this->saving = false;
        echo "✅ ذخیره انجام شد\n";
    }
}

$saver = new AutoSaver();
$saver->save();
$saver->save(); // درخواست دوم نادیده گرفته می‌شود
