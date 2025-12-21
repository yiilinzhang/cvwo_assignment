import { type RouteConfig, index, route } from "@react-router/dev/routes";

export default [
  index("routes/home.tsx"),
  route("posts/:id", "routes/posts.$id.tsx"),
  route("addposts", "routes/addposts.tsx"),
] satisfies RouteConfig;
