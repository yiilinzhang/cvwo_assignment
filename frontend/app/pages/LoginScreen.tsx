import { Button, Typography, TextField } from "@mui/material";
import { useNavigate } from "react-router";
import { Link } from "react-router";
import { type FormEvent } from "react";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import axios from "axios";
import { useToast } from "~/components/ToastProvider";

type LogInBody = {
  username: string;
  password: string;
}

export default function LoginPage() {
  const navigate = useNavigate();
  const queryClient = useQueryClient();
  const toast = useToast()
  
  const login = useMutation({
    mutationFn: async (body: LogInBody) => {
      await axios.post("http://localhost:8000/login", body, {
        withCredentials: true,
      });
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["me"] });
      toast("Successfully logged in.", "success");
      navigate("/");
    },
    onError: () => {
      toast("Failed to login account.", "error");
    },
  });

  const handleSubmit = async (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    const form = e.currentTarget;
    const formData = new FormData(form);
    const username = String(formData.get("username") || "").trim()
    const password = String(formData.get("password") || "").trim()
    if (!username || !password) {
      toast("Username and password cannot be empty", "error");
      return;
    }
    const body = {
      username: username,
      password: password,
    };
    login.mutate(body);

  };
  return (
    <form onSubmit={handleSubmit}>
      <div className="w-screen h-screen bg-[#F5F5F5] flex justify-center pt-40">
        <div className="flex flex-col w-96 rounded-2xl h-80 bg-white items-center py-8 gap-3">
          <Typography fontWeight={500} fontSize={20}>Sign In</Typography>
          <div className="flex flex-col">
            <Typography >username</Typography>
            <TextField required size="small" name="username" />
            <Typography >password</Typography>
            <TextField required size="small" name="password" type="password" />
          </div>
          <Button
            variant="contained"
            sx={{ background: "#9BE3FF" }}
            type="submit"
            disabled={login.isPending}
          >
            <Typography>{login.isPending ? "Loading" : "Enter!"}</Typography>
          </Button>
          <Button component={Link} to="/sign-up">
            <Typography>Sign-up instead</Typography>
          </Button>
        </div>
      </div>
    </form>
  );
}
