# VkAccelerationStructureGeometryTrianglesDataKHR(3) Manual Page

## Name

VkAccelerationStructureGeometryTrianglesDataKHR - Structure specifying a
triangle geometry in a bottom-level acceleration structure



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkAccelerationStructureGeometryTrianglesDataKHR` structure is
defined as:

``` c
// Provided by VK_KHR_acceleration_structure
typedef struct VkAccelerationStructureGeometryTrianglesDataKHR {
    VkStructureType                  sType;
    const void*                      pNext;
    VkFormat                         vertexFormat;
    VkDeviceOrHostAddressConstKHR    vertexData;
    VkDeviceSize                     vertexStride;
    uint32_t                         maxVertex;
    VkIndexType                      indexType;
    VkDeviceOrHostAddressConstKHR    indexData;
    VkDeviceOrHostAddressConstKHR    transformData;
} VkAccelerationStructureGeometryTrianglesDataKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `vertexFormat` is the [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) of each vertex
  element.

- `vertexData` is a device or host address to memory containing vertex
  data for this geometry.

- `maxVertex` is the number of vertices in `vertexData` minus one.

- `vertexStride` is the stride in bytes between each vertex.

- `indexType` is the [VkIndexType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkIndexType.html) of each index
  element.

- `indexData` is a device or host address to memory containing index
  data for this geometry.

- `transformData` is a device or host address to memory containing an
  optional reference to a
  [VkTransformMatrixKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkTransformMatrixKHR.html) structure describing
  a transformation from the space in which the vertices in this geometry
  are described to the space in which the acceleration structure is
  defined.

## <a href="#_description" class="anchor"></a>Description

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>Unlike the stride for vertex buffers in <a
href="VkVertexInputBindingDescription.html">VkVertexInputBindingDescription</a>
for graphics pipelines which must not exceed
<code>maxVertexInputBindingStride</code>, <code>vertexStride</code> for
acceleration structure geometry is instead restricted to being a 32-bit
value.</p></td>
</tr>
</tbody>
</table>

Valid Usage

- <a
  href="#VUID-VkAccelerationStructureGeometryTrianglesDataKHR-vertexStride-03735"
  id="VUID-VkAccelerationStructureGeometryTrianglesDataKHR-vertexStride-03735"></a>
  VUID-VkAccelerationStructureGeometryTrianglesDataKHR-vertexStride-03735  
  `vertexStride` **must** be a multiple of the size in bytes of the
  smallest component of `vertexFormat`

- <a
  href="#VUID-VkAccelerationStructureGeometryTrianglesDataKHR-vertexStride-03819"
  id="VUID-VkAccelerationStructureGeometryTrianglesDataKHR-vertexStride-03819"></a>
  VUID-VkAccelerationStructureGeometryTrianglesDataKHR-vertexStride-03819  
  `vertexStride` **must** be less than or equal to 2<sup>32</sup>-1

- <a
  href="#VUID-VkAccelerationStructureGeometryTrianglesDataKHR-vertexFormat-03797"
  id="VUID-VkAccelerationStructureGeometryTrianglesDataKHR-vertexFormat-03797"></a>
  VUID-VkAccelerationStructureGeometryTrianglesDataKHR-vertexFormat-03797  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#resources-buffer-view-format-features"
  target="_blank" rel="noopener">format features</a> of `vertexFormat`
  **must** contain
  `VK_FORMAT_FEATURE_ACCELERATION_STRUCTURE_VERTEX_BUFFER_BIT_KHR`

- <a
  href="#VUID-VkAccelerationStructureGeometryTrianglesDataKHR-indexType-03798"
  id="VUID-VkAccelerationStructureGeometryTrianglesDataKHR-indexType-03798"></a>
  VUID-VkAccelerationStructureGeometryTrianglesDataKHR-indexType-03798  
  `indexType` **must** be `VK_INDEX_TYPE_UINT16`,
  `VK_INDEX_TYPE_UINT32`, or `VK_INDEX_TYPE_NONE_KHR`

Valid Usage (Implicit)

- <a
  href="#VUID-VkAccelerationStructureGeometryTrianglesDataKHR-sType-sType"
  id="VUID-VkAccelerationStructureGeometryTrianglesDataKHR-sType-sType"></a>
  VUID-VkAccelerationStructureGeometryTrianglesDataKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_GEOMETRY_TRIANGLES_DATA_KHR`

- <a
  href="#VUID-VkAccelerationStructureGeometryTrianglesDataKHR-pNext-pNext"
  id="VUID-VkAccelerationStructureGeometryTrianglesDataKHR-pNext-pNext"></a>
  VUID-VkAccelerationStructureGeometryTrianglesDataKHR-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the
  `pNext` chain **must** be either `NULL` or a pointer to a valid
  instance of
  [VkAccelerationStructureGeometryMotionTrianglesDataNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureGeometryMotionTrianglesDataNV.html),
  [VkAccelerationStructureTrianglesDisplacementMicromapNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureTrianglesDisplacementMicromapNV.html),
  or
  [VkAccelerationStructureTrianglesOpacityMicromapEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureTrianglesOpacityMicromapEXT.html)

- <a
  href="#VUID-VkAccelerationStructureGeometryTrianglesDataKHR-sType-unique"
  id="VUID-VkAccelerationStructureGeometryTrianglesDataKHR-sType-unique"></a>
  VUID-VkAccelerationStructureGeometryTrianglesDataKHR-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique

- <a
  href="#VUID-VkAccelerationStructureGeometryTrianglesDataKHR-vertexFormat-parameter"
  id="VUID-VkAccelerationStructureGeometryTrianglesDataKHR-vertexFormat-parameter"></a>
  VUID-VkAccelerationStructureGeometryTrianglesDataKHR-vertexFormat-parameter  
  `vertexFormat` **must** be a valid [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) value

- <a
  href="#VUID-VkAccelerationStructureGeometryTrianglesDataKHR-indexType-parameter"
  id="VUID-VkAccelerationStructureGeometryTrianglesDataKHR-indexType-parameter"></a>
  VUID-VkAccelerationStructureGeometryTrianglesDataKHR-indexType-parameter  
  `indexType` **must** be a valid [VkIndexType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkIndexType.html) value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_acceleration_structure](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_acceleration_structure.html),
[VkAccelerationStructureGeometryDataKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureGeometryDataKHR.html),
[VkDeviceOrHostAddressConstKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceOrHostAddressConstKHR.html),
[VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html), [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html),
[VkIndexType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkIndexType.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkAccelerationStructureGeometryTrianglesDataKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
