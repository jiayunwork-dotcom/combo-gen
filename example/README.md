# combo-gen examples

Offline usage examples (no network required).

Build first:

```bash
export GOTOOLCHAIN=local CGO_ENABLED=0
go build -o /tmp/combo-gen .
```

All orderings of three items (3! = 6 lines):

```bash
/tmp/combo-gen perm "a,b,c"
```

```text
a b c
a c b
b a c
b c a
c a b
c b a
```

Ordered selections of length 2 (3!/1! = 6 lines). The positional list may come
before the flag:

```bash
/tmp/combo-gen perm "a,b,c" -k 2
```

Choose 2 of the indices 0..4, unordered (C(5,2) = 10 lines):

```bash
/tmp/combo-gen comb -n 5 -k 2
```

```text
0 1
0 2
0 3
0 4
1 2
1 3
1 4
2 3
2 4
3 4
```

Same shape but an index may repeat (C(6,2) = 15 lines):

```bash
/tmp/combo-gen comb -n 5 -k 2 -rep
```

Cartesian product of two sets (2 * 2 = 4 lines):

```bash
/tmp/combo-gen product "a,b" "1,2"
```

```text
a 1
a 2
b 1
b 2
```

Bad input exits non-zero and prints to stderr instead of panicking:

```bash
/tmp/combo-gen comb -n 2 -k 5   # k > n
/tmp/combo-gen perm ""          # empty list
/tmp/combo-gen product          # no sets
/tmp/combo-gen bogus            # unknown command
echo $?                         # 1
```
