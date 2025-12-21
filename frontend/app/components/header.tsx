import { SideBar } from "./sidebar";
import { useQuery } from "@tanstack/react-query";
import { PlusCircleIcon } from "@phosphor-icons/react";
import { Link } from "react-router";

export function Header() {
  const { isLoading, data } = useQuery({
    queryKey: [`topics`],
    queryFn: getTopics,
  });
  if (isLoading) return <div>Loading...</div>;
  return (
    <div className="h-20">
      <div className="h-20 bg-[#9BE3FF] ps-4 flex items-center gap-4">
        <SideBar topicsList={data?.payload.data} />
        <Link to="/">
          <text className="text-white font-bold text-5xl">CVWO</text>
        </Link>
        <div className="w-full flex justify-end px-4">
          <Link to="addposts">
            <button className="hover:cursor-pointer hover:bg-stone-300 p-1 rounded-2xl">
              <PlusCircleIcon size="50" color="white" weight="bold" />
            </button>
          </Link>
        </div>
      </div>
    </div>
  );
}

const getTopics = async () => {
  const response = await fetch("http://localhost:8000/topics");
  return await response.json();
};
