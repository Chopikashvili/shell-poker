package ux

import "github.com/AlecAivazis/survey/v2"

func SurveySettings(is *survey.IconSet) {
	is.Question.Text = ""
	is.SelectFocus.Format = "red"
}
