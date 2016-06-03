package engine

import (
	"errors"
	"reflect"
	"github.com/itsbalamurali/heyasha/models"
	"log"
)


// NewMsg builds a message struct with Tokens, Stems, and a Structured Input.
func NewMsg(u *models.User, msg string) *models.Message {
	tokens := TokenizeSentence(msg)
	stems := StemTokens(tokens)
	//si := ner.classifyTokens(tokens)

	// Get the intents as determined by each plugin
	for domainID, c := range bClassifiers {
		scores, idx, _ := c.ProbScores(stems)
		log.Println("intent score", domainIntents[domainID][idx],
			scores[idx])
		if scores[idx] > 0.7 {
			//si.Intents = append(si.Intents, string(pluginIntents[pluginID][idx]))
		}
	}

	m := &models.Message{
		User:            u,
		Sentence:        msg,
		Tokens:          tokens,
		Stems:           stems,
		//StructuredInput: si,
	}
	/*
		m, err = addContext(db, m)
		if err != nil {
			log.Debug(err)
		}
	*/
	return m
}

//http://stackoverflow.com/questions/18017979/golang-pointer-to-function-from-string-functions-name
func Call(fn map[string]interface{}, command string, params ...interface{}) (result []reflect.Value, err error) {
	f := reflect.ValueOf(fn[command])
	if len(params) != f.Type().NumIn() {
		err = errors.New("The number of params is not adapted.")
		return
	}
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	result = f.Call(in)
	return
}

/*
// /Functions map
funcs := map[string]interface{} {
	"command": function,
	"command2": function2,
	//so on and so forth
}

*/
