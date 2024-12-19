# 도구: snyk-scm-contributors-count

## 이 도구는 무엇을 하는가?

이 도구는 다음 SCMs 중 하나에 대한 지난 90일 동안의 기여자 수를 계산하고 요약하여 출력합니다:

- Azure Devops
- Bitbucket Cloud
- Bitbucket Server
- GitHub
- GitHub Enterprise
- GitLab
- GitLab Server

{% hint style="info" %}
각 SCM 간에 명명 규칙에 약간의 차이가 있습니다. 예를 들어, GitHub의 "Orgs"는 Azure의 "Projects" 및 Bitbucket의 "Workspaces"로 나타낼 수 있습니다. 이러한 차이는 도구가 각 SCM에 대해 받아 들이는 명령에 반영됩니다.
{% endhint %}

## **SCM-Contributors-Count 도구 작동 방식**

이 도구는 **온보딩 전 사용범위** 및 **Snyk 라이선스 사용량** 두 가지 모드로 작동합니다.

- **온보딩 전 사용범위:** Snyk에 온보딩하려는 사용자들을 위해 SCM을 통해 개발자 수를 추정하고 싶어하는 사용자들을 위한 것입니다.\
  이 모드에서는 사용자가 제공한 자격 증명을 사용하여 도구가 모든 정보를 SCM에서 직접 가져옵니다.
- **Snyk 라이선스 사용량 (Bitbucket과 Azure에만 해당):** 기존 Snyk 계정을 보유한 사용자들이 라이선스 소비에 대한 명확성과 세부 정보를 원하는 경우 (기여자 수, 이름, 이메일 등).\
  이 모드에서는 도구가 Snyk에서 모니터링하는 SCM 관련 프로젝트를 가져와 해당 프로젝트를 SCM의 저장소와 일치시킨 후 해당 저장소/프로젝트에 대한 기여자만을 계산합니다.

{% hint style="info" %}
현재 도구는 "noreply.github.com"으로 끝나는 이메일을 **계산하지 않습니다**.
{% endhint %}

## 도구 다운로드

다음을 실행하십시오:

```
npm i -g snyk-scm-contributors-count
```

또한 [릴리스 페이지](https://github.com/snyk-tech-services/snyk-scm-contributors-count/releases)에서 **바이너리**를 얻을 수 있습니다.