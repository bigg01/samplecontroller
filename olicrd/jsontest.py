import json

from pprint import pprint

jj = """
{
    "apiVersion": "v1",
    "items": [
        {
            "apiVersion": "o.guggenbuehl.local/v1",
            "kind": "NetGate",
            "metadata": {
                "clusterName": "",
                "creationTimestamp": "2018-06-17T18:37:08Z",
                "labels": {
                    "netzone": "v12"
                },
                "name": "oracle1",
                "namespace": "default",
                "resourceVersion": "204833",
                "selfLink": "/apis/o.guggenbuehl.local/v1/namespaces/default/netgates/oracle1",
                "uid": "710c8878-725d-11e8-bc98-8a9fa73b8f2c"
            },
            "spec": {
                "hostname": "oracle.external",
                "namespace": "egress-v12",
                "netzone": "v12",
                "port": 1521
            }
        }
    ],
    "kind": "List",
    "metadata": {
        "resourceVersion": "",
        "selfLink": ""
    }
}
"""

#print(json.dumps(jj))

bb = json.loads(jj)


for a in bb['items']:
    #pprint(bb[a])
    pprint(a['spec']['hostname'])
    pprint(a['spec'])
   
