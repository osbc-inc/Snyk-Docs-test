# SDLC의 정책 사용

개발자의 로컬 개발 환경부터 IDE 또는 CLI, Git 기반의 워크플로 및 CI/CD 및 프로덕션 단계까지 SDLC의 모든 단계에 정책을 적용할 수 있습니다.

이러한 다수의 보안 및 규정 준수 제어는 개발 프로세스에서 문제가 가능한 일찍 식별되어 가장 적게 비용이 들고 시간이 소요되는 시기에 수정될 수 있도록 보장합니다.

{% hint style="info" %}
또한 `.snyk` 파일은 Snyk가 CLI 및 CI/CD 플러그인을 위한 패치를 지정하고 특정 분석 동작을 정의하는 데 사용하는 정책 파일입니다. 자세한 내용은 [The .snyk file](the-.snyk-file.md)을 참조하십시오.
{% endhint %}

## 프로젝트 또는 조직에 정책 할당하기

[보안 정책](security-policies/) 및 [라이선스 정책](license-policies/) 모두 프로젝트 속성 및 조직에 정책을 적용할 수 있습니다. 이를 통해 프로젝트에 정책을 할당하고 조직에 정책을 할당할 수 있습니다. 자세한 내용은 [프로젝트에 정책 할당](assign-policies-to-projects.md) 및 [조직에 정책 할당](assign-a-policy-to-an-organization.md)을 참조하십시오.

### 예: 프로젝트에 라이선스 정책 할당하기

시나리오:\
회사의 법무팀에서 비즈니스 크리티컬한 프론트엔드 애플리케이션에 엄격한 라이선스 준수 제어가 필요하지만 내부 개발 프로젝트에는 큰 관심이 없습니다.

이 요구 사항을 충족하기 위해 먼저 이 정책을 적용하려는 Snyk 프로젝트에 `Critical`, `Production`, 및 `Frontend` 속성을 추가하십시오:

<figure><img src="../../.gitbook/assets/image (1) (3).png" alt="문제 탭에서 프로젝트에 관련 속성 추가"><figcaption><p>문제 탭에서 프로젝트에 관련 속성 추가</p></figcaption></figure>

다음으로 새로운 라이선스 정책을 만들고 해당 속성에 정책을 적용하십시오:

<figure><img src="../../.gitbook/assets/image (7) (1) (1).png" alt="선택된 속성에 라이선스 정책 적용"><figcaption><p>선택된 속성에 라이선스 정책 적용</p></figcaption></figure>

{% hint style="info" %}
정책 자체에서 프로젝트에 식별된 Copyleft 라이선스인 [**GPL-3.0**](https://snyk.io/learn/what-is-gpl-license-gplv3-explained/), [**AGPL-3.0 라이선스**](https://snyk.io/learn/agpl-license/) 등에 높은 심각도가 적용될 수 있습니다.\
라이선스 정책을 만들 때 Snyk는 테스트가 실패하는 이유를 설명하도록 권장합니다. 따라서 예를 들어 GPL 라이선스로 빌드가 실패하면 개발자들이 설명을 볼 수 있고 취해야 할 조치를 알 수 있습니다. 자세한 내용은 [라이선스 정책 및 규칙 만들기](license-policies/create-a-license-policy-and-rules.md)를 참조하십시오.
{% endhint %}

이 정책은 선택된 속성이 적용된 모든 프로젝트에 할당되었으며 다음 번 Snyk이 해당 프로젝트를 스캔할 때 적용됩니다.

자세한 내용은 [라이선스 정책](license-policies/)을 참조하십시오.

### 예: **프로젝트에 보안 정책 할당하기**

이전 예제와 유사한 프로세스를 사용하여 `FrontEnd` 환경의 알려진 취약점이 없는 `Medium` 심각도 취약점을 자동으로 무시하는 보안 정책을 정의할 수 있습니다:

<div align="left"><figure><img src="../../.gitbook/assets/image (14) (3).png" alt="Snyk 보안 정책 - 무시할 취약점 지정"><figcaption><p>Snyk 보안 정책 - 무시할 취약점 지정</p></figcaption></figure></div>

이 정책은 선택된 속성이 적용된 모든 프로젝트에 할당되었으며 다음 번 Snyk이 해당 프로젝트를 스캔할 때 적용됩니다.

자세한 내용은 [보안 정책](security-policies/)을 참조하십시오.

## GitHub 레포지토리에서 정책 적용하기

Snyk이 모니터링하는 GitHub 프로젝트에서 기여하는 개발자가 제출하는 모든 새로운 풀 리퀘스트는 해당 프로젝트에 할당된 정책을 확인합니다. 이를 통해 정책 위반 코드를 저장소에 커밋할 수 없도록 보장합니다.

{% hint style="info" %}
Snyk의 PR Checks 기능에 대한 자세한 내용은 [PR Checks](../../scan-with-snyk/pull-requests/pull-request-checks/)를 참조하십시오.
{% endhint %}

다음은 JavaScript 패키지 라이선스에 대한 PR 체크의 예시입니다.

이 예시는 JavaScript 애플리케이션에 `fullpage.js` 패키지를 추가하기 위한 풀 리퀘스트를 보여줍니다. 이 변경은 최신 버전의 패키지에 알려진 취약점이 없으므로 보안 정책을 통과하지만, 회사의 라이선스 정책 위반으로 인해 GPLv3 라이선스가 포함되어 있어 라이선스 정책 체크에 실패합니다.

<figure><img src="../../.gitbook/assets/image (5) (1) (1) (2).png" alt="라이선스 준수 실패에 대한 PR 체크"><figcaption><p>라이선스 준수 실패에 대한 PR 체크</p></figcaption></figure>

## CI/CD에서 정책 적용하기

할당된 정책은 CI/CD에서 효과를 발휘하여 빌드가 보안 및 규정 준수 기준을 준수하도록합니다.
