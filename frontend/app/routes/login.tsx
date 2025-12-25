import { Button, Typography, TextField, Box } from "@mui/material";

export function LoginPage() {
  return (
    <div className="w-screen h-screen bg-[#F5F5F5] flex justify-center pt-40">
      <div className="flex flex-col w-96 rounded-2xl h-80 bg-white items-center py-8 gap-5">
        <text className="font-bold text-2xl">Sign In</text>
        <div className="flex flex-col">
          <text className="text-l">username</text>
          <TextField required size="small" />
          <text className="text-l">password</text>
          <TextField required size="small" />
        </div>
        <Button variant="contained" sx={{ background: "#9BE3FF" }}>
          <Typography>Enter!</Typography>
        </Button>
      </div>
    </div>
  );
}
