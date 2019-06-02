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
//go:generate gofmt -s -w ../
var (
	keywords   = []string{"Describe", "List"}
	SDKName    = "github.com/aws/aws-sdk-go-v2"
	errorNoSDK = errors.New("The AWS SDK v2 is not a dependency from this project")
	timestamp  = time.Now()
	tplFuncs   = template.FuncMap{
		"ToLower": strings.ToLower,
		"GetPrefix": func(s string) string {
			for _, k := range keywords {
				if strings.HasPrefix(s, k) {
					return k
				}
			}
			return ""
		},
		"Clean": func(s string) string {
			for _, k := range keywords {
				s = strings.TrimPrefix(s, k)
			}
			s = strings.TrimSuffix(s, "Input")
			return s
		},
	}
	test = make([]Typ, 0)
)

type Typ struct {
	Name      string
	Resources []string
}

func main() {
	template := template.Must(template.New("").Funcs(tplFuncs).ParseGlob("*.tpl"))

	version, err := getSDKVersion()
	die(err)

	files := getFiles(getPathnameCode(version))

	for _, file := range files {
		if strings.Contains(file, "cloudformation") {
			continue
		}
		typ := getInfoFromFile(file)

		var name string
		resources := make([]string, 0)

		for _, r := range typ.Resources {
			r = strings.TrimRightFunc(r, func(r rune) bool {
				return r != ')'
			})

			r = strings.TrimLeftFunc(r, func(r rune) bool {
				return r != '('
			})

			r = strings.TrimLeftFunc(r, func(r rune) bool {
				return r != '.'
			})

			r = strings.Trim(r, "().")

			resources = append(resources, r)

			sliced := strings.Split(typ.Name, string(" "))
			name = strings.TrimSuffix(sliced[len(sliced)-3], "iface")
			name = strings.TrimSuffix(name, "API")
		}

		if name == "" && len(resources) == 0 {
			continue
		}

		typ.Name = name
		typ.Resources = resources

		test = append(test, *typ)
	}

	generateInit(template)
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

func getInfoFromFile(filename string) *Typ {
	f := openFile(filename)

	data := bytes.NewReader(f)
	scanner := bufio.NewScanner(data)

	r := new(Typ)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "//") {
			continue
		}

		line = strings.TrimSpace(line)

		if strings.Contains(line, "type") {
			r.Name = line
		}

		for _, keyword := range keywords {
			if strings.HasPrefix(line, keyword) {
				r.Resources = append(r.Resources, line)
				continue
			}
		}
	}

	return r
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

func die(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func generateInit(tpl *template.Template) {
	f, err := os.Create("../init.go")
	die(err)
	defer f.Close()

	err = tpl.ExecuteTemplate(f, "init.tpl", struct {
		Timestamp      time.Time
		ResourcesTypes []Typ
	}{
		Timestamp:      timestamp,
		ResourcesTypes: test,
	})
	die(err)
}

func generateResourcesFiles(tpl *template.Template) {
	for _, typ := range test {

		templates := map[string]string{
			"resource.tpl":      "../" + strings.ToLower(typ.Name) + ".go",
			"resource_data.tpl": "../" + strings.ToLower(typ.Name) + "_data.go",
		}

		for key := range templates {
			f, err := os.Create(templates[key])
			die(err)
			defer f.Close()

			err = tpl.ExecuteTemplate(f, key, struct {
				Timestamp time.Time
				Resource  string
			}{
				Timestamp: timestamp,
				Resource:  typ.Name,
			})
			die(err)

		}
	}
}

func generateTypeFile(tpl *template.Template) {
	f, err := os.Create("../types.go")
	die(err)
	defer f.Close()

	err = tpl.ExecuteTemplate(f, "types.tpl", struct {
		Timestamp time.Time
		Resources []Typ
	}{
		Timestamp: timestamp,
		Resources: test,
	})
	die(err)
}
