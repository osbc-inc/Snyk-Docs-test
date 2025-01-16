# 스캔 시작

Snyk를 사용하여 코드를 수동 및 자동으로 스캔할 수 있습니다. [Snyk CLI](start-scanning.md#scan-using-the-cli), [Snyk 웹 UI](start-scanning.md#scan-using-the-web-ui), [Snyk API](start-scanning.md#scan-using-the-api) 및 [PR 체크](start-scanning.md#using-pr-checks)를 통해 진행할 수 있습니다.

{% hint style="info" %}
결제 요금제에 따라 계정당 제한된 스캔(테스트)이 있을 수 있습니다. 자세한 내용은 [테스트로 간주되는 것은 무엇인가요?](../working-with-snyk/what-counts-as-a-test.md)를 참조하십시오.
{% endhint %}

<table><thead><tr><th width="220">기능</th><th width="126">Snyk 웹 UI</th><th width="111">Snyk CLI</th><th width="135">Snyk API</th><th>PR 체크</th></tr></thead><tbody><tr><td>자동 스캔</td><td><span data-gb-custom-inline data-tag="emoji" data-code="2714">✔️</span></td><td><span data-gb-custom-inline data-tag="emoji" data-code="2714">✔️</span></td><td><span data-gb-custom-inline data-tag="emoji" data-code="2714">✔️</span></td><td><span data-gb-custom-inline data-tag="emoji" data-code="2714">✔️</span></td></tr><tr><td>수동 스캔</td><td><span data-gb-custom-inline data-tag="emoji" data-code="2714">✔️</span></td><td><span data-gb-custom-inline data-tag="emoji" data-code="2714">✔️</span></td><td><span data-gb-custom-inline data-tag="emoji" data-code="2714">✔️</span></td><td><span data-gb-custom-inline data-tag="emoji" data-code="2796">➖</span></td></tr><tr><td>로컬 스캔</td><td><span data-gb-custom-inline data-tag="emoji" data-code="2796">➖</span></td><td><span data-gb-custom-inline data-tag="emoji" data-code="2714">✔️</span></td><td><span data-gb-custom-inline data-tag="emoji" data-code="2796">➖</span></td><td><span data-gb-custom-inline data-tag="emoji" data-code="2796">➖</span></td></tr><tr><td>CI/CD 파이프라인에 통합</td><td><span data-gb-custom-inline data-tag="emoji" data-code="2796">➖</span></td><td><span data-gb-custom-inline data-tag="emoji" data-code="2714">✔️</span></td><td><span data-gb-custom-inline data-tag="emoji" data-code="2796">➖</span></td><td><span data-gb-custom-inline data-tag="emoji" data-code="2796">➖</span></td></tr><tr><td>프로젝트 취약점 및 구성을 정확히 반영한 결과 획득</td><td><span data-gb-custom-inline data-tag="emoji" data-code="2714">✔️</span></td><td><span data-gb-custom-inline data-tag="emoji" data-code="2714">✔️</span></td><td><span data-gb-custom-inline data-tag="emoji" data-code="2714">✔️</span></td><td><span data-gb-custom-inline data-tag="emoji" data-code="2714">✔️</span></td></tr></tbody></table>

## CLI를 사용한 스캔

{% hint style="info" %}
자세한 내용은 [CLI로 시작하기](../snyk-cli/getting-started-with-the-snyk-cli.md)를 참조하십시오.
{% endhint %}

특정 스캔 방법에 대해 Snyk [CLI 명령](../snyk-cli/cli-commands-and-options-summary.md)을 다음과 같이 사용하십시오:

<table><thead><tr><th width="190">명령어</th><th width="236">기능</th><th>자세한 내용</th></tr></thead><tbody><tr><td><a href="../snyk-cli/commands/test.md">snyk test</a></td><td>오픈 소스 코드 스캔</td><td><a href="../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-open-source/">CLI에서 Snyk 오픈 소스 사용</a></td></tr><tr><td><a href="../snyk-cli/commands/code.md">snyk code test</a></td><td>애플리케이션 코드 스캔</td><td><a href="../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-code/">CLI에서 Snyk 코드 사용</a></td></tr><tr><td><a href="../snyk-cli/commands/container.md">snyk container test</a></td><td>컨테이너 이미지 스캔</td><td><a href="../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-container/">CLI에서 Snyk 컨테이너 사용</a></td></tr><tr><td><a href="../snyk-cli/commands/iac.md">snyk iac test</a></td><td>인프라스트럭처 코드 (IaC) 파일 스캔</td><td><a href="../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-iac/">IaC용 Snyk CLI</a></td></tr><tr><td><a href="../snyk-cli/commands/monitor.md">snyk monitor</a> 및 <a href="../snyk-cli/commands/container-monitor.md">snyk container monitor</a></td><td>새로운 취약점을 방지하기 위해 프로젝트를 지속적으로 모니터링합니다.</td><td><a href="../snyk-cli/scan-and-maintain-projects-using-the-cli/monitor-your-projects-at-regular-intervals.md">정기적으로 프로젝트 모니터링하기</a></td></tr></tbody></table>

## 웹 UI를 사용한 스캔

Snyk 프로젝트를 가져오거나 [프로젝트를 가져와 문제 식별](../getting-started/#import-a-project-to-scan-and-identify-issues)를 클릭하면 스캔이 실행됩니다. Snyk는 그 가져온 프로젝트에 대해 정기적인 스캔을 자동으로 실행하여 코드가 새로 공개된 취약점에 영향을 받았는지 확인합니다.

[웹 UI를 통해 Snyk 탐색](../getting-started/snyk-web-ui.md)을 참조하십시오.

기본 스캔 빈도와 사용 가능한 빈도는 프로젝트 유형에 따라 다르며, 자세한 내용은 [사용 설정](../snyk-admin/groups-and-organizations/usage-settings.md)을 참조하십시오. 프로젝트 **설정**에서 빈도를 설정하거나 [프로젝트 ID로 프로젝트 업데이트](../snyk-api/reference/projects.md#orgs-org_id-projects-project_id)용 API 엔드포인트를 사용할 수도 있습니다.

## API를 사용한 스캔

Snyk API는 코드 테스트를 위한 일련의 엔드포인트를 제공합니다. 호출이 테스트 엔드포인트로 이루어지면 스캔이 계산됩니다.

자세한 내용은 API [테스트](../snyk-api/reference/test-v1.md) 엔드포인트 문서를 참조하십시오.

## PR 체크 사용

Snyk는 모니터링 중인 저장소에 제출된 모든 새로운 풀 리퀘스트(PR)를 스캔하여 코드베이스에 새로운 취약점이 추가되는 것을 방지할 수 있습니다.

자세한 내용은 [풀 리퀘스트 체크](pull-requests/pull-request-checks/)를 참조하십시오.
