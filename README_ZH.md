# TMDB CLI

[English](README.md) | [简体中文](README_ZH.md)


一个功能强大的命令行工具，用于查询 [The Movie Database (TMDB)](https://www.themoviedb.org/) API。使用 Go 语言开发，支持搜索电影、电视剧、人物，并支持多种导出格式，如 JSON、Markdown 和 NFO (兼容 Kodi/Jellyfin)。

## 功能特性

- **搜索**: 支持搜索电影、电视剧和人物（支持多重搜索）。
- **详细信息**: 获取电影、电视剧集、季以及单个剧集的完整详情。
- **热门推荐**: 查看全球范围内的实时热门内容。
- **合集**: 获取系列电影（Collection）的详细资料。
- **通过外部 ID 查找**: 支持通过 IMDb、TVDB 等第三方 ID 查找对应的 TMDB 条目。
- **多种输出格式**:
  - `json`: 标准 JSON 输出（支持字段筛选）。
  - `markdown`: 格式精美的 Markdown 预览。
  - `nfo`: 兼容 Kodi/Jellyfin 的 XML 元数据文件。
  - `table`: 清洁的命令行表格视图。
- **海报下载**: 在导出 NFO 时可自动下载电影/系列海报到本地。
- **持久化配置**: 支持保存 API Token 和首选响应语言。

## 安装指南

克隆仓库并编译二进制文件：

```bash
git clone https://github.com/gahoolee/tmdb-cli.git
cd tmdb-cli
go build -o tmdb
```

## 配置说明

在使用工具前，需要设置您的 TMDB API 令牌 (v4 Read Access Token 或 v3 API Key)。

```bash
# 设置 API 令牌
./tmdb config set-auth 您的_TMDB_API_TOKEN

# 设置默认语言 (可选，如 zh-CN, en-US)
./tmdb config set-lang zh-CN
```

配置将保存在 `~/.tmdb.json` 文件中。

## 使用示例

### 通用选项

- `--format`: 输出格式 (json, markdown, nfo, table)。默认: `json`。
- `--output`, `-o`: 将输出保存到指定文件。
- `--fields`, `-f`: 以逗号分隔的字段列表 (如 `title,overview,budget`)。
- `--language`, `-l`: 覆盖该次请求的默认语言。
- `--poster`: 下载本地海报 (仅在 `nfo` 格式下生效)。

### 常用命令示例

#### 搜索电影
```bash
./tmdb search "盗梦空间" --type movie --format table
```

#### 获取电影详情并下载 NFO 及海报
```bash
./tmdb movie 27205 --format nfo --poster
```

#### 获取电视剧详情
```bash
./tmdb tv 60625 --format markdown
```

#### 获取特定的季或剧集详情
```bash
# 获取某剧集的第 1 季
./tmdb tv 60625 --season 1

# 获取第 1 季的第 1 集
./tmdb tv 60625 --season 1 --episode 1 --format markdown
```

#### 通过外部 ID (IMDb) 查找
```bash
./tmdb find tt0133093 --source imdb_id --format table
```

#### 查看热门电影
```bash
./tmdb trending --type movie --time day --format table
```

#### 获取电影合集详情
```bash
./tmdb collection 10 --format markdown
```

## 输出模式说明

### NFO 生成
当使用 `--format nfo` 时，工具将生成与 Kodi、Jellyfin 和 Emby 等媒体管理器兼容的 XML 元数据文件：
- 电影: 生成 `.nfo` 文件。
- 电视剧: 生成 `tvshow.nfo`。
- 剧集季: 生成 `season.nfo`。
- 单集: 生成对应的 `.nfo` 文件。

配合 `--poster` 标志，可以帮助您快速搭建本地媒体库。

## 环境要求
- Go 1.26 或更高版本。

## 开源协议
MIT License
