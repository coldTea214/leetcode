// topo sort
func canFinish(numCourses int, prerequisites [][]int) bool {
	inDegree := make([]int, numCourses)
	afterCourses := make([][]int, numCourses)
	for _, courses := range prerequisites {
		inDegree[courses[0]]++
		afterCourses[courses[1]] = append(afterCourses[courses[1]], courses[0])
	}

	canLearn := make([]int, 0, numCourses)
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			canLearn = append(canLearn, i)
		}
	}

	for i := 0; i < len(canLearn); i++ {
		course := canLearn[i]
		courses := afterCourses[course]
		for _, c := range courses {
			inDegree[c]--
			if inDegree[c] == 0 {
				canLearn = append(canLearn, c)
			}
		}
	}
	return len(canLearn) == numCourses
}
