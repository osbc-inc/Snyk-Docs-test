# Docker에서 인증서 확인 비활성화하기

인증서 확인을 비활성화하려면, 예를 들어 자체 서명된 인증서의 경우, 다음을 **docker run** 명령에 추가합니다:

```
-e NODE_TLS_REJECT_UNAUTHORIZED=0
```