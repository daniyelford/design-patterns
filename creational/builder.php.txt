<?php

class Report
{
    public string $title;
    public string $author;
    public string $content;
    public string $date;

    public function show()
    {
        echo "📄 {$this->title}\n";
        echo "✍️ Author: {$this->author}\n";
        echo "📅 Date: {$this->date}\n";
        echo "---------------------------\n";
        echo "{$this->content}\n";
    }
}

// Builder
class ReportBuilder
{
    private Report $report;

    public function __construct()
    {
        $this->report = new Report();
    }

    public function setTitle(string $title): self
    {
        $this->report->title = $title;
        return $this;
    }

    public function setAuthor(string $author): self
    {
        $this->report->author = $author;
        return $this;
    }

    public function setDate(string $date): self
    {
        $this->report->date = $date;
        return $this;
    }

    public function setContent(string $content): self
    {
        $this->report->content = $content;
        return $this;
    }

    public function build(): Report
    {
        return $this->report;
    }
}

class ReportDirector
{
    public function buildSimpleReport(ReportBuilder $builder): Report
    {
        return $builder
            ->setTitle("Daily Report")
            ->setAuthor("System")
            ->setDate(date('Y-m-d'))
            ->setContent("All systems operational.")
            ->build();
    }
}

// استفاده
$builder = new ReportBuilder();
$director = new ReportDirector();

$report = $director->buildSimpleReport($builder);
$report->show();
