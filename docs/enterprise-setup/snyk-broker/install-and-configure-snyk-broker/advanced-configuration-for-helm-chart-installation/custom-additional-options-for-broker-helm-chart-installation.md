# Broker Helm 차트 설치를 위한 사용자 정의 추가 옵션

환경 변수를 사용하여 추가 옵션을 주입해야 하는 경우, `override.yaml` 값을 파일로 사용하여 필요한 추가 환경 변수를 추가하십시오.

`--values override.yaml`를 추가하면 해당 값이 배포에 로드됩니다. 예를 들면:

```
helm install snyk-broker-chart snyk-broker/snyk-broker \
             --set scmType=github-com \
             --set brokerToken=<ENTER_BROKER_TOKEN> \
             --set scmToken=<ENTER_REPO_TOKEN> \
             --set brokerClientUrl=<ENTER_BROKER_CLIENT_URL>:<ENTER_BROKER_CLIENT_PORT> \
             --values override.yaml \
             -n snyk-broker --create-namespace
```

더 편리하다면 `override.yaml` 파일을 사용하지 않고 동일한 작업을 인라인으로 수행할 수 있습니다.

```
helm install snyk-broker-chart snyk-broker/snyk-broker \
             --set scmType=github-com \
             --set brokerToken=<ENTER_BROKER_TOKEN> \
             --set scmToken=<ENTER_REPO_TOKEN> \
             --set brokerClientUrl=<ENTER_BROKER_CLIENT_URL>:<ENTER_BROKER_CLIENT_PORT> \
             --set env[0].name=myEnvVarName \
             --set env[0].value=myEnvVarValue \
             --set env[1].name=myOtherEnvVarName \
             --set env[1].value=myOtherEnvVarValue \
             -n snyk-broker --create-namespace
```

값 파일에 추가하여 차트에 사용자 정의 Kubernetes 자원 및 객체를 추가할 수 있습니다.

Kubernetes 옵션 및 객체의 다양한 조합이 사양 계층의 여러 수준에서 사용 가능합니다. 배포, 컨테이너 및 pod에 따라 다릅니다. 이러한 정보는 [기본 `values.yaml` 파일](https://github.com/snyk/snyk-broker-helm/blob/a805f97235ba6b004df7a38c93ee94e399b699b7/charts/snyk-broker/values.yaml#L403)의 `extraObjects`, `extraVolumes`, `extraVolumeMounts`, 및 `extraPodSpecs`로 나열됩니다.

적합한 구문을 사용하고 `helm template` 명령을 사용하여 렌더링된 `yaml`을 유효성 검사하는 것을 잊지 마십시오.