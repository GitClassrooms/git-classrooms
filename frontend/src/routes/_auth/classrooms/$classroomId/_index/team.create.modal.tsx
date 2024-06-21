import { CreateTeamForm } from "@/components/createTeamForm.tsx";
import { Dialog, DialogContent } from "@/components/ui/dialog";
import { createFileRoute, useNavigate } from "@tanstack/react-router";

export const Route = createFileRoute("/_auth/classrooms/$classroomId/_index/team/create/modal")({
  component: CreateTeamModal,
});

function CreateTeamModal() {
  const { classroomId } = Route.useParams();
  const navigate = useNavigate();
  return (
    <Dialog
      defaultOpen
      onOpenChange={(open) => {
        if (!open) {
          navigate({
            to: "/classrooms/$classroomId",
            params: { classroomId },
            replace: true,
          });
        }
      }}
    >
      <DialogContent>
        <CreateTeamForm classroomId={classroomId} />
      </DialogContent>
    </Dialog>
  );
}