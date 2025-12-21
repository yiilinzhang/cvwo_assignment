import { List } from "@phosphor-icons/react";
import { Link } from "react-router";
import { useState, useEffect } from "react";


export function SideBar({ topicsList = [] }: { topicsList?: any[] }) {

    const [expanded, setExpanded] = useState(false)
    const expandSidebar = () => setExpanded(!expanded)
    return (
    <main>
        <div className={`${expanded ? `visible` : `invisible`}`}>
            <div className="bg-gray-600 opacity-20 flex h-screen w-screen fixed top-20 left-0"></div>
            <div className="bg-white top-20 h-screen fixed w-80 left-0 shadow flex flex-col">
            {topicsList.map((item) => {
                return (
                    <Link to={`/posts/${item.topic_id}` }>
                    <button className="hover:cursor-pointer hover:bg-[#6AD5FF] rounded-xl p-3 w-full" onClick={expandSidebar}>
                        <text>{item.title}</text>
                    </button></Link>
            )})}
            </div>
        </div>
        <button onClick={expandSidebar} className="hover:cursor-pointer hover:bg-[#6AD5FF] rounded-4xl p-3">
            <List size={40} color="white" weight="bold"/>
        </button>
        
    </main>)}