# Kustomize 파일들

쿠버네티스 매니페스트 파일을 빌드하고 Snyk CLI `iac test` 명령을 사용하여 스캔하여 Kustomize 템플릿을 검사할 수 있습니다.

```bash
kustomize build > kubernetes.yaml
snyk iac test kubernetes.yaml
```

Kustomize 템플릿에 따라 빌드 인수 뒤에 이름을 제공해야 할 수 있습니다.