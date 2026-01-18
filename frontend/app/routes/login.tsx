import { Button, Typography, TextField } from "@mui/material";
import { useNavigate } from "react-router";
import { Link } from "react-router";
import { type FormEvent } from "react";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import axios from "axios";

export default function LoginPage() {
  const navigate = useNavigate();
  const queryClient = useQueryClient();

  //TODO update type from any
  const login = useMutation({
    mutationFn: async (body: any) => {
      await axios.post("http://localhost:8000/login", body, {
        withCredentials: true,
      });
    },
    onSuccess: () => {
      alert("Successfully logged in.");
      navigate("/");
    },
    onError: () => {
      alert("Failed to login account.");
    },
  });

  const handleSubmit = async (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    const form = e.currentTarget;
    const formData = new FormData(form);
    const username = String(formData.get("username") || "").trim()
    const password = String(formData.get("password") || "").trim()
    if (!username || !password) {
      alert("Username and password cannot be empty");
      return;
    }
    const body = {
      username: username,
      password: password,
    };
    login.mutate(body);

    //TODO add code to save jwt
  };
  return (
    <form onSubmit={handleSubmit}>
      <div className="w-screen h-screen bg-[#F5F5F5] flex justify-center pt-40">
        <div className="flex flex-col w-96 rounded-2xl h-80 bg-white items-center py-8 gap-3">
          <Typography >Sign In</Typography>
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
