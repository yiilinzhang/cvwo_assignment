import { List } from "@phosphor-icons/react";
import { useState } from "react";

const topicList = ["Tech", "Game", "Fashion"];

export function SideBar(){

    const [expanded, setExpanded] = useState(false)
    const expandSidebar = () => setExpanded(!expanded)
    return (
    <main>
        <div className={`${expanded ? `visible` : `invisible`}`}>
            <div className="bg-gray-600 opacity-20 flex h-screen w-screen fixed top-20 left-0"></div>
            <div className="bg-white top-20 h-screen fixed w-80 left-0 shadow flex flex-col">
            {topicList.map((item) => {
                return (
                    <button className="hover:cursor-pointer hover:bg-[#6AD5FF] rounded-xl p-3 ">
                        <text>{item}</text>
                    </button>
            )})}</div>
        </div>
        <button onClick={expandSidebar} className="hover:cursor-pointer hover:bg-[#6AD5FF] rounded-4xl p-3">
            <List size={40} color="white" weight="bold"/>
        </button>
        
    </main>)
}
