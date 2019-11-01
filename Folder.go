package msgraph

import (
	"fmt"
)

// User represents a user from the ms graph API
type Folder struct {
	ID                			string   `json:"id"`
	DisplayName       			string   `json:"displayName"`
	ParentFolderId         	string   `json:"parentFolderId"`
	ChildFolderCount        string   `json:"childFolderCount"`
	UnreadItemCount       	string   `json:"unreadItemCount"`
	TotalItemCount 					string   `json:"totalItemCount"`


	//activePhone string       // private cache for the active phone number
	graphClient *GraphClient // the graphClient that called the user
}

func (u *Folder) String() string {
	return fmt.Sprintf("Folder(ID: \"%v\", DisplayName: \"%v\", ParentFolderId: \"%v\", ChildFolderCount: \"%v\", "+
		"UnreadItemCount: \"%v\", TotalItemCount: \"%v\")",
		u.ID, u.DisplayName, u.ParentFolderId, u.ChildFolderCount, u.UnreadItemCount, u.TotalItemCount)
}

// setGraphClient sets the graphClient instance in this instance and all child-instances (if any)
func (u *Folder) setGraphClient(gC *GraphClient) {
	u.graphClient = gC
}
