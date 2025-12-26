import { Button, Typography, TextField, Box } from "@mui/material";
import { Link } from "react-router";
export default function SignUpPage() {
  return (
    <div className="w-screen h-screen bg-[#F5F5F5] flex justify-center pt-40">
      <div className="flex flex-col w-96 rounded-2xl h-80 bg-white items-center py-8 gap-4">
        <text className="font-bold text-2xl">Sign Up</text>
        <div className="flex flex-col">
          <text className="text-l">username</text>
          <TextField required size="small" />
          <text className="text-l">password</text>
          <TextField required size="small" />
        </div>
        <Button variant="contained" sx={{ background: "#9BE3FF" }}>
          <Typography>Enter!</Typography>
        </Button>
        <Button component={Link}
            to="/sign-in"
        ><Typography>Sign-in instead</Typography></Button>
      </div>
    </div>
  );
}
