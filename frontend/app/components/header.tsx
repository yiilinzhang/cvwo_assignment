
import { SideBar } from "./sidebar";

export function Header(){
    return <header>
        <div className = "h-20 bg-[#9BE3FF] ps-4 flex items-center gap-4">
            <SideBar/>
            <text className = "text-white font-bold text-5xl">CVWO</text>
        </div>

    </header>
}

function expandSidebar(){}