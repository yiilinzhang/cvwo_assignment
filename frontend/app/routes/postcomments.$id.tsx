import { Comments } from "~/components/comments";
import { Post } from "../components/post";
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { useAuth } from "~/hooks/useAuth";
import { Button, TextField, Typography } from "@mui/material";
import { PlusIcon, ArrowLeftIcon } from "@phosphor-icons/react";
import { useState, type FormEvent } from "react";
import axios from "axios";
import { useNavigate } from "react-router";
import type { Route } from "./+types/postcomments.$id";

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

type PostsResponse = { payload?: { data?: PostItem[] } };
//Should accept the post details so i can abstract the post id and selet comments
//TODO chaneg this from cache to a new query
export default function PostComments({ params }: Route.ComponentProps) {
  const navigate = useNavigate();
  const { user, isLoading: userLoading } = useAuth();
  const queryClient = useQueryClient();
  const [isEditing, setIsEditing] = useState(false);
  const postId = Number(params.id);
  if (!Number.isFinite(postId)) {
    return <div>Invalid post id.</div>;
  }
  //Fetch all comments for the particular post
  const { isLoading, data } = useQuery<CommentsResponse>({
    queryKey: [`comments`, params.id],
    queryFn: async () => {
      const response = await axios.get(
        `http://localhost:8000/posts/${postId}/comments`
      );
      return await response.data;
    },
  });
  const cachedPosts = queryClient.getQueryData<PostsResponse>(["posts", "all"]);
  const { isLoading: postLoading, data: postData } = useQuery<PostsResponse>({
    queryKey: [`posts`, postId],
    queryFn: async () => {
      const url = "http://localhost:8000/posts";
      const response = await axios.get(url);

      return response.data;
    },
    initialData: cachedPosts,
  });

  const currPost = postData?.payload?.data?.find((p) => p.post_id === postId);

  const createComment = useMutation({
    mutationFn: async (body:{content : String}) =>
      await axios.post(`http://localhost:8000/posts/${postId}/comments`, body, {
        withCredentials: true,
      }),
    onSuccess: (body ) => {
      alert("Successfully created comment.");
      setIsEditing(false);
      queryClient.invalidateQueries({ queryKey: [`comments`, postId] });
    },
    onError:()=> {
      alert("Failed to create comment.")
    }
  });
  const handleSubmit = (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    const form = e.currentTarget;
    const formData = new FormData(form);
    const content = String(formData.get("comment") || "").trim();
    if (!content) {
      alert("Comment body cannot be empty");
      return;
    }
    createComment.mutate({content})
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
          isOwner={Number(currPost.user_id) === user?.payload.data}
          showChat={false}
        />
      )}
      <hr />
      {/* TODO check how i should pass in post prop requery or drill */}
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
              user ? setIsEditing(true) : alert("login to leave a comment");
            }}
          >
            <PlusIcon />
            <Typography>Add a Comment</Typography>
          </Button>
        )}
      </div>
      {data?.payload?.data?.map((comment) => (
        <Comments
          key={comment.comment_id}
          username={comment.name}
          content={comment.content}
          isOwner={Number(comment.user_id) === user?.payload.data}
          id={comment.comment_id}
          postId={postId}
        />
      ))}
    </div>
  );
}
