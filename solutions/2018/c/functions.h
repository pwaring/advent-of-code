gint compare_int(gconstpointer a, gconstpointer b)
{
  int32_t a_comp = GPOINTER_TO_INT(a);
  int32_t b_comp = GPOINTER_TO_INT(b);

  if (a_comp < b_comp)
  {
    return -1;
  }
  else if (a_comp == b_comp)
  {
    return 0;
  }
  else
  {
    return 1;
  }
}
