# Helm을 사용하여 컨테이너 레지스트리 에이전트용 브로커 설치

[Broker Container Registry Agent를 Docker를 사용하여 설치](./)하려면 `CR_AGENT_URL` 매개변수가 필요하지만 Helm을 사용하여 설치할 때는 필요하지 않습니다. 환경 변수는 [Docker로 설치하는 경우](https://docs.snyk.io/snyk-admin/snyk-broker/snyk-broker-container-registry-agent#configuring-and-running-the-broker-client-for-container-registry-agent)에 정의되며 Helm을 사용하여 설치할 때도 적용됩니다.

```
helm install snyk-broker-chart snyk-broker/snyk-broker \
             --set scmType=container-registry-agent \
             --set brokerToken=<ENTER_BROKER_TOKEN> \
             --set crType=<ENTER_CR_TYPE> \
             --set crBase=<ENTER_CR_BASE_URL> \
             --set crUsername=<ENTER_CR_USERNAME> \
             --set crPassword=<ENTER_CR_PASSWORD> \
             -n snyk-broker --create-namespace
```

`crType`의 허용되는 값:

```
acr
artifactory-cr
digitalocean-cr
docker-hub
ecr
gcr
google-artifact-cr
github-cr
gitlab-cr
harbor-cr
quay-cr
nexus-cr
```

이후 섹션에서 설명된대로 `Elastic Container Registry` 및 `Digital Ocean Container Registry`는 특정 매개변수를 요구합니다.

## **Elastic Container Registry (ecr)**

* crRoleArn
* crRegion
* crExternalId

```
helm install snyk-broker-chart snyk-broker/snyk-broker \
             --set scmType=container-registry-agent \
             --set brokerToken=<ENTER_BROKER_TOKEN> \
             --set crType=ecr \
             --set crRoleArn=<ENTER_CR_ROLE_ARN> \
             --set crRegion=<ENTER_CR_REGION> \
             --set crExternalId=<ENTER_CR_EXTERNAL_ID> \
             -n snyk-broker --create-namespace
```

## **DigitalOcean Container Registry (digitalocean-cr)**

* crToken

```
helm install snyk-broker-chart snyk-broker/snyk-broker \
             --set scmType=container-registry-agent \
             --set brokerToken=<ENTER_BROKER_TOKEN> \
             --set crType=digitalocean-cr \
             --set crBase=<ENTER_CR_BASE_URL> \
             --set crToken=<ENTER_CR_TOKEN> \
             -n snyk-broker --create-namespace
```  