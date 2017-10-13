# samplecontroller repository

[![](https://images.microbadger.com/badges/image/karmab/samplecontroller.svg)](https://microbadger.com/images/karmab/samplecontroller "Get your own image badge on microbadger.com")

This is a simple controller to demonstrate how to interact within kubernetes using python api and custom resource definitions

## Requisites

- a running kubernetes/openshift cluster
- If running on openshift, enough privileges as i m checking guitars cluster wide

```
oc new-project guitarcenter
oc adm policy add-cluster-role-to-user cluster-admin -z default -n guitarcenter
oc new-app karmab/samplecontroller
```

Note that the guitar custom resource definition gets created when launching the controller

## How to use

Create some guitars and see the review made for you

```
oc create -f crd/stratocaster.yml
oc create -f crd/lespaul.yml
oc get guitars -o yaml
```

## TODO

- build a basic flask ui to create and view guitar objects

## Copyright

Copyright 2017 Karim Boumedhel

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

## Problems?

Send me a mail at [karimboumedhel@gmail.com](mailto:karimboumedhel@gmail.com) !

Mac Fly!!!

karmab