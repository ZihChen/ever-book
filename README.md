# EverBook

![GitHub last commit](https://img.shields.io/github/last-commit/ZihChen/ever-book)
![Github repo size](https://img.shields.io/github/repo-size/ZihChen/ever-book)
![GitHub top language](https://img.shields.io/github/languages/top/ZihChen/ever-book)
![Lines of code](https://img.shields.io/tokei/lines/github/ZihChen/ever-book)

### 啟動腳本
到 LINE Developers 創建一個新的 Channel，產生出對應的 Channel secret 及  Channel access token，並帶入啟動腳本對應的兩個參數當中
```bash
cd ever-book
sh DeployService.sh [CHANNEL_SECRET] [CHANNEL_TOKEN]
```

### 功能
- 記帳
- 刪除紀錄
- 查看過去三個月的記帳統計
- 群組帳本