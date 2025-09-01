# VkPrimitiveTopology(3) Manual Page

## Name

VkPrimitiveTopology - Supported primitive topologies



## [](#_c_specification)C Specification

The primitive topologies defined by [VkPrimitiveTopology](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPrimitiveTopology.html) are:

```c++
// Provided by VK_VERSION_1_0
typedef enum VkPrimitiveTopology {
    VK_PRIMITIVE_TOPOLOGY_POINT_LIST = 0,
    VK_PRIMITIVE_TOPOLOGY_LINE_LIST = 1,
    VK_PRIMITIVE_TOPOLOGY_LINE_STRIP = 2,
    VK_PRIMITIVE_TOPOLOGY_TRIANGLE_LIST = 3,
    VK_PRIMITIVE_TOPOLOGY_TRIANGLE_STRIP = 4,
    VK_PRIMITIVE_TOPOLOGY_TRIANGLE_FAN = 5,
    VK_PRIMITIVE_TOPOLOGY_LINE_LIST_WITH_ADJACENCY = 6,
    VK_PRIMITIVE_TOPOLOGY_LINE_STRIP_WITH_ADJACENCY = 7,
    VK_PRIMITIVE_TOPOLOGY_TRIANGLE_LIST_WITH_ADJACENCY = 8,
    VK_PRIMITIVE_TOPOLOGY_TRIANGLE_STRIP_WITH_ADJACENCY = 9,
    VK_PRIMITIVE_TOPOLOGY_PATCH_LIST = 10,
} VkPrimitiveTopology;
```

## [](#_description)Description

- `VK_PRIMITIVE_TOPOLOGY_POINT_LIST` specifies a series of [separate point primitives](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#drawing-point-lists).
- `VK_PRIMITIVE_TOPOLOGY_LINE_LIST` specifies a series of [separate line primitives](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#drawing-line-lists).
- `VK_PRIMITIVE_TOPOLOGY_LINE_STRIP` specifies a series of [connected line primitives](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#drawing-line-strips) with consecutive lines sharing a vertex.
- `VK_PRIMITIVE_TOPOLOGY_TRIANGLE_LIST` specifies a series of [separate triangle primitives](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#drawing-triangle-lists).
- `VK_PRIMITIVE_TOPOLOGY_TRIANGLE_STRIP` specifies a series of [connected triangle primitives](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#drawing-triangle-strips) with consecutive triangles sharing an edge.
- `VK_PRIMITIVE_TOPOLOGY_TRIANGLE_FAN` specifies a series of [connected triangle primitives](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#drawing-triangle-fans) with all triangles sharing a common vertex. If the `VK_KHR_portability_subset` extension is enabled, and [VkPhysicalDevicePortabilitySubsetFeaturesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevicePortabilitySubsetFeaturesKHR.html)::`triangleFans` is `VK_FALSE`, then triangle fans are not supported by the implementation, and `VK_PRIMITIVE_TOPOLOGY_TRIANGLE_FAN` **must** not be used.
- `VK_PRIMITIVE_TOPOLOGY_LINE_LIST_WITH_ADJACENCY` specifies a series of [separate line primitives with adjacency](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#drawing-line-lists-with-adjacency).
- `VK_PRIMITIVE_TOPOLOGY_LINE_STRIP_WITH_ADJACENCY` specifies a series of [connected line primitives with adjacency](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#drawing-line-strips-with-adjacency), with consecutive primitives sharing three vertices.
- `VK_PRIMITIVE_TOPOLOGY_TRIANGLE_LIST_WITH_ADJACENCY` specifies a series of [separate triangle primitives with adjacency](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#drawing-triangle-lists-with-adjacency).
- `VK_PRIMITIVE_TOPOLOGY_TRIANGLE_STRIP_WITH_ADJACENCY` specifies [connected triangle primitives with adjacency](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#drawing-triangle-strips-with-adjacency), with consecutive triangles sharing an edge.
- `VK_PRIMITIVE_TOPOLOGY_PATCH_LIST` specifies [separate patch primitives](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#drawing-patch-lists).

Each primitive topology, and its construction from a list of vertices, is described in detail below with a supporting diagram, according to the following key:

image/svg+xml

Vertex

A point in 3-dimensional space. Positions chosen within the diagrams are arbitrary and for illustration only.

image/svg+xml 5

Vertex Number

Sequence position of a vertex within the provided vertex data.

image/svg+xml

Provoking Vertex

Provoking vertex within the main primitive. The tail is angled towards the relevant primitive. Used in [flat shading](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vertexpostproc-flatshading).

image/svg+xml

Primitive Edge

An edge connecting the points of a main primitive.

image/svg+xml

Adjacency Edge

Points connected by these lines do not contribute to a main primitive, and are only accessible in a [geometry shader](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#geometry).

image/svg+xml

Winding Order

The relative order in which vertices are defined within a primitive, used in the [facing determination](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#primsrast-polygons-basic). This ordering has no specific start or end point.

The diagrams are supported with mathematical definitions where the vertices (v) and primitives (p) are numbered starting from 0; v0 is the first vertex in the provided data and p0 is the first primitive in the set of primitives defined by the vertices and topology.

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkPipelineInputAssemblyStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineInputAssemblyStateCreateInfo.html), [vkCmdSetPrimitiveTopology](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetPrimitiveTopology.html), [vkCmdSetPrimitiveTopology](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetPrimitiveTopology.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPrimitiveTopology).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0