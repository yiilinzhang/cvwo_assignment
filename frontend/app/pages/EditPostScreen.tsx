import { useQuery, useQueryClient, useMutation } from "@tanstack/react-query";
import { useNavigate } from "react-router";
import { Button, Typography, TextField } from "@mui/material";
import axios from "axios";
import type { Route } from "./+types/EditPostScreen";
import type { FormEvent } from "react";
import { useToast } from "~/components/ToastProvider";

type PostItem = {
  post_id: number;
  user_id: number;
  title: string;
  content: string;
  topic_id: number;
}

type PostResponse = { payload?: {data?: PostItem[]}}
type Topic = { topic_id: number; title: string };
type TopicsResponse = { payload?: { data?: Topic[] } };
export default function EditPosts({ params }: Route.ComponentProps) {
  const navigate = useNavigate();
  const queryClient = useQueryClient();
  const cachedPosts = queryClient.getQueryData<PostResponse>(["posts", "all"]);
  const cachedTopics = queryClient.getQueryData<TopicsResponse>(["topics"]);
  const postID = params?.id;
  const toast = useToast()

  const { isLoading, data: postData } = useQuery<PostResponse>({
    queryKey: [`posts`, "all"],
    queryFn: async () => {
      const response = await axios.get("http://localhost:8000/posts")
      return response.data
    },
    initialData: cachedPosts,
  });

  const { data: topicsData } = useQuery<TopicsResponse>({
    queryKey: ["topics"],
    queryFn: async () => {
      const response = await axios.get("http://localhost:8000/topics");
      return response.data;
    },
    initialData: cachedTopics,
  });

  const currPost = postData?.payload?.data?.find(
    (p) => p.post_id === Number(postID)
  );
  const currTopicTitle = topicsData?.payload?.data?.find(
    (t) => t.topic_id === currPost?.topic_id
  )?.title;

  const handleSubmit = async (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    const form = e.currentTarget;
    const formData = new FormData(form);
    const content = String(formData.get("content") || "").trim()
    updatePost.mutate({content});
  };
  const updatePost = useMutation({
    mutationFn: async (body: {content: string}) => {
      await axios.patch(`http://localhost:8000/posts/${postID}`, body, {
        withCredentials: true,
      });
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["posts"] });
      toast("Successfully edited post.", "success");
      navigate(-1);
    },
    onError: () => toast("Failed to edit post.", "error"),
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
              value={currTopicTitle ?? String(currPost.topic_id)}
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
