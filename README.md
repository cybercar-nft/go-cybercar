# go-cybercar
CyberCar NFT CLI written in Go

### Commands

- `user`: 普通用户命令
  - `airdropQuota` 查询空投额度
  - `paused` 查询暂停状态
  - `phase` 运营活动阶段
  - `mintQuota` 白名单额度
  - `claim`: 认领空投
  - `mint`: 认购
- `admin`: 管理命令
  - `addAirdrop`: 添加空投名单
  - `pause`: 暂停
  - `unpause`: 恢复
  - `setPhase`: 设置运营阶段
  - `addWhitelist`: 添加白名单