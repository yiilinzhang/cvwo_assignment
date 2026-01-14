import { type RouteConfig, index, route } from "@react-router/dev/routes";

export default [
  index("routes/home.tsx"),
  route("posts/:id", "routes/posts.$id.tsx"),
  route("postcomments/:id", "routes/postcomments.$id.tsx"),
  route("addposts", "routes/addposts.tsx"),
  route("addtopics", "routes/addtopics.tsx"),
  route("sign-in", "routes/login.tsx"),
  route("sign-up", "routes/signup.tsx"),
] satisfies RouteConfig;
