# PR Checks 문제 해결

{% hint style="info" %}
PR 설명에 `###`을 사용하면 차단되어 PR 체크가 이루어지지 않습니다.
{% endhint %}

## PR Checks에 대한 일반적인 문제 해결

다음 표는 PR Checks의 일반적인 문제와 해결 방법을 나열합니다.

| 시나리오 | 설명 | 조치 |
| --- | --- | --- |
| PR 체크가 트리거되지 않음. | 저장소는 Snyk에 가져왔지만 PR이 올라가면 PR 체크가 트리거되지 않음. | <ol><li>PR 체크가 구성되어 있는지 확인하기 위해 <a href="configure-pull-request-checks.md">프로젝트 및 통합 설정</a>을 확인하세요.</li><li>웹훅이 없다는 배너가 표시된 경우, 통합 페이지의 프로젝트 설정을 확인하세요.
 <br><br>통합 권한 범위를 다시 확인하세요(자세한 내용은 <a href="../../../scm-ide-and-ci-cd-integrations/snyk-scm-integrations/">Git 저장소</a> 참조).
 <br><br>웹훅을 생성할 수 없는 경우, <a href="https://support.snyk.io">지원팀에 문의</a>하세요.</li></ol> |
| PR 체크가 예상대로이지만 실행되지 않음. | PR 체크가 Git 저장소 (SCM)에 예상된대로 나열되지만 완료되지 않음. | <p>이 문제는 일반적으로 PR 체크를 필요로 하는 브랜치 보호 규칙에 의해 발생합니다. 프로젝트가 비활성화되거나 Snyk에서 제거된 경우 PR 체크가 실행되지 않지만 아직까지 브랜치 보호 규칙이 존재하여 제거하거나 편집될 때까지 계속 적용됩니다.
<br><br>브랜치 보호 규칙을 확인하고 프로젝트가 가져와져 활성화되어 있는지 확인하세요.</p> |
| 하나의 PR에 여러 보안 및 라이선스 PR 체크가 실행됨. | PR이 제출되면 동일한 유형의 다수의 Snyk PR 체크가 실행되어 다른 결과로 실행될 수 있음. | <p>저장소가 여러 Snyk 조직으로 가져와진 경우 PR 체크는 구성된 조직에서 저장소에 대해 실행됩니다.<br><br>PR 체크의 이름을 확인하십시오. 조직 이름이 체크가 실행되는 대상에 표시됩니다. 또는 PR 체크 세부 정보를 선택하면 해당 조직의 결과로 이동합니다.</p> |

## 오픈소스 및 라이선스 체크

잘못된 양성 또는 음성 결과가 발견된 경우 조치를 취하여 문제를 진단하고 보고할 수 있습니다.&#x20;

| 시나리오 | 설명 | 조치 |
| --- | --- | --- |
| 잘못된 양성 | Snyk가 문제가 실제로 존재하지 않는 것으로 확인했지만 실패로 표시한 결과입니다. | <p>Snyk가 패키지 버전에 대한 문제를 잘못 식별한 경우 의존성 버전을 업데이트하도록 <a href="https://support.snyk.io">지원팀에 문의</a>하십시오.</p><p>변경 사항을 계속 진행하려면 결과를 성공으로 표시할 수 있습니다. 자세한 내용은 <a href="analyze-pr-checks-results.md#example-fix-dependency-issues-with-pr-checks">예시: PR 체크로 의존성 문제 해결</a>을 참조하세요.</p> |
| 잘못된 음성 | Snyk가 실제로 존재하는 문제를 감지하지 못해 통과로 표시한 결과입니다. | <p>문제를 해결하려면 <a href="https://snyk.io/vulnerability-disclosure/">취약점 폭로</a>를 제출할 수 있습니다.<br><br>Snyk가 패키지를 잘못 식별하거나 코드 추적이 없어 취약성을 감지하지 못한 경우에는 <a href="https://support.snyk.io">지원팀에 문의</a>하십시오.</p> |

## 코드 분석 체크

다음 표는 코드 분석 오류와 처리 방법을 나열합니다.

| 오류 | 설명 | 조치 |
| --- | --- | --- |
| 코드 분석 시작 실패 | <p>에러 원인:</p><ul><li>PR 체크가 데이터베이스에 생성되지 못함.</li><li>커밋 상태를 전송할 수 없음.</li></ul> | 잠시 기다린 후 [다시 시도](troubleshoot-pr-checks.md#re-run-pr-checks-results)하세요. |
| PR 분석을 완료하지 못했습니다. | PR 체크 결과가 예기치 않은 상태입니다. | 잠시 기다린 후 [다시 시도](troubleshoot-pr-checks.md#re-run-pr-checks-results)하거나 [성공으로 표시](troubleshoot-pr-checks.md#mark-as-successful)하세요. |
| 코드 분석 실패 | <p>에러 원인:</p><ul><li>분석을 완료할 수 없음.</li><li>커밋 상태를 전송할 수 없음.</li></ul> | 잠시 기다린 후 [다시 시도](troubleshoot-pr-checks.md#re-run-pr-checks-results)하거나 [성공으로 표시](troubleshoot-pr-checks.md#mark-as-successful)하세요. |
| 코드 분석 중 상위 소스 속도 제한 트리거 | Git 서버 속도 제한이 도달되었고 저장소를 읽을 수 없음. | 잠시 기다린 후 [다시 시도](troubleshoot-pr-checks.md#re-run-pr-checks-results)하거나 [성공으로 표시](troubleshoot-pr-checks.md#mark-as-successful)하세요. |
| 코드 분석을 수행할 수 있는 유효한 자격 증명 없음. | 개인 액세스 토큰 또는 OAuth가 인식되지 않거나 사용자 액세스가 제공되지 않습니다. | 자격 증명 문제가 있는 경우 Git 저장소 쪽 구성을 검토하세요. |

## 오류 발생 시 조치 방법

### PR 체크 결과 다시 실행

PR 체크 결과를 다시 실행하려면 다음을 수행하세요:

- Snyk 링크에서 [성공으로 표시](troubleshoot-pr-checks.md#mark-as-successful)할 수 있는 특정 사용자 또는 역할 지정.
- PR 체크와 "성공으로 표시"를 선택하십시오.&#x20;