package context

type ctxKey int

const ridKey ctxKey = ctxKey(0)
const claimsKey ctxKey = ctxKey(1)
const timingKey ctxKey = ctxKey(2)
