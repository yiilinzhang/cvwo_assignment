import type { Route } from "./+types/home";
import Posts from "./TopicPostsScreen.$id";

export function meta({}: Route.MetaArgs) {
  return [
    { title: "CVWO assignment" },
    { name: "description", content: "A CVWO chat forum!" },
  ];
}

export default function Home() {
  return <Posts/>;
}
