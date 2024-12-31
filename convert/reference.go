package convert

import (
	"strings"

	"github.com/go-git/go-git/v5/plumbing"
)

type ReferenceJSON struct {
	ShortName string `json:"shortName"`
	Name      string `json:"name"`
	Hash      string `json:"hash"`
	Type      string `json:"type"`
	Category  string `json:"category"` // 具体类型 (branch, tag, remote, etc.)
}

func ConvertToReferenceJSON(ref *plumbing.Reference) ReferenceJSON {
	// 提取 Reference 的字段
	refJSON := ReferenceJSON{
		ShortName: ref.Name().Short(),
		Name:      ref.Name().String(),
		Hash:      ref.Hash().String(),
		Type:      ref.Type().String(),
		Category:  getReferenceCategory(ref),
	}

	return refJSON
}

func getReferenceCategory(ref *plumbing.Reference) string {
	if ref.Name().IsBranch() {
		return "branch"
	} else if ref.Name().IsTag() {
		return "tag"
	} else if ref.Name().IsRemote() {
		return "remote"
	} else if strings.HasPrefix(ref.Name().String(), "refs/stash") {
		return "stash"
	} else {
		return "other"
	}
}
