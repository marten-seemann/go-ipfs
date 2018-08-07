package coreunix

import (
	"context"

	core "github.com/ipfs/go-ipfs/core"
	uio "gx/ipfs/QmXBooHftCHoCUmwuxSibWCgLzmRw2gd2FBTJowsWKy9vE/go-unixfs/io"
	path "gx/ipfs/Qmet61sdJdo6ZHDc59ZbffN5ED79K5kqfqkRjBx7ncaDg4/go-path"
	resolver "gx/ipfs/Qmet61sdJdo6ZHDc59ZbffN5ED79K5kqfqkRjBx7ncaDg4/go-path/resolver"
)

func Cat(ctx context.Context, n *core.IpfsNode, pstr string) (uio.DagReader, error) {
	r := &resolver.Resolver{
		DAG:         n.DAG,
		ResolveOnce: uio.ResolveUnixfsOnce,
	}

	dagNode, err := core.Resolve(ctx, n.Namesys, r, path.Path(pstr))
	if err != nil {
		return nil, err
	}

	return uio.NewDagReader(ctx, dagNode, n.DAG)
}
