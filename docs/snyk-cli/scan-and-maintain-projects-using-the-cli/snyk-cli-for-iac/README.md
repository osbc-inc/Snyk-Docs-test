# Snyk IaC용 CLI

{% hint style="info" %}
[IaC+](../../../scan-with-snyk/snyk-iac/getting-started-with-iac+-and-cloud-scans/) 버전의 Snyk CLI를 사용하려면 Snyk CLI v1.1022.0 이상을 설치하십시오.
{% endhint %}

## 개요

CLI를 사용하려면 먼저 [설치](../../install-or-update-the-snyk-cli/)하고, 그리고 [인증](../../commands/auth.md)해야 합니다.

Snyk Infrastructure as Code를 사용하면 CLI에서 직접 구성 파일을 테스트할 수 있습니다. 자세한 내용은 다음 페이지를 참조하십시오:

* [IaC 파일 테스트](test-your-iac-files/)
* [CLI 결과를 Snyk 웹 UI와 공유](share-cli-results-with-the-snyk-web-ui.md)
* [`.snyk` 정책 파일 사용한 IaC 무시](iac-ignores-using-the-.snyk-policy-file.md)
* [명령줄을 사용한 IaC 예외 설정](iac-exclusions-using-the-command-line.md)
* [IaC CLI 테스트 결과 이해](understand-the-iac-cli-test-results/) (리포트에 관한 정보 포함)

## 정기적인 IaC 파일 테스트

Snyk Infrastructure as Code는 CLI가 IaC 소스 파일을 플랫폼으로 주기적으로 보내지 않기 때문에 `snyk monitor`와 동일한 명령이 없습니다.

IaC CLI 결과가 Snyk 웹 UI에 나타나려면 [`snyk iac test --report`](https://docs.snyk.io/products/snyk-infrastructure-as-code/share-cli-results-with-the-snyk-web-ui)를 사용하여 일회성 스냅샷을 캡처하십시오. 선택적으로 이 명령을 주기적으로 실행하여 IaC 파일을 정기적으로 테스트할 수 있습니다.

또는 [SCM 통합](../../../scm-ide-and-ci-cd-integrations/snyk-scm-integrations/)을 추가할 수 있으며, Snyk는 주기적으로 주어진 Git 리포지토리를 모니터링하고 테스트할 것입니다.

## 프록시를 통한 Snyk 사용

프록시를 사용하는 경우 [Snyk CLI용 프록시 구성](../../configure-the-snyk-cli/proxy-configuration-for-snyk-cli.md)를 참조하십시오.

특히 IaC 스캔의 경우, \*.snyk.io 주소를 화이트리스트에 추가해야 하며, 이에 대한 설명은 [Snyk IP 주소를 허용목록에 추가하는 방법](https://support.snyk.io/s/article/How-can-we-allowlist-Snyk-IP-addresses)?에서 확인할 수 있습니다.