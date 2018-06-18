import requests
url = 'https://192.168.64.2:8443/apis/o.guggenbuehl.local/v1/namespaces/default/netgates'

token = 'sdmnY-K0kdiewWLVTEfjMHmJJMiz1JfO4VYDrSiwpKU'

headers = {'Authorization': 'Bearer ' + token}
re = requests.get(url, headers=headers, verify=False)

print re



