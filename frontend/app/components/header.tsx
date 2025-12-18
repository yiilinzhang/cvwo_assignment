
import { SideBar } from "./sidebar";
import { useQuery } from "@tanstack/react-query";


export function Header(){
    const {isLoading, data} = useQuery({
            queryKey: [`topics`], 
            queryFn: getTopics
        })
    if (isLoading) return <div>Loading...</div>;    
    return <header>
        <div className = "h-20 bg-[#9BE3FF] ps-4 flex items-center gap-4">
           <SideBar topicsList = {data?.payload.data}/>
            <text className = "text-white font-bold text-5xl">CVWO</text>
        </div>

    </header>
}

const getTopics = async () => {
    const response = await fetch("http://localhost:8000/topics") 
    return await response.json()
}
