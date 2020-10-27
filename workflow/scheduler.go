package workflow

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
	"github.com/souleb/buildeploy/app"
	"github.com/souleb/buildeploy/workflow/internal/dag"
)

// Ensure SchedulerService implements app.SchedulerService.
var _ app.SchedulerService = (*SchedulerService)(nil)

// Ensure that a Vdefined type is hashable.
var _ dag.VertexHashable = (*JobVertex)(nil)

type JobVertex app.Job

func (j *JobVertex) Hashcode() interface{} {
	jType := reflect.TypeOf(j)

	return fmt.Sprintf("%s/%s", j.Name, jType.String())
}

// SchedulerService take a workflow as input
// It turns it into a Dag
// Then performs a topological sort
// Finally add tasks to a queue
type SchedulerService struct {
	GraphMap map[string]*dag.Graph
}

func NewSchedulerService() *SchedulerService {
	return &SchedulerService{
		GraphMap: make(map[string]*dag.Graph),
	}
}

// Schedule take a workflow and defines how to run it.
func (s *SchedulerService) Schedule(workflow *app.Workflow) error {
	g, err := s.convertToGraph(workflow)
	if err != nil {
		return errors.Wrap(err, "creating a graph of this workflow failed.")
	}
	fmt.Println("Here the scheduler take action")
	fmt.Println(g)
	fmt.Println("And the topological sort")
	topOrder, err := g.TopologicalSort()
	if err != nil {
		fmt.Println(err)
		return errors.Wrap(err, "Cannot Schedule a workflow with cycles.")
	}
	fmt.Println(*topOrder)
	s.GraphMap[workflow.Name] = g
	fmt.Println("\nScheduler has finished bye!!!")
	return nil
}

func (s *SchedulerService) convertToGraph(workflow *app.Workflow) (*dag.Graph, error) {
	// Add all jobs to the graph as vertices
	// mantains a mapping of hashes and job names
	// If a job has a dependency that is not in the graph, return an error
	var g dag.Graph
	hashMap := make(map[string]string)
	for _, job := range workflow.Jobs {
		job := JobVertex(job)
		hashMap[job.Name] = job.Hashcode().(string)
		if needs := job.Needs; needs != nil {
			for _, name := range needs {
				v, ok := g.Vertex(hashMap[name])
				if !ok {
					// We go ahead and perform an early creation of the vertex,
					// It will be overwritten later on in the for loop.
					v = &JobVertex{
						Name: name,
					}
				}
				g.AddEdge(v.(*JobVertex), &job, 1)
			}
		}
		// We create the vertex in any case to overwrite early creation
		g.Add(&job)
	}

	return &g, nil

}
