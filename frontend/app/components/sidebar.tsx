import {
  ListIcon,
  PencilSimpleLineIcon,
  TrashIcon,
  HouseIcon,
  ArrowCircleUpRightIcon,
} from "@phosphor-icons/react";
import { Link } from "react-router";
import { useState, useEffect } from "react";
import { Button, IconButton, Typography } from "@mui/material";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import axios from "axios";
import { useAuth } from "../hooks/useAuth";

//SIidebar component for topics list. Used in Header
export function SideBar({ topicsList = [] }: { topicsList?: any[] }) {
  const [expanded, setExpanded] = useState(false);
  const queryClient = useQueryClient();
  const { user, isLoading } = useAuth();
  const expandSidebar = () => setExpanded(!expanded);
  const deleteTopic = useMutation({
    mutationFn: async (id) => {
      await axios.delete(`http://localhost:8000/topics/${id}`, {
        withCredentials: true,
      });
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["topics"] });
      alert("Topics sucessfully deleted");
    },
  });
  return (
    <main>
      <div className={`${expanded ? `visible` : `invisible`}`}>
        <div className="bg-gray-600 opacity-20 flex fixed top-20 left-0 w-screen h-screen z-40"
        onClick={expandSidebar}
        ></div>
        <div className="bg-white top-20 gap-2 h-screen fixed w-80 left-0 px-3 shadow flex flex-col z-50">
          <Button
            component={Link}
            to={`/`}
            onClick={expandSidebar}
            startIcon={<HouseIcon />}
            sx={{ color: "black" }}
          >
            Home
          </Button>
          <Button
            startIcon={<ArrowCircleUpRightIcon />}
            sx={{ color: "black" }}
          >
            Popular
          </Button>

          <hr />
          <Typography fontSize={25}>Others</Typography>
          {topicsList.map((item) => {
            return (
              <div className="flex flex-row  ">
                <Button
                  component={Link}
                  to={`/posts/${item.topic_id}`}
                  variant="text"
                  onClick={expandSidebar}
                  sx={{ textTransform: "none", color: "black", py: 1.5 }}
                >
                  <Typography sx={{ fontSize: "20px" }}>
                    {item.title}
                  </Typography>
                </Button>
                {/* TODO update db to num and chaneg this */}
                {Number(item.user_id) === Number(user?.payload.data) && (
                    <IconButton
                      aria-label="delete_post"
                      onClick={() => deleteTopic.mutate(item.topic_id)}
                    >
                      <TrashIcon size={20} color="black" weight="bold" />
                    </IconButton>
                )}
              </div>
            );
          })}
        </div>
      </div>
      <IconButton
        disableRipple
        onClick={expandSidebar}
        sx={{
          borderRadius: "50%",
          width: "60px",
          height: "60px",
          "&:hover": {
            backgroundColor: "#6AD5FF", // Background color on hover
          },
        }}
      >
        <ListIcon size={40} color="white" weight="bold" />
      </IconButton>
    </main>
  );
}
