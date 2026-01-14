import {
  ChatsIcon,
  PencilSimpleLineIcon,
  TrashIcon,
} from "@phosphor-icons/react";
import { IconButton } from "@mui/material";
import { Link } from "react-router";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import axios from "axios";



type PostProps = {
  id: number;
  title: string;
  content: string;
  isOwner: boolean;
  showChat: boolean;
};

//Component to render a post.
export function Post({ id, title, content, isOwner, showChat }: PostProps) {
 const queryClient = useQueryClient();
  const deletePost = useMutation({
    mutationFn: async () => {
      await axios.delete(`http://localhost:8000/posts/${id}`,
        {withCredentials: true}
      );
    },
    onSuccess: () =>{ 
       queryClient.invalidateQueries({ queryKey: ['posts'] });
       alert("Post sucessfully deleted")}
  })

  return (
    <div className="bg-[#D9D9D9] w-full flex flex-col p-4 gap-2">
      <text className="text-2xl font-semibold">{title}</text>
      <text className="text-xl ">{content}</text>
      <div className="flex justify-end">
       {showChat &&
        <IconButton
          aria-label="comment"
          component={Link}
          to={`/postcomments/${id}`}
        >
          <ChatsIcon size={30} color="black" />
        </IconButton>}

        {isOwner && (
          <>
            <IconButton aria-label="edit_post">
              <PencilSimpleLineIcon size={30} color="black" />
            </IconButton>
            <IconButton aria-label="delete_post" onClick={() => deletePost.mutate()}>
              <TrashIcon size={30} color="black" />
            </IconButton>
          </>
        )}
      </div>
    </div>
  );
}
