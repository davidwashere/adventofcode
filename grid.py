class Grid:
    def __init__(self, default=0, size=100, step=100):
        """
        default : object, optional
            The initial value for each grid item [default=0]

        size : int, optional
            The initial size of the grid in all directions from 0, 0, which results in grid being size * 4 [default=100]
        
        step : int, optional
            Increase size of grid quandrant by this amount each time a value is encountered beyond the quandrant's extents [default=100]
        """
        self.default = default
        self.size = size
        self.step = step

        self.grid_posX_posY = [[default for i in range(size)] for j in range(size)]
        self.grid_posX_negY = [[default for i in range(size)] for j in range(size)]
        self.grid_negX_posY = [[default for i in range(size)] for j in range(size)]
        self.grid_negX_negY = [[default for i in range(size)] for j in range(size)]


    def get(self, x, y):
        return self._get_grid_from_coord(x, y)

    def set(self, x, y, value):
        grid = self._get_grid_from_coord(x, y)

        grid[abs(x)][abs(y)] = value

    def _get_grid_from_coord(self, x, y):
        if abs(x) > self.size or abs(y) > self.size:
            raise Exception(f"{x},{y} is Out of Bounds, current max: {self.size}")

        if x < 0 and y < 0:
            return self.grid_negX_negY
        elif x >= 0 and y < 0:
            return self.grid_posX_negY
        elif x < 0 and y >= 0:
            return self.grid_negX_posY
        elif x >= 0 and y >= 0:
            return self.grid_posX_posY

    def __str__(self):
        output = ""

        tl = self.grid_negX_posY
        tr = self.grid_posX_posY
        bl = self.grid_negX_negY
        br = self.grid_posX_negY

        for row in tl:
            for col in row:
                output += f'{col}, '

            output += "\n"

        return output