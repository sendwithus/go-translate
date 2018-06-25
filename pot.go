package translate

import (
	"fmt"
	"strings"
)

var prefix = "# Translations template for PROJECT.\n" +
	"# Copyright (C) 2018 ORGANIZATION\n" +
	"# This file is distributed under the same license as the PROJECT project.\n" +
	"# FIRST AUTHOR <EMAIL@ADDRESS>, 2018.\n" +
	"#\n" +
	"#, fuzzy\n" +
	"msgid \"\"\n" +
	"msgstr \"\"\n" +
	"\"Project-Id-Version: PROJECT VERSION\"\n" +
	"\"Report-Msgid-Bugs-To: EMAIL@ADDRESS\"\n" +
	"\"PO-Revision-Date: YEAR-MO-DA HO:MI+ZONE\"\n" +
	"\"Last-Translator: FULL NAME <EMAIL@ADDRESS>\"\n" +
	"\"Language-Team: LANGUAGE <LL@li.org>\"\n" +
	"\"MIME-Version: 1.0\"\n" +
	"\"Content-Type: text/plain; charset=utf-8\"\n" +
	"\"Content-Transfer-Encoding: 8bit\"\n" +
	"\"Generated-By: go-translate 0.1\"\n\n"

type POT struct {
	translations []POTTranslation
}

type POTTranslation struct {
	original    string
	translation string
}

func (pot *POT) AddTrans(defaultText, transText string) {
	pot.translations = append(pot.translations, POTTranslation{defaultText, transText})
}

func (pot *POT) idLines(id string) []string {
	parts := strings.Split(id, "\n")
	if len(parts) == 1 {
		return parts
	}

	lines := []string{""}
	for i, part := range parts {
		if i != len(parts)-1 {
			part += "\\n"
		}
		lines = append(lines, part)
	}
	return lines
}
func (pot *POT) ToString() (string, error) {
	content := prefix

	first := true
	for _, trans := range pot.translations {
		if first {
			first = false
		} else {
			content += "\n"
		}
		jinjaTrans := NewJinjaTranslation(trans.original)
		if jinjaTrans.hasVars {
			content += "#, python-format\n"
		}
		content += fmt.Sprintf("msgid ")
		for _, line := range pot.idLines(jinjaTrans.toString()) {
			content += fmt.Sprintf("\"%s\"\n", line)
		}
		content += fmt.Sprintf("msgtxt \"%s\"\n", trans.translation)
	}
	return content, nil
}
