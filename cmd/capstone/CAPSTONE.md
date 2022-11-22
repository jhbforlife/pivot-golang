# Capstone Project
## Summary
For my capstone project, I have chosen to create a translation CLI tool using the [Google Cloud Translation API](https://cloud.google.com/translate/docs/reference/libraries/v2/go). It will be deployed using the [Google Cloud App Engine](https://cloud.google.com/appengine/docs/standard/go/runtime), and translation requests can be sent using HTTP GET requests to the service's public URL.  

In this project, I will demonstrate the following:
- [Building a REST API in Go](github.com/jhbforlife/pivot-golang/tree/capstone/translate)
- Integrating with a [remote API](https://cloud.google.com/translate/docs/reference/libraries/v2/go) for translations
- Using Go's [sql package](https://pkg.go.dev/database/sql) and a [sqlite driver](https://github.com/mattn/go-sqlite3) to cache translations that have already been fetched for a set period of time
- Deploying my REST API to a [third party PaaS ](https://cloud.google.com/appengine/docs/standard/go/runtime)
- [Building a user friendly CLI client application](github.com/jhbforlife/pivot-golang/tree/capstone/cmd/capstone) to interact with [my REST API](github.com/jhbforlife/pivot-golang/tree/capstone/translate)
- Error handling
- Testing in Go
- Working with the [Google Cloud Platform](https://cloud.google.com)
- Working with YAML files

## User Stories
### Number One
### As a user, I would like to invoke the CLI client application with a tag for the language to translate to, and the text to translate and have the translation printed out to the command line.  
**Acceptance Criteria**  
Given a valid language, I expect the CLI client application to print out the assumed from language with the text to translate, and the to language with the API provided translation.  
Example:
```
$ translate --to french hello
English: hello
French: bonjour
```
If the language is invalid or not supported by the translate service, I expect the appropriate error to be printed out instead of a provided translation.  
Example:  
```
$ translate --to fren hello
Error: Language not valid. Please try again
```
### Number Two
### As a user, if the translate service at first assumes the wrong language, or I would like to explicitly provide the language, I can do so with an additional tag.
**Acceptance Critera**  
Given a valid from language and to language, I expect the CLI client application to print out the from language and text to translate, along with the to language and translation provided by the API.  
Example:
```
$ translate -f english -t french this evening
English: this evening
French: ce soir
```
If either language is invalid or not supported by the translate service, I expect the appropriate answer to be printed out instead of a provided translation.  
Example:
```
$ translate -f -t french this evening
Error: From language not valid. Please try again
```
### Number Three
### As a user, I would like the CLI client application to provide usage help when I invoke it without a command so I may learn how to use it properly.
**Acceptance Criteria**  
Given the CLI client application is invoked without a command, usage help will be provided.  
Example:
```
$ translate
Usage: translate [options...] <text>

-f, --from
Provide the language you would like the service to translate from. If not provided, from language is assumed by the service.

-h, --help
Shows this usage information.

-t, --to
Provide the language you would like the service to translate to. If not provided, no translation will be provided.
```