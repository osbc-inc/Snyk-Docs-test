# Broker를 통한 Git 클론

{% hint style="warning" %}
Broker를 통해 Git Clone을 사용하도록 설정하면 조직과 관련된 모든 코드 에이전트 배포가 중단됩니다. 가져오기 및 반복 테스트가 중단되지만 PR 체크는 Code Agent에서 지원되지 않았기 때문에 계속됩니다. 원활한 전환을 위해 **이주 단계를 따라야 하며** 중개인을 통한 Git Clone(중개된 코드)로 마이그레이션하는 방법에 대한 통신이 고객에게 발송되었습니다. 이주 지침서를 받지 못했거나 도움이 필요한 경우 [Snyk 지원팀](https://support.snyk.io)에 문의하십시오.

**이주 단계를 완료한 후에**, 코드로 이동하여 다음과 같이 Code에 중개인을 통한 Git Clone을 활성화할 수 있습니다: \*\*설정 → 조직 설정 → Snyk 중개인 → 이미지 탐색기 사용(Git 중개인 활성화)\*\*를 토글합니다.이 기능을 활성화한 후 토글은 체크 표시가 있는 파란색으로 전환됩니다.
{% endhint %}

중개된 Snyk Code는 중개인이 코드 파일을 수락하도록 하는 것입니다. 중개인이 SCM 시스템과 Snyk 사이를 스캔합니다.

도커로 설치하는 경우 중개인이 리포지토리의 Git 클론을 수행하도록 중개인에 액세스 권한을 부여하려면 환경 변수를 추가하십시오: `ACCEPT_CODE=true`. 브로커 헬름 차트는 기본적으로 `ACCEPT_CODE=true`를 설정합니다.

도커로 설치하는 예시:

```
docker run --restart=always \
           -p 8000:8000 \
           -e BROKER_TOKEN=secret-broker-token \
           -e GITHUB_TOKEN=secret-github-token \
           -e PORT=8000 \
           -e BROKER_CLIENT_URL=http://my.broker.client:8000 \
           -e ACCEPT_CODE=true
       snyk/broker:github-com
```

이는 Git 서버를 위한 필요한 `accept` 규칙을 추가합니다.

이 작업을 완료한 후 SCM 시스템에 대한 중개인 지침을 따를 수 있습니다. 자세한 내용은 [Snyk Broker 설치 및 구성](install-and-configure-snyk-broker/)를 참조하십시오.
