import { type RouteConfig, index, route } from "@react-router/dev/routes";

export default [
    index("routes/home.tsx"),
    route("posts/:id", "routes/posts.$id.tsx"),
] satisfies RouteConfig;

