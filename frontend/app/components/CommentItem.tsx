import {
  PencilSimpleLineIcon,
  TrashIcon,
  ArrowBendUpLeftIcon,
  CaretDownIcon,
  CaretRightIcon
} from "@phosphor-icons/react";
import { Button, IconButton, TextField, Typography } from "@mui/material";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import axios from "axios";
import { useState, type FormEvent } from "react";

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
  const [isEditing, setIsEditing] = useState(false);
  const [isReplying, setIsReplying] = useState(false);

  const queryClient = useQueryClient();

  const createComment = useMutation({
    mutationFn: async (body: CreateCommentBody) =>
      await axios.post(`http://localhost:8000/posts/${postID}/comments`, body, {
        withCredentials: true,
      }),

    onSuccess: () => {
      alert("Successfully replied comment.");
      setIsReplying(false);
      queryClient.invalidateQueries({ queryKey: [`comments`, postID] });
    },

    onError: () => {
      alert("Failed to create reply.");
    },
  });
  const replyComment = async (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    const form = e.currentTarget;
    const formData = new FormData(form);
    const content = String(formData.get("reply") || "").trim();
    if (!content) {
      alert("Reply cannot be empty");
      return;
    }
    createComment.mutate({ content: content, parent_comment_id: commentID });
  };

  const updateComment = useMutation({
    mutationFn: async (body: { content: string }) => {
      await axios.patch(`http://localhost:8000/comments/${commentID}`, body, {
        withCredentials: true,
      });
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["comments", postID] });
      setIsEditing(false);
      alert("Successfully edited comment.");
    },

    onError: () => alert("Failed to edit comment."),
  });

  const editComment = async (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    const form = e.currentTarget;
    const formData = new FormData(form);
    const content = String(formData.get("comment") || "").trim();
    if (!content) {
      alert("Comment cannot be empty");
      return;
    }
    updateComment.mutate({ content });
  };

  const deleteComment = useMutation({
    mutationFn: async () => {
      await axios.delete(`http://localhost:8000/comments/${commentID}`, {
        withCredentials: true,
      });
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["comments", postID] });

      alert("Comment sucessfully deleted");
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
          <div className={`flex ${hasChildren? "justify-between": "justify-end"} ${isReplying ? "hidden" : ""}`}>
            {hasChildren && (
            <IconButton onClick={onToggleCollapsed}>
              {isCollapsed? <CaretDownIcon color="black"/>: <CaretRightIcon color="black"/>}
            </IconButton>
          )}
          <div className="flex flex-row">
            <IconButton
              aria-label="comment"
              onClick={() => setIsReplying(true)}
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
            )}</div>
          </div>
        </>
      )}
    </div>
  );
}
