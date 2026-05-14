package jump_game

func canJump(nums []int) bool {
	maxJump := 0

	for i := 0; i < len(nums); i++ {
		if i > maxJump {
			return false
		}

		if i+nums[i] > maxJump {
			maxJump = i + nums[i]
		}
	}

	return true
}
