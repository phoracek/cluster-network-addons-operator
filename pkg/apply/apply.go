package apply

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/api/equality"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	uns "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/types"
	k8sclient "sigs.k8s.io/controller-runtime/pkg/client"
)

// ApplyObject applies the desired object against the apiserver,
// merging it with any existing objects if already present.
func ApplyObject(ctx context.Context, client k8sclient.Client, obj *uns.Unstructured) error {
	name := obj.GetName()
	namespace := obj.GetNamespace()
	if name == "" {
		return errors.Errorf("object %s has no name", obj.GroupVersionKind().String())
	}
	gvk := obj.GroupVersionKind()
	// used for logging and errors
	objDesc := fmt.Sprintf("(%s) %s/%s", gvk.String(), namespace, name)
	log.Printf("reconciling %s", objDesc)

	if err := IsObjectSupported(obj); err != nil {
		return errors.Wrapf(err, "object %s unsupported", objDesc)
	}

	// Get existing
	existing := &uns.Unstructured{}
	existing.SetGroupVersionKind(gvk)
	err := client.Get(ctx, types.NamespacedName{Name: obj.GetName(), Namespace: obj.GetNamespace()}, existing)

	if err != nil && apierrors.IsNotFound(err) {
		log.Printf("does not exist, creating %s", objDesc)
		err := client.Create(ctx, obj)
		if err != nil {
			return errors.Wrapf(err, "could not create %s", objDesc)
		}
		log.Printf("successfully created %s", objDesc)
		return nil
	}
	if err != nil {
		return errors.Wrapf(err, "could not retrieve existing %s", objDesc)
	}

	// Merge the desired object with what actually exists
	if err := MergeObjectForUpdate(existing, obj); err != nil {
		return errors.Wrapf(err, "could not merge object %s with existing", objDesc)
	}
	if !equality.Semantic.DeepEqual(existing, obj) {
		if err := client.Update(ctx, obj); err != nil {
			// In older versions of the operator, we used daemon sets of type 'extensions/v1beta1', later we
			// changed that to 'apps/v1'. Because of this change, we are not able to seamlessly upgrade using
			// only Update methods. Following code handles this exception by deleting the old daemon set and
			// creating a new one.
			// TODO: Upgrade transaction should be handled by each component module separately. Once we make
			// that possible, this exception should be dropped.
			bridgeMarkerDaemonSetUpdateError := "DaemonSet.apps \"bridge-marker\" is invalid: spec.selector: Invalid value: v1.LabelSelector{MatchLabels:map[string]string{\"name\":\"bridge-marker\"}, MatchExpressions:[]v1.LabelSelectorRequirement(nil)}: field is immutable"
			if strings.Contains(err.Error(), bridgeMarkerDaemonSetUpdateError) {
				log.Print("update failed due to change in DaemonSet API group; removing original object and recreating")
				if err := client.Delete(ctx, existing); err != nil {
					return errors.Wrapf(err, "could not delete %s", objDesc)
				}
				if err := client.Create(ctx, obj); err != nil {
					return errors.Wrapf(err, "could not create %s", objDesc)
				}
				log.Print("update of conflicting DaemonSet was successful")
			}

			return errors.Wrapf(err, "could not update object %s", objDesc)
		}
		log.Print("update was successful")
	}

	return nil
}

// DeleteObject deletes an object in the apiserver
func DeleteObject(ctx context.Context, client k8sclient.Client, obj *uns.Unstructured) error {
	name := obj.GetName()
	namespace := obj.GetNamespace()
	if name == "" {
		return errors.Errorf("object %s has no name", obj.GroupVersionKind().String())
	}

	gvk := obj.GroupVersionKind()
	// used for logging and errors
	objDesc := fmt.Sprintf("(%s) %s/%s", gvk.String(), namespace, name)
	log.Printf("Handling deletion of %s", objDesc)

	// Get existing
	existing := &uns.Unstructured{}
	existing.SetGroupVersionKind(gvk)
	err := client.Get(ctx, types.NamespacedName{Name: obj.GetName(), Namespace: obj.GetNamespace()}, existing)
	if err != nil {
		if !apierrors.IsNotFound(err) {
			return errors.Wrapf(err, "could not retrieve existing %s", objDesc)
		}
		return nil
	}

	if err := client.Delete(ctx, existing); err != nil {
		return errors.Wrapf(err, "could not delete %s", objDesc)
	}

	return nil
}
