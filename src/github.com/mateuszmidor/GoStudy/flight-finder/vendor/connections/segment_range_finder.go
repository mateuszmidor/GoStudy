package connections

import (
	"airport"
	"segment"
)

type SegmentRangeFinder struct {
}

func (s *SegmentRangeFinder) ByFromAirport(segments segment.Segments, id airport.AirportID) (first, last segment.ID) {
	return segment.ID(lowerBound(segments, id)), segment.ID(upperBound(segments, id))
}

/*
template<class ForwardIt, class T>
ForwardIt lower_bound(ForwardIt first, ForwardIt last, const T& value)
{
    ForwardIt it;
    typename std::iterator_traits<ForwardIt>::difference_type count, step;
    count = std::distance(first, last);

    while (count > 0) {
        it = first;
        step = count / 2;
        std::advance(it, step);
        if (*it < value) {
            first = ++it;
            count -= step + 1;
        }
        else
            count = step;
    }
    return first;
}
*/
/*
template<class ForwardIt, class T>
ForwardIt upper_bound(ForwardIt first, ForwardIt last, const T& value)
{
    ForwardIt it;
    typename std::iterator_traits<ForwardIt>::difference_type count, step;
    count = std::distance(first, last);

    while (count > 0) {
        it = first;
        step = count / 2;
        std::advance(it, step);
        if (!(value < *it)) {
            first = ++it;
            count -= step + 1;
        }
        else
            count = step;
    }
    return first;
}
*/

// lowerBound is index of first matching element
func lowerBound(segments segment.Segments, id airport.AirportID) int {
	first := 0
	last := len(segments)
	count := last - first

	for count > 0 {
		i := first
		step := count / 2
		i += step
		if segments[i].From() < id {
			first = i + 1
			count -= step + 1
		} else {
			count = step
		}
	}
	return first
}

// upperBound is index of last matching element +1
func upperBound(segments segment.Segments, id airport.AirportID) int {
	first := 0
	last := len(segments)
	count := last - first

	for count > 0 {
		i := first
		step := count / 2
		i += step
		if segments[i].From() <= id {
			first = i + 1
			count -= step + 1
		} else {
			count = step
		}
	}
	return first
}
