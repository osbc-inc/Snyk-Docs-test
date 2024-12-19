# SCM, IDE 및 CI/CD 통합

이 문서의 이 섹션은 [Snyk SCM](snyk-scm-integrations/), [IDE](snyk-ide-plugins-and-extensions/) 및 [CI/CD 통합](snyk-ci-cd-integrations/)에 대한 정보를 제공합니다.

Snyk은 SCM, IDE 및 CI/CD 통합 방법을 지원하여 워크플로우 각 단계에서 보안을 구현할 수 있도록 합니다: 프로젝트 가져오기, 코드 작성, 빌드 및 배포.

{% hint style="info" %}
**기능 가용성**

API 및 Snyk AppRisk는 엔터프라이즈 플랜에서만 사용할 수 있습니다. 자세한 정보는 [플랜 및 가격](https://snyk.io/plans/)을 참조하십시오.
{% endhint %}

Snyk 환경에서 SCM 통합을 구현하는 두 가지 방법이 있습니다:

- **그룹 레벨** - 그룹 레벨에서 Snyk AppRisk를 위한 SCM 통합을 설정할 수 있습니다.
- **조직 레벨** - 조직 레벨에서 다른 모든 Snyk 제품 및 모든 Snyk 플랜을 위한 SCM 통합을 설정할 수 있습니다. 자세한 내용은 [통합 관리](../getting-started/snyk-web-ui.md)를 참조하십시오.

{% hint style="info" %}
그룹 및 조직 레벨에서 동일한 SCM 통합을 사용하려면 해당 통합을 두 레벨 모두에 설정해야 합니다.
{% endhint %}

## 통합 선택

엔터프라이즈 고객이라면 SDLC에 맞는 통합이 어떤 것인지 결정하기 위해 엔터프라이즈 구현 안내서의 [통합 롤아웃 선택](../implement-snyk/team-implementation-guide/phase-1-discovery-and-planning/choose-rollout-integrations.md)를 참조하십시오. 수입 전략 및 통합 선택에 대한 팁 및 고려 사항을 확인할 수 있습니다.

### GitHub vs GitHub Enterprise

**엔터프라이즈 플랜 사용자**로서 Snyk은 GitHub Enterprise 통합을 사용하는 것을 권장합니다. 개별 사용자 계정의 GitHub 서비스 계정 접근 토큰 (PAT)을 사용하는 대신 Snyk 조직 전체에서 단일 GitHub 서비스 계정 PAT를 사용할 수 있게 합니다. 이 통합은 GitHub Enterprise (GHE) 라이선스나 구독 여부에 관계없이 사용할 수 있습니다.

GitHub Enterprise 통합을 사용하면 새로운 Snyk 조직을 만드는 경우 통합 설정을 복제할 수 있습니다. 이로써 Snyk 그룹의 모든 조직에서 하나의 GitHub Enterprise 통합을 사용할 수 있습니다.

**무료 또는 팀 플랜 사용자**로서 Snyk은 GitHub 통합을 사용하는 것을 권장합니다. 이 통합은 개별 사용자 계정의 PAT만 필요하므로 이 수준에서 필요한 기능을 충족할 수 있습니다.

자체 호스팅된 GitHub Enterprise 제품을 사용하는 경우 GitHub Enterprise 통합을 사용해야 합니다.

{% hint style="info" %}
GitHub에서 GitHub Enterprise로 마이그레이션하는 자세한 단계는 [GitHub Enterprise로 마이그레이션하기](snyk-scm-integrations/github.md#migrate-to-the-github-enterprise-integration)를 참조하십시오.
{% endhint %}

### Bitbucket Cloud (PAT) vs Bitbucket Cloud 앱

일반적으로 Snyk은 새로운 Bitbucket Cloud 앱 통합을 사용하는 것을 권장합니다. 그러나 새로운 통합이 모든 경우에 적합하지는 않습니다. 이 섹션의 정보는 당신에게 적합한 통합을 결정하도록 돕기 위한 것입니다.

자세한 마이그레이션 지침은 [Snyk Bitbucket Cloud 앱으로 마이그레이션](snyk-scm-integrations/bitbucket-cloud.md#migrate-to-the-snyk-bitbucket-cloud-app)에서 확인할 수 있습니다.

#### 새로운 앱 통합에서 열리는 주요 기능

- Bitbucket의 [IP 주소 화이트리스트](https://support.atlassian.com/bitbucket-cloud/docs/control-access-to-your-private-content/) 프리미엄 티어 기능을 사용할 수 있음.
- Bitbucket Cloud의 여러 워크스페이스에 리포지토리를 분산하는 회사들에 대한 속도 제한 문제 처리를 돕습니다.
- Bitbucket Cloud의 첫 번째 파티 인터페이스 지원( Snyk의 보안 탭)에서 처음부터 지원되므로 Bitbucket Cloud에서 Snyk 보안 통찰력을 설치 및 유지 관리할 필요가 없습니다.

#### 새로운 앱 통합의 제한 사항

- 새로운 앱 통합에서 각 Snyk 조직은 Bitbucket Cloud의 하나의 워크스페이스에만 연결할 수 있습니다. Bitbucket의 여러 워크스페이스에서 동일한 조직의 단일 조직으로 프로젝트를 가져오려면 PAT 통합을 사용해야 합니다.
- 새로운 앱 통합은 아직 Snyk Multi-Tenant EU, Snyk Multi-Tenant AUS 및 Snyk Single-Tenant 클라우드 배포를 지원하지 않습니다.
- 새 앱 통합을 사용하려면 기본 브랜치 이외의 브랜치에서 프로젝트를 가져올 수 없습니다. 비기본 브랜치를 가져오려면 Snyk.io 가져오기 모달에서 수행해야 합니다.

#### Personal Access Token (PAT) 통합의 생애주기 종료 계획이 있나요?

아니요, Personal Access Token Bitbucket Cloud 통합은 완전히 지원되며 지원 중단 계획은 없습니다.

그러나 개발자가 Bitbucket 인터페이스에서 Snyk의 결과를 소비할 수 있도록 하는 PAT 통합의 확장 레이어로 작동하는 첫 번째 파티 인터페이스 _확장_ 앱이 있습니다. 이 확장 앱은 외부 계약업체가 개발 및 지원한 것입니다. 이 기능은 지금 새로운 앱 통합의 일부가 되었기 때문에 확장 앱은 더 이상 지원되지 않는 모드로 전환되었으며, PAT 통합과 함께 첫 번째 파티 인터페이스 기능을 사용하는 고객은 새로운 앱 통합으로 마이그레이션하여 처음부터 지원을 받아야 합니다.

## Snyk 통합을 위한 풀 리퀘스트

Snyk은 여러 Snyk 통합과 호환되는 종속성을 업그레이드하기 위해 스캔 결과를 기반으로 귀하를 대신하여 자동으로 풀 리퀘스트(PR)를 생성할 수 있습니다. 자세한 내용은 [자동 PR을 통한 오픈 소스 종속성 업그레이드](../scan-with-snyk/pull-requests/snyk-pull-or-merge-requests/upgrade-dependencies-with-automatic-prs-upgrade-prs/upgrade-open-source-dependencies-with-automatic-prs.md)을 참조하십시오.