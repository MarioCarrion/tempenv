// Copyright © 2017 Mario Carrion <mario@carrion.ws>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package parser

import (
	"bytes"
	"io/ioutil"
	"os"
	"path"
	"text/template"
)

func ParseFile(filename string) (string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}

	return Parse(filename, string(data))
}

func Parse(filename, text string) (string, error) {
	tmpl, err := template.New(path.Base(filename)).Funcs(template.FuncMap{"getEnv": getEnv}).Parse(text)
	if err != nil {
		return "", err
	}

	var writer bytes.Buffer
	err = tmpl.Execute(&writer, nil)
	if err != nil {
		return "", err
	}

	return writer.String(), nil
}

func getEnv(name string) string {
	return os.Getenv(name)
}
