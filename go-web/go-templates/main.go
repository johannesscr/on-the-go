package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"
	"time"
)

var tpl *template.Template

func init() {
	// here we add the func map
	// go allows you to pipe between functions
	tpl = template.Must(template.New("").Funcs(funcMap).ParseGlob("templates/*.html"))
}

func main() {
	// parsing files from text template package
	textTemplate()
	// parsing files from text template package
	textGlob()
	// preload templates
	preloadTemplates()
	// pass data to template
	passDataToTemplate()
	variablesInTemplate()
	compositeVariablesInTemplate()
	functionInTemplates()
	timeInGo()
	pipelineBetweenFunctions()
	predefinedGlobalFunctions()
	nestedTemplates()
	methodsTemplates()
}

func textTemplate() {
	fmt.Print("### Parsing Files ###\n\n")
	tpl2, err := template.ParseFiles("templates/base.html")
	if err != nil {
		log.Fatalln(err)
	}
	// take a writer and data and writes to the writer
	err = tpl2.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
	// add more file to the pointer of all the template files
	tpl2, err = tpl2.ParseFiles(
		"templates/base2.html",
		"templates/base3.html")
	err = tpl2.ExecuteTemplate(os.Stdout, "base2.html", nil)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Print("\n\n### END Parsing Files ###\n\n")
}

func textGlob() {
	fmt.Print("### Parsing Glob ###\n\n")
	tpl3, err := template.New("").Funcs(funcMap).ParseGlob("templates/*.html")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl3.ExecuteTemplate(os.Stdout, "base3.html", nil)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Print("\n\n### END Parsing Glob ###\n\n")
}

func preloadTemplates() {
	fmt.Print("### Preload Glob ###\n\n")
	err := tpl.ExecuteTemplate(os.Stdout, "base.html", nil)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Print("\n\n### END Preload Glob ###\n\n")
}

func passDataToTemplate() {
	fmt.Print("### Pass Data to Template ###\n\n")
	err := tpl.ExecuteTemplate(os.Stdout, "data.html", 42)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Print("\n\n### END Pass Data to Template ###\n\n")
}

func variablesInTemplate() {
	fmt.Print("### Variables in Template ###\n\n")
	err := tpl.ExecuteTemplate(os.Stdout, "variable.html", 42)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Print("\n\n### END Variables in Template ###\n\n")
}

func compositeVariablesInTemplate() {
	fmt.Print("### Composite Variables in Template ###\n\n")

	fmt.Print(">>> Slice <<<\n\n")
	xs := []string{"ben", "job", "sam", "freddy"}
	err := tpl.ExecuteTemplate(os.Stdout, "sliceVariable.html", xs)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Print(">>> Map <<<\n\n")
	ms := map[string]string{
		"america": "usa",
		"south africa": "rsa",
		"europe": "eu",
		"australia": "aus",

	}
	err = tpl.ExecuteTemplate(os.Stdout, "mapVariable.html", ms)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Print(">>> Map <<<\n\n")
	st := struct{
		FirstName string
		LastName string
		Age int
		Motto string
	}{"James", "Bond", 12, "die another day"}
	err = tpl.ExecuteTemplate(os.Stdout, "structVariable.html", st)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Print("\n\n### END Composite Variables in Template ###\n\n")
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

var funcMap = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
	"fdateYMD": formatDateYMD,
}

type sage struct{
	Name string
	Motto string
}

func functionInTemplates() {
	fmt.Print("### Functions in Template ###\n\n")

	b := sage{
		Name: "ben",
		Motto: "the belief of no beliefs",
	}
	g := sage{
		Name: "george",
		Motto: "be the change",
	}
	m := sage{
		Name: "martin",
		Motto: "hatred never ceases with hatred",
	}
	data := struct{
		Wisdom []sage
	}{
		Wisdom: []sage{b, g, m},
	}
	err := tpl.ExecuteTemplate(os.Stdout, "funcInTemplate.html", data)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Print("\n\n### END Functions in Template ###\n\n")
}

func formatDateYMD(t time.Time) string {
	return t.Format("2006-01-02")
}

func timeInGo() {
	fmt.Print("### Time in Template ###\n\n")
	t := time.Now()
	err := tpl.ExecuteTemplate(os.Stdout, "timeInTemplate.html", t)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Print("\n\n### END Time in Template ###\n\n")
}

func pipelineBetweenFunctions() {
	fmt.Print("### Pipeline in Template ###\n\n")
	err := tpl.ExecuteTemplate(os.Stdout, "pipelineTemplate.html", "test string")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Print("\n\n### END Pipeline in Template ###\n\n")
}

func predefinedGlobalFunctions() {
	fmt.Print("### Predefined Functions in Template ###\n\n")
	data := struct{
		Xs []string
		Score1 int
		Score2 int
	}{
		Xs: []string{"zero", "one", "two", "three"},
		Score1: 3,
		Score2: 8,
	}
	err := tpl.ExecuteTemplate(os.Stdout, "predefinedFunctionsTemplate.html", data)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Print("\n\n### END Predefined Functions in Template ###\n\n")
}

func nestedTemplates() {
	fmt.Print("### Nested in Template ###\n\n")
	err := tpl.ExecuteTemplate(os.Stdout, "modularTemplate.html", 12)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Print("\n\n### END Nested in Template ###\n\n")
}

type person struct {
	Name string
	Age int
}
func (p person) SomeProcessing() int {
	return 7
}
func (p person) AgeDbl() int {
	return 2 * p.Age
}
func (p person) TakesArgs(x int) int {
	return x + p.Age
}

func methodsTemplates() {
	fmt.Print("### Methods in Template ###\n\n")
	p := person{
		Name: "Jeff",
		Age: 71,
	}
	err := tpl.ExecuteTemplate(os.Stdout, "methodsTemplate.html", p)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Print("\n\n### END Methods in Template ###\n\n")
}