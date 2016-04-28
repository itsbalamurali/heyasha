//Project : Hey Asha!
//Copyright (C) Balamurali Pandranki - All Rights Reserved
//Unauthorized copying of this file, via any medium is strictly prohibited
//Proprietary and confidential
//Written by Balamurali Pandranki <balamurali@live.com>, 29/4/2016 2:18 AM
package stt

import (
	"net/http"
	"golang.org/x/text/language"
	"golang.org/x/net/context"
)

type langkey int

var (
	englishBase, _  = language.AmericanEnglish.Base()
	lkey            langkey = 0
)

// Detect Language from accept-language
func DetectLanguage(ctx context.Context, w http.ResponseWriter, r *http.Request) context.Context {
	lang := ""
	if tags, _, err := language.ParseAcceptLanguage(r.Header.Get("Accept-Language")); err == nil {
		if len(tags) > 0 {
			t := tags[0]
			base, _, _ := t.Raw() // base, sscript, region
			if base == englishBase {
				lang = "en-US"
			} else if t == language.BritishEnglish {
				lang = "en-GB"
			}
		}
	}
	ctx = LangWithContext(lang, ctx)
	return ctx
}

func LangWithContext(lang string, ctx context.Context) context.Context {
	return context.WithValue(ctx, lkey, lang)
}

func LangFromContext(ctx context.Context) string {
	lang, _ := ctx.Value(lkey).(string)
	return lang
}

