package engine


import (
	"github.com/jbrukh/bayesian"
	"os"
	"log"
	"github.com/itsbalamurali/heyasha/models"
	"github.com/itsbalamurali/heyasha/core/database"
)

// bClassifiers holds the trained bayesian classifiers for our domains. The key
// is the domain ID to which the trained classifier belongs.
var bClassifiers = map[uint64]*bayesian.Classifier{}

// domainIntents holds the intents for which each domain has been trained. The
// outer map divides the intents for each domain by domain ID.
var domainIntents = map[uint64][]bayesian.Class{}


func Boot()  {
	go func() {
		if os.Getenv("ENV") != "test" {
			log.Println("training classifiers")
		}
		if err := trainClassifiers(); err != nil {
			log.Println("could not train classifiers", err)
		}
	}()
}

// trainClassifiers trains classifiers for each plugin.
func trainClassifiers() error {
		ss := fetchTrainingSentences()
		// Assemble list of Bayesian classes from all trained intents
		// for this plugin. m is used to keep track of the classes
		// already taught to each classifier.
		m := map[string]struct{}{}
		for _, s := range ss {
			_, ok := m[s.Intent]
			if ok {
				continue
			}
			log.Println("learning intent", s.Intent)
			m[s.Intent] = struct{}{}
			domainIntents[s.DomainID] = append(domainIntents[s.DomainID],
				bayesian.Class(s.Intent))
		}

		// Build classifier from complete sets of intents
		for _, s := range ss {
			intents := domainIntents[s.DomainID]
			// Calling bayesian.NewClassifier() with 0 or 1
			// classes causes a panic.
			if len(intents) == 0 {
				break
			}
			if len(intents) == 1 {
				intents = append(intents, bayesian.Class("__no_intent"))
			}
			c := bayesian.NewClassifier(intents...)
			bClassifiers[s.DomainID] = c
		}

		// With classifiers initialized, train each of them on a
		// sentence's stems.
		for _, s := range ss {
			tokens := TokenizeSentence(s.Sentence)
			stems := StemTokens(tokens)
			c, exists := bClassifiers[s.DomainID]
			if exists {
				c.Learn(stems, bayesian.Class(s.Intent))
			}
		}
	return nil
}

func fetchTrainingSentences() ([]models.Intent) {
	var Intents = []models.Intent{}
	db := database.MysqlCon()
	db.Find(&Intents)
	log.Println("Fetching training sentences")
	return Intents
}