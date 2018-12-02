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

gint compare_char(gconstpointer a, gconstpointer b)
{
  char a_comp = GPOINTER_TO_CHAR(a);
  char b_comp = GPOINTER_TO_CHAR(b);

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
