import type { Route } from "./+types/home";
import Posts from "./posts.$id";

export function meta({}: Route.MetaArgs) {
  return [
    { title: "CVWO assignment" },
    { name: "description", content: "A CVWO chat forum!" },
  ];
}

export default function Home() {
  return <Posts params={null} />;
}
