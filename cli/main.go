package main

import (
	"embed"
	"flag"
	"fmt"
	iofs "io/fs"
	"os"

	"knack/internal/decisions"
	"knack/internal/glossary"
	"knack/internal/instructions"
	"knack/internal/queue"
	"knack/internal/skills"
	"knack/internal/status"
)

//go:embed all:embedded/skills
var embeddedSkills embed.FS

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}
	switch os.Args[1] {
	case "skills":
		skillsCmd(os.Args[2:])
	case "validate":
		validateCmd(os.Args[2:])
	case "decisions":
		decisionsCmd(os.Args[2:])
	case "status":
		statusCmd(os.Args[2:])
	case "glossary":
		glossaryCmd(os.Args[2:])
	case "instructions":
		instructionsCmd(os.Args[2:])
	default:
		fmt.Fprintf(os.Stderr, "unknown command: %s\n", os.Args[1])
		usage()
		os.Exit(1)
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: knack <command> [args]\n")
	fmt.Fprintln(os.Stderr, "commands:")
	fmt.Fprintln(os.Stderr, "  skills init [--target DIR]")
	fmt.Fprintln(os.Stderr, "  skills check [--dir DIR]")
	fmt.Fprintln(os.Stderr, "  validate <queue-file>")
	fmt.Fprintln(os.Stderr, "  decisions list|show NNNN|check")
	fmt.Fprintln(os.Stderr, "  status")
	fmt.Fprintln(os.Stderr, "  glossary check")
	fmt.Fprintln(os.Stderr, "  instructions <work-unit|adr|glossary-entry>")
}

func skillsCmd(args []string) {
	if len(args) < 1 {
		skillsUsage()
		os.Exit(1)
	}
	switch args[0] {
	case "init":
		skillsInitCmd(args[1:])
	case "check":
		skillsCheckCmd(args[1:])
	default:
		fmt.Fprintf(os.Stderr, "unknown skills command: %s\n", args[0])
		skillsUsage()
		os.Exit(1)
	}
}

func skillsUsage() {
	fmt.Fprintf(os.Stderr, "usage: knack skills <subcommand> [args]\n")
	fmt.Fprintln(os.Stderr, "subcommands:")
	fmt.Fprintln(os.Stderr, "  init [--target DIR]")
	fmt.Fprintln(os.Stderr, "  check [--dir DIR]")
}

func skillsInitCmd(args []string) {
	flags := flag.NewFlagSet("init", flag.ExitOnError)
	target := flags.String("target", ".", "target directory for skill scaffolding")
	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(os.Stderr, "init: %v\n", err)
		os.Exit(1)
	}

	skillFS, err := iofs.Sub(embeddedSkills, "embedded/skills")
	if err != nil {
		fmt.Fprintf(os.Stderr, "init: embedded skills: %v\n", err)
		os.Exit(1)
	}

	wrote, skipped, err := skills.Init(skillFS, *target)
	if err != nil {
		fmt.Fprintf(os.Stderr, "init: %v\n", err)
		os.Exit(1)
	}
	for _, name := range wrote {
		fmt.Printf("wrote skill %s\n", name)
	}
	for _, name := range skipped {
		fmt.Printf("skipped existing skill %s\n", name)
	}
}

func skillsCheckCmd(args []string) {
	flags := flag.NewFlagSet("check", flag.ExitOnError)
	dir := flags.String("dir", ".agents/skills", "directory containing skills")
	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(os.Stderr, "check: %v\n", err)
		os.Exit(1)
	}

	findings, err := skills.Check(os.DirFS(*dir))
	if err != nil {
		fmt.Fprintf(os.Stderr, "check: %v\n", err)
		os.Exit(1)
	}
	for _, f := range findings {
		fmt.Println(f)
	}
	if len(findings) > 0 {
		os.Exit(1)
	}
}

func validateCmd(args []string) {
	if len(args) < 1 {
		fmt.Fprintln(os.Stderr, "usage: knack validate <queue-file>")
		os.Exit(1)
	}
	path := args[0]

	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "validate: %v\n", err)
		os.Exit(1)
	}
	results := queue.Validate(string(data))
	if len(results) == 0 {
		fmt.Fprintf(os.Stderr, "validate: no work units found in %s\n", path)
		os.Exit(1)
	}
	for _, r := range results {
		fmt.Println(queue.Format(r))
	}
	if !queue.AllValid(results) {
		os.Exit(1)
	}
}

func decisionsCmd(args []string) {
	if len(args) < 1 {
		decisionsUsage()
		os.Exit(1)
	}
	switch args[0] {
	case "list":
		decisionsListCmd()
	case "show":
		decisionsShowCmd(args[1:])
	case "check":
		decisionsCheckCmd()
	default:
		fmt.Fprintf(os.Stderr, "unknown decisions command: %s\n", args[0])
		decisionsUsage()
		os.Exit(1)
	}
}

func decisionsUsage() {
	fmt.Fprintf(os.Stderr, "usage: knack decisions <subcommand> [args]\n")
	fmt.Fprintln(os.Stderr, "subcommands:")
	fmt.Fprintln(os.Stderr, "  list")
	fmt.Fprintln(os.Stderr, "  show NNNN")
	fmt.Fprintln(os.Stderr, "  check")
}

func decisionsListCmd() {
	adrs, err := decisions.List(os.DirFS("."), "decisions")
	if err != nil {
		fmt.Fprintf(os.Stderr, "decisions list: %v\n", err)
		os.Exit(1)
	}
	for _, adr := range adrs {
		fmt.Printf("%s: %s (%s)\n", adr.Number, adr.Title, adr.Status)
	}
}

func decisionsShowCmd(args []string) {
	if len(args) < 1 {
		fmt.Fprintln(os.Stderr, "usage: knack decisions show NNNN")
		os.Exit(1)
	}
	data, err := decisions.Show(os.DirFS("."), "decisions", args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "decisions show: %v\n", err)
		os.Exit(1)
	}
	fmt.Print(string(data))
}

func decisionsCheckCmd() {
	findings, err := decisions.Check(os.DirFS("."), "decisions", os.DirFS("."), ".loop")
	if err != nil {
		fmt.Fprintf(os.Stderr, "decisions check: %v\n", err)
		os.Exit(1)
	}
	for _, f := range findings {
		fmt.Println(f)
	}
	if len(findings) > 0 {
		os.Exit(1)
	}
}

func statusCmd(args []string) {
	r, err := status.Generate(os.DirFS("."))
	if err != nil {
		fmt.Fprintf(os.Stderr, "status: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("queue: %d pending, %d done, %d failed\n", r.Pending, r.Done, r.Failed)
	fmt.Printf("evidence: %d\n", r.Evidence)
	fmt.Printf("adrs: %d\n", r.ADRs)
}

func glossaryCmd(args []string) {
	if len(args) < 1 {
		glossaryUsage()
		os.Exit(1)
	}
	switch args[0] {
	case "check":
		glossaryCheckCmd(args[1:])
	default:
		fmt.Fprintf(os.Stderr, "unknown glossary command: %s\n", args[0])
		glossaryUsage()
		os.Exit(1)
	}
}

func glossaryUsage() {
	fmt.Fprintf(os.Stderr, "usage: knack glossary <subcommand> [args]\n")
	fmt.Fprintln(os.Stderr, "subcommands:")
	fmt.Fprintln(os.Stderr, "  check")
}

func glossaryCheckCmd(args []string) {
	flags := flag.NewFlagSet("check", flag.ExitOnError)
	file := flags.String("file", "glossary.md", "path to glossary file")
	if err := flags.Parse(args); err != nil {
		fmt.Fprintf(os.Stderr, "glossary check: %v\n", err)
		os.Exit(1)
	}

	findings, err := glossary.Check(os.DirFS("."), *file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "glossary check: %v\n", err)
		os.Exit(1)
	}
	for _, f := range findings {
		fmt.Println(f.String())
	}
	if len(findings) > 0 {
		os.Exit(1)
	}
}

func instructionsCmd(args []string) {
	if len(args) < 1 {
		instructionsUsage()
		os.Exit(1)
	}
	if err := instructions.Print(os.Stdout, args[0]); err != nil {
		fmt.Fprintf(os.Stderr, "instructions: %v\n", err)
		os.Exit(1)
	}
}

func instructionsUsage() {
	fmt.Fprintf(os.Stderr, "usage: knack instructions <artifact>\n")
	fmt.Fprintln(os.Stderr, "artifacts:")
	fmt.Fprintln(os.Stderr, "  work-unit")
	fmt.Fprintln(os.Stderr, "  adr")
	fmt.Fprintln(os.Stderr, "  glossary-entry")
}
