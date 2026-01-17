import { useQueryClient } from "@tanstack/react-query";
import { useNavigate } from "react-router";
import { Button, Typography, TextField } from "@mui/material";
//TODO add typing later
//TODO use MUI alert for a prettier alert
export default function AddTopics() {
  const navigate = useNavigate();
  const queryClient = useQueryClient();
  const handleSubmit = async (e) => {
    e.preventDefault();

    const form = e.target;
    const formData = new FormData(form);
    const body = {
      title: formData.get("title"),
    };
    try {
      //TODO change to axios
      const response = await fetch("http://localhost:8000/topics", {
        method: "POST",
        credentials: "include",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(body),
      });
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      const result = await response.json();
    } catch (error) {
      console.error("Error:", error);
      alert("Failed to create topic.");
      return;
    }
    queryClient.invalidateQueries({ queryKey: ["topics"] });
    alert("Successfully created topic.");
    navigate("/");
  };

  return (
    <form method="post" onSubmit={handleSubmit}>
      <div className="flex flex-col gap-12 items-center py-12 ">
        <Typography color="black" fontWeight={500} fontSize="2rem">
          Create a new Topic today!
        </Typography>

        <div className="flex flex-col justify-start gap-2">
          <Typography fontWeight={500} fontSize="1.5rem">
            Title
          </Typography>
          <TextField
            name="title"
            size="small"
            variant="outlined"
            required
            sx={{ width: 300 }}
          />
        </div>

        <Button
          type="submit"
          variant="contained"
          sx={{ background: "#9BE3FF", mt: 2 }}
        >
          <Typography sx={{ fontSize: "20px" }}>Add now!</Typography>
        </Button>
      </div>
    </form>
  );
}
