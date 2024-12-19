# GitHub 공급업체 구성

## GitHub 공급업체의 인증

`iac describe`를 사용하려면 GitHub에 인증된 요청을 만들기 위한 자격 증명을 설정해야 합니다. Snyk은 [환경 변수](https://registry.terraform.io/providers/integrations/github/latest/docs#argument-reference)에서 구성 정보를 검색합니다.

GitHub 토큰은 [이 GitHub 페이지](https://github.com/settings/tokens/)에서 생성할 수 있습니다.

```
$ GITHUB_TOKEN=14758f1afd44c09b7992073ccf00b43d \
  GITHUB_ORGANIZATION=my-org \
  snyk iac describe --to="github+tf"
```

## 최소한의 권한 정책 <a href="#least-privileged-policy" id="least-privileged-policy"></a>

`iac describe`가 모든 GitHub 지원 리소스를 스캔하려면 다음 GitHub 토큰 범위가 필요합니다.

```
# 사적 저장소 열거에 필요함
repo

# 조직 팀 목록 및 기타 조직 관련 리소스 나열에 필요함
read:org
```

## **저장소 권한**

저장소 목록의 오류를 보려면 토큰에 `repo` 권한을 설정해야 합니다. 이 권한이 없으면 모든 사적 저장소가 원격으로 삭제된 것으로 나타납니다.