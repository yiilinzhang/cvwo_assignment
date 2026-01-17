import {
  PencilSimpleLineIcon,
  TrashIcon,
  ArrowBendUpLeftIcon
} from "@phosphor-icons/react";
import { IconButton } from "@mui/material";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import axios from "axios";


type CommentProps = {
  username: string;
  content: string;
  isOwner: boolean;
  id: number;
}

export function Comments({ username, content, isOwner, id } : CommentProps) {
  const queryClient = useQueryClient();
  const deleteComment = useMutation({
    mutationFn: async () => {
      await axios.delete(`http://localhost:8000/comments/${id}`,
        {withCredentials: true}
      );
    },
    onSuccess: () =>{ 
       queryClient.invalidateQueries({ queryKey: ['comments'] });
       alert("Post sucessfully deleted")}
  })
  return (
    <div className="bg-[#D9D9D9] w-full flex flex-col p-4 gap">
      <text className="text-2xl font-semibold">{username}</text>
      <text className="text-xl ">{content}</text>
      <div className="flex justify-end">
        <IconButton aria-label="comment">
          <ArrowBendUpLeftIcon size={27} color="black" />
        </IconButton>
        
{isOwner ? 
    <div><IconButton aria-label="edit_comment">
      <PencilSimpleLineIcon size={27} color="black" />
    </IconButton>

    <IconButton aria-label="delete_comment"
    onClick={() => deleteComment.mutate()}>
      <TrashIcon size={27} color="black" />
    </IconButton></div> : <></>}
      </div>
    </div>
  );
}
