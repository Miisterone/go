package student

import "github.com/01-edu/z01"

func PrintCombN(n int) {

	if n == 1 {
		for nb1 := '0'; nb1 <= '9'; nb1++ {
			z01.PrintRune(nb1)
			if nb1 < '9' {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}

	if n == 2 {
		for nb1 := '0'; nb1 <= '9'; nb1++ {
			for nb2 := nb1 + 1; nb2 <= '9'; nb2++ {
				z01.PrintRune(nb1)
				z01.PrintRune(nb2)
				if nb1 < '9' {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
		z01.PrintRune('\n')
	}

	if n == 3 {
		for nb1 := '0'; nb1 <= '7'; nb1++ {
			for nb2 := nb1 + 1; nb2 <= '8'; nb2++ {
				for nb3 := nb2 + 1; nb3 <= '9'; nb3++ {
					if nb1 == '7' && nb2 == '8' && nb3 == '9' {
						z01.PrintRune(nb1)
						z01.PrintRune(nb2)
						z01.PrintRune(nb3)
						z01.PrintRune(' ')
					} else {
						z01.PrintRune(nb1)
						z01.PrintRune(nb2)
						z01.PrintRune(nb3)
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
			}

		}
		z01.PrintRune('\n')
	}

	if n == 4 {
		for nb1 := '0'; nb1 <= '7'; nb1++ {
			for nb2 := nb1 + 1; nb2 <= '8'; nb2++ {
				for nb3 := nb2 + 1; nb3 <= '9'; nb3++ {
					for nb4 := nb3 + 1; nb4 <= '9'; nb4++ {
						if nb1 == '7' && nb2 == '8' && nb3 == '9' {
							z01.PrintRune(nb1)
							z01.PrintRune(nb2)
							z01.PrintRune(nb3)
							z01.PrintRune(nb4)
							z01.PrintRune(' ')
						} else {
							z01.PrintRune(nb1)
							z01.PrintRune(nb2)
							z01.PrintRune(nb3)
							z01.PrintRune(nb4)
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}
					}
				}
			}

		}
		z01.PrintRune('\n')
	}

	if n == 5 {
		for nb1 := '0'; nb1 <= '7'; nb1++ {
			for nb2 := nb1 + 1; nb2 <= '8'; nb2++ {
				for nb3 := nb2 + 1; nb3 <= '9'; nb3++ {
					for nb4 := nb3 + 1; nb4 <= '9'; nb4++ {
						for nb5 := nb4 + 1; nb5 <= '9'; nb5++ {
							if nb1 == '7' && nb2 == '8' && nb3 == '9' {
								z01.PrintRune(nb1)
								z01.PrintRune(nb2)
								z01.PrintRune(nb3)
								z01.PrintRune(nb4)
								z01.PrintRune(nb5)
								z01.PrintRune(' ')
							} else {
								z01.PrintRune(nb1)
								z01.PrintRune(nb2)
								z01.PrintRune(nb3)
								z01.PrintRune(nb4)
								z01.PrintRune(nb5)
								z01.PrintRune(',')
								z01.PrintRune(' ')
							}
						}
					}
				}
			}

		}
		z01.PrintRune('\n')
	}

	if n == 6 {
		for nb1 := '0'; nb1 <= '7'; nb1++ {
			for nb2 := nb1 + 1; nb2 <= '8'; nb2++ {
				for nb3 := nb2 + 1; nb3 <= '9'; nb3++ {
					for nb4 := nb3 + 1; nb4 <= '9'; nb4++ {
						for nb5 := nb4 + 1; nb5 <= '9'; nb5++ {
							for nb6 := nb5 + 1; nb6 <= '9'; nb6++ {
								if nb1 == '7' && nb2 == '8' && nb3 == '9' {
									z01.PrintRune(nb1)
									z01.PrintRune(nb2)
									z01.PrintRune(nb3)
									z01.PrintRune(nb4)
									z01.PrintRune(nb5)
									z01.PrintRune(nb6)
									z01.PrintRune(' ')
								} else {
									z01.PrintRune(nb1)
									z01.PrintRune(nb2)
									z01.PrintRune(nb3)
									z01.PrintRune(nb4)
									z01.PrintRune(nb5)
									z01.PrintRune(nb6)
									z01.PrintRune(',')
									z01.PrintRune(' ')
								}
							}
						}
					}
				}
			}

		}
		z01.PrintRune('\n')
	}

	if n == 7 {
		for nb1 := '0'; nb1 <= '7'; nb1++ {
			for nb2 := nb1 + 1; nb2 <= '8'; nb2++ {
				for nb3 := nb2 + 1; nb3 <= '9'; nb3++ {
					for nb4 := nb3 + 1; nb4 <= '9'; nb4++ {
						for nb5 := nb4 + 1; nb5 <= '9'; nb5++ {
							for nb6 := nb5 + 1; nb6 <= '9'; nb6++ {
								for nb7 := nb6 + 1; nb7 <= '9'; nb7++ {
									if nb1 == '7' && nb2 == '8' && nb3 == '9' {
										z01.PrintRune(nb1)
										z01.PrintRune(nb2)
										z01.PrintRune(nb3)
										z01.PrintRune(nb4)
										z01.PrintRune(nb5)
										z01.PrintRune(nb6)
										z01.PrintRune(nb7)
										z01.PrintRune(' ')
									} else {
										z01.PrintRune(nb1)
										z01.PrintRune(nb2)
										z01.PrintRune(nb3)
										z01.PrintRune(nb4)
										z01.PrintRune(nb5)
										z01.PrintRune(nb6)
										z01.PrintRune(nb7)
										z01.PrintRune(',')
										z01.PrintRune(' ')
									}
								}
							}
						}
					}
				}
			}

		}
		z01.PrintRune('\n')
	}

	if n == 8 {
		for nb1 := '0'; nb1 <= '7'; nb1++ {
			for nb2 := nb1 + 1; nb2 <= '8'; nb2++ {
				for nb3 := nb2 + 1; nb3 <= '9'; nb3++ {
					for nb4 := nb3 + 1; nb4 <= '9'; nb4++ {
						for nb5 := nb4 + 1; nb5 <= '9'; nb5++ {
							for nb6 := nb5 + 1; nb6 <= '9'; nb6++ {
								for nb7 := nb6 + 1; nb7 <= '9'; nb7++ {
									for nb8 := nb7 + 1; nb8 <= '9'; nb8++ {
										if nb1 == '7' && nb2 == '8' && nb3 == '9' {
											z01.PrintRune(nb1)
											z01.PrintRune(nb2)
											z01.PrintRune(nb3)
											z01.PrintRune(nb4)
											z01.PrintRune(nb5)
											z01.PrintRune(nb6)
											z01.PrintRune(nb7)
											z01.PrintRune(nb8)
											z01.PrintRune(' ')
										} else {
											z01.PrintRune(nb1)
											z01.PrintRune(nb2)
											z01.PrintRune(nb3)
											z01.PrintRune(nb4)
											z01.PrintRune(nb5)
											z01.PrintRune(nb6)
											z01.PrintRune(nb7)
											z01.PrintRune(nb8)
											z01.PrintRune(',')
											z01.PrintRune(' ')
										}
									}
								}
							}
						}
					}
				}
			}

		}
		z01.PrintRune('\n')
	}

	if n == 9 {
		for nb1 := '0'; nb1 <= '9'; nb1++ {
			for nb2 := nb1 + 1; nb2 <= '9'; nb2++ {
				for nb3 := nb2 + 1; nb3 <= '9'; nb3++ {
					for nb4 := nb3 + 1; nb4 <= '9'; nb4++ {
						for nb5 := nb4 + 1; nb5 <= '9'; nb5++ {
							for nb6 := nb5 + 1; nb6 <= '9'; nb6++ {
								for nb7 := nb6 + 1; nb7 <= '9'; nb7++ {
									for nb8 := nb7 + 1; nb8 <= '9'; nb8++ {
										for nb9 := nb8 + 1; nb9 <= '9'; nb9++ {
											if nb1 == '1' && nb2 == '2' && nb3 == '3' && nb4 == '4' && nb5 == '5' && nb6 == '6' && nb7 == '7' && nb8 == '8' && nb9 == '9' {
												z01.PrintRune(nb1)
												z01.PrintRune(nb2)
												z01.PrintRune(nb3)
												z01.PrintRune(nb4)
												z01.PrintRune(nb5)
												z01.PrintRune(nb6)
												z01.PrintRune(nb7)
												z01.PrintRune(nb8)
												z01.PrintRune(nb9)
												z01.PrintRune(' ')
											} else {
												z01.PrintRune(nb1)
												z01.PrintRune(nb2)
												z01.PrintRune(nb3)
												z01.PrintRune(nb4)
												z01.PrintRune(nb5)
												z01.PrintRune(nb6)
												z01.PrintRune(nb7)
												z01.PrintRune(nb8)
												z01.PrintRune(nb9)
												z01.PrintRune(',')
												z01.PrintRune(' ')
											}
										}
									}
								}
							}
						}
					}
				}
			}

		}
		z01.PrintRune('\n')
	}
}
