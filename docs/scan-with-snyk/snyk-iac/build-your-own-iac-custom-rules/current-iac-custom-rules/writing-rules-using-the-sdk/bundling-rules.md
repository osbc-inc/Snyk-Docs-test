# 번들링 규칙

준비가 되면 다음 명령어를 실행하여 **사용자 정의 규칙 번들**을 **빌드할 수** 있습니다:

```
snyk-iac-rules build
```

현재 폴더에 생성된 규칙 이외에 더 많은 것이 있다면, `--ignore` 옵션을 사용하여 프로덕션에 적합하지 않은 폴더와 파일을 제외하는 것을 고려해 보세요. 이렇게 하면 프로세스 속도를 높이고 생성된 번들 크기를 작게 유지할 수 있습니다.

**기본 진입점을 덮어쓸 수** 있습니다. `deny`가 아닌 규칙 이름을 `allow` 또는 `violation` 등으로 선택한 경우, 다음과 같이 실행하여 덮어쓸 수 있습니다:

```
snyk-iac-rules build --entrypoint "<패키지 이름>/<함수 이름>"
```

마지막으로, 추출하지 않고 번들 내용을 확인할 수 있습니다. 다음을 실행하여:

```
tar -tf bundle.tar.gz
```

출력은 번들에 포함된 모든 파일입니다:

```
/data.json
/lib/main.rego
/rules/MY_RULE/main.rego
/policy.wasm
/.manifest
```

지금 새로 빌드된 사용자 정의 번들로 [snyk iac test를 실행할 수 있습니다.](../use-iac-custom-rules-with-cli/)