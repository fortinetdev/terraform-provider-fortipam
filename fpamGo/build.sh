#!/bin/bash
java -jar swagger-codegen-cli-2.4.32.jar generate -i openapi_spec/FortiPAM_swagger_document.json -l go
find . -name '*.go' -exec sed -i -e 's/package swagger/package fpam_go/g' {} \;
find . -name '*.md' -exec sed -i -e 's/import ".\/swagger"/package .\/fpam_go/g' {} \;