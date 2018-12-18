#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
#include <string.h>
#include <inttypes.h>

#include "glib_indirect.h"

static const size_t MAXIMUM_LINE_LENGTH = 20;
static const char location_format[] = "%" SCNu16 ", %" SCNu16;

static intmax_t manhatten_distance(uint16_t p_x, uint16_t p_y, uint16_t q_x, uint16_t q_y)
{
  intmax_t x_distance = imaxabs(p_x - q_x);
  intmax_t y_distance = imaxabs(p_y - q_y);

  return (x_distance + y_distance);
}

struct location
{
  uint16_t x;
  uint16_t y;
  bool infinite_area;
};

struct coordinate
{
  uint16_t x;
  uint16_t y;
  uint16_t closest_location_count;
  intmax_t closest_location_distance;
  intmax_t total_location_distance;
  struct location *closest_location;
};

int main(void)
{
  uint16_t x_current = 0;
  uint16_t y_current = 0;
  uint16_t x_min = UINT16_MAX;
  uint16_t x_max = 0;
  uint16_t y_min = UINT16_MAX;
  uint16_t y_max = 0;

  GSList *location_list = NULL;
  GSList *location_list_item = NULL;
  struct location *current_location = NULL;

  GSList *coordinate_list = NULL;
  GSList *coordinate_list_item = NULL;
  struct coordinate *current_coordinate = NULL;

  char *current_line = calloc(MAXIMUM_LINE_LENGTH, sizeof(char));

  while ((current_line = fgets(current_line, MAXIMUM_LINE_LENGTH, stdin)) != NULL)
  {
    // Remove trailing line feeds and carriage returns
    current_line[strcspn(current_line, "\r\n")] = '\0';
    sscanf(current_line, location_format, &x_current, &y_current);

    x_min = (x_current < x_min) ? x_current : x_min;
    x_max = (x_current > x_max) ? x_current : x_max;
    y_min = (y_current < y_min) ? y_current : y_min;
    y_max = (y_current > y_max) ? y_current : y_max;

    current_location = calloc(1, sizeof(struct location));
    current_location->x = x_current;
    current_location->y = y_current;
    current_location->infinite_area = false;

    location_list = g_slist_append(location_list, current_location);
  }

  printf("x_min: %" PRIu16 "\n", x_min);
  printf("x_max: %" PRIu16 "\n", x_max);
  printf("y_min: %" PRIu16 "\n", y_min);
  printf("y_max: %" PRIu16 "\n", y_max);

  for (location_list_item = location_list; location_list_item != NULL; location_list_item = location_list_item->next)
  {
    current_location = (struct location *) location_list_item->data;

    if (current_location->x == x_min || current_location->x == x_max || current_location->y == y_min || current_location->y == y_max)
    {
      current_location->infinite_area = true;
    }

    printf("Current location: %" PRIu16 ",%" PRIu16 ". Infinite area: %s\n", current_location->x, current_location->y, current_location->infinite_area ? "true" : "false");
  }

  // Build a list of coordinates, from x_min,y_min to x_max,y_max. Effectively
  // a square but implemented as a list.
  for (x_current = x_min; x_current <= x_max; x_current++)
  {
    for (y_current = y_min; y_current <= y_max; y_current++)
    {
      current_coordinate = calloc(1, sizeof(struct coordinate));
      current_coordinate->x = x_current;
      current_coordinate->y = y_current;
      current_coordinate->closest_location_count = 0;
      current_coordinate->closest_location_distance = INT_MAX;
      current_coordinate->closest_location = NULL;

      // Iterate over list of locations and find which one is 'closest'
      for (location_list_item = location_list; location_list_item != NULL; location_list_item = location_list_item->next)
      {
        current_location = (struct location *) location_list_item->data;

        intmax_t location_distance = manhatten_distance(current_location->x, current_location->y, current_coordinate->x, current_coordinate->y);

        if (location_distance < current_coordinate->closest_location_distance)
        {
          // This location is closer than any others, so reset the closest location
          current_coordinate->closest_location_count = 1;
          current_coordinate->closest_location_distance = location_distance;
          current_coordinate->closest_location = current_location;
        }
        else if (location_distance == current_coordinate->closest_location_distance)
        {
          // This location is equally close to others
          current_coordinate->closest_location_count++;

          // We no longer have a single closest location, so set this to NULL
          current_coordinate->closest_location = NULL;
        }
      }

      // Add coordinate to the list
      coordinate_list = g_slist_append(coordinate_list, current_coordinate);
    }
  }

  struct location *largest_area_location = NULL;
  uint16_t largest_area_size = 0;
  uint16_t current_area_size = 0;

  // Find the location with the highest number of closest coordinates, excluding
  // infinite locations
  for (location_list_item = location_list; location_list_item != NULL; location_list_item = location_list_item->next)
  {
    current_location = (struct location *) location_list_item->data;
    current_area_size = 0;

    if (!current_location->infinite_area)
    {
      for (coordinate_list_item = coordinate_list; coordinate_list_item != NULL; coordinate_list_item = coordinate_list_item->next)
      {
        current_coordinate = (struct coordinate *) coordinate_list_item->data;

        // Assumption: Both closest location and current location point to same memory address.
        // This should be the case if they are 'equal' as we hold pointers to each location
        // rather than the location data itself.
        if (current_coordinate->closest_location == current_location)
        {
          current_area_size++;
        }
      }

      if (current_area_size > largest_area_size)
      {
        largest_area_location = current_location;
        largest_area_size = current_area_size;
      }
    }
  }

  printf("Largest area size: %" PRIu16 "\n", largest_area_size);

  // Populate the distance to each location of each coordinate
  const uint16_t maximum_location_distance = 10000;
  intmax_t maximum_distance_count = 0;

  for (coordinate_list_item = coordinate_list; coordinate_list_item != NULL; coordinate_list_item = coordinate_list_item->next)
  {
    current_coordinate = (struct coordinate *) coordinate_list_item->data;

    for (location_list_item = location_list; location_list_item != NULL; location_list_item = location_list_item->next)
    {
      current_location = (struct location *) location_list_item->data;

      intmax_t location_distance = manhatten_distance(current_location->x, current_location->y, current_coordinate->x, current_coordinate->y);

      current_coordinate->total_location_distance += location_distance;
    }

    if (current_coordinate->total_location_distance < maximum_location_distance)
    {
      maximum_distance_count++;
    }
  }

  printf("Maximum distance count (part B): %" PRIdMAX "\n", maximum_distance_count);
}
