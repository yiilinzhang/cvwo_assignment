import { Button, Typography, TextField } from "@mui/material";
import { useMutation } from "@tanstack/react-query";
import axios from "axios";
import type { FormEvent } from "react";
import { Link } from "react-router";
import { useNavigate } from "react-router";

type SignUpBody = {
  username: string;
  password: string;
}

export default function SignUpPage() {
  const navigate = useNavigate();

  //TODO update type
  const signup = useMutation({
    mutationFn: async (body: SignUpBody) => {
      await axios.post("http://localhost:8000/users", body)
    },
    onSuccess: () => {//TODO add code to save jwt
    alert("Successfully created account.");
    navigate("/");},
    onError: () => {
      alert("Failed to create account.");
    }
  }) 

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
    signup.mutate(body);
  };

  return (
    <form onSubmit={handleSubmit}>
      <div className="w-screen h-screen bg-[#F5F5F5] flex justify-center pt-40">
        <div className="flex flex-col w-96 rounded-2xl h-80 bg-white items-center py-8 gap-4">
          <text className="font-bold text-2xl">Sign Up</text>
          <div className="flex flex-col">
            <text className="text-l">username</text>
            <TextField required size="small" name="username" />
            <text className="text-l">password</text>
            <TextField required size="small" name="password" type="password" />
          </div>
          <Button
            variant="contained"
            sx={{ background: "#9BE3FF" }}
            type="submit"
            disabled={signup.isPending}
          >
            <Typography>
              {signup.isPending ? "Loading" : "Enter!"}</Typography>
          </Button>
          <Button component={Link} to="/sign-in"

          >
            <Typography>Sign-in instead</Typography>
          </Button>
        </div>
      </div>
    </form>
  );
}
