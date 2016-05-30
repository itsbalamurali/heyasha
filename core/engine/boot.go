package engine

/*
import (
	"github.com/jbrukh/bayesian"
	"os"
	"log"
	"net/http"
	"time"
	"strconv"
	"encoding/json"
)

var pluginsGo = []dt.PluginConfig{}


// bClassifiers holds the trained bayesian classifiers for our plugins. The key
// is the plugin ID to which the trained classifier belongs.
var bClassifiers = map[uint64]*bayesian.Classifier{}

// pluginIntents holds the intents for which each plugin has been trained. The
// outer map divides the intents for each plugin by plugin ID.
var domainIntents = map[uint64][]bayesian.Class{}


type tSentence struct {
	ID       uint64
	Sentence string
	Intent   string
	DomainID uint64
}

func Boot()  {
	go func() {
		if os.Getenv("ASHA_ENV") != "test" {
			log.Info("training classifiers")
		}
		if err = trainClassifiers(); err != nil {
			log.Info("could not train classifiers", err)
		}
	}()
}

// trainClassifiers trains classifiers for each plugin.
func trainClassifiers() error {
	for _, pconf := range pluginsGo {
		ss, err := fetchTrainingSentences(pconf.ID, pconf.Name)
		if err != nil {
			return err
		}

		// Assemble list of Bayesian classes from all trained intents
		// for this plugin. m is used to keep track of the classes
		// already taught to each classifier.
		m := map[string]struct{}{}
		for _, s := range ss {
			_, ok := m[s.Intent]
			if ok {
				continue
			}
			//log.Debug("learning intent", s.Intent)
			m[s.Intent] = struct{}{}
			pluginIntents[s.PluginID] = append(pluginIntents[s.PluginID],
				bayesian.Class(s.Intent))
		}

		// Build classifier from complete sets of intents
		for _, s := range ss {
			intents := pluginIntents[s.PluginID]
			// Calling bayesian.NewClassifier() with 0 or 1
			// classes causes a panic.
			if len(intents) == 0 {
				break
			}
			if len(intents) == 1 {
				intents = append(intents, bayesian.Class("__no_intent"))
			}
			c := bayesian.NewClassifier(intents...)
			bClassifiers[s.PluginID] = c
		}

		// With classifiers initialized, train each of them on a
		// sentence's stems.
		for _, s := range ss {
			tokens := TokenizeSentence(s.Sentence)
			stems := StemTokens(tokens)
			c, exists := bClassifiers[s.PluginID]
			if exists {
				c.Learn(stems, bayesian.Class(s.Intent))
			}
		}
	}
	return nil
}

func fetchTrainingSentences(domain uint64, name string) ([]tSentence, error) {
	c := &http.Client{Timeout: 10 * time.Second}
	pid := strconv.FormatUint(pluginID, 10)
	u := os.Getenv("ITSABOT_URL") + "/api/plugins/train/" + pid
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err = resp.Body.Close(); err != nil {
			//log.Info("failed to close response body.", err)
		}
	}()
	ss := []tSentence{}

	// This occurs when the plugin has not been published, which we should
	// ignore on boot.
	if resp.StatusCode == http.StatusBadRequest {
		//log.Infof("warn: plugin %s has not been published", name)
		return ss, nil
	}
	err = json.NewDecoder(resp.Body).Decode(&ss)
	return ss, err
}
*/