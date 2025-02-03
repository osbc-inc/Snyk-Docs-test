# Sysdig와 통합

워크로드 정보를 감지할 때 기능을 향상시키기 위해 Snyk이Sysdig와 파트너십을 맺었습니다. 이 통합은 Snyk이감지하는 워크로드 문제를 Sysdig가 제공하는 런타임 데이터로 보강합니다.

## Sysdig 통합 활성화

Sysdig와 성공적으로 통합하려면, Snyk Controller는 `snyk-monitor` 네임스페이스에 추가 Sysdig Secret가 필요합니다. Sysdig Secret의 이름은 `snyk-sysdig-secret`입니다.

{% hint style="info" %}
Sysdig 설치 후 아래 명령을 실행하여 클러스터 내에서 Snyk Controller가 Sysdig를 감지할 수 있도록 설정하십시오.
{% endhint %}

`snyk-monitor` 네임스페이스에 `snyk-sysdig-secret` 생성하기:

```
kubectl create secret generic snyk-sysdig-secret -n snyk-monitor \
  --from-literal=token=$SYSDIG_RISK_SPOTLIGHT_TOKEN \
  --from-literal=endpoint=$SYSDIG_ENDPOINT_URL \
  --from-literal=cluster=$SYSDIG_AGENT_CLUSTER
```

SYSDIG\_RISK\_SPOTLIGHT\_TOKEN은 "Risk Spotlight Integrations Token"이며 Sysdig UI를 통해 생성해야 합니다. 이 API 토큰을 생성하려면 [Sysdig Risk Spotlight 가이드](https://docs.sysdig.com/en/docs/sysdig-secure/integrations-for-sysdig-secure/risk-spotlight-integrations/#generate-a-token-for-the-integration)를 참조하십시오.

SYSDIG\_ENDPOINT\_URL은 Sysdig SaaS 어플리케이션과 지역과 관련이 있습니다. 확인하려면 [SaaS Regions and IP Ranges](https://docs.sysdig.com/en/docs/administration/saas-regions-and-ip-ranges)를 참조하십시오. 예를 들어, US West(Oregon)의 경우, 도메인은 [us2.app.sysdig.com](https://us2.app.sysdig.com/)입니다("https://" 접두사는 빼야 합니다).

SYSDIG\_AGENT\_CLUSTER는 [Sysdig Agent 설치](https://docs.sysdig.com/en/docs/installation/sysdig-secure/install-agent-components/kubernetes/#parameter-definitions) 시 구성한 것들입니다 - global.clusterConfig.name.

Snyk이 Sysdig와 통합하여 런타임에서 실행된 패키지에 대한 정보를 수집할 수 있도록 하려면, Snyk Controller를 설치할 때 `--set sysdig.enabled=true`를 사용하십시오:

```
helm upgrade --install snyk-monitor snyk-charts/snyk-monitor \
  --namespace snyk-monitor \
  --set clusterName="Production cluster" \
  --set sysdig.enabled=true
```

이제 Snyk Controller는 30분마다 Sysdig로부터 데이터를 수집합니다.

## Snyk 취약점 데이터와 우선순위 점수 보강

Snyk은 실행된 패키지를 활용하여 감지한 취약점의 우선순위 점수를 높이고 있습니다. 이를 통해 Snyk은 어떤 취약점을 먼저 해결해야 하는지를 더 잘 우선순위를 매길 수 있습니다. 이 우선순위 점수는 **프로젝트** 페이지와 [Snyk public API](https://snyk.docs.apiary.io/#reference/projects/aggregated-project-issues/list-all-aggregated-issues)에서 확인할 수 있습니다.

![런타임에서 실행된 패키지들](<../../../.gitbook/assets/image (113) (1) (2) (1) (1) (2) (1) (1) (1).png>)

런타임에서 실행된 패키지를 확인하려면, 다음 일일 스캔을 기다리거나 수동으로 Snyk에 워크로드를 가져와야 합니다.

Sysdig 통합을 활성화한 후에는 수동으로 워크로드를 가져오기 전에 30분간 기다려야 합니다. 완료된 패키지의 수집에 관련한 타이밍 고려 사항으로 인합니다:

* Snyk Controller는 실행된 패키지에 대한 데이터를 30분마다 수집합니다.
* Snyk는 매일 새로운 취약점을 위해 가져온 Kubernetes 프로젝트를 다시 스캔합니다.

## 애플리케이션 지원

애플리케이션 취약점의 경우, Snyk은 현재 다음 언어를 지원합니다:

* Java
* JavaScript
* Go

지원하는 언어 목록을 확인하려면 [컨테이너 이미지에서 애플리케이션 취약점 검색](../use-snyk-container/detect-application-vulnerabilities-in-container-images.md)을 참조하십시오.
