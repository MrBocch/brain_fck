,[                         read char into current cell
  [->+>+<<]                duplicate it to next two cells
  >>[-<<+>>]               move one copy back
  <<[                      test if zero (EOF)
    >>[->+>+<<]            duplicate char again
    >>[-<<+>>]             move one copy back
    <<                     back to original
    -[                     while nonzero
      >>+<<                add to helper
      -                    decrement
    ]                      now we have copy in helper
    >>[                    on helper
      -[                   while nonzero
        <<+>>              add to orig
        -                  decrement
      ]
    ]
    <<.                    print original (transformed)
    ,                      read next char
  ]
]
