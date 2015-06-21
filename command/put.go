package command

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/hashicorp/consul/api"
	"github.com/mitchellh/cli"
)

type Put struct {
	Ui cli.Ui
}

func (c *Put) Help() string {
	return "Usage: consulkv put [-datacenter=] [-flags=0] KEY [VALUE]"
}

func (c *Put) Synopsis() string {
	return "Put key/value"
}

func (c *Put) Run(args []string) int {
	var datacenter string
	var flags uint64
	cmdFlags := flag.NewFlagSet("put", flag.ContinueOnError)
	cmdFlags.Usage = func() { c.Ui.Output(c.Help()) }
	cmdFlags.StringVar(&datacenter, "datacenter", "", "")
	cmdFlags.Uint64Var(&flags, "flags", 0, "")
	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}
	args = cmdFlags.Args()

	var key string
	var value []byte
	var err error

	switch len(args) {
	case 0:
		c.Ui.Error("Key must be specified")
		return 1
	case 1:
		key = args[0]
		value, err = ioutil.ReadAll(os.Stdin)
		if err != nil {
			c.Ui.Error(fmt.Sprintf("Error reading data from Stdin: %s", err))
			return 1
		}
	case 2:
		key = args[0]
		value = []byte(args[1])
	default:
		c.Ui.Error("Too many arguments")
		return 1
	}

	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		c.Ui.Error(fmt.Sprintf("Error connecting to Consul agent: %s", err))
		return 1
	}
	kv := client.KV()

	pair := api.KVPair{Key: key, Value: value, Flags: flags}
	_, err = kv.Put(&pair, &api.WriteOptions{Datacenter: datacenter})
	if err != nil {
		c.Ui.Error(fmt.Sprintf("Error setting data: %s", err))
		return 1
	}

	return 0
}
