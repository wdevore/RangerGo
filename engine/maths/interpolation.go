package maths

// Lerp returns a the value between min and max given t = 0->1
// func Lerp(min::T, max::T, t::T) where {T <: AbstractFloat}
//     min * (1.0 - t) + max * t
// end

// As per:
// https://gamedev.stackexchange.com/questions/18615/how-do-i-linearly-interpolate-between-two-vectors
// func Lerp(min::Math.Vector2D{T}, max::Math.Vector2D{T}, t::T, out::Math.Vector2D{T}) where {T <: AbstractFloat}
//     Math.scale!(min, (1.0 - t), v1)
//     Math.scale!(max, t, v2)
//     Math.add!(v1, v2, out)
// end

// TODO need to review for negative ranges.
// `linear` returns 0->1 for a "value" between min and max.
// Generally used to map from view-space to unit-space
// func Linear(min::T, max::T, value::T) where {T <: AbstractFloat}
//     if min < 0.0
//         (value - max) / (min - max)
//     else
//         (value - min) / (max - min)
//     end
// end
