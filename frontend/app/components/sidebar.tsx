import { ListIcon, PencilSimpleLineIcon, TrashIcon } from "@phosphor-icons/react";
import { Link } from "react-router";
import { useState, useEffect } from "react";
import { Button, IconButton, Typography } from "@mui/material";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import axios from "axios";
import { useAuth } from "../hooks/useAuth"


//SIidebar component for topics list. Used in Header
export function SideBar({ topicsList = [] }: { topicsList?: any[] }) {
  const [expanded, setExpanded] = useState(false);
    const queryClient = useQueryClient();
    const {user, isLoading} = useAuth()
  const expandSidebar = () => setExpanded(!expanded);
   const deleteTopic = useMutation({
    mutationFn: async () => {
      await axios.delete(`http://localhost:8000/topics/${id}`,
        {withCredentials: true}
      );
    },
    onSuccess: () =>{ 
       queryClient.invalidateQueries({ queryKey: ['topics'] });
       alert("Topics sucessfully deleted")}
  })
  return (
    <main>
      <div className={`${expanded ? `visible` : `invisible`}`}>
        <div className="bg-gray-600 opacity-20 flex h-screen w-screen fixed top-20 left-0"></div>
        <div className="bg-white top-20 h-screen fixed w-80 left-0 shadow flex flex-col">
          {topicsList.map((item) => {
            return (
              <>

              <Button
                component={Link}
                to={`/posts/${item.topic_id}`}
                variant="text"
                onClick={expandSidebar}
                sx={{ textTransform: "none", color: "black", py: 1.5 }}
              >
                <Typography sx={{ fontSize: "20px" }}>{item.title}</Typography>
              </Button>
              {item.userId === user?.payload.data && <>
            <IconButton aria-label="edit_post">
              <PencilSimpleLineIcon size={30} color="black" />
            </IconButton>
            <IconButton aria-label="delete_post" onClick={() => deleteTopic.mutate()}>
              <TrashIcon size={30} color="black" />
            </IconButton>
          </>}</>
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
