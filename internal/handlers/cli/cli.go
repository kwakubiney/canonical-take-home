package cli

import (
	"flag"
	"log"

	"github.com/kwakubiney/canonical-take-home/command"
	"github.com/kwakubiney/canonical-take-home/internal/domain/model"
	"github.com/kwakubiney/canonical-take-home/internal/domain/repository"
)

type Options struct {
	Method       string
	TypeOfObject string
	Fields       string
	Help		 *bool
}

type UserService struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) *UserService {
	return &UserService{
		repository: repository,
	}
}

// find all
func (u *UserService) FindByKeys(keys []string) (*model.User, error) {
	res  := u.repository.FindByKeys(keys)
	return res, nil
}

type CliHandler struct {
	options Options
	service UserService
}


func NewCliHandler(opts Options, service UserService) *CliHandler{
	return &CliHandler{
		options: opts,
		service: service,
	}
}

func (s *CliHandler) Dispatch() {
	if *s.options.Help || s.options.Method == "" {
		flag.Usage()
		return
	}
	if s.options.TypeOfObject == "user" {
		switch s.options.Method {
		case "create":
			if s.options.TypeOfObject== "" || s.options.Fields == "" {
				flag.Usage()
				return
			}
			createFieldKeys := command.ParseFields(s.options.Fields)
			log.Println(createFieldKeys)
			if !command.ValidateCreateandUpdateUserFields(s.options.Method, createFieldKeys) {
				flag.Usage()
				return
			}
			//TODO: pass to create handler based on method.
		case "update":
			{
				if s.options.TypeOfObject == "" || s.options.Fields == "" {
					flag.Usage()
					return
				}
				updateFieldKeys := command.ParseFields(s.options.Fields)
				if !command.ValidateCreateandUpdateUserFields(s.options.Method, updateFieldKeys) {
					flag.Usage()
					return
				}
				log.Println(updateFieldKeys)
				//pass to update handler based on method.
			}

		case "delete":
			{
				if s.options.TypeOfObject == "" || s.options.Fields == "" {
					flag.Usage()
					return
				}
				deleteFieldKeys := command.ParseFields(s.options.Fields)
				if !command.ValidateCreateandUpdateUserFields(s.options.Method, deleteFieldKeys) {
					flag.Usage()
					return
				}
				log.Println(deleteFieldKeys)
				//pass to delete handler based on method.
			}

		default:
			flag.Usage()
			return
		}
}
}