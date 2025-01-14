# SCM 통합을 위한 워크스페이스

{% hint style="info" %}
Workspaces는 특히 대규모 기업 환경에서 Snyk의 SCM 통합 결과의 정확성과 신뢰성을 향상시킵니다. 이 기능은 미래에 계획 중인 추가 기능 및 개선을 지원합니다. 이러한 이유로 Snyk는 귀하의 그룹 수준에서 [이 옵션을 기본적으로 활성화](workspaces-for-scm-integrations.md#manage-workspaces)하는 것을 강력히 권장합니다.
{% endhint %}

## Workspaces의 중요성: 신뢰성과 정확성 향상

Snyk은 Workspaces가 가장 **신뢰성 있고** **정확한** 취약점 탐지를 제공하는 데 상당한 발전 단계라고 믿습니다.

### Workspaces가 신뢰성 있는 결과를 지원하는 방법

기존에 Snyk은 SCM API를 사용하여 저장소 콘텐츠에 액세스했는데, 이는 주요 및 보조 속도 제한 및 콘텐츠 제한을 부과합니다. 예를 들어 GitHub.com API는 시간당 일정 수의 요청만 허용하도록 속도 제한이되며 Git 데이터베이스에서 검색된 트리 항목 수에도 제한이 있습니다.

이러한 API를 통해 저장소 콘텐츠가 검색되면 이러한 제한으로 인해 다양한 방법으로 분석이 완전하지 않는 경우가 많습니다. 특히 아주 많은 수의 저장소 또는 10만 개 이상의 파일을 포함하는 저장소, 일명 모노레포에서는 더욱 그렇습니다.

Workspaces를 통해 이러한 제한이 제거됩니다.

### Workspaces가 정확한 결과를 지원하는 방법

결과의 정확성은 복제를 통해 다양한 방법으로 향상됩니다. 특정 커밋 SHA에 대한 소스 코드 저장소의 완전한 뷰에 접근할 수 있기 때문에, 10만 개 이상의 파일을 포함하는 저장소를 포함한 경우에도 복제를 통해 분석이 더 완전해집니다.

## Snyk 데이터 수집

Workspaces가 활성화되면 Snyk은 주어진 커밋에서 저장소 콘텐츠의 임시 및 얕은 클론 및 모든 커밋 메타데이터(커밋 메시지, 작성자 및 타임스탬프 포함)를 SCM 통합을 통해 흡수합니다.

{% hint style="info" %}
Snyk 데이터 처리 및 Git 리포지토리 클론에 대한 캐시 보유 기간에 대한 자세한 내용은 [취약점 소스 데이터와 관련된 캐시 보유 기간](../../../working-with-snyk/how-snyk-handles-your-data.md#cache-retention-period-related-to-vulnerability-source-data) 및 [데이터가 안전하다는 것을 보장하기 위한 Snyk의 보호장치](../../../working-with-snyk/how-snyk-handles-your-data.md#safeguards-snyk-puts-in-place-to-ensure-data-is-secure)를 참조하십시오.
{% endhint %}

## Workspaces에서 사용되는 Git 프로토콜

저장소는 HTTPS를 사용하여 복제됩니다. SSH 기반 복제는 현재 사용할 수 없습니다.

## Snyk 브로커 상호작용

Git 작업이 브로커를 통해 허용되는 경우 브로커드 연결이 지원됩니다.

{% hint style="warning" %}
이는 `accept.json`에서의 제한을 재정의합니다. 자세한 내용은 [도커를 위한 브로커를 통한 복제 기능](../../../enterprise-setup/snyk-broker/git-clone-through-broker.md)을 참조하십시오.
{% endhint %}

## Workspaces 관리

Workspaces 기능은 그룹 또는 조직 수준에서 **통합 설정** 페이지를 통해 관리됩니다. 이를 위해서는 Snyk 그룹 관리자 또는 Snyk 조직 관리자여야 합니다.

그룹 수준에서 활성화된 경우 해당 그룹 내 모든 조직은 Workspaces가 활성화됩니다.

Workspaces 기능은 그룹 수준에서 비활성화되어 있더라도 해당 그룹 내 개별 조직을 위해 활성화할 수 있습니다.

Workspaces 기능을 비활성화할 수는 있지만, 이렇게 하면 다음 두 가지 특정 방식으로 저장소를 스캔하는 Snyk의 능력에 제약이 생길 수 있다는 점을 이해하는 것이 중요합니다:

1. **스캔 빈도**: Workspaces가 비활성화된 경우 Snyk는 SCM API를 통해 저장소 콘텐츠를 분석하며, 일반적으로 주요 및 보조 속도 제한 및 콘텐츠 제한을 부과합니다. 예를 들어, GitHub.com API는 한 시간에 특정 수의 요청만 허용하므로, 1시간에 많은 수의 저장소를 분석하는 능력이 제한됩니다.
2. **스캔 완전성**: Workspaces 기능이 비활성화된 경우 Snyk는 SCM API를 통해 저장소 콘텐츠를 분석하는데, 일반적으로 Git 데이터베이스에서 검색될 수 있는 트리 항목 수가 제한되어 놓아 취약점이 누락될 수 있습니다. 이는 파일이 많은 저장소 또는 모노레포로 불리는 저장소의 분석에 영향을 줍니다.
