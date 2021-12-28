`go doc hero`
```sh
type Hero struct {
        // Hero Name
        Name string
        // Hero Alias
        Alias string
}
    Hero struct form make new heroes

func (h Hero) Attack()
func (h Hero) Fly(kms int)
```

`go doc hero.Attack`

```sh
func (h Hero) Attack()
    Attack is a method set, permit the hero launch attack
```

`go doc hero.Fly`

```sh
func (h Hero) Fly(kms int)
    Fly is a method set, permit the hero fly
```