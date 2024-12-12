package discover



type DiscoverPage struct {
  visble bool
}

// MinSize returns the minimum size this object needs to be drawn.
func (dp *DiscoverPage) MinSize() Size{}


// Move moves this object to the given position relative to its parent.
// This should only be called if your object is not in a container with a layout manager.
func (dp *DiscoverPage) Move(Position){}


// Position returns the current position of the object relative to its parent.
func (dp *DiscoverPage) Position() Position{}


// Resize resizes this object to the given size.
// This should only be called if your object is not in a container with a layout manager.
func (dp *DiscoverPage) Resize(Size){}


// Size returns the current size of this object.
func (dp *DiscoverPage) Size() Size{}


// Hide hides this object.
func (dp *DiscoverPage) Hide(){
  dp.visible = false
}


// Visible returns whether this object is visible or not.
func (dp *DiscoverPage) Visible() bool{
  return dp.visible
}


// Show shows this object.
func (dp *DiscoverPage) Show(){
  db.visible = true
}


// Refresh must be called if this object should be redrawn because its inner state changed.
func (dp *DiscoverPage) Refresh(){}

// --- New Functions -----------------------------------------------
