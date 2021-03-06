package driver

import (
	"fmt"

	"github.com/sjindal-arista/Arista-QA/utils"
)

type Question struct {
	q       string
	quesID  string
	user    *User
	answers []*Answer
}

func (ques *Question) Matches(query string) bool {
	// for now keep it simple
	return utils.SubsequenceMatcher(ques.q, query)
}

func (d *driver) AddNewQuestion(newQ string, u *User) error {
	// assert user already exists
	if !d.userExists(u) {
		return fmt.Errorf("User does not exist")
	}
	questionObj := &Question{
		q:      newQ,
		quesID: utils.GenerateUUID("Ques"),
		user:   u,
	}
	d.listQues = append(d.listQues, questionObj)
	return nil
}

func (d *driver) SearchQuestion(query string) []*Question {
	results := make([]*Question, 0)
	for _, ques := range d.listQues {
		if ques.Matches(query) {
			results = append(results, ques)
		}
	}
	return results
}

func (d *driver) GetQuestion(qID string) *Question {
	for _, ques := range d.listQues {
		if ques.quesID == qID {
			return ques
		}
	}
	return nil
}
