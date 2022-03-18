package service

import (
	"github.com/hiteshrepo/token_project/internal/app/model"
	"github.com/pkg/errors"
	"log"
	"strconv"
	"strings"
)

type criteria struct {
	paramCount int
}

var eligibleActions = map[string]criteria{
	"-create": {paramCount: 6},
	"-write":  {paramCount: 14},
	"-read":   {paramCount: 6},
	"-drop":   {paramCount: 6},
}

type config struct {
	host string
	port string
}

const (
	eligibleApp      = "tokenclient"
	fixedInputCount  = 2
	idParamName      = "-id"
	nameParamName    = "-name"
	lowParamName     = "-low"
	midParamName     = "-mid"
	highParamName    = "-high"
	hostParamName    = "-host"
	portParamName    = "-port"
	actionNameCreate = "-create"
	actionNameRead   = "-read"
	actionNameDrop   = "-drop"
	actionNameWrite   = "-write"
)

type ParserService struct {}

func ProvideParserService() *ParserService {
	return &ParserService{}
}

func (ps *ParserService) GetAction(parts []string) string {
	if len(parts) < 2 {
		return ""
	}

	app := strings.ToLower(strings.TrimSpace(parts[0]))
	action := strings.ToLower(strings.TrimSpace(parts[1]))

	if _, ok := eligibleActions[action]; ok && app == eligibleApp {
		return action
	}

	return ""
}

func (ps *ParserService) GetInput(parts []string, action string) (*config, *model.Token, error) {
	cnf := &config{}
	input := &model.Token{}
	if c, _ := eligibleActions[action]; c.paramCount != len(parts)-fixedInputCount {
		return nil, nil, errors.New("invalid input")
	}

	if action == actionNameCreate || action == actionNameRead || action == actionNameDrop {
		getParams(parts, input, cnf)
	}

	if action == actionNameWrite {
		getWriteParams(parts, input, cnf)
	}

	return cnf, input, nil
}

func getParams(parts []string, input *model.Token, cnf *config) {
	assignCommonParams(parts, input, cnf)
}

func getWriteParams(parts []string, input *model.Token, cnf *config) {
	count := 2
	assignCommonParams(parts, input, cnf)

	for count < len(parts) {
		if parts[count] == nameParamName {
			input.Name = parts[count+1]
		}

		if parts[count] == lowParamName {
			low, err := getNumConverted(parts, count)
			checkErr(err)
			input.Low = int64(low)
		}

		if parts[count] == highParamName {
			high, err := getNumConverted(parts, count)
			checkErr(err)
			input.High = int64(high)
		}

		if parts[count] == midParamName {
			mid, err := getNumConverted(parts, count)
			checkErr(err)
			input.Mid = int64(mid)
		}

		count += 2
	}
}

func assignCommonParams(parts []string, input *model.Token, cnf *config) {
	count := 2
	for count < len(parts) {
		if parts[count] == idParamName {
			id, err := getNumConverted(parts, count)
			checkErr(err)
			input.Id = int64(id)
		}

		if parts[count] == hostParamName {
			cnf.host = parts[count+1]
		}

		if parts[count] == portParamName {
			cnf.port = parts[count+1]
		}
		count += 2
	}
}

func getNumConverted(parts []string, count int) (int, error) {
	id, err := strconv.Atoi(parts[count+1])
	if err != nil {
		return -1, err
	}

	return id, nil
}

func checkErr(err error) {
	if err != nil {
		log.Fatalf(err.Error())
	}
}
