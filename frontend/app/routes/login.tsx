import { Button, Typography, TextField } from "@mui/material";
import { useNavigate } from "react-router";
import { Link } from "react-router";
import { type FormEvent } from "react";

export default function LoginPage() {
  const navigate = useNavigate();

  const handleSubmit = async (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    const form = e.currentTarget;
    const formData = new FormData(form);
    const body = {
      username: formData.get("username"),
      password: formData.get("password"),
    };
    try {
      const response = await fetch("http://localhost:8000/login", {
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
      alert("Failed to login account.");
      return;
    }
    //TODO add code to save jwt
    alert("Successfully logged in.");
    navigate("/");
  };
  return (
    <form method="post" onSubmit={handleSubmit}>
      <div className="w-screen h-screen bg-[#F5F5F5] flex justify-center pt-40">
        <div className="flex flex-col w-96 rounded-2xl h-80 bg-white items-center py-8 gap-3">
          <text className="font-bold text-2xl">Sign In</text>
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
          >
            <Typography>Enter!</Typography>
          </Button>
          <Button component={Link} to="/sign-up">
            <Typography>Sign-up instead</Typography>
          </Button>
        </div>
      </div>
    </form>
  );
}
