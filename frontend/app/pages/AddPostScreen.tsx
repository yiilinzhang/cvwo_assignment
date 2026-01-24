import { useMutation, useQuery } from "@tanstack/react-query";
import { useNavigate } from "react-router";
import { Button, Typography, TextField, MenuItem } from "@mui/material";
import { api } from "~/lib/api";
import { useToast } from "~/components/ToastProvider";

type Topic = { topic_id: number; title: string };
type TopicsResponse = { payload?: { data?: Topic[] } };

type CreatePostBody = {
  title: string;
  content: string;
  topic_id: number;
};

//Page where the users can make a new post
export default function AddPosts() {
  const toast = useToast()
  const { isLoading, data } = useQuery<TopicsResponse>({
    queryKey: [`topiclist`],
    queryFn: async () => {
      const response = await api.get("/topics");
      return response.data;
    },
  });
  const navigate = useNavigate();

  const createPost = useMutation({
    mutationFn: async (body: CreatePostBody) =>
      await api.post("/posts", body),
    onSuccess: () => {
      toast("Successfully created post.", "success");
      navigate("/");
    },
    onError: () => {
      toast("Failed to create post.", "error");
    },
  });

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    const form = e.currentTarget;
    const formData = new FormData(form);
    const title = String(formData.get("title") || "").trim();
    const content = String(formData.get("content") || "").trim();
    const topicID = Number(formData.get("topic_id"));
    if (!content || !title || !Number.isFinite(topicID)) return;
    const body: CreatePostBody = {
      title: title,
      content: content,
      topic_id: topicID,
    };
    createPost.mutate(body);
  };

  return (
    <form onSubmit={handleSubmit}>
      <div className="flex flex-col gap-1 items-center py-12">
        <Typography color="black" fontWeight={500} fontSize={22}>
          Create a new post today!
        </Typography>

        <div className="flex flex-col justify-start gap-2">
          <Typography fontWeight={500} fontSize={18}>
            Topic
          </Typography>

          <TextField
            select
            name="topic_id"
            size="small"
            sx={{ width: 160 }}
            required
            disabled={isLoading}
          >
            {isLoading ? (
              <MenuItem>Loading topics....</MenuItem>
            ) : (
              data?.payload?.data?.map((title) => (
                <MenuItem key={title.topic_id} value={title.topic_id}>
                  {title.title}
                </MenuItem>
              ))
            )}
          </TextField>

          <Typography fontWeight={500} fontSize={18}>
            Title
          </Typography>
          <TextField
            slotProps={{
              input: {
                inputProps: { maxLength: 50 },
              },
            }}
            name="title"
            size="small"
            variant="outlined"
            required
            sx={{ width: 300 }}
          />

          <Typography fontWeight={500} fontSize={18}>
            Content
          </Typography>
          <TextField
            name="content"
            size="medium"
            variant="outlined"
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
          disabled={createPost.isPending}
        >
          <Typography fontSize={18}>
            {createPost.isPending ? "Posting..." : "Post now!"}
          </Typography>
        </Button>
      </div>
    </form>
  );
}
