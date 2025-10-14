package main

import (
    "fmt"
    "time"
)

// Product
type Report struct {
    Title   string
    Author  string
    Content string
    Date    string
}

func (r Report) Show() {
    fmt.Println("📄", r.Title)
    fmt.Println("✍️ Author:", r.Author)
    fmt.Println("📅 Date:", r.Date)
    fmt.Println("---------------------------")
    fmt.Println(r.Content)
}

// Builder
type ReportBuilder struct {
    report Report
}

func NewReportBuilder() *ReportBuilder {
    return &ReportBuilder{}
}

func (b *ReportBuilder) SetTitle(title string) *ReportBuilder {
    b.report.Title = title
    return b
}

func (b *ReportBuilder) SetAuthor(author string) *ReportBuilder {
    b.report.Author = author
    return b
}

func (b *ReportBuilder) SetDate(date string) *ReportBuilder {
    b.report.Date = date
    return b
}

func (b *ReportBuilder) SetContent(content string) *ReportBuilder {
    b.report.Content = content
    return b
}

func (b *ReportBuilder) Build() Report {
    return b.report
}

// Director
type ReportDirector struct{}

func (d ReportDirector) BuildSimpleReport(builder *ReportBuilder) Report {
    return builder.
        SetTitle("Daily Report").
        SetAuthor("System").
        SetDate(time.Now().Format("2006-01-02")).
        SetContent("All systems operational.").
        Build()
}

func main() {
    builder := NewReportBuilder()
    director := ReportDirector{}

    report := director.BuildSimpleReport(builder)
    report.Show()
}
