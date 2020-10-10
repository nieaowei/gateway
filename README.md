# gateway

[![Code](https://nekilc.coding.net/badges/gateway/job/335633/build.svg)](https://nekilc.coding.net/p/gateway/ci/job)
[![Build](https://nekilc.coding.net/badges/gateway/job/336986/build.svg)](https://nekilc.coding.net/p/gateway/ci/job)

go网关

使用coding进行持续集成，coding仓库：[这里](https://nekilc.coding.net/public/gateway/server/git/files)

|服务|代码扫描|构建状态|连接|
|:---:|:---:|:---:|:---:|
|Dashboard|[![构建状态](https://nekilc.coding.net/badges/gateway/job/335633/build.svg)](https://nekilc.coding.net/p/gateway/ci/job)|[![构建状态](https://nekilc.coding.net/badges/gateway/job/336986/build.svg)](https://nekilc.coding.net/p/gateway/ci/job)|[链接](http://code.nekilc.com:8880)|
|Web|[![构建状态](https://nekilc.coding.net/badges/gateway/job/335634/build.svg)](https://nekilc.coding.net/p/gateway/ci/job)|[![构建状态](https://nekilc.coding.net/badges/gateway/job/338147/build.svg)](https://nekilc.coding.net/p/gateway/ci/job)|[链接](http://code.nekilc.com:8888)|
|Http Proxy|[![构建状态](https://nekilc.coding.net/badges/gateway/job/335633/build.svg)](https://nekilc.coding.net/p/gateway/ci/job)|[![构建状态](https://nekilc.coding.net/badges/gateway/job/336986/build.svg)](https://nekilc.coding.net/p/gateway/ci/job)|[链接](http://code.nekilc.com:8800)|
|Grpc Proxy|||
|Tcp Proxy|||

# todo

- [x] Add a series of data statistics interface. -- 2020.8.16
- [x] Add avatar acquisition interface. -- 2020.8.16
- [x] Waiting to fix the problem of incomplete data on GetServiceDetail -- 2020.8.16
- [x] Refactor DTO. - 2020.8.17
- [ ] Add avatar upload interface. - 2020.8.18
- [x] Organize validator and swag documents.
- [x] The part of DTO is wait to be completed. - 2020。8.19
- [x] Some uniqueness checks in the DAO module.
- [ ] Fix data validation(ex. MetaDataHeader.).
- [ ] Complete the interfaces related to tenant management.- 2020.8.24
- [x] Refactor the GetServiceStatistical and add GetServiceAmount.
- [ ] Integrated log.