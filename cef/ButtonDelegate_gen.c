// Code created from "callback.c.tmpl" - don't edit by hand

#include "ButtonDelegate_gen.h"
#include "_cgo_export.h"

void gocef_set_button_delegate_proxy(cef_button_delegate_t *self) {
	// Casts to (void *) added to avoid warnings since Go callbacks can't have
	// some modifiers, such as 'const' applied to their parameter signatures.
	self->on_button_pressed = (void *)&gocef_button_delegate_on_button_pressed;
	self->on_button_state_changed = (void *)&gocef_button_delegate_on_button_state_changed;
}
