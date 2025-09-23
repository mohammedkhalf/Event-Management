package main

import (
	"net/http"
	"rest-api-in-gin/internal/database"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Create Event
func (app *application) createEvent(c *gin.Context) {

	var event database.Event

	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := app.models.Events.Insert(&event)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Event"})
		return
	}

	c.JSON(http.StatusCreated, event)

}

// Get All Events

func (app *application) getAllEvents(c *gin.Context) {

	events, err := app.models.Events.GetAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve Events"})
		return
	}

	c.JSON(http.StatusOK, events)

}

// Get Event
func (app *application) getEvent(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id")) //extracts the event ID from the URL parameters

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Event Id"})
		return
	}

	event, err := app.models.Events.Get(id)

	if event == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve Event"})
		return
	}

	c.JSON(http.StatusOK, event)

}

// Update Event
func (app *application) updateEvent(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Event Id"})
		return
	}

	existingEvent, err := app.models.Events.Get(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve Event"})
		return
	}

	if existingEvent == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event Not Found"})
		return
	}

	updateEvent := &database.Event{}

	if err := c.ShouldBindJSON(&updateEvent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateEvent.Id = id

	if err := app.models.Events.Update(updateEvent); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update Event"})
		return
	}

	c.JSON(http.StatusOK, updateEvent)
}

// Delete Event
func (app *application) deleteEvent(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Event Id"})
		return
	}

	if err := app.models.Events.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed To delete Event"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
