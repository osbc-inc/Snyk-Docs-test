# Kubernetes 비밀 및 Helm 차트 설치

Snyk 브로커 헬름 차트의 버전 `2.8.0`부터 외부 시크릿이 지원됩니다.&#x20;

이 기능을 활성화하려면 `values.yaml`에서 `useExternalSecrets`를 `true`로 설정하거나 `--set externalSecrets=true`를 사용하십시오.&#x20;

필요한 시크릿 목록을 얻으려면 Helm 설치의 드라이 런을 수행하십시오. 이 작업은 사용 중인 쿠버네티스 환경에 대한 변경 사항을 만들지 않지만 다음을 필요로 합니다:

```
helm install snyk-broker-chart \
  snyk-broker/snyk-broker \
  --set externalSecrets=true \
  --set scmType=<your-scm-type> \
  --dry-run=client
```

시크릿 목록과 해당되는 이름 및 값이 생성됩니다. 다음 예시에서는 `scmType=nexus`를 사용합니다:

```
### 시크릿 생성 비활성화 ###

브로커가 설치될 동일한 네임스페이스에 다음 네 개의 시크릿이 존재해야 합니다. 각 시크릿은 하나의 키-값 쌍을 포함합니다. `< >` 문자로 나타낸 값은 귀하의 비밀 데이터를 추가해야 함을 나타냅니다.

-> NAME:KEY <VALUE>
-> nexus-broker-token-snyk-broker-chart:nexus-broker-token-key <your-broker-token>
-> nexus-base-nexus-url-snyk-broker-chart:nexus-base-nexus-url <BASE_NEXUS_URL>
-> nexus-nexus-url-snyk-broker-chart:nexus-nexus-url <NEXUS_URL>
-> nexus-broker-client-validation-url-snyk-broker-chart:nexus-broker-client-validation-url <BROKER_CLIENT_VALIDATION_URL>
```

이 예시에서는 네 개의 시크릿이 필요하며, 각 시크릿은 같은 네임스페이스 내에 존재해야 하며 각각 하나의 키-값 쌍을 포함해야 합니다. `< >` 문자로 나타낸 값은 귀하의 비밀 데이터를 추가해야 합니다.

## 시크릿 및 키 이름 변경

다음의 각 Helm 값은 `name`과 `key`를 지원하여 Snyk 브로커 헬름 차트가 해당 시크릿 이름과 시크릿 내의 특정 키를 참조할 수 있게 합니다:

* `externalCredentialSecret` (`artifactory`, `nexus` 또는 `nexus2`가 아닌 모든 브로커 유형에 사용됨: 해당 브로커 유형과 관련된 비밀번호 또는 PAT)
* `brokerTokenSecret` (브로커 토큰에 사용됨)
* `scmTokenPoolSecret` ([자격 증명 풀링](../advanced-configuration-for-snyk-broker-docker-installation/credential-pooling-with-docker-and-helm.md)이 활성화된 경우 사용됨)
* `artifactoryUrlSecret` (`artifactory` 전용)
* `baseNexusUrlSecret` (`nexus` 및 `nexus2` 전용)
* `nexusUrlSecret` (`nexus` 및 `nexus2` 전용)
* `brokerClientValidationUrlSecret` (`nexus` 및 `nexus2` 전용, `artifactory`에 대해 선택적으로 사용됨)

예를 들어, 귀하의 Kubernetes 클러스터에 다음 형식의 브로커 토큰이 포함된 시크릿이 있는 경우:

```
apiVersion: v1
kind: Secret
metadata:
  name: snyk-broker-secrets
type: Opaque
data:
  org-x-broker-token: <broker-token-here>
```

다음과 같이 설정하십시오:

```
useExternalSecrets: true

brokerTokenSecret:
  name: snyk-broker-secrets
  key: org-x-broker-token
```

헬름 차트는 시크릿 `snyk-broker-secrets`의 `org-x-broker-token` 키에 포함된 내용을 브로커 토큰에 대해 참조합니다.

## 부분 외부 시크릿

`useExternalSecrets`가 true인 경우, 브로커 헬름 차트는 시크릿에 대한 값이 제공되었는지 확인합니다 (예: `brokerToken=<your-broker-token>`)

* 값이 존재하는 경우, 시크릿을 평소대로 생성합니다.
* 값이 없는 경우, 외부 시크릿을 찾습니다.

이 방법으로, 일부 시크릿은 브로커 헬름 차트에 의해 제어되고, 다른 시크릿은 외부에서 제어될 수 있습니다:

```
scmType: github-com
brokerToken: <my-broker-token>
useExternalSecrets: true
githubToken: ""
```

이 값 세트는 다음 작업을 수행합니다:

* 제공된 브로커 토큰에 대한 시크릿을 생성합니다.
* 필수 GitHub 토큰에 대한 외부 시크릿을 참조합니다.

Helm 설치의 드라이 런을 수행하면 필요한 시크릿 이름과 키가 제공됩니다:

```
### 시크릿 생성 비활성화 ###

브로커 헬름 차트가 설치되는 동일한 네임스페이스에 다음 시크릿이 존재하는지 확인하십시오:

-> NAME:KEY <VALUE>
-> github-com-token-snyk-broker-chart:github-com-token-key <GITHUB_TOKEN>
```

브로커 토큰 시크릿은 브로커 헬름 차트에 직접 값을 제공하기 때문에 이 목록에서 제외됩니다.

## 다중 키가 있는 단일 외부 시크릿 사용

하나의 Kubernetes 시크릿에는 Snyk 브로커가 작동하는 데 필요한 모든 자격 증명이 포함될 수 있습니다. 예를 들어, `nexus` 유형의 브로커를 사용하는 경우 Kubernetes에 다음 시크릿이 있는 것으로 가정하십시오:

```
apiVersion: v1
kind: Secret
metadata:
  name: snyk-broker-secrets
type: Opaque
data:
  nexus-broker-token-key: <broker-token-here>
  nexus-nexus-url: https://user:pass@nexus.tld/myrepository
  nexus-base-nexus-url: https://user:pass@nexus.tld
  nexus-broker-client-validation-url: https://user:pass@nexus.tld/service/rest/v1/status/check
```

`scmType=nexus`의 모든 필요한 값을 위해 이 시크릿을 지정하려면 다음을 설정하세요:

```
brokerTokenSecret:
  name: snyk-broker-secrets

nexusUrlSecret:
  name: snyk-broker-secrets

baseNexusUrlSecret:
  name: snyk-broker-secrets

brokerClientValidationUrlSecret:
  name: snyk-broker-secrets
```