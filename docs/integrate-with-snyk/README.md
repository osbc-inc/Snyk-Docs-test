# Snyk과 통합하기

## Snyk의 통합

Snyk 내에서 타사 기능을 사용하고 다른 도구와의 사용을 도울 많은 통합이 제공됩니다. 통합 및 해당 워크플로우를 완료하는 데 사용할 수 있는 기타 방법에 대한 정보는 [SCM, IDE 및 CI/CD 워크플로우 및 통합](../scm-ide-and-ci-cd-integrations/)에서 확인하세요.

이 페이지는 추가적인 Snyk 통합 및 해당 위치를 식별합니다.

Snyk는 리포지토리 게이트키퍼용 플러그인 및 패키지 리포지토리에 연결하기 위한 통합을 제공합니다:

* [게이트키퍼 플러그인](../scan-with-snyk/snyk-open-source/manage-vulnerabilities/gatekeeper-plugins/)
* [패키지 리포지토리 통합](../scan-with-snyk/snyk-open-source/package-repository-integrations/)

Snyk Container 및 Snyk Iac를 지원하는 통합이 있습니다:

* [Snyk Container 통합](../scan-with-snyk/snyk-container/container-registry-integrations/)
* [클라우드 플랫폼 통합](cloud-platforms-integrations/)

[이벤트 전달 통합](event-forwarding/)을 통해 Snyk 플랫폼 이벤트를 직접 다른 플랫폼의 특정 제품에 전달하여 사용자 정의 경보 설정, 고유 보고서 작성, 자동화 트리거 등을 설정할 수 있습니다.

[알림 및 티켓 시스템 통합](jira-and-slack-integrations/)을 통해 Jira 및 Slack에서 Snyk를 사용할 수 있습니다.

Snyk가 취약점 관리 도구와 어떻게 작동하는지에 대한 정보도 제공합니다. 

Snyk는 대체 보고서 도구를 제공합니다. 자세한 내용은 [보고서 및 BI 통합](../manage-risk/reporting/reporting-and-bi-integrations-snowflake-data-share/)을 참조하세요.

## Snyk AppRisk를 위한 통합

통합 페이지에는 모든 활성 통합이 표시되며 기존 Snyk 조직에서 자동 동기화된 모든 데이터가 포함되어 있으며 통합 허브에 액세스할 수 있습니다.

다음은 자동으로 동기화된 Snyk 데이터 예시입니다: Snyk 오픈 소스, Snyk 코드, Snyk IaC, Snyk 컨테이너.

{% hint style="info" %}
[Snyk AppRisk를 위한 타사 통합](../manage-risk/snyk-apprisk/integrations-for-snyk-apprisk/connect-a-third-party-integration.md)로 이동하여 사용 가능한 통합 및 설정 세부 정보를 확인하세요.
{% endhint %}

각 연결된 통합을 통해 데이터 동기화를 일시 중지하거나 통합 프로필 및 구성을 수정하거나 통합을 삭제하거나 통합이 마지막으로 동기화된 시간 및 다음 동기화가 예정된 시간을 확인할 수 있습니다.

### Snyk AppRisk 통합 생태계

아래 표를 참조하여 Snyk AppRisk에 대한 모든 통합의 가용성과 호환성을 확인할 수 있습니다. 통합은 유형으로 분류되며 이름으로 나열되며, Snyk AppRisk Essentials 및 Snyk AppRisk Pro에 대해 사용 가능 여부가 표시됩니다.

<table><thead><tr><th width="172">통합 유형</th><th width="164">통합 이름</th><th width="198">Snyk AppRisk Essentials</th><th>Snyk AppRisk Pro</th></tr></thead><tbody><tr><td>SCM</td><td><ul><li><a href="../scm-ide-and-ci-cd-integrations/snyk-scm-integrations/github.md#group-level-snyk-apprisk-integrations">GitHub</a></li><li><a href="../scm-ide-and-ci-cd-integrations/snyk-scm-integrations/bitbucket-cloud.md#group-level-snyk-apprisk-integrations">BitBucket</a></li><li><a href="../scm-ide-and-ci-cd-integrations/snyk-scm-integrations/gitlab.md#group-level-snyk-apprisk-integrations">GitLab</a></li><li><a href="../scm-ide-and-ci-cd-integrations/snyk-scm-integrations/azure-repositories-tfs.md#group-level-snyk-apprisk-integrations">Azure DevOps</a></li></ul></td><td>                <span data-gb-custom-inline data-tag="emoji" data-code="2714">✔️</span></td><td>                   <span data-gb-custom-inline data-tag="emoji" data-code="2714">✔️</span></td></tr><tr><td>개발 포털 및 서비스 카달로그</td><td><ul><li><a href="../scm-ide-and-ci-cd-integrations/snyk-scm-integrations/application-context-for-scm-integrations/">Backstage 카달로그</a></li><li><a href="../manage-risk/snyk-apprisk/integrations-for-snyk-apprisk/connect-a-third-party-integration.md#servicenow-cmdb-setup-guide">ServiceNow CMDB</a></li><li><a href="../scm-ide-and-ci-cd-integrations/snyk-scm-integrations/application-context-for-scm-integrations/#atlassian-compass">Atlassian Compass</a></li><li><a href="../scm-ide-and-ci-cd-integrations/snyk-scm-integrations/application-context-for-scm-integrations/#harness">Harness</a></li><li><a href="../scm-ide-and-ci-cd-integrations/snyk-scm-integrations/application-context-for-scm-integrations/#opslevel">OpsLevel</a></li><li><a href="../scm-ide-and-ci-cd-integrations/snyk-scm-integrations/application-context-for-scm-integrations/#datadog-org-context-service-catalog">Datadog Org Context (Service Catalog)</a></li></ul></td><td>               <span data-gb-custom-inline data-tag="emoji" data-code="2714">✔️</span></td><td>                     <span data-gb-custom-inline data-tag="emoji" data-code="2714">✔️</span></td></tr><tr><td>리스크 관리 협업</td><td><ul><li><a href="jira-and-slack-integrations/slack-integration.md">Slack</a></li><li>이메일</li></ul></td><td>                <span data-gb-custom-inline data-tag="emoji" data-code="2714">✔️</span></td><td>                    <span data-gb-custom-inline data-tag="emoji" data-code="2714">✔️</span></td></tr><tr><td>AST</td><td><ul><li><a href="../manage-risk/snyk-apprisk/integrations-for-snyk-apprisk/connect-a-third-party-integration.md#nightfall-setup-guide">NightFall</a></li><li><a href="../manage-risk/snyk-apprisk/integrations-for-snyk-apprisk/connect-a-third-party-integration.md#gitguardian-setup-guide">GitGuardian</a></li></ul></td><td>               <span data-gb-custom-inline data-tag="emoji" data-code="2716">✖️</span></td><td>                     <span data-gb-custom-inline data-tag="emoji" data-code="2714">✔️</span></td></tr><tr><td>런타임 보안 및 가시성</td><td><ul><li><a href="../manage-risk/snyk-apprisk/integrations-for-snyk-apprisk/snyk-runtime-sensor.md">Snyk 런타임 센서</a></li><li><a href="../manage-risk/snyk-apprisk/integrations-for-snyk-apprisk/connect-a-third-party-integration.md#sysdig-setup-guide">Sysdig</a></li><li><a href="../manage-risk/snyk-apprisk/integrations-for-snyk-apprisk/connect-a-third-party-integration.md#dynatrace-setup-guide">Dynatrace</a></li></ul></td><td>               <span data-gb-custom-inline data-tag="emoji" data-code="2716">✖️</span></td><td>                     <span data-gb-custom-inline data-tag="emoji" data-code="2714">✔️</span></td></tr></tbody></table>

### Integration Hub 사용

[Integration Hub](../getting-started/snyk-web-ui.md#manage-integrations-for-asset-discovery-asset-coverage-and-issues-from-third-party-vendors) 페이지를 사용하여 통합을 도입하고 SCM 도구에서 Snyk AppRisk에 데이터를 제공하세요.

다음 단계를 따라 통합을 추가할 수 있습니다:

1. AppRisk를 열고 통합 페이지로 이동합니다.
2. **Integration Hub**을 클릭합니다.
3. 연결하려는 통합을 선택하고 **Add**를 클릭합니다.
4. 연결을 구성하고 **Done**을 클릭합니다.

통합을 설정하는 방법에 대한 단계별 자세한 내용은 [Snyk SCM 통합](../scm-ide-and-ci-cd-integrations/snyk-scm-integrations/#group-level-snyk-apprisk-scm-integrations) 페이지 또는 [Snyk AppRisk를 위한 타사 통합](../manage-risk/snyk-apprisk/integrations-for-snyk-apprisk/connect-a-third-party-integration.md) 페이지를 참조하세요.

통합이 확인되면, 통합 페이지에 카드가 표시되어 연결을 활성화 또는 비활성화하거나 설정을 편집하거나 구성에서 연결을 제거할 수 있습니다.

<figure><img src="../.gitbook/assets/image (11) (4).png" alt="AppRisk - Integration status" width="375"><figcaption><p>Snyk AppRisk - 통합 상태</p></figcaption></figure>

### Snyk 브로커 사용

SCM 인스턴스가 공개적으로 액세스할 수 없는 경우 Snyk 브로커가 필요합니다. Docker 또는 Helm을 사용하여 Snyk 브로커를 설치 및 구성할 수 있습니다. Snyk 브로커에 대한 자세한 내용은 [Snyk 브로커 - AppRisk](../enterprise-setup/snyk-broker/snyk-broker-apprisk.md)를 참조하세요.

{% hint 스타일 = "info" %}
명령을 실행하기 전에 Snyk 브로커 배포 환경에서 Snyk AppRisk 플래그를 활성화하세요.
{% endhint %}

* GitHub - Snyk 브로커 설치 및 구성
  * [Docker 사용하기](../enterprise-setup/snyk-broker/install-and-configure-snyk-broker/github-prerequisites-and-steps-to-install-and-configure-broker/github-install-and-configure-using-docker.md#docker-run-command-to-set-up-a-broker-client-for-github)
  * [Helm 사용하기](../enterprise-setup/snyk-broker/install-and-configure-snyk-broker/github-prerequisites-and-steps-to-install-and-configure-broker/github-install-and-configure-using-helm.md)
  * [환경 변수](../enterprise-setup/snyk-broker/install-and-configure-snyk-broker/github-prerequisites-and-steps-to-install-and-configure-broker/github-environment-variables-for-snyk-broker.md)
* GitHub Enterprise - Snyk 브로커 설치 및 구성:
  * [Docker 사용하기](../enterprise-setup/snyk-broker/install-and-configure-snyk-broker/github-enterprise-prerequisites-and-steps-to-install-and-configure-broker/github-enterprise-install-and-configure-using-docker.md#docker-run-command-to-set-up-a-broker-client-for-github-enterprise)
  * [Helm 사용하기](../enterprise-setup/snyk-broker/install-and-configure-snyk-broker/github-enterprise-prerequisites-and-steps-to-install-and-configure-broker/github-enterprise-install-and-configure-using-helm.md)
  * [환경 변수](../enterprise-setup/snyk-broker/install-and-configure-snyk-broker/github-enterprise-prerequisites-and-steps-to-install-and-configure-broker/github-enterprise-environment-variables-for-snyk-broker.md)
* BitBucket - Snyk 브로커 설치 및 구성:
  * [Docker 사용하기](../enterprise-setup/snyk-broker/install-and-configure-snyk-broker/bitbucket-server-data-center-prerequisites-and-steps-to-install-and-configure-broker/data-center.md#docker-run-command-to-set-up-a-broker-client-for-bitbucket)
  * [Helm 사용하기](../enterprise-setup/snyk-broker/install-and-configure-snyk-broker/bitbucket-server-data-center-prerequisites-and-steps-to-install-and-configure-broker/bitbucket-server-data-center-install-and-configure-using-helm.md)
  * [환경 변수](../enterprise-setup/snyk-broker/install-and-configure-snyk-broker/bitbucket-server-data-center-prerequisites-and-steps-to-install-and-configure-broker/bitbucket-server-data-center-environment-variables-for-snyk-broker-basic-auth.md)
* GitLab - Snyk 브로커 설치 및 구성:
  * [Docker 사용하기](../enterprise-setup/snyk-broker/install-and-configure-snyk-broker/gitlab-prerequisites-and-steps-to-install-and-configure-broker/setup-broker-with-gitlab.md#docker-run-command-to-set-up-a-broker-client-for-gitlab)
  * [Helm 사용하기](../enterprise-setup/snyk-broker/install-and-configure-snyk-broker/gitlab-prerequisites-and-steps-to-install-and-configure-broker/gitlab-install-and-configure-using-helm.md)
  * [환경 변수](../enterprise-setup/snyk-broker/install-and-configure-snyk-broker/gitlab-prerequisites-and-steps-to-install-and-configure-broker/gitlab-environment-variables-for-snyk-broker.md)
* Azure - Snyk 브로커 설치 및 구성:
  * [Docker 사용하기](../enterprise-setup/snyk-broker/install-and-configure-snyk-broker/azure-repos-prerequisites-and-steps-to-install-and-configure-broker/setup-broker-with-azure-repos.md#docker-run-command-to-set-up-a-broker-client-for-azure-repos)
  * [Helm 사용하기](../enterprise-setup/snyk-broker/install-and-configure-snyk-broker/azure-repos-prerequisites-and-steps-to-install-and-configure-broker/azure-repos-install-and-configure-and-configure-using-helm.md)
  * [환경 변수](../enterprise-setup/snyk-broker/install-and-configure-snyk-broker/azure-repos-prerequisites-and-steps-to-install-and-configure-broker/azure-repos-environment-variables-for-snyk-broker.md)

통합에 액세스할 수 있는 허용된 end-point 목록이 포함된 최신 `.json` 파일은 [GitHub](https://github.com/snyk/broker/tree/565242baf003f06f445489dd96cc68c8386ede38/defaultFilters/apprisk)에서 확인할 수 있습니다.