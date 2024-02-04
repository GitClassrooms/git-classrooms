import { assignmentsQueryOptions, classroomMemberQueryOptions, classroomQueryOptions, classroomsQueryOptions } from "@/api/classrooms";
import { Button } from "@/components/ui/button";
import { Card, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Table, TableBody, TableCaption, TableCell, TableHead, TableHeader, TableRow } from "@/components/ui/table";
import { Assignment } from "@/types/assignments";
import { Classroom } from "@/types/classroom";
import { User } from "@/types/user";
import { useSuspenseQuery } from "@tanstack/react-query";
import { Link, createFileRoute } from "@tanstack/react-router";
import { Loader2 } from "lucide-react";

export const Route = createFileRoute("/_auth/classrooms/$classroomId")({
  component: ClassroomDetail,
  loader: ({ context, params }) => {
    const classroom = context.queryClient.ensureQueryData(classroomQueryOptions(params.classroomId))
    const assignments = context.queryClient.ensureQueryData(assignmentsQueryOptions(params.classroomId))
    const members = context.queryClient.ensureQueryData(classroomMemberQueryOptions(params.classroomId))

    return { classroom, assignments, members }

  },
  pendingComponent: () => <Loader2 className="animate-spin h-8 w-8" />,
});

function ClassroomDetail() {
  const { classroomId } = Route.useParams()
  const { data: classroom } = useSuspenseQuery(classroomQueryOptions(classroomId))
  const { data: assignments } = useSuspenseQuery(assignmentsQueryOptions(classroomId))
  const { data: members } = useSuspenseQuery(classroomMemberQueryOptions(classroomId))

  return (
    <div className="p-2 space-y-6">
      <div className="flex flex-row justify-between">
        <h1 className="text-xl font-bold">Classroom Details </h1>
      </div>

      <Card>
        <CardHeader>
          <CardTitle>{classroom.classroom.name}</CardTitle>
          <CardDescription>{classroom.classroom.description}</CardDescription>
        </CardHeader>
      </Card>

      <div className="flex flex-row justify-between">
        <h1 className="text-xl font-bold">Assignments</h1>
        <Button variant="default">
          Create assignment
        </Button>
      </div>
      <AssignmentTable assignments={assignments} />

      <div className="flex flex-row justify-between">
        <h1 className="text-xl font-bold">Members</h1>
        <Button variant="default">
          Invite members
        </Button>
      </div>
      <MemberTable members={members} />
    </div>
  );
}


function AssignmentTable({ assignments }: { assignments: Assignment[] }) {
  return (
    <Table>
      <TableCaption>Assignments</TableCaption>
      <TableHeader>
        <TableRow>
          <TableHead>Name</TableHead>
          <TableHead>Due date</TableHead>
          <TableHead className="text-right">Actions</TableHead>
        </TableRow>
      </TableHeader>
      <TableBody>
        {assignments.map(a =>
          <TableRow key={a.id}>
            <TableCell>{a.name}</TableCell>
            <TableCell>{a.dueDate}</TableCell>
            <TableCell className="text-right">
            </TableCell>
          </TableRow>
        )}
      </TableBody>
    </Table>
  )
}

function MemberTable({ members }: { members: User[] }) {
  return (
    <Table>
      <TableCaption>Members</TableCaption>
      <TableHeader>
        <TableRow>
          <TableHead>Name</TableHead>
          <TableHead>E-Mail</TableHead>
          <TableHead className="text-right">Actions</TableHead>
        </TableRow>
      </TableHeader>
      <TableBody>
        {members.map(m =>
          <TableRow key={m.id}>
            <TableCell>{m.name}</TableCell>
            <TableCell>{m.gitlabEmail}</TableCell>
            <TableCell className="text-right">
            </TableCell>
          </TableRow>
        )}
      </TableBody>
    </Table>
  )
}
