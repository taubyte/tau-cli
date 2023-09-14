package dreamland

import (
	"context"
	"fmt"
	"strings"
	"time"

	dreamland "github.com/taubyte/dreamland/service"
	"github.com/taubyte/dreamland/service/api"
	"github.com/taubyte/tau-cli/env"
)
//dream_client is a variable that dream_client that is of the the pointer to dreamlan.Client. The variable is a pointer due to the *
var dream_client *dreamland.Client
//Client takes input of variable ctx of type context.Context and returns variables of type dreamland.Client pointer and error. 
//The function is a singleton that is designed to instantiate a dreamland.Client pointer or return an error if applicable
func Client(ctx context.Context) (*dreamland.Client, error) {
	//if dream_client is nil, its not instantiated and needs to be instantiated. Otherwise just return the dream_client var as *dreamland.Client and error as nil
	//also take note that we are using the global variable dream_client here instead of creating a seperate local variable
	if dream_client == nil {
		//create error variable error
		var err error
		//set dream_client and err to a new dreamland Object with the New contrstructor in the dreamland class
		dream_client, err = dreamland.New(ctx, dreamland.URL("http://127.0.0.1:1421"), dreamland.Timeout(15*time.Second))
		//if err is nil, no error
		//if err is not nil, there is an error that needs to be returned. dreamland.Client pointer is returned as nil, instead err as error
		if err != nil {
			return nil, err
		}
	}
	return dream_client, nil
}
//Status function takes input of ctx of type context.Context and returns an error api.Echart named echart
func Status(ctx context.Context) (echart api.Echart, err error) {
	//create a new dreamland.Client pointer named dreamClient
	var dreamClient *dreamland.Client
	//initailize dreamClient to the Client function with input ctx, while also checking for erros
	dreamClient, err = Client(ctx)
	//major coding error, we are not returning anything
	//we should  return nil,err
	//also this is supposed to be an error check
	if err != nil {
		return
	}
	//selectedUniverse var is assigned to env.GetCustomNetworkUrl()
	selectedUniverse, _ := env.GetCustomNetworkUrl()
	//set universe to the result of dreamClient.Universe() method (or possibly constructor of a subclass) working with selectedUniverse as input
	universe := dreamClient.Universe(selectedUniverse)
	//set echart and err to Universe.Status() result
	echart, err = universe.Status()
	//why is nothing returned
	//there is no error check here
	return
}
//HTTPPort function takes a ctx of type context.Context and string name, and returns an int and error)
//this is also a regular function not a singleton
//returns HTTPPort number or error given a name 
func HTTPPort(ctx context.Context, name string) (int, error) {
	//set echart,err to the results of a call to Status with ctx input. Howevere Status does not return anything in the function as it should have, something will break
	echart, err := Status(ctx)
	//we will assume that Status returns what it should from now on
	//error check for Status(ctx) call, if it breaks return 0 and err
	if err != nil {
		return 0, err
	}
	//a for loop goes though each node in echart.Nodes if error check succeeds
	for _, node := range echart.Nodes {
		//if name in input and node.Name name match, declare new httpPort valr with node.Value["http"] as value and error variable ok
		if strings.Contains(node.Name, name) {
			httpPort, ok := node.Value["http"]
			//error check if ok fails, return 0 for port and a fmt.Errorf which prints an error indicating that an http port is not set up
			if !ok {
				return 0, fmt.Errorf("http port for `%s` not set", name)
			}
			//otherwise return httpPort and nil for error
			return httpPort, nil
		}
	}
	//return 0 and an atomatic print message if string name is not found in echart
	return 0, fmt.Errorf("node `%s` not found", name)
}
