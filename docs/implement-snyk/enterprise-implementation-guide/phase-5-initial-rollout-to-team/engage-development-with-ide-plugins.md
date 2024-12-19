# IDE 플러그인으로 개발 활동에 참여하기

## 개발자들에게 Snyk IDE 플러그인 소개하기

[Snyk IDE 플러그인](../../../scm-ide-and-ci-cd-integrations/snyk-ide-plugins-and-extensions/)을 설치하도록 개발자들을 격려하는 것은 개발자의 채택을 향상시키고 "좌측으로 이동(shift left)"을 달성하는 중요한 단계입니다. 여기서 좌측으로 이동이란, 개발자들이 풀 리퀘스트에 통합되거나 빌드에 추가되기 전에 문제를 수정하는 데 더 책임있고 능동적인 역할을 하는 것입니다.

Snyk IDE 플러그인은 개발자들이 보안 취약점과 문제를 찾아 해결하는 데 도움을 줍니다. 이는 개별 개발자들이 보안 검토를 통과하고 개발 중에 비용이 많이 드는 수정을 피하며, 보안 코드를 개발하고 전달하는 시간을 줄일 수 있도록 합니다.

웨비나 [IDE 및 CLI에서 Snyk 사용하는 개발자를 위한 워크플로우](https://www.youtube.com/watch?v=jzUJS6S6H48)는 IDE 사용에 대해 상세히 다루고, IDE 플러그인을 사용하여 문제를 검토하는 데 대한 데모가 포함되어 있습니다.

## 구현 팁

보안에 대해 시작한지 얼마 되지 않은 기업은 수정해야 할 문제가 무엇인지 명확히 지시하는 가시성 및 예방 단계로 시작해야 합니다. 개발자들이 우선순위가 매겨진 문제를 해결하거나 새 작업을 시작할 때 IDE 플러그인을 도입하는 자연스러운 기회가 발생하며, 이를 통해 보안 게이트로 인해 발생할 수 있는 잠재적인 마찰을 제거하기 위해 선행적으로 테스트할 수 있습니다. 이것은 애플리케이션 보안 프로그램이 종종 반응적인 방법에서 선행적인 방법으로 전환하는 지점입니다.

성숙한 애플리케이션 보안 프로그램을 갖춘 기업은 이미 보안 개념에 익숙하기 때문에 개발자들에게 Snyk IDE 플러그인에 더 일찍 접근할 수 있습니다. 따라서 롤아웃은 뒷받침문제를 해결하는 것이 아니라 새 문제를 예방하고 수정 확인하는 것에 더 중점을 두게 됩니다.

## 사용 가능한 플러그인

다음과 같은 Snyk 플러그인 및 확장이 있습니다:

- [Eclipse 플러그인](../../../scm-ide-and-ci-cd-integrations/snyk-ide-plugins-and-extensions/eclipse-plugin/)
- [JetBrains 플러그인](../../../scm-ide-and-ci-cd-integrations/snyk-ide-plugins-and-extensions/jetbrains-plugins/)
- [Visual Studio 확장 프로그램](../../../scm-ide-and-ci-cd-integrations/snyk-ide-plugins-and-extensions/visual-studio-extension/)
- [Visual Studio Code 확장 프로그램](../../../scm-ide-and-ci-cd-integrations/snyk-ide-plugins-and-extensions/visual-studio-code-extension/)

학습 세션 [IDE에서 Snyk 사용 소개](https://learn.snyk.io/lesson/snyk-in-an-ide/)을 통해 추가 지침을 받을 수 있습니다.

{% hint style="info" %}
Snyk 애플리케이션이 Snyk EU 또는 AU 데이터 센터에 호스팅되는 경우, 지역에 따른 URL로 IDE 플러그인 설정에서 이를 명시해야 합니다.\
자세한 내용은 [지역 호스팅 및 데이터 보유](../../../working-with-snyk/regional-hosting-and-data-residency.md)를 참조하세요.
{% endhint %}