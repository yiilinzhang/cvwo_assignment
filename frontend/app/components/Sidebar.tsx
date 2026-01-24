import {
  ListIcon,
  HouseIcon,
  ArrowCircleUpRightIcon,
} from "@phosphor-icons/react";
import { Link } from "react-router";
import { useEffect, useState } from "react";
import { Button, IconButton, Typography } from "@mui/material";
import { useQueryClient } from "@tanstack/react-query";
type Topic = { topic_id: number; title: string; user_id: number };

//SIidebar component for topics list. Used in Header
export function SideBar({ topicsList = [] }: { topicsList?: Topic[] }) {
  const [expanded, setExpanded] = useState(false);
  const queryClient = useQueryClient();
  const toggle = () => setExpanded(!expanded);
  useEffect(() => {
    document.body.style.overflow = expanded ? "hidden" : "";
  }, [expanded]);
  const sidebarItem = {
    justifyContent: "flex-start",
    textTransform: "none",
    color: "black",
    px: 1.5,
    py: 1,
    gap: 1,
    "&:hover": { backgroundColor: "rgba(0,0,0,0.06)" },
  };
  return (
    <main>
      {expanded && (
        <div className=" inset-0 top-20 fixed">
          <div
            className={`bg-zinc-400 h-screen ${expanded ? "visible opacity-20" : "invisible pointer-events-none"}`}
            onClick={toggle}
          />
          <div
            className={`bg-white duration-300 transition-transform ease-in-out top-20 gap-2 h-screen fixed w-80 px-3 flex flex-col overflow-y-auto ${
              expanded ? "translate-x-0" : "-translate-x-full"
            }`}
          >
            <Button
              component={Link}
              to={`/`}
              onClick={toggle}
              startIcon={<HouseIcon />}
              sx={sidebarItem}
            >
              <Typography fontSize={17}>Home</Typography>
            </Button>

            <hr />
            {topicsList.map((item) => {
              return (
                <div className="flex flex-col" key={item.topic_id}>
                  <Button
                    component={Link}
                    to={`/posts/${item.topic_id}`}
                    onClick={toggle}
                    sx={sidebarItem}
                  >
                    <Typography fontSize={17}>{item.title}</Typography>
                  </Button>
                </div>
              );
            })}
          </div>
        </div>
      )}
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
