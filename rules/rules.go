package rules

import (
	"gitlab.com/tbacompany/gitector/reader"
	"strings"
)

type GitError struct {
	ErrorType   interface{}
	Title       string
	Description string
	Commit      reader.GitCommit
}

func Rules(description []reader.GitCommit, directory string) []GitError {
	config := ReadConfig(directory)
	var errors []GitError
	for _, elem := range description {
		foundErrors := singleCommit(elem, config)
		errors = append(errors, foundErrors...)
	}
	return errors
}

func singleCommit(description reader.GitCommit, config ProjectConfig) []GitError {
	var errors []GitError

	// Skip if merge request
	if strings.HasPrefix(description.Title, "Merge") {
		return errors
	}

	if !MaxCharacters(description.Title, config.maxCharacters) {
		errors = append(errors, MaxCharactersError(description.Title, config.maxCharacters, description))
	}

	if !CheckForValidEmail(description.Signature.Email, config.emailDomains) {
		errors = append(errors, CheckForValidEmailError(description.Signature.Email, config.emailDomains, description))
	}

	if !ContainsTicketNumber(description.Title, config.ticketRegexp) {
		errors = append(errors, ContainsTicketNumberError(description.Title, config.ticketRegexp, description))
	}

	if !EmptyLineBetween(description.Description) {
		errors = append(errors, EmptyLineBetweenError(description))
	}

	if config.noTrailingPunctuationInTitle && !CheckForTrailingPunctuationInTitle(description.Title) {
		errors = append(errors, CheckForTrailingPunctuationInTitleError(description))
	}

	if !CheckForMaxFiles(description.FilesCount, config.maxFiles) {
		errors = append(errors, CheckForMaxFilesError(description.FilesCount, config.maxFiles, description))
	}

	if !StartsWithBaseVerb(description.Title) {
		errors = append(errors, StartsWithBaseVerbError(description.Title, description))
	}

	return errors
}
