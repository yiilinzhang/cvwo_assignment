import type { Route } from "./+types/home";
import { LoginPage } from "./login";
import Posts from "./posts.$id";

export function meta({}: Route.MetaArgs) {
  return [
    { title: "New React Router App" },
    { name: "description", content: "Welcome to React Router!" },
  ];
}

export default function Home() {
  return //<LoginPage/>
  <Posts params={null} />;
}
