// +build !skippackr
// Code generated by github.com/gobuffalo/packr/v2. DO NOT EDIT.

// You can use the "packr2 clean" command to clean up this,
// and any other packr generated files.
package packrd

import (
	"github.com/gobuffalo/packr/v2"
	"github.com/gobuffalo/packr/v2/file/resolver"
)

var _ = func() error {
	const gk = "58cbd89fd8df298a2f4c690a94262881"
	g := packr.New(gk, "")
	hgr, err := resolver.NewHexGzip(map[string]string{
		"2576b02ce4cf56084d3ef2b6714559cd": "1f8b08000000000000ff548c31aec23010446bcf295629be922f487a6a6a1a4ee0388bb120b6b55e835094bb53100ada99f75eb6ee663d53617904c740987312a516a6f141af75ec5d9a87a02c21faa1d4b1381bf7f95e7d8843d124d673830ed057663a7f325454aa535a6026daa0fe681356e052a3a3133fdb9fa7a3ffafbbc0086b95487fdbb4c098e940d30e66c58a77000000ffff7abf77f9b4000000",
		"4adc860d2724dbd31a7da30e60cd476c": "1f8b08000000000000ff648f314ec4301045ebcc2946ae9255484e404f45b1141488226b0d598b8d6d66c641c8f2dd91376e1095a5ef3fefe9c7c57e2e2be1553502b82d0656eca133394fcfcb46a5cc42bc3b4b061011cdeaf49a2e930ddbec94d8f975967411bbf887784babf33387a4c4adee49e7ca363000ec0b57b6ec164f8d3abd1c6ffdfe48dee2f97eddcbbfc6806fef077a7a528d980fd0230a744c9ad8ff2d64e8ba6c9444cd88f2234a5b19a12b509ae9c8faeffbf6e94c1283177ae53a6b44c653cbbf12890e48cc81abb5c9bcbb41f90d0000ffff3d3fccb43f010000",
		"8cfce0dcb361a8f7da25bc989bab49b8": "1f8b08000000000000ff94924f6b1b3110c5cfd6a71874285270e5f61af0a1ae4deb434c69a097108a563b598bec4acb6876dbb2ec772ff2bfd8294dd39376f4f6bdf763506bdda3ad101aeb8310be69233128319195e76d5718179b9967241faa59ea8ae46c78dbd65de5c32c71245ba17cd5cf143b46920200400e83d9d806c771b6656effb84c48bd77cf83d336b6a9dd6597e87c636b29b410bd2548d4c3d5c1656ef7a710fcab4518067313cb8fb54d691c2131758e611093120ef86669a3188578e882830dfe501aae2e4d839810724701de5c08c378f229fbcca4611d3c2f6d5417453a8765da391c7173a535a516136b4a9843993f6e7c459651e97f357c666e9586bbfbfd7e4d9e9f78f376cdd79da212f52fa67da1e830a5d54f261f9277aaa8a37b84ab23fd228f53c0a3fea49c2c53c01e0327b8bb3f69f9420312453a5b64f0f56b58b2f96f1c593b63c8e3141e10e1f038cc727ffe77f969f7bb6752c0f51cac29cd72a1749ecd2db29255a4e69a6d51e3f7d8b28f21c929c8d5e6d37ab39aaf4388cb85d4e643c7f11827265abc54fb0d29f91894ce8fd487ea0c58be33efa518c5ef000000ffff86605978ac030000",
		"901dee514c4b917fa63490110021898e": "1f8b08000000000000ff2a484cce4e4c4f55c8cd4f49cde1e22aa92c4855a8aed6f3cd4f71ce492c2eaead55282e292a4d2e51a8e6aae502040000ffffcb815bad2d000000",
		"9040efce6e173ddbe5aa65b5ee616485": "1f8b08000000000000ff74cc3b4ec4301080e19a39c59450646c030d97e00e8e333883fcc28f1444b9fb6ad345daadff4f7fcccb088cfb4edf36f27100f88c86cc0740e5bf2195f1151011bdf475cce47254d2b94af2aa8db9399ba6128697849b264d065e9ec85e6de7c916996a718ff0afa4ff75289f6bc4cdd01799cf4b6f6b2ead9cbf859d441beeec9df4458dfa6337562ec85935697803b8050000ffff40556b30e7000000",
		"e2180af47102e9758aad571e6e107783": "1f8b08000000000000ff2a484cce4e4c4f554849cce7e24a2bcd4b5670c9d7d0ace6e2aa05040000ffffd318594c1a000000",
	})
	if err != nil {
		panic(err)
	}
	g.DefaultResolver = hgr

	func() {
		b := packr.New("all", "./templates/plugin")
		b.SetResolver("dao/dao.go.tmpl", packr.Pointer{ForwardBox: gk, ForwardPath: "e2180af47102e9758aad571e6e107783"})
		b.SetResolver("go.mod.tmpl", packr.Pointer{ForwardBox: gk, ForwardPath: "9040efce6e173ddbe5aa65b5ee616485"})
		b.SetResolver("http/http.go.tmpl", packr.Pointer{ForwardBox: gk, ForwardPath: "4adc860d2724dbd31a7da30e60cd476c"})
		b.SetResolver("main.go.tmpl", packr.Pointer{ForwardBox: gk, ForwardPath: "8cfce0dcb361a8f7da25bc989bab49b8"})
		b.SetResolver("model/model.go.tmpl", packr.Pointer{ForwardBox: gk, ForwardPath: "901dee514c4b917fa63490110021898e"})
		b.SetResolver("service/service.go.tmpl", packr.Pointer{ForwardBox: gk, ForwardPath: "2576b02ce4cf56084d3ef2b6714559cd"})
	}()

	return nil
}()
