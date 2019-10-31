package spa

import (
	"context"
	"reflect"

	ahorav1alpha1 "github.com/ahora/spa-operator/pkg/apis/ahora/v1alpha1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/api/extensions/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/intstr"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var log = logf.Log.WithName("controller_spa")

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new SPA Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileSPA{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("spa-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource SPA
	err = c.Watch(&source.Kind{Type: &ahorav1alpha1.SPA{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	/*
		// TODO(user): Modify this to be the types you create that are owned by the primary resource
		// Watch for changes to secondary resource Pods and requeue the owner SPA
		err = c.Watch(&source.Kind{Type: &appsv1.Deployment{}}, &handler.EnqueueRequestForOwner{
			IsController: true,
			OwnerType:    &ahorav1alpha1.SPA{},
		})
	*/
	if err != nil {
		return err
	}

	return nil
}

// blank assignment to verify that ReconcileSPA implements reconcile.Reconciler
var _ reconcile.Reconciler = &ReconcileSPA{}

// ReconcileSPA reconciles a SPA object
type ReconcileSPA struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme
}

// Reconcile reads that state of the cluster for a SPA object and makes changes based on the state read
// and what is in the SPA.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  This example creates
// a Pod as an example
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.

func (r *ReconcileSPA) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling SPA")

	// Fetch the SPA instance
	instance := &ahorav1alpha1.SPA{}
	err := r.client.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	// Define a new Deployment object
	newDeploymentInstance := newDeploymentForCR(instance)
	newServiceInstance := newServiceForCR(instance)
	newIngressInstance := newIngress(instance)

	// Set SPA deployment, Service & ingress as the owner and controller
	if err := controllerutil.SetControllerReference(instance, newDeploymentInstance, r.scheme); err != nil {
		return reconcile.Result{}, err
	}

	if err := controllerutil.SetControllerReference(instance, newServiceInstance, r.scheme); err != nil {
		return reconcile.Result{}, err
	}

	if err := controllerutil.SetControllerReference(instance, newIngressInstance, r.scheme); err != nil {
		return reconcile.Result{}, err
	}

	// Check if this Deployment already exists
	foundDeployment := &appsv1.Deployment{}
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: newDeploymentInstance.Name, Namespace: newDeploymentInstance.Namespace}, foundDeployment)
	if err != nil && errors.IsNotFound(err) {
		reqLogger.Info("Creating a new Deployment", "Deployment.Namespace", newDeploymentInstance.Namespace, "deployment.Name", newDeploymentInstance.Name)
		err = r.client.Create(context.TODO(), newDeploymentInstance)
		if err != nil {
			return reconcile.Result{}, err
		}
	} else if err != nil {
		return reconcile.Result{}, err
	} else {

		if !reflect.DeepEqual(newDeploymentInstance.Spec, foundDeployment.Spec) {
			ctx := context.TODO()
			err := r.client.Update(ctx, newDeploymentInstance)

			if err != nil {
				return reconcile.Result{}, err
			} else {
				reqLogger.Info("Updated a new Deployment", "Deployment.Namespace", newDeploymentInstance.Namespace, "deployment.Name", newDeploymentInstance.Name)
			}
		}
	}

	// Check if this Ingress already exists
	foundIngress := &v1beta1.Ingress{}
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: newIngressInstance.Name, Namespace: newIngressInstance.Namespace}, foundIngress)
	if err != nil && errors.IsNotFound(err) {
		reqLogger.Info("Creating a new Ingress", "Ingress.Namespace", newIngressInstance.Namespace, "Ingress.Name", newIngressInstance.Name)
		err = r.client.Create(context.TODO(), newIngressInstance)
		if err != nil {
			return reconcile.Result{}, err
		}
	} else if err != nil {
		return reconcile.Result{}, err
	}

	// Check if this Service already exists
	foundService := &corev1.Service{}
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: newServiceInstance.Name, Namespace: newServiceInstance.Namespace}, foundService)
	if err != nil && errors.IsNotFound(err) {
		reqLogger.Info("Creating a new service", "Service.Namespace", newServiceInstance.Namespace, "Service.Name", newServiceInstance.Name)
		err = r.client.Create(context.TODO(), newServiceInstance)
		if err != nil {
			return reconcile.Result{}, err
		}

		// Deployment created successfully - don't requeue
		return reconcile.Result{}, nil
	} else if err != nil {
		return reconcile.Result{}, err
	}

	return reconcile.Result{}, nil
}

func newDeploymentForCR(cr *ahorav1alpha1.SPA) *appsv1.Deployment {
	labels := map[string]string{
		"app": cr.Name,
	}
	return &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      cr.Name + "-deployment",
			Namespace: cr.Namespace,
			Labels:    labels,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: cr.Spec.Replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: labels,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labels,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Resources:      cr.Spec.Resources,
							ReadinessProbe: cr.Spec.ReadinessProbe,
							LivenessProbe:  cr.Spec.LivenessProbe,
							Name:           "spa",
							Image:          "ahora/spa:0.0.1",
							Env: []corev1.EnvVar{
								{
									Name:  "SPAARCHIVEURL",
									Value: cr.Spec.SPAArchiveURL,
								},
							},
						},
					},
				},
			},
		},
	}
}

func newServiceForCR(cr *ahorav1alpha1.SPA) *corev1.Service {
	labels := map[string]string{
		"app": cr.Name,
	}
	return &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      cr.Name,
			Namespace: cr.Namespace,
			Labels:    labels,
		},
		Spec: corev1.ServiceSpec{
			Selector: labels,
			Ports: []corev1.ServicePort{
				{
					Name:     cr.Name,
					Protocol: corev1.ProtocolTCP,
					Port:     80,
				},
			},
		},
	}
}

func newIngress(cr *ahorav1alpha1.SPA) *v1beta1.Ingress {
	labels := map[string]string{
		"app": cr.Name,
	}

	//create paths for all paths from CR.
	paths := []v1beta1.HTTPIngressPath{
		{
			Backend: v1beta1.IngressBackend{
				ServiceName: cr.Name,
				ServicePort: intstr.FromInt(80),
			},
		},
	}

	for _, path := range cr.Spec.Paths {
		paths = append(paths, path)
	}

	//Create rules for all gosts
	var rules []v1beta1.IngressRule
	for _, host := range cr.Spec.Hosts {
		rule := v1beta1.IngressRule{
			Host: host,
			IngressRuleValue: v1beta1.IngressRuleValue{
				HTTP: &v1beta1.HTTPIngressRuleValue{
					Paths: paths,
				},
			},
		}
		rules = append(rules, rule)
	}

	return &v1beta1.Ingress{
		ObjectMeta: metav1.ObjectMeta{
			Name:        cr.Name,
			Namespace:   cr.Namespace,
			Labels:      labels,
			Annotations: cr.Annotations,
		},
		Spec: v1beta1.IngressSpec{
			TLS:   cr.Spec.TLS,
			Rules: rules,
		},
	}
}
