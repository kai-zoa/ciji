# ciji

CIJI displays various information indicators.

# example

```
$ ciji '{{.Battery.Percentage}}% {{.Battery.VProgressBar}}'
42% ▃
```

```
$ ciji 'Phase of the Moon: {{.MoonPhase.EMOJI}}'
Phase of the Moon: 🌘

```

```
# tmux.conf
set -g status-right "#(ciji '{{.Battery.TMUXColor}}{{.Battery.Percentage}}%%{{.Battery.VProgressBar}}') #[fg=cyan]%a %H:%M #(ciji '{{.MoonPhase.EMOJI}}') "
```
