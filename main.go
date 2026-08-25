package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"combo-gen/internal/comb"
	"combo-gen/internal/perm"
	"combo-gen/internal/product"
	"combo-gen/internal/server"
)

func fatal(format string, args ...any) {
	fmt.Fprintf(os.Stderr, "combo-gen: "+format+"\n", args...)
	os.Exit(1)
}

func reorderFlags(args []string) []string {
	var flags, pos []string
	i := 0
	for i < len(args) {
		a := args[i]
		if strings.HasPrefix(a, "-") {
			if i+1 < len(args) && !strings.HasPrefix(args[i+1], "-") && !strings.Contains(a, "=") {
				flags = append(flags, a, args[i+1])
				i += 2
			} else {
				flags = append(flags, a)
				i++
			}
		} else {
			pos = append(pos, a)
			i++
		}
	}
	return append(flags, pos...)
}

func splitList(s string) []string {
	parts := strings.Split(s, ",")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		if p = strings.TrimSpace(p); p != "" {
			out = append(out, p)
		}
	}
	return out
}

func printResults(results [][]string) {
	for _, r := range results {
		fmt.Println(strings.Join(r, " "))
	}
}

func main() {
	if len(os.Args) < 2 {
		runServer(nil)
		return
	}
	cmd := os.Args[1]
	switch cmd {
	case "serve":
		runServer(os.Args[2:])
	case "perm":
		args := reorderFlags(os.Args[2:])
		runPerm(args)
	case "comb":
		args := reorderFlags(os.Args[2:])
		runComb(args)
	case "product":
		args := reorderFlags(os.Args[2:])
		runProduct(args)
	default:
		args := reorderFlags(os.Args[1:])
		runPerm(args)
	}
}

func runServer(args []string) {
	addr := ":8080"
	for i := 0; i < len(args)-1; i++ {
		if args[i] == "-addr" || args[i] == "--addr" {
			addr = args[i+1]
			break
		}
	}
	cfg := server.Config{Addr: addr}
	fmt.Fprintf(os.Stdout, "combo-gen server listening on %s\n", server.FormatAddr(addr))
	if err := server.ListenAndServe(cfg); err != nil {
		fmt.Fprintf(os.Stderr, "server error: %v\n", err)
		os.Exit(1)
	}
}

func runPerm(args []string) {
	fs := flag.NewFlagSet("perm", flag.ContinueOnError)
	k := fs.Int("k", -1, "length of each permutation (default: use all items)")
	if err := fs.Parse(args); err != nil {
		fatal("%v", err)
	}
	if fs.NArg() < 1 {
		fatal("perm requires a comma-separated list argument")
	}
	items := splitList(fs.Arg(0))
	if len(items) == 0 {
		fatal("perm requires at least one item")
	}
	if *k < 0 {
		printResults(perm.Permutations(items))
		return
	}
	res, err := perm.PermutationsK(items, *k)
	if err != nil {
		fatal("%v", err)
	}
	printResults(res)
}

func runComb(args []string) {
	fs := flag.NewFlagSet("comb", flag.ContinueOnError)
	n := fs.Int("n", 0, "number of indices to choose from")
	k := fs.Int("k", 0, "size of each combination")
	rep := fs.Bool("rep", false, "allow an index to be chosen more than once")
	if err := fs.Parse(args); err != nil {
		fatal("%v", err)
	}
	if *n < 0 {
		fatal("comb: n must be non-negative")
	}
	items := make([]string, *n)
	for i := range items {
		items[i] = strconv.Itoa(i)
	}
	var (
		res [][]string
		err error
	)
	if *rep {
		res, err = comb.CombinationsWithRepetition(items, *k)
	} else {
		res, err = comb.Combinations(items, *k)
	}
	if err != nil {
		fatal("%v", err)
	}
	printResults(res)
}

func runProduct(args []string) {
	fs := flag.NewFlagSet("product", flag.ContinueOnError)
	if err := fs.Parse(args); err != nil {
		fatal("%v", err)
	}
	if fs.NArg() < 1 {
		fatal("product requires at least one comma-separated set")
	}
	sets := make([][]string, 0, fs.NArg())
	for _, a := range fs.Args() {
		set := splitList(a)
		if len(set) == 0 {
			fatal("product: set %q has no items", a)
		}
		sets = append(sets, set)
	}
	res, err := product.CartesianProduct(sets...)
	if err != nil {
		fatal("%v", err)
	}
	printResults(res)
}
