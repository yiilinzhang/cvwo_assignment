package comments

// Function to turn comment data into a nested structure. For front end processing
func MakeCommentTree(rows []QueryCommentResponse) []*CommentParent {
	nodes := make(map[int]*CommentParent)

	for _, r := range rows {
		nodes[r.ID] = &CommentParent{
			ID:       r.ID,
			Content:  r.Content,
			UserID:   r.UserID,
			UserName: r.UserName,
			ParentID: r.ParentID,
			Children: nil,
		}
	}

	res := make([]*CommentParent, 0)

	for _, r := range rows {
		n := nodes[r.ID]
		if n.ParentID == nil {
			res = append(res, n)
			continue
		}
		parent := nodes[*n.ParentID]
		parent.Children = append(parent.Children, n)
	}
	return res

}