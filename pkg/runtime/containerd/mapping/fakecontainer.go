package mapping

import (
	"context"

	"github.com/containerd/containerd"
	"github.com/containerd/containerd/containers"
	"github.com/containerd/containerd/errdefs"
	prototypes "github.com/gogo/protobuf/types"
	specs "github.com/opencontainers/runtime-spec/specs-go"
	"github.com/pkg/errors"
)

// FakeContainer is fake representation of model.Container
type FakeContainer struct {
	id        string
	labels    map[string]string
	isRunning bool
}

// ID fake impl.
func (c *FakeContainer) ID() string {
	return c.id
}

// Info fake impl.
func (c *FakeContainer) Info() containers.Container {
	return containers.Container{
		Labels: c.labels,
	}
}

// Delete fake impl.
func (c *FakeContainer) Delete(context.Context, ...containerd.DeleteOpts) error {
	return nil
}

// NewTask fake impl.
func (c *FakeContainer) NewTask(context.Context, containerd.IOCreation, ...containerd.NewTaskOpts) (task containerd.Task, err error) {
	return task, err
}

// Spec fake impl.
func (c *FakeContainer) Spec() (*specs.Spec, error) {
	return &specs.Spec{
		Process: &specs.Process{},
	}, nil
}

// Task fake impl.
func (c *FakeContainer) Task(context.Context, containerd.IOAttach) (task containerd.Task, err error) {
	if c.isRunning {
		return task, nil
	}
	return task, errors.Wrapf(errdefs.ErrNotFound, "Task not found")
}

// Image fake impl.
func (c *FakeContainer) Image(context.Context) (image containerd.Image, err error) {
	return image, nil
}

// Labels fake impl.
func (c *FakeContainer) Labels(context.Context) (labels map[string]string, err error) {
	return labels, nil
}

// SetLabels fake impl.
func (c *FakeContainer) SetLabels(context.Context, map[string]string) (labels map[string]string, err error) {
	return labels, nil
}

// Extensions fake impl.
func (c *FakeContainer) Extensions() map[string]prototypes.Any {
	return map[string]prototypes.Any{}
}

func fakeRunningContainer(namespace, podName, containerName string) containerd.Container {
	return newFakeContainer(namespace, podName, containerName, true)
}

func fakeCreatedContainer(namespace, podName, containerName string) containerd.Container {
	return newFakeContainer(namespace, podName, containerName, false)
}

func newFakeContainer(namespace, podName, containerName string, isRunning bool) containerd.Container {
	labels := map[string]string{}
	labels[buildLabelKeyFor(podNameLabel)] = podName

	return &FakeContainer{
		id:        containerName,
		labels:    labels,
		isRunning: isRunning,
	}
}