package repositories_test

import (
	"encoder/application/repositories"
	"encoder/domain"
	"encoder/framework/database"
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestJobRepositoryDbInsert(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	repoVideo := repositories.VideoRepositoryDb{Db: db}
	repoVideo.Insert(video)

	job, err := domain.NewJob("output_path", "Pending", video)

	require.Nil(t, err)

	repoJob := repositories.JobRepositoryDb{Db: db}
	repoJob.Insert(job)

	foundJob, err := repoJob.Find(job.ID)

	require.NotEmpty(t, foundJob.ID)
	require.Nil(t, err)
	require.Equal(t, foundJob.ID, job.ID)
	require.Equal(t, foundJob.VideoID, job.VideoID)
}

func TestJobRepositoryDbUpdate(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	repoVideo := repositories.VideoRepositoryDb{Db: db}
	repoVideo.Insert(video)

	job, err := domain.NewJob("output_path", "Pending", video)

	require.Nil(t, err)

	repoJob := repositories.JobRepositoryDb{Db: db}
	repoJob.Insert(job)

	job.Status = "Complete"

	repoJob.Update(job)

	foundJob, err := repoJob.Find(job.ID)

	require.NotEmpty(t, foundJob.ID)
	require.Nil(t, err)
	require.Equal(t, foundJob.Status, job.Status)
}
