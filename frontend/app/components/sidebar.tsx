import {
  ListIcon,
  TrashIcon,
  HouseIcon,
  ArrowCircleUpRightIcon,
} from "@phosphor-icons/react";
import { Link } from "react-router";
import { useState } from "react";
import { Button, IconButton, Typography } from "@mui/material";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import axios from "axios";
type Topic = { topic_id: number; title: string; user_id: number };

//SIidebar component for topics list. Used in Header
export function SideBar({
  topicsList = [],
  userID,
}: {
  topicsList?: Topic[];
  userID: number | null;
}) {
  const [expanded, setExpanded] = useState(false);
  const queryClient = useQueryClient();
  const toggle = () => setExpanded(!expanded);
  const deleteTopic = useMutation({
    mutationFn: async (id: number) => {
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
      <div
        className={`fixed inset-0 top-20 z-40 ${expanded ? "pointer-events-auto" : "pointer-events-none"}`}
      >
        <div
          className={`bg-zinc-400 transition-opacity duration-300 ease-out h-screen ${expanded ? "opacity-40 backdrop-blur-sm" : "opacity-0"}`}
          onClick={toggle}
        />
        <div
          className={`bg-white duration-300 ease-in-out top-20 gap-2 h-screen fixed w-80 px-3 flex flex-col ${expanded ? "translate-x-0" : "-translate-x-full"}`}
        >
          <Button
            component={Link}
            to={`/`}
            onClick={toggle}
            startIcon={<HouseIcon />}
            sx={{ color: "black" }}
          >
            <Typography fontSize={17}>Home</Typography>
          </Button>
          <Button
            startIcon={<ArrowCircleUpRightIcon />}
            sx={{ color: "black" }}
          >
            <Typography fontSize={17}>Popular</Typography>
          </Button>

          <hr/>
          {topicsList.map((item) => {
            return (
              <div className="flex flex-row" key={item.topic_id}>
                <Button
                  component={Link}
                  to={`/posts/${item.topic_id}`}
                  variant="text"
                  onClick={toggle}
                  sx={{ textTransform: "none", color: "black", py: 1.5 }}
                >
                  <Typography fontSize={20}>{item.title}</Typography>
                </Button>
                {userID !== null && item.user_id === userID && (
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
        onClick={toggle}
        sx={{
          borderRadius: "50%",
          width: "60px",
          height: "60px",
          "&:hover": {
            backgroundColor: "#6AD5FF",
          },
        }}
      >
        <ListIcon size={40} color="white" weight="bold" />
      </IconButton>
    </main>
  );
}
