CustomResource
==========================

CRDの作成
-----------------------

`crd.yaml`


コントローラの作成
-------------------------

bashスクリプトを使用するのでWSL内で実行

```bash
$ go get k8s.io/gengo/args
$ go get k8s.io/gengo/examples/deepcopy-gen/generators
$ go get k8s.io/gengo/types
$ go get k8s.io/gengo/generator
$ go get k8s.io/apimachinery/pkg/apis/meta/v1
$ go get k8s.io/apimachinery/pkg/runtime
```

1. `foo` カスタムリソース

`pkg/apis/foo/v1alpha`の下に `doc.go`, `register.go`, `types.go` を作成

以下を実行。GOPATHからの相対パス
```bash
$ bash $GOPATH/src/k8s.io/code-generator/generate-groups.sh client,deepcopy github.com/dtamura/samplecontroller/pkg/client github.com/dtamura/samplecontroller/pkg/apis foo:v1alpha
```


1. コントローラの実装 `main.go`

1. ビルド・コンテナの作成

1. コントローラのデプロイ用のマニフェスト
`deploy.yaml`

```bash
$ kubectl apply -f deploy.yaml
role.rbac.authorization.k8s.io/foo-reader unchanged
rolebinding.rbac.authorization.k8s.io/foo-reader-rolebinding unchanged
deployment.apps/controller-main configured
```

1. fooリソースの作成
```bash
$ kubectl apply -f example-cr.yaml
foo.samplecontroller.k8s.io/foo-001 created
foo.samplecontroller.k8s.io/foo-002 created

$ kubectl get foo
NAME      AGE
foo-001   49s
foo-002   49s
```


1. コントローラのPodのログ
```
&{TypeMeta:{Kind: APIVersion:} ListMeta:{SelfLink:/apis/samplecontroller.k8s.io/v1alpha/namespaces/default/foos ResourceVersion:58359 Continue: RemainingItemCount:<nil>} Items:[{TypeMeta:{Kind:Foo APIVersion:samplecontroller.k8s.io/v1alpha} ObjectMeta:{Name:foo-001 GenerateName: Namespace:default SelfLink:/apis/samplecontroller.k8s.io/v1alpha/namespaces/default/foos/foo-001 UID:3c667369-7706-11ea-9ec0-4201ac100013 ResourceVersion:58256 Generation:1 CreationTimestamp:2020-04-05 06:25:25 +0000 UTC DeletionTimestamp:<nil> DeletionGracePeriodSeconds:<nil> Labels:map[] Annotations:map[kubectl.kubernetes.io/last-applied-configuration:{"apiVersion":"samplecontroller.k8s.io/v1alpha","kind":"Foo","metadata":{"annotations":{},"name":"foo-001","namespace":"default"},"spec":{"deploymentName":"deploy-foo-001","replicas":1}}
] OwnerReferences:[] Finalizers:[] ClusterName: ManagedFields:[]} Status:{Name:} Spec:{Name:}} {TypeMeta:{Kind:Foo APIVersion:samplecontroller.k8s.io/v1alpha} ObjectMeta:{Name:foo-002 GenerateName: Namespace:default SelfLink:/apis/samplecontroller.k8s.io/v1alpha/namespaces/default/foos/foo-002 UID:3c785e44-7706-11ea-9ec0-4201ac100013 ResourceVersion:58257 Generation:1 CreationTimestamp:2020-04-05 06:25:25 +0000 UTC DeletionTimestamp:<nil> DeletionGracePeriodSeconds:<nil> Labels:map[] Annotations:map[kubectl.kubernetes.io/last-applied-configuration:{"apiVersion":"samplecontroller.k8s.io/v1alpha","kind":"Foo","metadata":{"annotations":{},"name":"foo-002","namespace":"default"},"spec":{"deploymentName":"deploy-foo-002","hoge":"huga","replicas":1}}
] OwnerReferences:[] Finalizers:[] ClusterName: ManagedFields:[]} Status:{Name:} Spec:{Name:}}]}
&{TypeMeta:{Kind: APIVersion:} ListMeta:{SelfLink:/apis/samplecontroller.k8s.io/v1alpha/namespaces/default/foos ResourceVersion:58395 Continue: RemainingItemCount:<nil>} Items:[{TypeMeta:{Kind:Foo APIVersion:samplecontroller.k8s.io/v1alpha} ObjectMeta:{Name:foo-001 GenerateName: Namespace:default SelfLink:/apis/samplecontroller.k8s.io/v1alpha/namespaces/default/foos/foo-001 UID:3c667369-7706-11ea-9ec0-4201ac100013 ResourceVersion:58256 Generation:1 CreationTimestamp:2020-04-05 06:25:25 +0000 UTC DeletionTimestamp:<nil> DeletionGracePeriodSeconds:<nil> Labels:map[] Annotations:map[kubectl.kubernetes.io/last-applied-configuration:{"apiVersion":"samplecontroller.k8s.io/v1alpha","kind":"Foo","metadata":{"annotations":{},"name":"foo-001","namespace":"default"},"spec":{"deploymentName":"deploy-foo-001","replicas":1}}
] OwnerReferences:[] Finalizers:[] ClusterName: ManagedFields:[]} Status:{Name:} Spec:{Name:}} {TypeMeta:{Kind:Foo APIVersion:samplecontroller.k8s.io/v1alpha} ObjectMeta:{Name:foo-002 GenerateName: Namespace:default SelfLink:/apis/samplecontroller.k8s.io/v1alpha/namespaces/default/foos/foo-002 UID:3c785e44-7706-11ea-9ec0-4201ac100013 ResourceVersion:58257 Generation:1 CreationTimestamp:2020-04-05 06:25:25 +0000 UTC DeletionTimestamp:<nil> DeletionGracePeriodSeconds:<nil> Labels:map[] Annotations:map[kubectl.kubernetes.io/last-applied-configuration:{"apiVersion":"samplecontroller.k8s.io/v1alpha","kind":"Foo","metadata":{"annotations":{},"name":"foo-002","namespace":"default"},"spec":{"deploymentName":"deploy-foo-002","hoge":"huga","replicas":1}}
] OwnerReferences:[] Finalizers:[] ClusterName: ManagedFields:[]} Status:{Name:} Spec:{Name:}}]}
```


参考文献
--------------------
https://github.com/takaishi/hello2019/tree/master/hello-custom-controller
