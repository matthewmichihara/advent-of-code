import Node.Composite
import Node.Literal
import kotlin.math.ceil
import kotlin.math.max

sealed interface Node {
    var parent: Composite?

    fun isLeftChild(): Boolean {
        return parent?.left == this
    }

    fun isRightChild(): Boolean {
        return parent?.right == this
    }

    class Composite(
        var left: Node,
        var right: Node,
    ) : Node {
        override var parent: Composite? = null
        override fun toString() = "Composite(left=$left, right=$right, parent=${parent != null})"
    }

    class Literal(
        var value: Int,
    ) : Node {
        override var parent: Composite? = null
        override fun toString() = "Literal(value=$value, parent=${parent != null})"
    }
}

fun main() {
    val input = """
        [[[3,[8,6]],[6,1]],[[[1,1],2],[[1,0],0]]]
        [[[1,[7,3]],1],9]
        [[[2,6],[[3,1],[0,9]]],[[7,[4,8]],[[2,7],3]]]
        [[[3,[0,4]],[[8,4],[1,9]]],[7,[2,[5,7]]]]
        [[[4,5],[[0,7],1]],[9,[0,4]]]
        [[5,[[1,5],[3,6]]],8]
        [[3,[[9,3],9]],9]
        [2,[[[2,1],[0,5]],[9,9]]]
        [[2,[6,9]],[[[4,1],0],[3,4]]]
        [[[[6,8],0],[[8,8],9]],[[[4,2],3],[3,[7,3]]]]
        [[3,7],9]
        [[[[2,5],8],[2,5]],[[0,[5,7]],[[2,5],4]]]
        [[[8,[6,6]],0],[4,[[5,6],[8,4]]]]
        [[[1,[8,2]],[[0,4],[2,6]]],[[3,4],0]]
        [[1,[[9,2],[6,0]]],[[[0,9],5],[[8,0],[1,5]]]]
        [[2,[[2,3],[1,8]]],[3,[[7,2],[0,7]]]]
        [[5,4],5]
        [[[[4,2],[4,8]],[7,3]],[0,[[8,9],6]]]
        [[[6,7],0],5]
        [[2,[[9,0],[8,4]]],[[[7,4],[3,4]],0]]
        [[[9,[8,9]],1],[[5,[6,7]],3]]
        [[2,[0,0]],[3,[[2,5],[1,4]]]]
        [[0,1],[0,[[8,8],[8,3]]]]
        [[[0,2],[2,8]],[1,[[7,0],0]]]
        [[[[5,4],3],[[7,5],[2,6]]],[[5,8],[0,1]]]
        [0,[0,0]]
        [[5,[[5,6],0]],[[[2,7],9],[7,9]]]
        [[[[0,8],2],[[2,5],[7,6]]],[[9,7],[[8,7],[9,2]]]]
        [[[0,[4,6]],[[6,3],[4,4]]],[8,[[4,8],[4,8]]]]
        [[[[8,9],[3,8]],8],[[[7,9],6],[9,[2,7]]]]
        [[[[8,9],[1,6]],0],[[[8,7],4],[9,[1,4]]]]
        [5,7]
        [[[[1,5],[3,6]],[[5,5],4]],[[3,3],[4,[4,0]]]]
        [[[0,6],[5,[5,3]]],[[4,[0,0]],8]]
        [7,[6,8]]
        [[[[8,5],9],[[3,2],7]],[[[6,6],5],2]]
        [[[[4,4],[0,4]],9],0]
        [[0,[3,[9,3]]],[9,[[8,0],[0,9]]]]
        [[[[4,0],0],[1,[1,7]]],[[3,[3,0]],[[1,3],6]]]
        [[9,4],[3,[[7,1],6]]]
        [[[[3,7],7],1],[[4,3],[[6,9],[6,9]]]]
        [[[8,[2,5]],[[8,4],4]],[[[3,4],[6,7]],[5,[8,5]]]]
        [2,[4,[[3,2],7]]]
        [[[[3,1],[5,6]],[[2,7],7]],[4,[8,[7,4]]]]
        [[7,8],[[[3,9],7],2]]
        [[[[8,8],[5,8]],[[1,0],[6,0]]],[[[1,2],6],[[4,2],[5,5]]]]
        [[1,[0,9]],[[[2,1],1],1]]
        [[6,[8,1]],[4,[[7,8],5]]]
        [[[1,[1,6]],[1,[5,7]]],[[[2,8],6],0]]
        [9,1]
        [[[0,[6,5]],[[8,5],2]],[[[2,4],[7,3]],[[1,5],[9,2]]]]
        [[[2,7],[0,[3,6]]],[[[1,0],[9,6]],[1,[0,4]]]]
        [6,[[[5,9],8],[0,2]]]
        [7,[[[9,4],[8,6]],[[1,1],1]]]
        [[[2,1],0],8]
        [1,[[6,[1,4]],[[0,0],[1,9]]]]
        [[[1,[7,9]],2],8]
        [[[[0,9],2],[[8,4],9]],[0,[[7,7],[4,8]]]]
        [[1,[2,[1,8]]],[[[3,6],[2,1]],[3,[5,0]]]]
        [[3,3],[3,5]]
        [[[[9,3],[4,3]],[5,[8,1]]],[[6,[5,0]],9]]
        [0,[[9,[3,5]],3]]
        [[[9,1],0],[[[5,9],[8,0]],[7,[4,8]]]]
        [[[[7,7],8],3],[[[6,6],[6,5]],[6,4]]]
        [[[[3,7],1],[9,[4,2]]],[[9,[2,5]],[[9,0],5]]]
        [5,[[0,2],6]]
        [[[[2,7],[5,3]],[1,8]],2]
        [[[8,[7,7]],[9,[0,0]]],4]
        [[[4,[1,4]],0],[[[8,7],8],[[4,1],7]]]
        [[[[0,6],0],[[3,2],[9,8]]],[[9,[4,5]],[[7,7],[0,8]]]]
        [[[[6,3],3],[[1,5],7]],[[0,1],[7,7]]]
        [[[[2,0],2],[3,[3,5]]],[[[0,8],[8,2]],[[0,6],5]]]
        [[[6,[5,3]],[[5,5],9]],[[5,9],[[8,7],[3,7]]]]
        [[[[1,7],[3,4]],[9,2]],1]
        [[[[8,2],6],1],[[5,[2,7]],[3,9]]]
        [5,[5,7]]
        [[[[9,8],[3,4]],[[2,5],[5,6]]],[[[2,7],7],[9,[8,7]]]]
        [[[1,4],[[6,1],[1,3]]],[1,[7,[1,7]]]]
        [[[[1,4],8],[[5,1],8]],[[[1,3],[6,9]],[6,[3,3]]]]
        [[[[4,0],[0,7]],[4,5]],[4,2]]
        [3,8]
        [7,[[[7,6],5],[[6,6],5]]]
        [[[5,[0,5]],[4,4]],[3,[[4,2],[7,0]]]]
        [[[[7,9],8],[9,6]],[5,0]]
        [[[[3,0],[5,2]],1],[[[6,9],[5,3]],[[2,5],[6,3]]]]
        [7,[[[7,7],[4,5]],[9,2]]]
        [[7,[[4,2],[9,3]]],[7,[6,1]]]
        [7,9]
        [[[8,[8,1]],[[7,3],1]],[[9,8],[2,[8,3]]]]
        [[[9,3],3],3]
        [[[8,[5,7]],[[2,1],[1,3]]],[[[3,5],2],0]]
        [[[8,8],0],[[1,4],[[8,6],9]]]
        [[9,[3,[3,0]]],[1,7]]
        [1,[[[8,8],1],[2,[0,5]]]]
        [[0,[1,5]],[9,[0,[9,0]]]]
        [1,[[[1,1],[8,3]],[1,8]]]
        [[5,[[7,7],[3,3]]],[[[6,6],[7,8]],[1,[0,0]]]]
        [[[[6,7],1],[0,2]],[[[4,2],[7,6]],[[8,4],[4,9]]]]
        [[6,[[3,3],[9,0]]],[1,[[4,5],4]]]
        [[[[3,4],7],[9,0]],[[[4,5],1],[[5,1],[9,3]]]]
    """.trimIndent()

    part1(input)
    part2(input)
}

fun part1(input: String) {
    val numbers = input.lines().map { parse(it).node }
    var number = numbers[0]
    for (n in numbers.drop(1)) {
        number = add(number, n)
        reduce(number)
    }
    println(magnitude(number))
}

fun part2(input: String) {
    val numbers = input.lines()
    var maxMagnitude = 0
    for ((i, a) in numbers.withIndex()) {
        for ((j, b) in numbers.withIndex()) {
            if (i == j) continue
            val (nodeA, _) = parse(a)
            val (nodeB, _) = parse(b)
            val nodeC = add(nodeA, nodeB)
            reduce(nodeC)
            maxMagnitude = max(magnitude(nodeC), maxMagnitude)
        }
    }
    println(maxMagnitude)
}

fun magnitude(node: Node): Int = when (node) {
    is Literal -> node.value
    is Composite -> 3 * magnitude(node.left) + 2 * magnitude(node.right)
}

fun reduce(node: Node) {
    if (node is Literal) return
    val composite = node as Composite

    while (true) {
        val needsExplode = findDeepComposite(composite, 4)
        if (needsExplode != null) {
            explode(needsExplode)
            continue
        }

        val needsSplit = findLiteralToSplit(composite)
        if (needsSplit != null) {
            split(needsSplit)
            continue
        }

        break
    }
}

fun explode(composite: Composite) {
    // Root node, nothing to do.
    val parent = composite.parent ?: return

    val leftLiteral = findLeftLiteral(composite)
    if (leftLiteral != null) {
        // assuming composite has two literals
        leftLiteral.value += (composite.left as Literal).value
    }
    val rightLiteral = findRightLiteral(composite)
    if (rightLiteral != null) {
        rightLiteral.value += (composite.right as Literal).value
    }

    if (composite.isLeftChild()) {
        parent.left = Literal(0)
        parent.left.parent = parent
    } else if (composite.isRightChild()) {
        parent.right = Literal(0)
        parent.right.parent = parent
    }
}

fun split(literal: Literal) {
    val parent = literal.parent!!
    val left = Literal(literal.value / 2)
    val right = Literal(ceil(literal.value.toDouble() / 2).toInt())
    val replacement = Composite(left, right)
    replacement.parent = parent

    left.parent = replacement
    right.parent = replacement

    if (literal.isLeftChild()) {
        parent.left = replacement
    } else if (literal.isRightChild()) {
        parent.right = replacement
    }
}

fun findLeftLiteral(node: Composite): Literal? {
    var n: Composite = node
    while (n.parent != null && n.isLeftChild()) {
        n = n.parent!!
    }

    if (n.parent == null) return null

    var left = n.parent!!.left
    while (left is Composite) {
        left = left.right
    }

    return left as Literal
}

fun findRightLiteral(node: Composite): Literal? {
    var n: Composite = node
    while (n.parent != null && n.isRightChild()) {
        n = n.parent!!
    }

    if (n.parent == null) return null

    var right = n.parent!!.right
    while (right is Composite) {
        right = right.left
    }

    return right as Literal
}

fun findDeepComposite(node: Node, depth: Int): Composite? {
    when (node) {
        is Literal -> return null
        is Composite -> {
            if (depth <= 0) return node
            val leftRes = findDeepComposite(node.left, depth - 1)
            if (leftRes != null) return leftRes

            val rightRes = findDeepComposite(node.right, depth - 1)
            if (rightRes != null) return rightRes

            return null
        }
    }
}

fun findLiteralToSplit(node: Node): Literal? {
    when (node) {
        is Literal -> {
            if (node.value >= 10) return node
            return null
        }
        is Composite -> {
            val left = findLiteralToSplit(node.left)
            if (left != null) return left

            val right = findLiteralToSplit(node.right)
            if (right != null) return right

            return null
        }
    }
}

data class ParseResult(
    val node: Node, val offset: Int
)

fun parse(s: String): ParseResult {
    var i = 0
    when (s[i]) {
        '[' -> {
            i += 1
            val (left, offset1) = parse(s.substring(i))
            i += offset1 + 1
            val (right, offset2) = parse(s.substring(i))
            i += offset2 + 1
            val node = Composite(left, right)
            left.parent = node
            right.parent = node
            return ParseResult(node, i)
        }
        else -> {
            var digitBuffer = ""
            while (s[i] != ',' && s[i] != ']') {
                digitBuffer += s[i]
                i += 1
            }
            val node = Literal(digitBuffer.toInt())
            return ParseResult(node, i)
        }
    }
}

@Suppress("unused")
fun serialize(node: Node): String = when (node) {
    is Literal -> node.value.toString()
    is Composite -> "[${serialize(node.left)},${serialize(node.right)}]"
}

fun add(a: Node, b: Node): Node {
    val node = Composite(left = a, right = b)
    a.parent = node
    b.parent = node
    return node
}