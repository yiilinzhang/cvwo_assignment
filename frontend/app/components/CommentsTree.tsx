import { CommentItem } from "./CommentItem";

type CommentNode = {
  comment_id: number;
  user_id: number;
  name: string;
  content: string;
  children?: CommentNode[] | null;
};

type CommentsTreeProps = {
  comments?: CommentNode[] | null;
  userID: number | null;
  postID: number;
  collapsed: Set<number>;
  onToggleCollapsed: (commentId: number) => void;
};

//Recursively renders children comments
export function CommentsTree({
  comments,
  userID,
  postID,
  collapsed,
  onToggleCollapsed,
}: CommentsTreeProps) {
  if (!comments || comments.length === 0) return null;

  return (
    <div className="flex flex-col gap-4">
      {comments.map((comment) => {
        const isCollapsed = collapsed.has(comment.comment_id);
        const hasChildren = !!comment.children?.length;
        return (
          <div key={comment.comment_id} className="flex flex-col gap-4">
            <CommentItem
              username={comment.name}
              content={comment.content}
              isOwner={userID !== null && Number(comment.user_id) === userID}
              commentID={comment.comment_id}
              postID={postID}
              hasChildren={hasChildren}
              isCollapsed={isCollapsed}
              onToggleCollapsed={() => onToggleCollapsed(comment.comment_id)}
            />
            {comment.children?.length && !isCollapsed ? (
              <div className="ml-6 border-l-2 border-black/20 pl-4">
                <CommentsTree
                  comments={comment.children}
                  userID={userID}
                  postID={postID}
                  collapsed={collapsed}
                  onToggleCollapsed={onToggleCollapsed}
                />
              </div>
            ) : null}
          </div>
        );
      })}
    </div>
  );
}
