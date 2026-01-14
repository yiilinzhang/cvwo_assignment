import {
  ChatsIcon,
  PencilSimpleLineIcon,
  TrashIcon,
  ArrowBendUpLeftIcon
} from "@phosphor-icons/react";
import { IconButton } from "@mui/material";

export function Comments({ username, content, owner }) {
  return (
    <div className="bg-[#D9D9D9] w-full flex flex-col p-4 gap-2">
      <text className="text-2xl font-semibold">{username}</text>
      <text className="text-xl ">{content}</text>
      <div className="flex justify-end">
        <IconButton aria-label="comment">
          <ArrowBendUpLeftIcon size={30} color="black" />
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
