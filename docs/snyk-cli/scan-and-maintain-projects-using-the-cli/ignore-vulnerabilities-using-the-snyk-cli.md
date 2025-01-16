# Snyk CLI를 사용하여 취약점 무시하기

{% hint style="info" %}
의 경우, 이러한 옵션들은 기본적으로 작동합니다.

의 경우, 이러한 옵션들도 작동하지만, 무시를 등록한 후 `snyk test`나 `snyk monitor`를 실행할 때 `--policy-path=` 옵션을 사용해야 합니다. 예를 들어, `snyk container test node --policy-path=.snyk.`와 같이 사용합니다.

[Snyk ](../../scan-with-snyk/snyk-iac/scan-your-iac-source-code/)의 경우, [snyk 정책 파일을 사용한 IaC 무시](snyk-cli-for-iac/iac-ignores-using-the-.snyk-policy-file.md)를 참조하십시오.

의 경우, [Snyk 코드 CLI 테스트에서 디렉터리와 파일 제외하기](snyk-cli-for-snyk-code/exclude-directories-and-files-from-snyk-code-cli-tests.md)를 참조하십시오.
{% endhint %}

가끔 Snyk이 업데이트나 Snyk 패치가 없는 취약점을 발견하거나 현재 응용 프로그램에서 쉽게 이용될 수 없다고 생각하는 경우, 해당 취약점을 일정 기간 동안 Snyk에 무시하도록 지시할 수 있습니다.

`snyk ignore --id=<ISSUE_ID> [--expiry=<EXPIRY>] [--reason=<REASON>] [--policy-path=<PATH_TO_POLICY_FILE>] [<OPTIONS>]`

`snyk ignore` 명령은 `.snyk` 파일을 업데이트하며 다음 옵션들을 지원합니다:

| **옵션**                                | **설명**                                                                                                                                                                                                                                                                                                                                                                                                                                          | **기본값** | **필수** |
| ------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------- | ------ |
| `--id`                                | <p>무시할 문제에 대한 Snyk ID. <code>snyk test</code>를 실행하여 특정 취약점에 대한 URL의 마지막 세그먼트를 가져와서 찾을 수 있습니다.</p><p><strong>예시:</strong> <a href="https://security.snyk.io/vuln/SNYK-DEBIAN10-NODETOUGHCOOKIE-5759362">https://security.snyk.io/vuln/SNYK-DEBIAN10-NODETOUGHCOOKIE-5759362</a>에서 발견된 취약점의 Snyk ID는 다음과 같습니다:</p><p><a href="https://security.snyk.io/vuln/SNYK-DEBIAN10-NODETOUGHCOOKIE-5759362">SNYK-DEBIAN10-NODETOUGHCOOKIE-5759362</a>.</p> | 없음      | 예      |
| `--expiry`                            | <p>YYYY-MM-DD 형식의 만료 날짜 (<a href="https://tools.ietf.org/html/rfc2822#page-14">RFC2822</a>와 <a href="https://www.iso.org/iso-8601-date-and-time-format.html">ISO 8601</a> 지원).</p><p>예시: <code>--expiry=2017-04-30</code>.</p>                                                                                                                                                                                                                  | 30 일    | 아니오    |
| `--reason`                            | 이 문제를 무시하는데 사용하는 읽기 쉬운 \<REASON>. 예시: `reason='현재 쉽게 이용되지 않음'`.                                                                                                                                                                                                                                                                                                                                                                                 | 없음      | 아니오    |
| `--policy-path=<PATH_TO_POLICY_FILE>` | 수동으로 전달할 .snyk 정책 파일의 경로.                                                                                                                                                                                                                                                                                                                                                                                                                       | 없음      | 아니오    |
| `--path`                              | 문제를 무시할 리소스의 경로. 예시: `path='tough-cookie@2.15.8'`                                                                                                                                                                                                                                                                                                                                                                                               | 모두      | 아니오    |
