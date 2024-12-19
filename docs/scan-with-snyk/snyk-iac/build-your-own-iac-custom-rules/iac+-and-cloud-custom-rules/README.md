# IaC+ 및 클라우드 사용자 정의 규칙

{% hint style="info" %}
**기능 가용성**\
IaC 사용자 정의 규칙은 엔터프라이즈 요금제에서만 사용 가능합니다. 더 자세한 정보는 [요금제 및 가격정보](https://snyk.io/plans/)를 참조하십시오.
{% endhint %}

## IaC에서 클라우드 사용자 정의 역할을 위한 전제 조건

다음을 PATH에 설치하십시오: Snyk CLI >= 1.1168.0\
자세한 내용은 [Snyk CLI 설치 또는 업데이트](../../../../snyk-cli/install-or-update-the-snyk-cli/)를 참조하십시오.

CLI를 적절하게 구성하고 기본 조직을 설정하십시오:\
`snyk config set org=<org id>`

다수의 Snyk 조직을 사용하는 경우, 원하는 조직을 지정하기 위해 명령에 `--org=<your org id>`를 추가할 수 있습니다.

## Snyk IaC에서 미리 정의된 보안 규칙

Snyk IaC에는 AWS, Azure, GCP 및 Kubernetes를 포함하는 보안 규칙 세트가 기본 제공되어 작동합니다. 이러한 규칙은 보안 연구, 모범 사례, 인정된 표준 및 기준에 기반하며 새로운 규칙이 정기적으로 출시됩니다. Snyk의 보안 엔지니어링 팀이 이들을 적극적으로 유지 보수합니다.

이러한 규칙은 대부분의 경우에 첫 번째 스캔에서 필요한 것을 충족하기 위해 의도되었지만, 시스템에 추가적인 보안 규칙을 강제 적용해야 할 수도 있습니다. 예를 들어 태깅 표준과 같은 것들입니다.

## 클라우드 사용자 정의 규칙의 목적

Snyk의 미리 정의된 규칙을 보완하기 위해 IaC에서 클라우드 사용자 정의 규칙을 사용하면 SDLC(소프트웨어 개발 라이프사이클) 전반에서 내부 보안 컨트롤을 강제할 수 있습니다. 클라우드 사용자 정의 규칙을 사용하여 다음을 식별하고 강조할 수 있습니다:

* 소스 코드 관리, CLI, Terraform Cloud 및 배포된 클라우드 환경을 포함한 SDLC 전반의 클라우드 구성 문제
* 클라우드 이외의 Terraform 공급자를 사용하는 Terraform IaC 구성에서의 문제, 예를 들어 GitHub 또는 Snowflake 구성과 같은 것들.

다음은 클라우드 사용자 정의 규칙 사용 시의 단계입니다:

1. [사용자 정의 규칙 프로젝트 생성](create-a-custom-rules-project.md)
2. [사용자 정의 규칙 생성](create-a-custom-rule.md)
3. [(권장) 사용자 정의 규칙 사양(Unit Test) 작성](recommended-write-a-custom-rule-spec-unit-test.md)
4. [(권장) 사용자 정의 규칙 테스트](recommended-test-your-custom-rule.md)
5. [사용자 정의 규칙 빌드, 번들 및 업로드](build-bundle-and-upload-your-custom-rules.md)
6. [사용자 정의 규칙 강제 적용](enforce-your-custom-rules.md)
7. [예시 사용자 정의 규칙 탐색](explore-example-custom-rules.md)