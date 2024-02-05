// This file is auto-generated by TanStack Router

// Import Routes

import { Route as rootRoute } from "./routes/__root"
import { Route as LoginImport } from "./routes/login"
import { Route as AuthImport } from "./routes/_auth"
import { Route as IndexImport } from "./routes/index"
import { Route as AuthClassroomsIndexImport } from "./routes/_auth/classrooms/index"
import { Route as AuthClassroomsCreateImport } from "./routes/_auth/classrooms/create"
import { Route as AuthClassroomsClassroomIdIndexImport } from "./routes/_auth/classrooms/$classroomId/index"
import { Route as AuthClassroomsClassroomIdInviteImport } from "./routes/_auth/classrooms/$classroomId/invite"
import { Route as AuthClassroomsClassroomIdAssignmentsCreateImport } from "./routes/_auth/classrooms/$classroomId/assignments/create"
import { Route as AuthClassroomsClassroomIdAssignmentsAssignmentIdIndexImport } from "./routes/_auth/classrooms/$classroomId/assignments/$assignmentId/index"

// Create/Update Routes

const LoginRoute = LoginImport.update({
  path: "/login",
  getParentRoute: () => rootRoute,
} as any)

const AuthRoute = AuthImport.update({
  id: "/_auth",
  getParentRoute: () => rootRoute,
} as any)

const IndexRoute = IndexImport.update({
  path: "/",
  getParentRoute: () => rootRoute,
} as any)

const AuthClassroomsIndexRoute = AuthClassroomsIndexImport.update({
  path: "/classrooms/",
  getParentRoute: () => AuthRoute,
} as any)

const AuthClassroomsCreateRoute = AuthClassroomsCreateImport.update({
  path: "/classrooms/create",
  getParentRoute: () => AuthRoute,
} as any)

const AuthClassroomsClassroomIdIndexRoute =
  AuthClassroomsClassroomIdIndexImport.update({
    path: "/classrooms/$classroomId/",
    getParentRoute: () => AuthRoute,
  } as any)

const AuthClassroomsClassroomIdInviteRoute =
  AuthClassroomsClassroomIdInviteImport.update({
    path: "/classrooms/$classroomId/invite",
    getParentRoute: () => AuthRoute,
  } as any)

const AuthClassroomsClassroomIdAssignmentsCreateRoute =
  AuthClassroomsClassroomIdAssignmentsCreateImport.update({
    path: "/classrooms/$classroomId/assignments/create",
    getParentRoute: () => AuthRoute,
  } as any)

const AuthClassroomsClassroomIdAssignmentsAssignmentIdIndexRoute =
  AuthClassroomsClassroomIdAssignmentsAssignmentIdIndexImport.update({
    path: "/classrooms/$classroomId/assignments/$assignmentId/",
    getParentRoute: () => AuthRoute,
  } as any)

// Populate the FileRoutesByPath interface

declare module "@tanstack/react-router" {
  interface FileRoutesByPath {
    "/": {
      preLoaderRoute: typeof IndexImport
      parentRoute: typeof rootRoute
    }
    "/_auth": {
      preLoaderRoute: typeof AuthImport
      parentRoute: typeof rootRoute
    }
    "/login": {
      preLoaderRoute: typeof LoginImport
      parentRoute: typeof rootRoute
    }
    "/_auth/classrooms/create": {
      preLoaderRoute: typeof AuthClassroomsCreateImport
      parentRoute: typeof AuthImport
    }
    "/_auth/classrooms/": {
      preLoaderRoute: typeof AuthClassroomsIndexImport
      parentRoute: typeof AuthImport
    }
    "/_auth/classrooms/$classroomId/invite": {
      preLoaderRoute: typeof AuthClassroomsClassroomIdInviteImport
      parentRoute: typeof AuthImport
    }
    "/_auth/classrooms/$classroomId/": {
      preLoaderRoute: typeof AuthClassroomsClassroomIdIndexImport
      parentRoute: typeof AuthImport
    }
    "/_auth/classrooms/$classroomId/assignments/create": {
      preLoaderRoute: typeof AuthClassroomsClassroomIdAssignmentsCreateImport
      parentRoute: typeof AuthImport
    }
    "/_auth/classrooms/$classroomId/assignments/$assignmentId/": {
      preLoaderRoute: typeof AuthClassroomsClassroomIdAssignmentsAssignmentIdIndexImport
      parentRoute: typeof AuthImport
    }
  }
}

// Create and export the route tree

export const routeTree = rootRoute.addChildren([
  IndexRoute,
  AuthRoute.addChildren([
    AuthClassroomsCreateRoute,
    AuthClassroomsIndexRoute,
    AuthClassroomsClassroomIdInviteRoute,
    AuthClassroomsClassroomIdIndexRoute,
    AuthClassroomsClassroomIdAssignmentsCreateRoute,
    AuthClassroomsClassroomIdAssignmentsAssignmentIdIndexRoute,
  ]),
  LoginRoute,
])
