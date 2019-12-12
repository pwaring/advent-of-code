<?php

declare(strict_types=1);
error_reporting(E_ALL);

require 'vendor/autoload.php';

use GraphDS\Graph\UndirectedGraph;
use GraphDS\Algo\Dijkstra;

// Read in the orbits as an undirected graph
$graph = new UndirectedGraph();

$fp = fopen('input', 'r');

while ($current_orbit = fgets($fp))
{
  $current_orbit = trim($current_orbit);
  list($orbiting, $object) = explode(')', $current_orbit);

  $graph->addVertex($orbiting);
  $graph->addVertex($object);

  $graph->addEdge($orbiting, $object, 1);
}

fclose($fp);

$algorithm = new Dijkstra($graph);

$algorithm->run('YOU');
$result = $algorithm->get('SAN');

// Subtract 2 from the distance because we do not count the hop between YOU
// and the object it orbits, and the hop between SAN and the object it orbits
$min_distance = $result['dist'] - 2;

print("$min_distance\n");
