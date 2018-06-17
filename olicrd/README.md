### Idea
We create custom crd for egress per project definitions. An write a controller for it.
Or an custom haproxy who reade when it sturt up egress definitions.
can be used with haproxy or with https://github.com/bigg01/tcpproxy.



```
$ oc create -f netguard.yml

$ oc describe customresourcedefinition netguards.o.guggenbuehl.local
Name:		netguards.o.guggenbuehl.local
Namespace:	
Labels:		<none>
Annotations:	<none>
API Version:	apiextensions.k8s.io/v1beta1
Kind:		CustomResourceDefinition
Metadata:
  Creation Timestamp:	2018-06-17T14:21:59Z
  Resource Version:	1111
  Self Link:		/apis/apiextensions.k8s.io/v1beta1/customresourcedefinitions/netguards.o.guggenbuehl.local
  UID:			cbe81b41-7239-11e8-9fbd-4437e637cdb5
Spec:
  Group:	o.guggenbuehl.local
  Names:
    Kind:	NetGuard
    List Kind:	NetGuardList
    Plural:	netguards
    Short Names:
      gu
    Singular:	netguard
  Scope:	Namespaced
  Version:	v1
Status:
  Accepted Names:
    Kind:	NetGuard
    List Kind:	NetGuardList
    Plural:	netguards
    Short Names:
      gu
    Singular:	netguard
  Conditions:
    Last Transition Time:	<nil>
    Message:			no conflicts found
    Reason:			NoConflicts
    Status:			True
    Type:			NamesAccepted
    Last Transition Time:	2018-06-17T14:21:59Z
    Message:			the initial names have been accepted
    Reason:			InitialNamesAccepted
    Status:			True
    Type:			Established
Events:				<none>

$ oc replace -f testnetguard1.yml 
netguard "oracle1" replaced

$ oc describe netguards
Name:		oracle1
Namespace:	myproject
Labels:		<none>
Annotations:	<none>
API Version:	o.guggenbuehl.local/v1
Kind:		NetGuard
Metadata:
  Cluster Name:		
  Creation Timestamp:	2018-06-17T14:23:44Z
  Generation:		0
  Resource Version:	1459
  Self Link:		/apis/o.guggenbuehl.local/v1/namespaces/myproject/netguards/oracle1
  UID:			0abde987-723a-11e8-9fbd-4437e637cdb5
Spec:
  Hostname:	oracle.external
  Namespace:	egress-v12
  Netzone:	v12
  Port:		1521
Events:		<none>

```

```jq
 oc get NetGuards   -o json | jq  '.items[]| .spec '
{
  "hostname": "oracle.external",
  "namespace": "egress-v12",
  "netzone": "v12",
  "port": 1521
}

```
