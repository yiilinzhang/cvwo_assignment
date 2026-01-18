import { SideBar } from "./sidebar";
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { PlusCircleIcon, UserCircleIcon } from "@phosphor-icons/react";
import { Link } from "react-router";
import { useNavigate } from "react-router";
import { Button, IconButton, Typography, Menu, MenuItem } from "@mui/material";
import { useAuth } from "../hooks/useAuth";
import { useState, type MouseEvent } from "react";
import axios from "axios";

type Topic = { topic_id: number; title: string; user_id: number };
type TopicsResponse = { payload?: { data?: Topic[] } };

//Header component for all pages
export function Header() {
  const navigate = useNavigate();
  const queryClient = useQueryClient();
  const { user, isLoading: userLoading } = useAuth();
  const [userAnchorEl, setUserAnchorEl] = useState<HTMLElement | null>(null);
  const [addAnchorEl, setAddAnchorEl] = useState<HTMLElement | null>(null);
  const getTopics = async (): Promise<TopicsResponse> =>
    await axios.get("http://localhost:8000/topics");

  const { isLoading, data } = useQuery<TopicsResponse>({
    queryKey: [`topics`],
    queryFn: getTopics,
  });

  const handleUserClick = (event: MouseEvent<HTMLElement>) => {
    setUserAnchorEl(event.currentTarget);
  };

  const handleUserClose = () => {
    setUserAnchorEl(null);
  };

  const handleAddClick = (event: MouseEvent<HTMLElement>) => {
    setAddAnchorEl(event.currentTarget);
  };

  const handleAddClose = () => {
    setAddAnchorEl(null);
  };

  const logout = useMutation({
    mutationFn: async () =>
      await axios.post("http://localhost:8000/logout", {
        withCredentials: true,
      }),
    onSuccess: () => {
      () => {
        queryClient.invalidateQueries({ queryKey: ["me"] });
        alert("Successfully logged out.");
        handleUserClose();
        navigate("/");
      };
    },
    onError: () => alert("Failed to logout"),
  });
  if (isLoading || userLoading) return <div>Loading...</div>;
  return (
    <div className="h-20 sticky top-0 z-50">
      <div className="h-20 bg-[#9BE3FF] ps-4 flex items-center gap-4">
        <SideBar topicsList={data?.payload?.data} userId={user?.payload?.data || null}/>
        <Link to="/">
          <Typography fontWeight={500} color="white" fontSize={46}>
            CVWO
          </Typography>
        </Link>
        <div className="w-full flex justify-end px-4">
          {/* TODO impmenet my post */}
          {user ? (
            <>
              <IconButton
                aria-controls="add-dropdown"
                onClick={handleAddClick}
                aria-haspopup="true"
                sx={{ borderRadius: 999 }}
              >
                <PlusCircleIcon size="60" color="white" weight="bold" />
              </IconButton>
              <Menu
                id="add-dropdown"
                anchorEl={addAnchorEl}
                keepMounted
                open={Boolean(addAnchorEl)}
                onClose={handleAddClose}
              >
                <MenuItem
                  component={Link}
                  to="addposts"
                  onClick={handleAddClose}
                >
                  Add Posts
                </MenuItem>
                <MenuItem
                  component={Link}
                  to="addtopics"
                  onClick={handleAddClose}
                >
                  Add Topics
                </MenuItem>
              </Menu>
              <IconButton
                aria-controls="user-dropdown"
                aria-haspopup="true"
                sx={{ borderRadius: 999 }}
                onClick={handleUserClick}
              >
                <UserCircleIcon size="60" color="white" weight="bold" />
              </IconButton>
              <Menu
                id="user-dropdown"
                anchorEl={userAnchorEl}
                keepMounted
                open={Boolean(userAnchorEl)}
                onClose={handleUserClose}
              >
                <MenuItem onClick={() => logout.mutate()}>Logout</MenuItem>
                <MenuItem>My Posts</MenuItem>
              </Menu>
            </>
          ) : (
            <Button
              component={Link}
              to="/sign-in"
              variant="text"
              sx={{ borderRadius: 999 }}
            >
              <Typography fontWeight={500} color="white" fontSize={28}>
                Login
              </Typography>
            </Button>
          )}
        </div>
      </div>
    </div>
  );
}
