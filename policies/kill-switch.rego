package fabric.kill_switch
default allow = false
allow {
    input.action == "trigger_kill_switch"
    input.blast_radius in ["user", "tenant", "app", "region", "global"]
    input.reason != ""
}
deny_unsafe_blast {
    input.blast_radius not in ["user", "tenant", "app", "region", "global"]
}
