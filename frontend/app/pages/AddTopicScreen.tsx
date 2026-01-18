import { useMutation, useQueryClient } from "@tanstack/react-query";
import { useNavigate } from "react-router";
import { Button, Typography, TextField } from "@mui/material";
import { type FormEvent } from "react";
import axios from "axios";


//TODO add typing later
//TODO use MUI alert for a prettier alert
export default function AddTopics() {
  const navigate = useNavigate();
  const queryClient = useQueryClient();
  const addTopics = useMutation({
    mutationFn: async (body: {title: string}) =>
      await axios.post("http://localhost:8000/topics", body, {
        withCredentials: true,
      }),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["topics"] });
      alert("Successfully created topic.");
      navigate("/");
    },
    onError: () => alert("Failed to create topic."),
  });

  const handleSubmit = async (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    const form = e.currentTarget;
    const formData = new FormData(form);
    const title = String(formData.get("title") || "").trim()
    if (!title) {
      alert("Title cannot be empty");
      return;
    }
    addTopics.mutate({title})
  };

  return (
    <form onSubmit={handleSubmit}>
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
          disabled={addTopics.isPending}
        >
          <Typography sx={{ fontSize: "20px" }}>{addTopics.isPending ? "Adding" : "Add now!"}</Typography>
        </Button>
      </div>
    </form>
  );
}
