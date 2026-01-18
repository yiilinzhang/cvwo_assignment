import { useQuery, useQueryClient, useMutation } from "@tanstack/react-query";
import { useNavigate } from "react-router";
import { Button, Typography, TextField } from "@mui/material";
import axios from "axios";
import type { Route } from "./+types/EditPostScreen";
import type { FormEvent } from "react";

type PostItem = {
  post_id: number;
  user_id: number;
  title: string;
  content: string;
  topic_id: number;
}

type PostResponse = { payload?: {data?: PostItem[]}}
//TODO add typing later
//TODO use MUI alert for a prettier alert
export default function EditPosts({ params }: Route.ComponentProps) {
  const navigate = useNavigate();
  const queryClient = useQueryClient();
  const cachedPosts = queryClient.getQueryData<PostResponse>(["posts", "all"]);
  const postId = params?.id;

  const { isLoading, data: postData } = useQuery<PostResponse>({
    queryKey: [`posts`, "all"],
    queryFn: async () => {
      const response = await axios.get("http://localhost:8000/posts")
      return response.data
    },
    initialData: cachedPosts,
  });

  const currPost = postData?.payload?.data?.find(
    (p) => p.post_id === Number(postId)
  );

  const handleSubmit = async (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    const form = e.currentTarget;
    const formData = new FormData(form);
    const content = String(formData.get("content") || "").trim()
    updatePost.mutate({content});
  };
  const updatePost = useMutation({
    mutationFn: async (body: {content: string}) => {
      await axios.patch(`http://localhost:8000/posts/${postId}`, body, {
        withCredentials: true,
      });
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["posts"] });
      alert("Successfully edited post.");
      navigate(-1);
    },
    onError: () => alert("Failed to edit post."),
  });

  return (
    <form onSubmit={handleSubmit}>
      {isLoading || !currPost ? (
        <div>Loading...</div>
      ) : (
        <div className="flex flex-col gap-1 items-center py-12">
          <Typography color="black" fontWeight={500} fontSize="2rem">
            Edit Post
          </Typography>

          <div className="flex flex-col justify-start gap-2">
            <Typography fontWeight={500} fontSize="1.5rem">
              Topic
            </Typography>

            <TextField
              size="small"
              variant="outlined"
              defaultValue={currPost.topic_id}
              slotProps={{
                input: {
                  readOnly: true,
                },
              }}
              helperText="Topics cannot be updated."
            />

            <Typography fontWeight={500} fontSize="1.5rem">
              Title
            </Typography>
            <TextField
              size="small"
              variant="outlined"
              sx={{ width: 300 }}
              defaultValue={currPost.title}
              slotProps={{
                input: {
                  readOnly: true,
                },
              }}
              helperText="Titles cannot be updated."
            />

            <Typography fontWeight={500} fontSize="1.5rem">
              Content
            </Typography>
            <TextField
              name="content"
              size="medium"
              variant="outlined"
              defaultValue={currPost.content}
              required
              multiline
              rows={4}
              sx={{ width: 300 }}
            />
          </div>

          <Button
            type="submit"
            variant="contained"
            sx={{ background: "#9BE3FF", mt: 2 }}
            disabled={updatePost.isPending}
          >
            <Typography sx={{ fontSize: "20px" }}>{updatePost.isPending ? "Saving..." : "Save"}</Typography>
          </Button>
        </div>
      )}
    </form>
  );
}
