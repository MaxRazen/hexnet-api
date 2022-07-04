package main

import (
	"bufio"
	"fmt"
	"golang.org/x/term"
	"hexnet/api/common"
	"hexnet/api/users"
	"os"
	"strings"
	"syscall"
)

type command struct {
	name     string
	flags    map[string]string
	argument string
}

const (
	CmdHelp     = "help"
	CmdMakeUser = "make:user"
	CmdMigrate  = "migrate"
)

const (
	CmdDescHelp     = "\t\t\tDisplay this help"
	CmdDescMakeUser = "\t\tCreate a user"
	CmdDescMigrate  = "\t\t\tMigrate DB"
)

func main() {
	cmd := resolveCliCommand(os.Args)

	switch cmd.name {
	case CmdHelp:
		helpHandler()
	case CmdMakeUser:
		userData := users.UserCreateData{
			Name:     stringPrompt("Enter name", false),
			Login:    stringPrompt("Enter login", false),
			Password: stringPrompt("Enter password", true),
		}
		fmt.Println()
		makeUserHandler(userData)
	case CmdMigrate:
		migrateHandler()
	default:
		helpHandler()
	}
}

func resolveCliCommand(cliArgs []string) command {
	if len(cliArgs) < 2 {
		return command{
			name: CmdHelp,
		}
	}

	args := cliArgs[1:]
	var flags = make(map[string]string)
	var arg string

	for i := 1; i < len(args); i++ {
		if strings.HasPrefix(args[i], "--") {
			tokens := strings.Split(args[i][2:], "=")
			flags[tokens[0]] = tokens[1]
		} else {
			arg = args[i]
		}
	}

	return command{
		name:     args[0],
		flags:    flags,
		argument: arg,
	}
}

func helpHandler() {
	commandList := map[string]string{
		CmdHelp:     CmdDescHelp,
		CmdMakeUser: CmdDescMakeUser,
		CmdMigrate:  CmdDescMigrate,
	}

	for k, v := range commandList {
		fmt.Printf("%v%v\n", k, v)
	}
}

func makeUserHandler(data users.UserCreateData) {
	interactWithDb()
	m, err := users.CreateUserAction(data)

	if err != nil {
		fmt.Printf("Fatal: %v\n", err.Error())
	}
	fmt.Printf("User %v was successfully created with ID: %v\n", m.Name, m.ID)
}

func migrateHandler() {
	interactWithDb()
	users.AutoMigrate()
	fmt.Printf("DB Migrated Successfully\n")
}

func stringPrompt(label string, hidden bool) string {
	var s string
	r := bufio.NewReader(os.Stdin)
	for {
		_, err := fmt.Fprint(os.Stderr, label+": ")

		if err != nil {
			panic(err.Error())
		}

		if hidden {
			b, _ := term.ReadPassword(syscall.Stdin)
			s = string(b)
		} else {
			s, _ = r.ReadString('\n')
		}

		if s != "" {
			break
		}
	}
	return strings.TrimSpace(s)
}

func interactWithDb() {
	config := common.LoadConfig("")
	common.InitDbConnection(config.Env.DB)
}
