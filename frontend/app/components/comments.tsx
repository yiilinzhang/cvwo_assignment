import {
  PencilSimpleLineIcon,
  TrashIcon,
  ArrowBendUpLeftIcon,
} from "@phosphor-icons/react";
import { Button, IconButton, TextField, Typography } from "@mui/material";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import axios from "axios";
import { useState, type FormEvent } from "react";

type CommentProps = {
  username: string;
  content: string;
  isOwner: boolean;
  id: number;
  postId: number;
};

export function Comments({ username, content, isOwner, id, postId }: CommentProps) {
  const [isEditing, setIsEditing] = useState(false);
  const queryClient = useQueryClient();
  //TODO fix type later
  const updateComment = useMutation({
    mutationFn: async (body: { content: string }) => {
      await axios.patch(`http://localhost:8000/comments/${id}`, body, {
        withCredentials: true,
      });
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["comments", postId] });
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
    updateComment.mutate({ content });
  };

  const deleteComment = useMutation({
    mutationFn: async () => {
      await axios.delete(`http://localhost:8000/comments/${id}`, {
        withCredentials: true,
      });
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["comments"] });

      alert("Post sucessfully deleted");
    },
  });
  return (
    <div className="bg-[#D9D9D9] w-full flex flex-col p-4 gap">
      <Typography >{username}</Typography>
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
          <Typography >{content}</Typography>
          <div className="flex justify-end">
            <IconButton aria-label="comment">
              <ArrowBendUpLeftIcon size={27} color="black" />
            </IconButton>

            {isOwner ? (
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
            ) : (
              <></>
            )}
          </div>
        </>
      )}
    </div>
  );
}
