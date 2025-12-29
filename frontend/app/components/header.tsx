import { SideBar } from "./sidebar";
import { useQuery } from "@tanstack/react-query";
import { PlusCircleIcon } from "@phosphor-icons/react";
import { Link } from "react-router";
import { Button, Typography } from "@mui/material";

export function Header() {
  const { isLoading, data } = useQuery({
    queryKey: [`topics`],
    queryFn: getTopics,
  });
  if (isLoading) return <div>Loading...</div>;
  return (
    <div className="h-20 sticky top-0">
      <div className="h-20 bg-[#9BE3FF] ps-4 flex items-center gap-4">
        <SideBar topicsList={data?.payload.data} />
        <Link to="/">
          <text className="text-white font-bold text-5xl">CVWO</text>
        </Link>
        <div className="w-full flex justify-end px-4">
          <Button
            component={Link}
            to="addposts"
            variant="text"
            sx={{ borderRadius: 999 }}
          >
            <PlusCircleIcon size="50" color="white" weight="bold" />
          </Button>
          {/* TODO should show user/login when user is signedin or not */}
          <Button
            component={Link}
            to="/sign-in"
            variant="text"
            sx={{ borderRadius: 999 }}
          >
            <Typography fontWeight={500} color="white" fontSize={28} >Login</Typography>
          </Button>
        </div>
      </div>
    </div>
  );
}

const getTopics = async () => {
  const response = await fetch("http://localhost:8000/topics");
  return await response.json();
};
