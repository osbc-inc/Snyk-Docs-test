# 사용자 정의 Rego 내장 함수

SDK는 테스트 중에 사용할 수 있는 몇 가지 도우미 Rego 함수를 등록합니다.

이러한 함수들은 다음과 같습니다:
* `hcl2.unmarshal_file`: 파일 경로를 가져와 HCL2 형식에서 JSON으로 변환합니다.
* `yaml.unmarshal_file`: 파일 경로를 가져와 YAML 형식에서 JSON으로 변환합니다.

이러한 함수들은 `lib/testing/main.rego`의 테스트 프레임워크에서 사용되며, JSON fixture를 직접 생성하는 대신 픽스처 파일의 경로를 전달할 수 있습니다.