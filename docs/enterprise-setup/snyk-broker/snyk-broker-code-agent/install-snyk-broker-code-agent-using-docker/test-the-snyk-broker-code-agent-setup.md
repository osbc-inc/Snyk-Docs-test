# Snyk 브로커 코드 에이전트 설정 테스트

{% hint style="warning" %}
**사용 중단됨**

코드 에이전트는 더 이상 유지되지 않으며 사용이 중단되었습니다.

Snyk 브로커를 통해 Snyk 코드 분석을 실행하는 권장 방법은 [브로커 코드](https://docs.snyk.io/enterprise-setup/snyk-broker/install-and-configure-snyk-broker/advanced-configuration-for-snyk-broker-docker-installation/snyk-code-clone-capability-with-broker-for-docker)를 사용하는 것입니다. 코드 에이전트는 이점이 없는 대안 방법입니다. 자세한 내용은 Snyk 통합 컨설턴트 또는 기술 지원 매니저에게 문의하시거나 [Snyk 지원팀](https://support.snyk.io)에 문의하십시오.

자동 [PR Checks](https://docs.snyk.io/scan-with-snyk/pull-requests/pull-request-checks) 기능은 Snyk 브로커 - 코드 에이전트에 대해 지원되지 않습니다.
{% endhint %}

Snyk 브로커 설정에서 코드 에이전트를 테스트하려면 리포지토리를 Snyk에 가져와서 코드 분석 프로젝트 및 결과를 수신했는지 확인하십시오. 선택한 리포지토리가 Snyk에 가져와지고 {{Snyk 코드}} 결과가 표시되었을 때, 코드 에이전트를 사용한 Snyk 브로커 설정이 성공적으로 작동 중인 것입니다.

설정을 테스트하려면 다음 단계를 따르십시오:

1. Snyk 웹 UI에서 코드 에이전트를 테스트하려는 조직을 엽니다.

2. 필요한 조직이 열리면 **프로젝트 추가**를 선택하십시오. 그런 다음 코드 에이전트를 설정한 SCM을 선택하십시오:

<figure><img src="../../../../.gitbook/assets/Code Agent - Test - Selecting SCM for import.png" alt="프로젝트 추가 및 SCM 선택"><figcaption><p>프로젝트 추가 및 SCM 선택</p></figcaption></figure>

3. **개인 및 조직 리포지토리** 페이지에서 Snyk 코드가 분석하도록 원하는 리포지토리를 선택하고 **선택한 리포지토리 추가** 버튼을 클릭하십시오:

<figure><img src="../../../../.gitbook/assets/Code Agent - Test - Selecting repos for import.png" alt="{Snyk 코드를 분석할 리포지토리 선택"><figcaption><p>{Snyk 코드를 분석할 리포지토리 선택</p></figcaption></figure>

선택한 리포지토리가 {{Snyk 코드}}에 가져와지며 **프로젝트** 페이지에 진행률 표시줄이 표시됩니다. 가져오기가 완료되면 **프로젝트** 페이지에 가져오기 성공 메시지가 나타납니다. 가져온 리포지토리는 **코드 분석** 결과를 포함하는 **코드 분석** 프로젝트를 포함한 대상 폴더로 나타납니다.

<figure><img src="../../../../.gitbook/assets/Code Agent - Test - Code Analysis Project.png" alt="확인 메시지 및 대상 폴더"><figcaption><p>확인 메시지 및 대상 폴더</p></figcaption></figure>

4. {{Snyk 코드}} 결과를 보려면 **코드 분석** 프로젝트를 선택하십시오.

**코드 분석** 페이지를 열어 선택한 리포지토리에서 발견된 취약점 문제의 목록과 세부 정보가 표시됩니다:

<figure><img src="../../../../.gitbook/assets/Code Agent - Test - Code Analysis page.png" alt="취약점이 표시된 코드 분석 페이지"><figcaption><p>취약점이 표시된 코드 분석 페이지</p></figcaption></figure>

5. 취약점 문제의 세부 정보를 보려면 **전체 세부 정보**를 선택하십시오.

문제의 세부 정보가 표시되며 브로커 클라이언트 설정 방식에 따라 코드 스니펫이 **데이터 플로우** 탭에 나타나거나 나타나지 않을 수 있습니다:

이 예시는 코드 스니펫이 표시된 {{Snyk 코드}} 결과를 보여줍니다.

<figure><img src="../../../../.gitbook/assets/Broker - Results - with code snippets (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (2).png" alt="코드 스니펫이 표시된 브로커 클라이언트 실행"><figcaption><p>코드 스니펫이 표시된 브로커 클라이언트 실행</p></figcaption></figure>

이 예시는 코드 스니펫이 없는 {{Snyk 코드}} 결과를 보여줍니다:

<figure><img src="../../../../.gitbook/assets/Broker - Results - without code snippets (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (1) (4) (1).png" alt="코드 스니펫이 표시되지 않은 브로커 클라이언트 실행"><figcaption><p>코드 스니펫이 표시되지 않은 브로커 클라이언트 실행</p></figcaption></figure>

Snyk 브로커 코드 에이전트 설정의 문제를 해결하는 방법에 대한 자세한 내용은 [브로커 문제 해결](../../troubleshooting-broker.md)을 참조하십시오.