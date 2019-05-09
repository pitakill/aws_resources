// +build ignore
package main

import (
	"bufio"
	"bytes"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"
)

//go:generate go run gen.go
//go:generate gofmt -s -w ../*.go
var (
	resourcesTypes = make(map[string][]string, 0)
	resourcesTyp   = make([]string, 0)
	SDKName        = "github.com/aws/aws-sdk-go-v2"
	errorNoSDK     = errors.New("The AWS SDK v2 is not a dependency from this project")
	timestamp      = time.Now()
	tplFuncs       = template.FuncMap{
		"ToLower": strings.ToLower,
	}
)

func main() {
	template := template.Must(template.New("").Funcs(tplFuncs).ParseGlob("*.tpl"))

	version, err := getSDKVersion()
	die(err)

	files := getFiles(getPathnameCode(version))

	for _, file := range files {
		if strings.Contains(file, "cloudformation") {
			continue
		}
		describes, types := getInfoFromFile(file)

		if len(describes) > 0 {
			resources := make([]string, 0)
			for _, structure := range describes {

				structure = strings.TrimRightFunc(structure, func(r rune) bool {
					return r != ')'
				})

				structure = strings.TrimLeftFunc(structure, func(r rune) bool {
					return r != '('
				})

				structure = strings.TrimLeftFunc(structure, func(r rune) bool {
					return r != '.'
				})

				structure = strings.Trim(structure, "().")

				resources = append(resources, structure)
			}

			sliced := strings.Split(file, string(filepath.Separator))
			typ := strings.TrimSuffix(sliced[len(sliced)-3], "iface")

			resourcesTypes[typ] = resources
		}

		if len(types) > 0 {
			for _, typ := range types {
				splitted := strings.Split(typ, " ")
				typCleaned := splitted[1]

				typCleaned = strings.TrimSuffix(typCleaned, "API")

				resourcesTyp = append(resourcesTyp, typCleaned)
			}
		}
	}

	generateInit("init", "code", template)
	generateResourcesFiles(template)
	generateTypeFile(template)
}

func getFiles(path string) []string {
	files, err := filepath.Glob(filepath.Join(path, "service/**/*iface/interface.go"))
	die(err)

	return files
}

func getPathnameCode(version string) string {
	goPath := os.Getenv("GOPATH")

	return filepath.Join(goPath, "pkg", "mod", SDKName+"@"+version)
}

func openFile(input string) []byte {
	f, err := ioutil.ReadFile(input)
	die(err)

	return f
}

func getInfoFromFile(filename string) (declarations, types []string) {
	f := openFile(filename)

	data := bytes.NewReader(f)
	scanner := bufio.NewScanner(data)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "//") {
			continue
		}

		line = strings.TrimSpace(line)

		if strings.Contains(line, "type") {
			types = append(types, line)
		}

		if strings.HasPrefix(line, "Describe") {
			declarations = append(declarations, line)
		}
	}

	return
}

func getSDKVersion() (string, error) {
	wd, err := os.Getwd()
	die(err)

	f := openFile(filepath.Join(wd, "..", "go.mod"))
	data := bytes.NewReader(f)
	scanner := bufio.NewScanner(data)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(scanner.Text(), SDKName) {
			splittedLine := strings.Split(line, " ")
			return splittedLine[len(splittedLine)-1], nil
		}
	}

	return "", errorNoSDK
}

func generateTypeFile(tpl *template.Template) {
	f, err := os.Create("../types.go")
	die(err)
	defer f.Close()

	err = tpl.ExecuteTemplate(f, "types.tpl", struct {
		Timestamp time.Time
		Resources []string
	}{
		Timestamp: timestamp,
		Resources: resourcesTyp,
	})
	die(err)
}

func generateResourcesFiles(tpl *template.Template) {
	for _, typ := range resourcesTyp {
		f, err := os.Create("../" + strings.ToLower(typ) + ".go")
		die(err)
		defer f.Close()

		err = tpl.ExecuteTemplate(f, "resource.tpl", struct {
			Timestamp time.Time
			Resource  string
		}{
			Timestamp: timestamp,
			Resource:  typ,
		})
		die(err)
	}
}

func generateInit(output, input string, tpl *template.Template) {
	f, err := os.Create("../" + output + ".go")
	die(err)
	defer f.Close()

	err = tpl.ExecuteTemplate(f, input+".tpl", struct {
		Timestamp      time.Time
		ResourcesTypes map[string][]string
	}{
		Timestamp:      timestamp,
		ResourcesTypes: resourcesTypes,
	})
	die(err)
}

func die(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
