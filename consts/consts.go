package consts

var HELP_INFO_TEXT string = `
kubectl controls the Kubernetes cluster manager.

Find more information at: https://kubernetes.io/docs/reference/kubectl/

Basic Commands (Intermediate):
  get             Display one or many resources
`

var HELP_INFO_GET_PODS string = `
Display one or many resources.

Prints a table of the most important information about the specified resources. You can filter the list using a label
selector and the --selector flag. If the desired resource type is namespaced you will only see results in your current
namespace unless you pass --all-namespaces.

 By specifying the output as 'template' and providing a Go template as the value of the --template flag, you can filter
the attributes of the fetched resources.

Use "kubectl api-resources" for a complete list of supported resources.

Examples:
  # List all pods in ps output format
  kubectl get pods
  
  # List all pods in ps output format in that namespace
  kubectl get pods -n you_namespac
`

var UNKOWN_COMMAND string = `
unkown command,see the help command to list the allowed commands.
`

var NO_RESOURCES string = `No resources found in internal namespace.`

var FLAG_NAMESPACE_WARGNING string = `error: flag needs an argument: 'n' in -n
See 'kubectl get --help' for usage.`
