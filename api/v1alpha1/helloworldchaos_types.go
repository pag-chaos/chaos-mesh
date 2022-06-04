package v1alpha1

import (
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:object:root=true
// +chaos-mesh:experiment
// +chaos-mesh:oneshot=true

// HelloWorldChaos is the Schema for the helloworldchaos API
type HelloWorldChaos struct {
    metav1.TypeMeta   `json:",inline"`
    metav1.ObjectMeta `json:"metadata,omitempty"`

    Spec   HelloWorldChaosSpec   `json:"spec"`
    Status HelloWorldChaosStatus `json:"status,omitempty"`
}

// HelloWorldChaosSpec is the content of the specification for a HelloWorldChaos
type HelloWorldChaosSpec struct {
    // ContainerSelector specifies target
    ContainerSelector `json:",inline"`

    // Duration represents the duration of the chaos action
    // +optional
    Duration *string `json:"duration,omitempty"`
}

// HelloWorldChaosStatus represents the status of a HelloWorldChaos
type HelloWorldChaosStatus struct {
    ChaosStatus `json:",inline"`
}

// GetSelectorSpecs is a getter for selectors
func (obj *HelloWorldChaos) GetSelectorSpecs() map[string]interface{} {
    return map[string]interface{}{
        ".": &obj.Spec.ContainerSelector,
    }
}
