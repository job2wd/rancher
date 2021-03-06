package v1

import (
	"github.com/rancher/norman/lifecycle"
	"k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

type RoleLifecycle interface {
	Create(obj *v1.Role) (*v1.Role, error)
	Remove(obj *v1.Role) (*v1.Role, error)
	Updated(obj *v1.Role) (*v1.Role, error)
}

type roleLifecycleAdapter struct {
	lifecycle RoleLifecycle
}

func (w *roleLifecycleAdapter) Create(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Create(obj.(*v1.Role))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *roleLifecycleAdapter) Finalize(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Remove(obj.(*v1.Role))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *roleLifecycleAdapter) Updated(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Updated(obj.(*v1.Role))
	if o == nil {
		return nil, err
	}
	return o, err
}

func NewRoleLifecycleAdapter(name string, clusterScoped bool, client RoleInterface, l RoleLifecycle) RoleHandlerFunc {
	adapter := &roleLifecycleAdapter{lifecycle: l}
	syncFn := lifecycle.NewObjectLifecycleAdapter(name, clusterScoped, adapter, client.ObjectClient())
	return func(key string, obj *v1.Role) (*v1.Role, error) {
		newObj, err := syncFn(key, obj)
		if o, ok := newObj.(*v1.Role); ok {
			return o, err
		}
		return nil, err
	}
}
