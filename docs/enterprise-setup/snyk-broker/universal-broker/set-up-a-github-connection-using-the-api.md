# API를 사용하여 GitHub 연결 설정하기

이 페이지는 유니버셜 브로커와 연결된 GitHub을 설정하기 위한 Snyk API 예제를 제공합니다. 필요한 통합별로 여러 연결을 반복적으로 설정할 수 있습니다.

`Snyk-broker-config` 명령어를 사용하는 것이 쉬운 경험을 제공합니다. API는 자동화 및 더 많은 제어를 허용합니다. 브로커 배포, 자격 증명, 및 연결에 대한 명확한 이해가 필요합니다.

{% hint style="info" %}
아래 명령어 중 어떤 것이든 필요한 경우 `api.snyk.io`를 지역 동등물로 교체하십시오. 예를 들어 `api.eu.snyk.io`와 같이.
{% endhint %}

## 조직을 위한 Broker 앱 설치 <a href="#id-1-install-the-broker-app-on-your-org" id="id-1-install-the-broker-app-on-your-org"></a>

조직 수준에서 유니버셜 브로커 앱을 설치합니다. 그룹 수준 설치는 지원되지 않습니다. [이 조직에 Snyk 앱 설치](../../../snyk-api/reference/apps.md#orgs-org_id-apps-installs) 엔드포인트를 사용합니다. 다음은 API를 호출할 때 사용할 App ID입니다:

`Snyk Broker App ID: cb43d761-bd17-4b44-9b6c-e5b8ad077d33`

## 배포 생성 <a href="#id-2-create-your-deployment" id="id-2-create-your-deployment"></a>

다음 명령을 사용하여 배포를 생성하세요.

```
curl --location --request POST 'https://api.snyk.io/rest/tenants/TENANT_ID/brokers/installs/INSTALL_ID/deployments?version=2024-02-08~experimental' \
--header 'Content-Type: application/vnd.api+json' \
--header 'Authorization: token YOUR_SNYK_TOKEN' \
--data-raw '{
    "data": {
        "type": "broker_deployment",
        "attributes": {
            "broker_app_installed_in_org_id":"ORG_ID_WHERE_APP_WAS_INSTALLED",
            "metadata": {
                "deployment_name": "My Universal Broker Deployment",
                "cluster": "Cluster X Region Y or whatever you need to not lose your deployment."
            }
        }
    }
}'
```

위 명령을 실행하면 `DEPLOYMENT_ID`(예: `data.id`)가 반환됩니다.

```
{
    ...
    "data": {
        "id": "12345678-1234-1234-1234-123456789012",
        "type": "broker_deployment",
        "attributes": {
            "install_id": "12345678-1234-1234-1234-123456789012",
            "metadata": {
                "deployment_name": "My Universal Broker Deployment",
                "cluster": "Cluster X Region Y or whatever you need to not lose your deployment."
            }
        }
    },
    ...
} 
```

이제 Broker 클라이언트를 실행할 준비가 되었습니다.

## Broker 배포 실행 <a href="#run-your-broker-deployment_1" id="run-your-broker-deployment_1"></a>

필요에 따라 `-e BROKER_SERVER_URL=https://broker.REGION.snyk.io \`와 같이 원하는 환경을 타겟팅하세요.

```
docker run --restart=always \
    -p 8000:8000 \
    -e ACCEPT_CODE=true \
    -e DEPLOYMENT_ID=<DEPLOYMENTID> \
    -e CLIENT_ID=<CLIENTID> \
    -e CLIENT_SECRET=<CLIENTSECRET> \
    -e UNIVERSAL_BROKER_ENABLED=true \
    -e PORT=8000 \
    -e BROKER_HA_MODE_ENABLED=true \
snyk/broker:universal
```

명령 실행 시 다음 메시지가 출력에 표시되어야 합니다.

```
{"name":"snyk-broker","hostname":"c918a4e1535a","pid":1,"level":30,"msg":"Found deployment 12345678-1234-1234-1234-123456789012. Waiting for connections. (polling)","time":"2024-06-18T20:15:20.572Z","v":0}
```

## 자격 증명 참조 생성 <a href="#id-3-create-your-credentials-references" id="id-3-create-your-credentials-references"></a>

```
curl --location --request POST 'https://api.snyk.io/rest/tenants/TENANT_ID/brokers/installs/INSTALL_ID/deployments/DEPLOYMENT_ID/credentials?version=2024-02-08~experimental' \
--header 'Content-Type: application/vnd.api+json' \
--header 'Authorization: token YOUR_SNYK_TOKEN' \
--data-raw '{
    "data": {
        "type": "deployment_credential",
        "attributes": [{
            "comment": "This is github token service account XYZ (example)",
            "environment_variable_name": "MY_GITHUB_TOKEN",
            "type": "github"
        }]
    }
}'
```

이 명령은 자격 증명 참조 ID(`data.id`)를 반환합니다.

```
{
    ...
    "data": {
        "id": "12345678-1234-1234-1234-123456789012",
        "type": "deployment_credential",
        "attributes": {
            "comment": "test cred",
            "deployment_id": "12345678-1234-1234-1234-123456789012",
            "environment_variable_name": "MY_GITHUB_TOKEN",
            "type": "github"
        }
    },
    ...
}
```

한 번의 호출에서 최대 10개의 자격 증명을 생성할 수 있으며, 서로 다른 유형의 객체를 추가할 수도 있습니다.

```
...
"attributes": [{
            "comment": "This is github token service account XYZ (example)",
            "environment_variable_name": "MY_GITHUB_TOKEN",
            "type": "github"
        },
        {
            "comment": "This is my other github token service account ABC (example)",
            "environment_variable_name": "MY_OTHER_GITHUB_TOKEN",
            "type": "github"
        }]
...
```

## 연결 생성 <a href="#id-4-create-your-connections" id="id-4-create-your-connections"></a>

```
curl --location --request POST 'https://api.snyk.io/rest/tenants/TENANT_ID/brokers/installs/INSTALL_ID/deployments/DEPLOYMENT_ID/connections?version=2024-02-08~experimental' \
--header 'Content-Type: application/vnd.api+json' \
--header 'Authorization: token YOUR_SNYK_TOKEN' \
--data-raw '{
    "data": {
        "type": "broker_connection",
        "attributes": {
            "name": "my github connection",
            "configuration": {
                "required": {
                    "github_token": "CRED_REFERENCE_UUID",
                    "broker_client_url":"http://your.broker.hostname:port"
                },
                "type": "github"
            },
            "deployment_id": "DEPLOYMENT_ID"
        }
    }
}'
```

이 호출은 연결 ID(`data.id`)를 반환하며, 예제에서 확인할 수 있습니다.

```
{
    ...
    "data": {
        "id": "12345678-1234-1234-1234-123456789012",
        "type": "broker_connection",
        "attributes": {
            "deployment_id": "12345678-1234-1234-1234-123456789012",
            "name": "my github connection",
            "secrets": {
                "primary": {
                    "encrypted": "not-yet-implemented",
                    "expires_at": "1970-01-01T00:00:00.000Z",
                    "nonce": "not-yet-implemented"
                },
                "secondary": {
                    "encrypted": "not-yet-implemented",
                    "expires_at": "1970-01-01T00:00:00.000Z",
                    "nonce": "not-yet-implemented"
                }
            },
            "configuration": {
                "required": {
                    "broker_client_url": "dummy",
                    "github_token": "${MY_GITHUB_TOKEN}"
                },
                "type": "github"
            }
        }
    },
    ...
}
```

자격 증명 참조가 누락된 경우 다음과 같은 메시지가 표시됩니다.

```
{"name":"snyk-broker","hostname":"029cda64bd98","pid":1,"level":50,"connection":"my github connection","msg":"Connection is missing environment variable value MY_GITHUB_TOKEN. Connection is disabled till value is provided. Restart broker once added.","time":"2024-06-18T14:29:06.910Z","v":0}

{"name":"my github connection","hostname":"029cda64bd98","pid":1,"level":50,"id":"12345678-1234-1234-1234-123456789012","msg":"Connection is disabled due to (a) missing environment variable(s). Please provide the value and restart the broker client.","time":"2024-06-18T14:29:06.911Z","v":0}
```
