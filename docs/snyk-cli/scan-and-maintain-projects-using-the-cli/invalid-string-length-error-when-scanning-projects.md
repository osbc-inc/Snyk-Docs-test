# 프로젝트를 스캔할 때 무효한 문자열 길이 오류가 발생하는 문제

무효한 문자열 길이 오류는 다음 상황에서 발생할 수 있습니다:

* `--all-projects` 옵션을 `--json` 또는 `--json-file-output=<OUTPUT_FILE_PATH>` 옵션과 함께 사용하여 **여러 프로젝트**를 한꺼번에 스캔하는 경우.
* `--json` 또는 `--json-file-output=<OUTPUT_FILE_PATH>` 옵션을 사용하여 **큰 프로젝트**를 스캔하는 경우.

무효한 문자열 길이 오류를 피하기 위해 다음 대처 방법을 사용할 수 있습니다.

* `--json` 또는 `--json-file-output=<OUTPUT_FILE_PATH>` 옵션을 제거하세요. CLI를 사용하는 통합에서는 이 옵션을 제거할 수 없을 수도 있습니다. 예를 들어 CI/CD 통합에서.
* JSON 출력이 필요한 경우 `--severity-threshold=<SEVERITY_LEVEL>` 옵션을 사용하여 심각도 임계값을 늘려보세요. 이렇게 하면 덜 중요한 취약점이 보고되는 수가 줄어들어 JSON 파일 출력의 크기가 줄어들 가능성이 높습니다.
* `--all-projects` 옵션을 사용하여 여러 프로젝트를 스캔하는 경우, 이 옵션을 제거하고 프로젝트를 개별적으로 스캔해보세요.
* JSON 출력을 줄이는 또 다른 방법으로 `--prune-repeated-subdependencies` 옵션을 추가할 수 있습니다. 이를 통해 결과 스캔의 크기가 충분히 줄어들어 JSON이 해결될 수 있을 것입니다.&#x20;

CLI 옵션에 대한 자세한 정보는 [CLI 테스트 명령 도움말](../commands/test.md)을 참조하십시오.