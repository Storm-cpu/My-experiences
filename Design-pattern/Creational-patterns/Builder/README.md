### Concept
Builder pattern cho phép xây dựng các đối tượng phức tạp theo từng bước.

Ví dụ: Để xây dựng một ngôi nhà đơn giản, ta cần xây dựng bốn bức tường, lát nền nhà, lắp cửa sổ, cửa ra vào và lợp mái nhà. Tuy nhiên, nếu muốn xây dựng một ngôi nhà lớn hơn, có sân sau và các tiện ích khác , việc sử dụng lớp con để extend sẽ trở nên phức tạp. Thay vào đó, Builder Pattern cho phép xây dựng ngôi nhà bằng cách sử dụng các bước như xây tường, lắp cửa, lợp mái, mà mình có thể gọi theo nhu cầu để tạo ra một cấu hình cụ thể của ngôi nhà

### Structure
Structure của Builder pattern bao gồm các thành phần:

![builder_structure](../../access/builder_structure.png)

- Director: Là một lớp hoặc đối tượng chịu trách nhiệm xây dựng đối tượng cuối cùng. Nó sử dụng một đối tượng Builder để thực hiện các bước xây dựng.
- Builder: Là một giao diện hoặc lớp trừu tượng chứa các phương thức để xây dựng từng phần của đối tượng cuối cùng. Các phương thức này có thể là buildStepA(), buildStepB(), và getResult().
- ConcreteBuilder1 và ConcreteBuilder2: Là các lớp cụ thể thực hiện giao diện Builder. Chúng xây dựng các phần khác nhau của đối tượng cuối cùng.
- Product1 và Product2: Là các lớp đại diện cho đối tượng cuối cùng. Chúng chứa các phương thức như buildStepA(), buildStepB(), getProduct1(), hoặc getProduct2().

### Example
```
type HouseBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor(numFloor int)
	setRoofType()
	setWallType()
	setHavePool()
	setHaveGarden()
	setHaveGarage()
	getHouse() House
}

func getHouseBuilder(builderType string) (HouseBuilder, error) {
	if builderType == "wooden" {
		return newWoodenBuilder(), nil
	}

	if builderType == "brick" {
		return newBrickBuilder(), nil
	}
	return nil, fmt.Errorf("invalid option")
}

type House struct {
	windowType string
	doorType   string
	floor      int
	roofType   string
	wallType   string
	havePool   bool
	haveGarden bool
	haveGarage bool
}

// ****************************************WOODEN BUILDER**************************************** //
type WoodenBuilder struct {
	windowType string
	doorType   string
	floor      int
	roofType   string
	wallType   string
	havePool   bool
	haveGarden bool
	haveGarage bool
}

func newWoodenBuilder() *WoodenBuilder {
	return &WoodenBuilder{}
}

func (b *WoodenBuilder) setWindowType() {
	b.windowType = "Wooden Window"
}

func (b *WoodenBuilder) setDoorType() {
	b.doorType = "Wooden Door"
}

func (b *WoodenBuilder) setRoofType() {
	b.roofType = "Wooden Roof"
}

func (b *WoodenBuilder) setWallType() {
	b.wallType = "Wooden Wall"
}

func (b *WoodenBuilder) setNumFloor(numFloor int) {
	b.floor = numFloor
}

func (b *WoodenBuilder) setHavePool() {
	b.havePool = true
}

func (b *WoodenBuilder) setHaveGarden() {
	b.haveGarden = true
}

func (b *WoodenBuilder) setHaveGarage() {
	b.haveGarage = true
}

func (b *WoodenBuilder) getHouse() House {
	return House{
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
		roofType:   b.roofType,
		wallType:   b.wallType,
		havePool:   b.havePool,
		haveGarden: b.haveGarden,
		haveGarage: b.haveGarage,
	}
}

// ****************************************BRICK BUILDER**************************************** //
type BrickBuilder struct {
	House
}

func newBrickBuilder() *BrickBuilder {
	return &BrickBuilder{}
}

func (b *BrickBuilder) setWindowType() {
	b.windowType = "Brick Window"
}

func (b *BrickBuilder) setDoorType() {
	b.doorType = "Brick Door"
}

func (b *BrickBuilder) setRoofType() {
	b.roofType = "Brick Roof"
}

func (b *BrickBuilder) setWallType() {
	b.wallType = "Brick Wall"
}

func (b *BrickBuilder) setNumFloor(numFloor int) {
	b.floor = numFloor
}

func (b *BrickBuilder) setHavePool() {
	b.havePool = true
}

func (b *BrickBuilder) setHaveGarden() {
	b.haveGarden = true
}

func (b *BrickBuilder) setHaveGarage() {
	b.haveGarage = true
}

func (b *BrickBuilder) getHouse() House {
	return House{
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
		roofType:   b.roofType,
		wallType:   b.wallType,
		havePool:   b.havePool,
		haveGarden: b.haveGarden,
		haveGarage: b.haveGarage,
	}
}

// ****************************************DIRECTOR**************************************** //
type Director struct {
	builder HouseBuilder
}

func newDirector(b HouseBuilder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) setBuilder(b HouseBuilder) {
	d.builder = b
}

func (d *Director) buildHouse(buildType string) (House, error) {
	if buildType == "simple" {
		d.builder.setNumFloor(1)
		d.builder.setWallType()
		d.builder.setDoorType()
		d.builder.setWindowType()
		d.builder.setRoofType()
		return d.builder.getHouse(), nil
	}
	if buildType == "fancy" {
		d.builder.setNumFloor(3)
		d.builder.setWallType()
		d.builder.setDoorType()
		d.builder.setWindowType()
		d.builder.setRoofType()
		d.builder.setHavePool()
		d.builder.setHaveGarden()
		d.builder.setHaveGarage()
		return d.builder.getHouse(), nil
	}
	return House{}, fmt.Errorf("invalid option")
}

// ****************************************MAIN**************************************** //
func main() {
	woodenBuilder, _ := getHouseBuilder("wooden")
	stoneBuilder, _ := getHouseBuilder("brick")

	director := newDirector(woodenBuilder)
	facyWoodenHouse, _ := director.buildHouse("fancy")

	fmt.Println()
	fmt.Println("====================FANCY WOODEN HOUSE====================")
	fmt.Printf("Wooden House Door Type: %s\n", facyWoodenHouse.doorType)
	fmt.Printf("Wooden House Window Type: %s\n", facyWoodenHouse.windowType)
	fmt.Printf("Wooden House Num Floor: %d\n", facyWoodenHouse.floor)
	fmt.Printf("Wooden House Wall Type: %s\n", facyWoodenHouse.wallType)
	fmt.Printf("Wooden House Roof Type: %s\n", facyWoodenHouse.roofType)
	fmt.Printf("Wooden House Have Pool: %t\n", facyWoodenHouse.havePool)
	fmt.Printf("Wooden House Have Garden: %t\n", facyWoodenHouse.haveGarden)
	fmt.Printf("Wooden House Have Garage: %t\n", facyWoodenHouse.haveGarage)

	director.setBuilder(stoneBuilder)
	simpleBrickHouse, _ := director.buildHouse("simple")

	fmt.Println()
	fmt.Println("====================SIMPLE BRICK HOUSE====================")
	fmt.Printf("Brick House Door Type: %s\n", simpleBrickHouse.doorType)
	fmt.Printf("Brick House Window Type: %s\n", simpleBrickHouse.windowType)
	fmt.Printf("Brick House Num Floor: %d\n", simpleBrickHouse.floor)
	fmt.Printf("Brick House Wall Type: %s\n", simpleBrickHouse.wallType)
	fmt.Printf("Brick House Roof Type: %s\n", simpleBrickHouse.roofType)
	fmt.Printf("Brick House Have Pool: %t\n", simpleBrickHouse.havePool)
	fmt.Printf("Brick House Have Garden: %t\n", simpleBrickHouse.haveGarden)
	fmt.Printf("Brick House Have Garage: %t\n", simpleBrickHouse.haveGarage)

}
```

### Applicability
Builder pattern thường được áp dụng khi:

- Khi cần tạo ra một đối tượng có nhiều thành phần hoặc thuộc tính phức tạp.
- Trì hoãn việc thực hiện các bước một số bước builder nhất định

### Pros and Cons
Ưu điểm của Builder:

- Tách biệt logic xây dựng
- Tùy chỉnh linh hoạt: có thể tùy chỉnh từng bước xây dựng để tạo ra các biểu diễn khác nhau của đối tượng.
- Giảm sự phức tạp khi khởi tạo đối tượng

Nhược điểm của Builder:

- Tăng số lượng class do phải tạo thêm các lớp Builder và Product
- Không phù hợp cho đối tượng đơn giản
