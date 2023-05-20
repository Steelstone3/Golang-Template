package world

import "github.com/Steelstone3/Star-Trek-Explorers/systems/generators"

type PlanetName struct {
	PlanetName string
}

func constructPlanetName(seed int64) PlanetName {
	planetNames := []string{
		"51 Pegasi b",
		"Proxima Centauri b",
		"Trappist-1 system (Trappist-1b to Trappist-1h)",
		"Kepler-452b",
		"Kepler-186f",
		"Kepler-10b",
		"Kepler-16b",
		"Kepler-22b",
		"Kepler-62e",
		"Kepler-62f",
		"Kepler-69c",
		"HD 189733b",
		"HD 209458b (Osiris)",
		"HD 219134b",
		"HD 40307g",
		"Gliese 436 b",
		"Gliese 1214 b",
		"Gliese 876 d",
		"Gliese 667Cc",
		"GJ 1214 b",
		"WASP-17b",
		"WASP-12b",
		"WASP-33b",
		"WASP-43b",
		"WASP-46b",
		"WASP-80b",
		"WASP-94Ab",
		"WASP-121b",
		"WASP-127b",
		"HAT-P-1b",
		"HAT-P-2b",
		"HAT-P-7b",
		"HAT-P-11b",
		"HAT-P-26b",
		"HAT-P-32b",
		"HAT-P-44b",
		"HAT-P-46b",
		"HAT-P-67b",
		"HAT-P-69b",
		"HAT-P-76b",
		"HAT-P-99b",
		"HAT-P-108b",
		"HAT-P-113b",
		"HAT-P-118b",
		"HAT-P-122b",
		"HAT-P-124b",
		"HAT-P-130b",
		"HAT-P-131b",
		"HAT-P-132b",
		"HAT-P-133b",
		"TOI-178b",
		"TOI-700d",
		"LHS 1140b",
		"Kepler-62f",
		"Kepler-442b",
		"Kepler-452b",
		"Kepler-1652b",
		"Kepler-283c",
		"Kepler-1649c",
		"Kepler-174d",
		"Kepler-61b",
		"Kepler-1090b",
		"Kepler-1060b",
		"Kepler-186f",
		"Kepler-438b",
		"Kepler-443b",
		"Kepler-1638b",
		"Kepler-452b",
		"Kepler-1551b",
		"Kepler-1653b",
		"Kepler-1649b",
		"Kepler-1600b",
		"Kepler-1537b",
		"Kepler-1697b",
		"Kepler-1589b",
		"Kepler-1540b",
		"Kepler-1507b",
		"Kepler-1651b",
		"Kepler-1604b",
		"Kepler-1704b",
		"Kepler-1708b",
		"Kepler-1710b",
		"Kepler-1720b",
		"Kepler-1675b",
		"Kepler-1677b",
		"Kepler-1743b",
		"Kepler-1792b",
		"Kepler-1805b",
		"Kepler-1794b",
		"Kepler-1777b",
		"Kepler-1796b",
		"Kepler-1810b",
		"Kepler-1864b",
		"Kepler-1875b",
		"Kepler-1903b",
		"Kepler-1904b",
		"Kepler-1954b",
		"Kepler-1965b",
		"Kepler-1958b",
		"Kepler-1966b",
		"Kepler-20b",
		"Kepler-20c",
		"Kepler-20d",
		"Kepler-20e",
		"Kepler-20f",
		"Kepler-21b",
		"Kepler-21c",
		"Kepler-22b",
		"Kepler-23b",
		"Kepler-23c",
		"Kepler-24b",
		"Kepler-24c",
		"Kepler-25b",
		"Kepler-25c",
		"Kepler-26b",
		"Kepler-26c",
		"Kepler-27b",
		"Kepler-27c",
		"Kepler-27d",
		"Kepler-28b",
		"Kepler-28c",
		"Kepler-29b",
		"Kepler-29c",
		"Kepler-30b",
		"Kepler-30c",
		"Kepler-30d",
		"Kepler-31b",
		"Kepler-31c",
		"Kepler-32b",
		"Kepler-32c",
		"Kepler-32d",
		"Kepler-33b",
		"Kepler-33c",
		"Kepler-33d",
		"Kepler-33e",
		"Kepler-33f",
		"Kepler-34b",
		"Kepler-34c",
		"Kepler-35b",
		"Kepler-35c",
		"Kepler-36b",
		"Kepler-36c",
		"Kepler-37b",
		"Kepler-37c",
		"Kepler-37d",
		"Kepler-37e",
		"Kepler-38b",
		"Kepler-39b",
		"Kepler-39c",
		"Kepler-39d",
	}

	planetName := generators.GetRandomString(seed, planetNames)

	return PlanetName{
		PlanetName: planetName,
	}
}
