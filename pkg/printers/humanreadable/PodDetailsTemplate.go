package humanreadable

// PodDetailsTemplate is go template for printing pod details
const PodDetailsTemplate = `Name: {{.Metadata.Name}}
Namespace:	{{.Metadata.Namespace}}
Containers:{{range .Spec.Containers}}
  {{.Name}}:
    ContainerID: {{.ID}}
    Image: {{.Image}}
{{else}}
  (No containers)
{{end}}
`