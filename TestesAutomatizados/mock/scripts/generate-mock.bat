@echo off
docker run --rm -v %cd%:/src -w /src vektra/mockery:v2.53.3 ^
  --name=HTTPClient ^
  --dir=repository ^
  --output=repository/mocks ^
  --outpkg=mocks