import { useQuery } from "@tanstack/react-query";
import { useNavigate } from "react-router";
import { Button, Typography, TextField, MenuItem } from "@mui/material";
import axios from "axios";
type Topic = { topic_id: number; title: string };
type TopicsResponse = { payload?: { data?: Topic[] } };

//Add post page
export default function AddPosts() {
  //TODO use this isloading
  const { isLoading, data } = useQuery<TopicsResponse>({
    queryKey: [`topiclist`],
    queryFn: async () => {
      const response = await fetch("http://localhost:8000/topics");
      return await response.json();
    },
  });
  const navigate = useNavigate();

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    const form = e.currentTarget;
    const formData = new FormData(form);
    const title = String(formData.get("title") || "").trim();
    if (!title) return;
    const content = String(formData.get("content") || "").trim();
    if (!content) return;
    const topicId = Number(formData.get("topic_id"));
    if (!Number.isFinite(topicId)) {
    }
    const body = {
      title: formData.get("title"),
      content: formData.get("content"),
      topic_id: Number(formData.get("topic_id")),
    };
    try {
      await axios.post("http://localhost:8000/posts", body, {
        withCredentials: true,
      });
    } catch (error) {
      console.error("Error:", error);
      alert("Failed to create post.");
      return;
    }
    alert("Successfully created post.");
    navigate("/");
  };

  return (
    <form onSubmit={handleSubmit}>
      <div className="flex flex-col gap-1 items-center py-12">
        <Typography color="black" fontWeight={500} fontSize="2rem">
          Create a new post today!
        </Typography>

        <div className="flex flex-col justify-start gap-2">
          <Typography fontWeight={500} fontSize="1.5rem">
            Topic
          </Typography>

          <TextField
            select
            name="topic_id"
            size="small"
            sx={{ width: 160 }}
            required
          >
            {data?.payload?.data?.map((title) => (
              <MenuItem key={title.topic_id} value={title.topic_id}>
                {title.title}
              </MenuItem>
            ))}
          </TextField>

          <Typography fontWeight={500} fontSize="1.5rem">
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

          <Typography fontWeight={500} fontSize="1.5rem">
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
        >
          <Typography sx={{ fontSize: "20px" }}>Post now!</Typography>
        </Button>
      </div>
    </form>
  );
}
