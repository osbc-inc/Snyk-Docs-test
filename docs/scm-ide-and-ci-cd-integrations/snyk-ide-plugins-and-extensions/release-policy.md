# Snyk IDE 플러그인 릴리스 및 지원 정책

이 문서는 Snyk IDE 플러그인의 릴리스 정책에 대해 상세히 설명합니다.

빌드는 포괄적인 테스트 후 구체적인 시점에 배포되어 안정적으로 간주됩니다.

**주기**: 매년 여섯 번의 버전, 대략 **8주**마다 한 번씩. 새로 발견된 보안 취약점에 대한 핫픽스 릴리스가 가능합니다.

## IDE 플러그인 버전 관리

모든 Snyk IDE 플러그인은 산업 표준 [Semantic Versioning](https://semver.org/)을 따릅니다. 이는 세 부분으로 구성된 표기를 가지며, 주어진 버전 번호, `주버전.부버전.패치`를 증가시킵니다:

1. **주버전**은 **파괴적인** 변경 사항이 있을 때 증가합니다; 아래는 예시입니다.
2. **부버전**은 역호환성을 유지하면서 기능을 추가할 때 증가합니다.
3. **패치** 버전은 역호환성 버그 수정 시 증가합니다.

릴리스 노트에서는 모든 변경 사항을 확인하며, 주로 새로운 기능이나 중요한 변경 사항에 중점이 둡니다.

**파괴적인** 변경 사항의 예시는 다음과 같습니다:

* 이전에 결과를 내던 기능을 폐기하면, 사용자는 릴리스 이후 결과를 더 이상 받지 못합니다.
* 새로운 기능을 도입하여 새로운 결과를 발생시킬 수 있습니다.
* 결과의 수에 영향을 줄 수 있는 필수 구성 변경을 도입합니다.

## 지원 정책

{% hint style="info" %}
이 정책은 2025년 6월 24일부터 효력이 발생합니다.
{% endhint %}

Snyk은 플러그인 버전의 최신 12개월을 지원하여 기능과 성능을 보장합니다. 이전 버전은 지원 종료 (End-of-Support, EOS)되어 버그 수정이나 문제 해결을 받지 못합니다.

Snyk은 새 버전에서만 수정을 제공하며, 이전 버전을 수정할 수 없습니다. 고객은 개선 사항을 얻기 위해 업그레이드해야 합니다.

이 정책은 혁신을 육성하면서 리소스를 최적화합니다.

도움이 필요하면 Snyk 지원에 [요청](https://support.snyk.io)을 제출하십시오.

## 미리보기 빌드

Snyk은 진행 중인 기능을 테스트하고자 하는 사용자를 위해 미리보기 버전을 제공합니다.

{% hint style="info" %}
이러한 버전에는 버그가 포함될 수 있으며, 공식적으로 지원되지 않습니다.
{% endhint %}

미리보기 빌드는 일일 최대 여러 번 배포되며 최신 변경 사항을 포함합니다. 그들은 {년}. {월}. {일}{시간} 패턴을 따르는 버전 번호가 부여됩니다.

IDE 플러그인의 미리보기 버전은 `(미리보기) Snyk Security` 형식으로 명명된 독립된 플러그인으로 배포됩니다. Snyk은 안정 버전과 함께 설치하지 않도록 권장합니다.
