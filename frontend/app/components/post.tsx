import {
  ChatsIcon,
  PencilSimpleLineIcon,
  TrashIcon,
} from "@phosphor-icons/react";
import { IconButton } from "@mui/material";

export function Post({ title, content, owner }) {
  return (
    <div className="bg-[#D9D9D9] w-full flex flex-col p-4 gap-2">
      <text className="text-2xl font-semibold">{title}</text>
      <text className="text-xl ">{content}</text>
      <div className="flex justify-end">
        <IconButton aria-label="comment">
          <ChatsIcon size={30} color="black" />
        </IconButton>

        <IconButton aria-label="edit_post">
          <PencilSimpleLineIcon size={30} color="black" />
        </IconButton>

        <IconButton aria-label="delete_post">
          <TrashIcon size={30} color="black" />
        </IconButton>
      </div>
    </div>
  );
}
