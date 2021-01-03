package main

import (
	"encoding/json"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
)

type SvcPermission struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SvcPermissionSpec `json:"spec,omitempty"`
}

type SvcPermissionSpec struct {

	//授权的用户id
	UserName string `json:"userName"`
}

type SvcPermissionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []SvcPermission `json:"items,omitempty"`
}

var versionResource = schema.GroupVersionResource{Group: "lstack.k8s.io", Version: "v1", Resource: "svcpermissions"}

type SvcPClient struct {
	client dynamic.Interface
}

func (s *SvcPClient) List(option metav1.ListOptions) (*SvcPermissionList, error) {
	unstructuredList, err := s.client.Resource(versionResource).List(option)
	if err != nil {
		return nil, err
	}
	list := &SvcPermissionList{}
	_ = MarshalResource(unstructuredList, list)
	return list, nil
}

func MarshalResource(src interface{}, v interface{}) error {
	bytes, err := json.Marshal(src)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bytes, v)
	if err != nil {
		return err
	}
	return nil
}
