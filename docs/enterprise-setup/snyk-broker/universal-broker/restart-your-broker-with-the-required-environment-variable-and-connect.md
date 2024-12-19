# 필요한 환경 변수와 함께 Broker를 다시 시작하고 연결하기

## 필요한 환경 변수를 사용하여 Broker를 다시 시작하기 <a href="#restart-your-broker-with-required-environment-variable" id="restart-your-broker-with-required-environment-variable"></a>

지역 Snyk 인스턴스에서는 `-e BROKER_SERVER_URL=https://broker.REGION.snyk.io \`를 사용합니다.

| <pre><code>docker run --restart=always \
    -p 8000:8000 \
    -e ACCEPT_CODE=true \
    -e DEPLOYMENT_ID=<DEPLOYMENTID> \
    -e CLIENT_ID=<CLIENTID> \
    -e CLIENT_SECRET=<CLIENTSECRET> \
    -e UNIVERSAL_BROKER_ENABLED=true \
    -e MY_GITHUB_TOKEN=GITHUB_TOKEN_VALUE \
    -e PORT=8000 \
    -e BROKER_HA_MODE_ENABLED=true \
snyk/broker:universal
</code></pre> |
| --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |

이 시점에서 Broker는 다음 메시지를 표시합니다:

| <pre><code>{"name":"my github connection","hostname":"ae7d64e0edac","pid":1,"level":30,"id":"12345678-1234-1234-1234-123456789012","msg":"Connection (my github connection) not in use by any orgs. Will check periodically and create connection when in use.","time":"2024-06-18T20:21:24.382Z","v":0}
</code></pre> |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |

이제 조직 통합에서 연결을 사용할 수 있습니다.

## 조직 통합에 연결하여 연결 사용하기 <a href="#id-5-connect-your-org-integration-to-use-a-connection" id="id-5-connect-your-org-integration-to-use-a-connection"></a>

| <pre><code>curl --location --request POST 'https://api.snyk.io/rest/tenants/TENANT_ID/brokers/connections/CONNECTION_ID/orgs/ORG_ID/integration?version=2024-02-08~experimental' \
--header 'Content-Type: application/vnd.api+json' \
--header 'Authorization: token YOUR_SNYK_TOKEN' \
--data-raw '{
    "data": {
                "integration_id": "INTEGRATION_ID",
                "type": "github"
    }
}'
</code></pre> |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |

필요에 따라 조직을 연결할 때마다 조직을 다시 연결합니다.