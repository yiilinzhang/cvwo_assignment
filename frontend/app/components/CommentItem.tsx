import {
  PencilSimpleLineIcon,
  TrashIcon,
  ArrowBendUpLeftIcon,
  CaretDownIcon,
  CaretRightIcon,
} from "@phosphor-icons/react";
import { Button, IconButton, TextField, Typography } from "@mui/material";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { api } from "~/lib/api";
import { useState, type FormEvent } from "react";
import { useToast } from "~/components/ToastProvider";
import { useAuth } from "~/hooks/useAuth";

type CommentProps = {
  username: string;
  content: string;
  isOwner: boolean;
  commentID: number;
  postID: number;
  hasChildren: boolean;
  isCollapsed: boolean;
  onToggleCollapsed: () => void;
};

type CreateCommentBody = {
  content: string;
  parent_comment_id?: number;
};

//DO NOT call by itself. Use CommentsTree instead.
export function CommentItem({
  username,
  content,
  isOwner,
  commentID,
  postID,
  hasChildren,
  isCollapsed,
  onToggleCollapsed,
}: CommentProps) {
  const { userID, isLoading: userLoading } = useAuth();
  const [isEditing, setIsEditing] = useState(false);
  const [isReplying, setIsReplying] = useState(false);

  const queryClient = useQueryClient();
  const toast = useToast();
  const createComment = useMutation({
    mutationFn: async (body: CreateCommentBody) =>
      await api.post(`/posts/${postID}/comments`, body),

    onSuccess: () => {
      toast("Successfully replied comment.", "success");
      setIsReplying(false);
      queryClient.invalidateQueries({ queryKey: [`comments`, postID] });
    },

    onError: () => {
      toast("Failed to create reply.", "error");
    },
  });
  const replyComment = async (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    const form = e.currentTarget;
    const formData = new FormData(form);
    const content = String(formData.get("reply") || "").trim();
    if (!content) {
      toast("Reply cannot be empty", "error");
      return;
    }
    createComment.mutate({ content: content, parent_comment_id: commentID });
  };

  const updateComment = useMutation({
    mutationFn: async (body: { content: string }) => {
      await api.patch(`/comments/${commentID}`, body);
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["comments", postID] });
      setIsEditing(false);
      toast("Successfully edited comment.", "success");
    },

    onError: () => toast("Failed to edit comment.", "error"),
  });

  const editComment = async (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    const form = e.currentTarget;
    const formData = new FormData(form);
    const content = String(formData.get("comment") || "").trim();
    if (!content) {
      toast("Comment cannot be empty", "error");
      return;
    }
    updateComment.mutate({ content });
  };

  const deleteComment = useMutation({
    mutationFn: async () => {
      await api.delete(`/comments/${commentID}`);
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["comments", postID] });

      toast("Comment sucessfully deleted", "success");
    },
  });
  return (
    <div className="bg-[#D9D9D9] w-full flex flex-col p-4 gap">
      <Typography fontWeight={600} fontSize={17}>
        {username}
      </Typography>
      {isEditing ? (
        <form onSubmit={editComment}>
          <TextField
            name="comment"
            size="medium"
            variant="outlined"
            defaultValue={content}
            required
            multiline
            rows={4}
          />
          <div className="mt-2">
            <Button
              variant="outlined"
              disableRipple
              disabled={updateComment.isPending}
              sx={{
                color: "black",
                borderColor: "black",
                borderRadius: 20,
                width: 200,
                marginRight: 2,
              }}
              type="submit"
            >
              {updateComment.isPending ? "Saving...." : "Save"}
            </Button>
            <Button
              variant="outlined"
              disableRipple
              sx={{
                color: "black",
                borderColor: "black",
                borderRadius: 20,
                width: 200,
              }}
              onClick={() => setIsEditing(false)}
              type="button"
            >
              Cancel
            </Button>
          </div>
        </form>
      ) : (
        <>
          <Typography>{content}</Typography>
          {isReplying && (
            <form onSubmit={replyComment}>
              <TextField
                name="reply"
                size="medium"
                variant="outlined"
                required
                multiline
                rows={4}
              />
              <div>
                <Button
                  variant="outlined"
                  disableRipple
                  disabled={createComment.isPending}
                  sx={{
                    color: "black",
                    borderColor: "black",
                    borderRadius: 20,
                    width: 200,
                    marginRight: 2,
                  }}
                  type="submit"
                >
                  {createComment.isPending ? "Replying...." : "Reply"}
                </Button>
                <Button
                  variant="outlined"
                  disableRipple
                  sx={{
                    color: "black",
                    borderColor: "black",
                    borderRadius: 20,
                    width: 200,
                  }}
                  type="button"
                  onClick={() => setIsReplying(false)}
                >
                  Cancel
                </Button>
              </div>
            </form>
          )}
          <div
            className={`flex ${hasChildren ? "justify-between" : "justify-end"} ${isReplying ? "hidden" : ""}`}
          >
            {hasChildren && (
              <IconButton onClick={onToggleCollapsed}>
                {isCollapsed ? (
                  <CaretDownIcon color="black" />
                ) : (
                  <CaretRightIcon color="black" />
                )}
              </IconButton>
            )}
            <div className="flex flex-row">
              <IconButton
                aria-label="comment"
                onClick={() => {
                  userID
                    ? setIsEditing(true)
                    : toast("Login to leave a comment", "error");
                }}
              >
                <ArrowBendUpLeftIcon size={27} color="black" />
              </IconButton>

              {isOwner && !isReplying && (
                <div>
                  <IconButton
                    aria-label="edit_comment"
                    onClick={() => setIsEditing(true)}
                  >
                    <PencilSimpleLineIcon size={27} color="black" />
                  </IconButton>

                  <IconButton
                    aria-label="delete_comment"
                    onClick={() => deleteComment.mutate()}
                  >
                    <TrashIcon size={27} color="black" />
                  </IconButton>
                </div>
              )}
            </div>
          </div>
        </>
      )}
    </div>
  );
}
