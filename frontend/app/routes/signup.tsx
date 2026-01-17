import { Button, Typography, TextField, Box } from "@mui/material";
import { Link } from "react-router";
import { useQuery } from "@tanstack/react-query";
import { useNavigate } from "react-router";
export default function SignUpPage() {
  const navigate = useNavigate();

  const handleSubmit = async (e) => {
    e.preventDefault();

    const form = e.target;
    const formData = new FormData(form);
    const body = {
      username: formData.get("username"),
      password: formData.get("password"),
    };
    try {
      const response = await fetch("http://localhost:8000/users", {
        method: "POST",
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
      alert("Failed to create account.");
      return;
    }
    //TODO add code to save jwt
    alert("Successfully created account.");
    navigate("/");
  };
  return (
    <form method="post" onSubmit={handleSubmit}>
      <div className="w-screen h-screen bg-[#F5F5F5] flex justify-center pt-40">
        <div className="flex flex-col w-96 rounded-2xl h-80 bg-white items-center py-8 gap-4">
          <text className="font-bold text-2xl">Sign Up</text>
          <div className="flex flex-col">
            <text className="text-l">username</text>
            <TextField required size="small" name="username" />
            <text className="text-l">
              password
            </text>
            <TextField required size="small" name="password" type="password"/>
          </div>
          <Button variant="contained" sx={{ background: "#9BE3FF" }} type="submit">
            <Typography>Enter!</Typography>
          </Button>
          <Button component={Link} to="/sign-in">
            <Typography>Sign-in instead</Typography>
          </Button>
        </div>
      </div>
    </form>
  );
}
