# CLI 테스트의 심각도 임계값

테스트를 보다 잘 통제하기 위해 `snyk test` 명령에 `--severity-threshold` 옵션을 사용할 수 있습니다. 지원되는 수준 중 하나를 선택하여 사용합니다: `low|medium|high|critical`. 해당 옵션을 사용하면 **지정된 수준 또는 그 이상의 취약점만이 보고됩니다.**

`$ snyk test --severity-threshold=medium`

{% hint style="info" %}
`--severity-threshold`를 `low`로 설정하면 임계값을 지정하지 않고 명령을 실행한 것과 동일한 효과가 있습니다. 모든 취약점이 보고됩니다.
{% endhint %}

참고: `--severity-threshold` 옵션은 `snyk test`, `snyk code`, `snyk container`, `snyk iac test` 명령에서 사용할 수 있습니다. 자세한 내용은 각 명령의 [CLI 명령 도움말](../commands/) 페이지를 참조하십시오.