import { ChatsIcon, PencilSimpleLineIcon, TrashIcon } from "@phosphor-icons/react";

export function Post({title, content, owner}){
    return (<div className = "bg-[#D9D9D9] w-5/6 flex flex-col p-4 gap-2">
        <text className="text-2xl font-semibold">{title}</text>
        <text className="text-xl ">{content}</text>
        <div className="flex justify-end">
        <button className="hover:cursor-pointer hover:bg-stone-300 p-1 rounded-2xl"><ChatsIcon size={30} /></button>
        <button className={`hover:cursor-pointer hover:bg-stone-300 p-1 rounded-2xl ${owner ? `visible` : `hidden`}`}><PencilSimpleLineIcon size={30} /></button>
        <button className={`hover:cursor-pointer hover:bg-stone-300 p-1 rounded-2xl ${owner ? `visible` : `hidden`}`}><TrashIcon size={30} /></button>
        </div>

    </div>)
}