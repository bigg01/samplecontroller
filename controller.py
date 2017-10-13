import json
import yaml
from kubernetes import client, config, watch
import os

DOMAIN = "kool.karmalabs.local"


def review_guitar(crds, event, obj):
    metadata = obj.get("metadata")
    goodguitars = ['fender', 'martin']
    if not metadata:
        print("No metadata in object, skipping: %s" % json.dumps(obj, indent=1))
        return
    name = metadata.get("name")
    namespace = metadata.get("namespace")
    obj["spec"]["review"] = True
    brand = obj["spec"]["brand"]
    if brand in goodguitars:
        obj["spec"]["comment"] = "this is a great instrument"
    else:
        obj["spec"]["comment"] = "this is shit"

    print("Updating: %s" % name)
    crds.replace_namespaced_custom_object(DOMAIN, "v1", namespace, "guitars", name, obj)


if __name__ == "__main__":
    if 'KUBERNETES_PORT' in os.environ:
        config.load_incluster_config()
    else:
        config.load_kube_config()
    v1 = client.ApiextensionsV1beta1Api()
    current_crds = [x['spec']['names']['kind'].lower() for x in v1.list_custom_resource_definition().to_dict()['items']]
    if 'guitar' not in current_crds:
        print("Creating guitar definition")
        with open('guitar.yml') as data:
            body = yaml.load(data)
        v1.create_custom_resource_definition(body)
    crds = client.CustomObjectsApi()

    print("Waiting for Guitars to come up...")
    stream = watch.Watch().stream(crds.list_cluster_custom_object, DOMAIN, "v1", "guitars")
    for event in stream:
        obj = event["object"]
        spec = obj.get("spec")
        if not spec:
            print("No 'spec' in object, skipping event: %s" % json.dumps(obj, indent=1))
            continue
        done = spec.get("review", False)
        if done:
            continue
        review_guitar(crds, event, obj)