TOKEN=$(oc whoami -t)

ENDPOINT="192.168.64.2:8443"

curl -k \
    -H "Authorization: Bearer ${TOKEN}" \
    -H 'Accept: application/json' \
    https://${ENDPOINT}/apis/apiextensions.k8s.io/v1beta1/customresourcedefinitions


curl -k \
    -H "Authorization: Bearer ${TOKEN}" \
    -H 'Accept: application/json' \
    https://${ENDPOINT}/apis/o.guggenbuehl.local/v1/namespaces/default/netgates


curl -k \
    -H "Authorization: Bearer ${TOKEN}" \
    -H 'Accept: application/json' \
    https://${ENDPOINT}/apis/o.guggenbuehl.local/v1/


