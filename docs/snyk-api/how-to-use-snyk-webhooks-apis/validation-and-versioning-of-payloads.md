# Payload의 유효성 검사 및 버전 관리

## Payload의 유효성 검사

웹훅에 전송된 모든 전송은 `X-Hub-Signature` 헤더를 포함하며, 해당 헤더에는 전송에 대한 해시 서명이 포함됩니다. 이 서명은 요청 본문의 HMAC 해시 다이제스트로, sha256을 사용하여 생성되며 `secret`를 HMAC 키로 사용합니다.

다음과 같은 Node.JS에서의 함수를 사용하여 이러한 서명을 확인할 수 있습니다:

```javascript
import * as crypto from 'crypto';

function verifySignature(request, secret) {
  const hmac = crypto.createHmac('sha256', secret);
  const buffer = JSON.stringify(request.body);
  hmac.update(buffer, 'utf8');

  const signature = `sha256=${hmac.digest('hex')}`;

  return signature === request.headers['x-hub-signature'];
}
```

## Payload 버전 관리

Payload는 시간이 지남에 따라 진화할 수 있으므로 버전이 지정됩니다. Payload 버전은 `X-Snyk-Event` 헤더의 접미사로 제공됩니다. 예를 들어, `project_snapshot/v0`는 `project_snapshot` 이벤트의 `v0` 버전임을 나타냅니다.

버전 번호는 주요 변경 사항이 발생할 때만 증가하며, 예를 들어 이전에 존재하던 필드를 제거하거나 필드 이름을 변경할 때입니다. 새로운 필드를 추가하는 것과 같은 추가적인 변경 사항이 있는 경우 버전 번호는 증가하지 않습니다.

베타 단계에서는 웹훅 페이로드의 구조가 언제든지 변경될 수 있으므로 Snyk은 페이로드 버전을 확인하는 것을 권장합니다.