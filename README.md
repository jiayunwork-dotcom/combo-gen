# combo-gen

`combo-gen` 是一个纯 Go 标准库实现的组合数学生成器。给定一组元素，它可以枚举出：

- 全排列（`n!` 种顺序）与 k-排列（长度为 k 的有序选取）；
- k-组合（无序、不可重复）与可重复 k-组合（无序、允许重复取同一元素）；
- 多个集合的笛卡尔积（每个集合各取一个元素的所有组合）。

不依赖任何第三方库、可离线运行。

## 构建与测试

```bash
export GOTOOLCHAIN=local CGO_ENABLED=0
go build ./...
go test ./...
```

## 用法

```bash
# 全排列：输出 6 行
combo-gen perm "a,b,c"

# k-排列：长度为 2 的有序选取，输出 6 行
combo-gen perm "a,b,c" -k 2

# 从 n 个下标（0..n-1）中选 k 个的组合，输出 C(5,2)=10 行
combo-gen comb -n 5 -k 2

# 可重复组合，输出 C(5+2-1,2)=15 行
combo-gen comb -n 5 -k 2 -rep

# 笛卡尔积：输出 4 行
combo-gen product "a,b" "1,2"
```

每行是一个结果，元素以空格分隔。位置参数可以出现在 flag 之前（内部由 `reorderFlags` 归一化）。输入非法时（未知子命令、空列表、`k < 0`、不可重复变体的 `k > n` 等）打印错误并以非 0 退出，不会 panic。

## 语义约定

- `Permutations` 对空输入返回**一个空排列**（`0! = 1`），而非空结果；
- `PermutationsK` / `Combinations` 在 `k == 0` 时返回**一个空选取**（选 0 个恰有 1 种方式）；
- `Combinations` / `PermutationsK` 在 `k > len(items)` 时返回错误；可重复组合允许 `k > len(items)`，但元素为空且 `k > 0` 时返回错误；
- `CartesianProduct` 传入 0 个集合时返回错误；任一集合为空时返回空结果（无错误）。

## 目录结构

```
internal/perm     排列：全排列 Permutations、k-排列 PermutationsK
internal/comb     组合：Combinations、CombinationsWithRepetition
internal/product  笛卡尔积：CartesianProduct（变参集合）
```
