import { useQuery } from "@tanstack/react-query";
import { useNavigate } from "react-router";
import { Button, Typography, TextField, MenuItem } from "@mui/material";
//TODO add typing later
export default function addPosts() {
  const { isLoading, data } = useQuery({
    queryKey: [`topiclist`],
    queryFn: async () => {
      const response = await fetch("http://localhost:8000/topics");
      return await response.json();
    },
  });
  const navigate = useNavigate();

  const handleSubmit = async (e) => {
    e.preventDefault();

    const form = e.target;
    const formData = new FormData(form);
    const body = {
      title: formData.get("title"),
      content: formData.get("content"),
      topic_id: Number(formData.get("topic_id")),
    };
    console.log(body);
    try {
      const response = await fetch("http://localhost:8000/posts", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(body),
      });
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      console.log(response);
      const result = await response.json();
    } catch (error) {
      console.error("Error:", error);
      alert("Failed to create post.");
      return;
    }
    alert("Successfully created post.");
    navigate("/");
  };

  return (
    <form method="post" onSubmit={handleSubmit}>
      <div className="flex flex-col gap-1 items-center py-2">
          <Typography color="black" fontWeight={500} fontSize="2rem">Create a new post today!</Typography>
          
        <div className="flex flex-col justify-start gap-2">
          <Typography fontWeight={500} fontSize="1.5rem">
            Topic
          </Typography>

          <TextField select name="topic_id" size="small" sx={{ width: 160 }}>
            {data?.payload.data.map((title) => (
              <MenuItem key={title.topic_id} value={title.topic_id}>
                {title.title}
              </MenuItem>
            ))}
          </TextField>

          <text className="font-medium text-2xl ">Title</text>
          <TextField
            name="title"
            size="small"
            variant="outlined"
            required
            sx={{ width: 300 }}
          />

          <text className="font-medium text-2xl">Content</text>
          <TextField
            name="title"
            size="medium"
            variant="outlined"
            required
            multiline
            rows={4}
            sx={{ width: 300 }}
          ></TextField>
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
