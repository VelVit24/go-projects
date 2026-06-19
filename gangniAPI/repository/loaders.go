package repository

import (
	"log"

	"github.com/VelVit24/models"
)

func (r *Repository) SelectLoaders() ([]models.Loader, error) {
	rows, err := r.db.Query("select * from loaders")
	if err != nil {
		return nil, err
	}
	loaders := []models.Loader{}
	for rows.Next() {
		l := models.Loader{}
		err := rows.Scan(&l.AutoWeight, &l.BatteryType, &l.BrakeType, &l.ChargingTime, &l.EngineType, &l.ForkLength,
			&l.FrontWheels, &l.Height, &l.HydraulicLiftingEngine, &l.Id, &l.Length, &l.LiftHeight, &l.LiftingAngle,
			&l.LiftingCylinder, &l.LongmenFrameMaterial, &l.MaxLiftWeight, &l.Name, &l.PicturePathLinux, &l.PicturePathWindows,
			&l.Price, &l.RearWheels, &l.SteeringMode, &l.TurningRadius, &l.Voltage, &l.WheelAxis, &l.Width, &l.WorkingHours,
		)
		if err != nil {
			log.Println(err.Error())
		}
		loaders = append(loaders, l)
	}
	return loaders, nil
}

func (r *Repository) SelectManualLoaders() ([]models.ManualLoader, error) {
	rows, err := r.db.Query("select * from manual_loaders")
	if err != nil {
		return nil, err
	}
	loaders := []models.ManualLoader{}
	for rows.Next() {
		l := models.ManualLoader{}
		err := rows.Scan(&l.BrakeType, &l.Control, &l.DriveGear, &l.ForkLength, &l.ForkWidth,
			&l.Id, &l.Length, &l.LiftingSpeed, &l.MaxLiftWeight, &l.MaxSpeed, &l.Name,
			&l.PicturePathLinux, &l.PicturePathWindows, &l.Price,
		)
		if err != nil {
			log.Println(err.Error())
		}
		loaders = append(loaders, l)
	}
	return loaders, nil
}

func (r *Repository) SelectLoaderImage(id int) (string, string, error) {
	row := r.db.QueryRow("select picturePathLinux, picturePathWindows from loaders where id=$1", id)
	var linux, windows string
	err := row.Scan(&linux, &windows)
	return linux, windows, err
}

func (r *Repository) SelectManualLoaderImage(id int) (string, string, error) {
	row := r.db.QueryRow("select picturePathLinux, picturePathWindows from manual_loaders where id=$1", id)
	var linux, windows string
	err := row.Scan(&linux, &windows)
	return linux, windows, err
}
