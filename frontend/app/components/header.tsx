import { SideBar } from "./sidebar";
import { useQuery, useQueryClient } from "@tanstack/react-query";
import { PlusCircleIcon, UserCircleIcon } from "@phosphor-icons/react";
import { Link } from "react-router";
import { useNavigate } from "react-router";
import { Button, IconButton, Typography, Menu, MenuItem } from "@mui/material";
import { useAuth } from "../hooks/useAuth";
import { useState } from "react";
import axios from "axios";

//Header component for all pages
export function Header() {
  const navigate = useNavigate();
  const queryClient = useQueryClient();
  const { user, isLoading: userLoading } = useAuth();
  const [userAnchorEl, setUserAnchorEl] = useState(null);
  const [addAnchorEl, setAddAnchorEl] = useState(null);
  const { isLoading, data } = useQuery({
    queryKey: [`topics`],
    queryFn: getTopics,
  });
  if (isLoading || userLoading) return <div>Loading...</div>;
  const handleUserClick = (event) => {
    setUserAnchorEl(event.currentTarget);
  };
  const handleUserClose = () => {
    setUserAnchorEl(null);
  };
  const handleAddClick = (event) => {
    setAddAnchorEl(event.currentTarget);
  };
  const handleAddClose = () => {
    setAddAnchorEl(null);
  };

  const logout = () => {
    axios
      .post("http://localhost:8000/logout", {}, { withCredentials: true })
      .then(() => {
        queryClient.invalidateQueries({ queryKey: ["me"] });
        alert("Successfully logged out.");
        handleUserClose();
        navigate("/");
      });
  };
  return (
    <div className="h-20 sticky top-0">
      <div className="h-20 bg-[#9BE3FF] ps-4 flex items-center gap-4">
        <SideBar topicsList={data?.payload.data} />
        <Link to="/">
          <Typography fontWeight={500} color="white" fontSize={46}>
            CVWO
          </Typography>
        </Link>
        <div className="w-full flex justify-end px-4">
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
            <MenuItem component={Link} to="addposts" onClick={handleAddClose}>
              Add Posts
            </MenuItem>
            <MenuItem component={Link} to="addtopics" onClick={handleAddClose}
            >Add Topics</MenuItem>
          </Menu>

          {/* TODO add drop down to log out */}
          {user ? (
            <>
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
                <MenuItem onClick={logout}>Logout</MenuItem>
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
//TODO change to axios

const getTopics = async () => {
  const response = await fetch("http://localhost:8000/topics");
  return await response.json();
};
