import { CommentItem } from "~/components/CommentItem";
import { Post } from "../components/Post";
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { useAuth } from "~/hooks/useAuth";
import { Button, TextField, Typography } from "@mui/material";
import { PlusIcon, ArrowLeftIcon } from "@phosphor-icons/react";
import { useState, type FormEvent } from "react";
import axios from "axios";
import { useNavigate } from "react-router";
import type { Route } from "./+types/PostCommentsScreen.$id";
import { CommentsTree } from "~/components/CommentsTree";
import { useToast } from "~/components/ToastProvider";

type Comment = {
  comment_id: number;
  user_id: number;
  content: string;
  name: string;
};

type CommentsResponse = { payload?: { data?: Comment[] } };

type PostItem = {
  post_id: number;
  user_id: number;
  title: string;
  content: string;
};

type PostsResponse = { payload?: { data?: PostItem } };

export default function PostComments({ params }: Route.ComponentProps) {
  const navigate = useNavigate();
  const { userID, isLoading: userLoading } = useAuth();
  const queryClient = useQueryClient();
  const [isEditing, setIsEditing] = useState(false);
  const [collapsed, setCollapsed] = useState<Set<number>>(new Set());
  const toast = useToast()

  const postID = Number(params.id);
  if (!Number.isFinite(postID)) {
    return <div>Invalid post id.</div>;
  }

  const toggleCollapsed = (commentId: number) => {
    setCollapsed((prev) => {
      const next = new Set(prev);
      if (next.has(commentId)) {
        next.delete(commentId);
      } else {
        next.add(commentId);
      }
      return next;
    });
  };

  //Fetch all comments for the particular post
  const { isLoading, data } = useQuery<CommentsResponse>({
    queryKey: [`comments`, postID],
    queryFn: async () => {
      const response = await axios.get(
        `http://localhost:8000/posts/${postID}/comments`,
      );
      return await response.data;
    },
  });

  const { isLoading: postLoading, data: postData } = useQuery<PostsResponse>({
    queryKey: [`posts`, postID],
    queryFn: async () => {
      const url = `http://localhost:8000/posts/${postID}`;
      const response = await axios.get(url);

      return response.data;
    },
  });
  const currPost = postData?.payload?.data;

  const createComment = useMutation({
    mutationFn: async (body: { content: string }) =>
      await axios.post(`http://localhost:8000/posts/${postID}/comments`, body, {
        withCredentials: true,
      }),
    onSuccess: (body) => {
      toast("Successfully created comment.", "success");
      setIsEditing(false);
      queryClient.invalidateQueries({ queryKey: [`comments`, postID] });
    },
    onError: () => {
      toast("Failed to create comment.", "error");
    },
  });

  const handleSubmit = (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    const form = e.currentTarget;
    const formData = new FormData(form);
    const content = String(formData.get("comment") || "").trim();
    if (!content) {
      toast("Comment body cannot be empty", "error");
      return;
    }
    createComment.mutate({ content });
  };

  return (
    <div className="flex flex-col gap-6 py-6 px-20">
      <Button
        startIcon={<ArrowLeftIcon />}
        size="large"
        onClick={() => navigate(-1)}
        sx={{ color: "black", width: 90, borderRadius: "50px" }}
      >
        Back
      </Button>
      {postLoading || !currPost ? (
        <div></div>
      ) : (
        <Post
          id={currPost.post_id}
          title={currPost.title}
          content={currPost.content}
          isOwner={Number(currPost.user_id) === userID}
          showChat={false}
        />
      )}
      <hr />
      <div className="flex flex-row items-center gap-2">
        {isEditing ? (
          <form onSubmit={handleSubmit} className="w-full flex flex-col gap-2">
            <TextField
              fullWidth={true}
              multiline={true}
              rows={3}
              required
              name="comment"
            />
            <div>
              <Button
                variant="outlined"
                disableRipple
                sx={{
                  color: "black",
                  borderColor: "black",
                  borderRadius: 20,
                  width: 200,
                  marginRight: 2,
                }}
                type="submit"
              >
                Post Comment
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
              >
                Cancel
              </Button>
            </div>
          </form>
        ) : (
          <Button
            variant="outlined"
            disableRipple
            sx={{ color: "black", borderColor: "black", borderRadius: 20 }}
            onClick={() => {
              userID ? setIsEditing(true) : toast("Login to leave a comment", "error");
            }}
          >
            <PlusIcon />
            <Typography>Add a Comment</Typography>
          </Button>
        )}
      </div>
      {isLoading ? (
        <Typography>Loading comments...</Typography>
      ) : (
        <CommentsTree
          comments={data?.payload?.data}
          postID={postID}
          userID={userID}
          collapsed={collapsed}
          onToggleCollapsed={toggleCollapsed}
        />
      )}
    </div>
  );
}
