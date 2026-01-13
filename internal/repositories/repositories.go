package repositories

import (
	"errors"
	"sync"

	"github.com/kodra-pay/reconciliation-service/internal/models"
)

var ErrRunNotFound = errors.New("reconciliation run not found")

// RunRepository defines the persistence boundary for reconciliation runs.
type RunRepository interface {
	Create(run *models.ReconciliationRun) (*models.ReconciliationRun, error)
	Get(id int) (*models.ReconciliationRun, error)
	List() ([]*models.ReconciliationRun, error)
}

// InMemoryRunRepository is a simple thread-safe store; replace with DB-backed impl later.
type InMemoryRunRepository struct {
	mu   sync.Mutex
	seq  int
	data map[int]*models.ReconciliationRun
}

func NewInMemoryRunRepository() *InMemoryRunRepository {
	return &InMemoryRunRepository{
		data: make(map[int]*models.ReconciliationRun),
	}
}

func (r *InMemoryRunRepository) Create(run *models.ReconciliationRun) (*models.ReconciliationRun, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.seq++
	run.ID = r.seq
	r.data[run.ID] = run
	return run, nil
}

func (r *InMemoryRunRepository) Get(id int) (*models.ReconciliationRun, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	run, ok := r.data[id]
	if !ok {
		return nil, ErrRunNotFound
	}
	return run, nil
}

func (r *InMemoryRunRepository) List() ([]*models.ReconciliationRun, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	runs := make([]*models.ReconciliationRun, 0, len(r.data))
	for _, run := range r.data {
		runs = append(runs, run)
	}
	return runs, nil
}
