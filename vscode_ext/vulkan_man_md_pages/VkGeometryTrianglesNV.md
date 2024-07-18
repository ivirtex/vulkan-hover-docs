# VkGeometryTrianglesNV(3) Manual Page

## Name

VkGeometryTrianglesNV - Structure specifying a triangle geometry in a
bottom-level acceleration structure



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkGeometryTrianglesNV` structure specifies triangle geometry in a
bottom-level acceleration structure and is defined as:

``` c
// Provided by VK_NV_ray_tracing
typedef struct VkGeometryTrianglesNV {
    VkStructureType    sType;
    const void*        pNext;
    VkBuffer           vertexData;
    VkDeviceSize       vertexOffset;
    uint32_t           vertexCount;
    VkDeviceSize       vertexStride;
    VkFormat           vertexFormat;
    VkBuffer           indexData;
    VkDeviceSize       indexOffset;
    uint32_t           indexCount;
    VkIndexType        indexType;
    VkBuffer           transformData;
    VkDeviceSize       transformOffset;
} VkGeometryTrianglesNV;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `vertexData` is the buffer containing vertex data for this geometry.

- `vertexOffset` is the offset in bytes within `vertexData` containing
  vertex data for this geometry.

- `vertexCount` is the number of valid vertices.

- `vertexStride` is the stride in bytes between each vertex.

- `vertexFormat` is a [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) describing the format of
  each vertex element.

- `indexData` is the buffer containing index data for this geometry.

- `indexOffset` is the offset in bytes within `indexData` containing
  index data for this geometry.

- `indexCount` is the number of indices to include in this geometry.

- `indexType` is a [VkIndexType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkIndexType.html) describing the format
  of each index.

- `transformData` is an optional buffer containing an
  [VkTransformMatrixNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkTransformMatrixNV.html) structure defining a
  transformation to be applied to this geometry.

- `transformOffset` is the offset in bytes in `transformData` of the
  transform information described above.

## <a href="#_description" class="anchor"></a>Description

If `indexType` is `VK_INDEX_TYPE_NONE_NV`, then this structure describes
a set of triangles determined by `vertexCount`. Otherwise, this
structure describes a set of indexed triangles determined by
`indexCount`.

Valid Usage

- <a href="#VUID-VkGeometryTrianglesNV-vertexOffset-02428"
  id="VUID-VkGeometryTrianglesNV-vertexOffset-02428"></a>
  VUID-VkGeometryTrianglesNV-vertexOffset-02428  
  `vertexOffset` **must** be less than the size of `vertexData`

- <a href="#VUID-VkGeometryTrianglesNV-vertexOffset-02429"
  id="VUID-VkGeometryTrianglesNV-vertexOffset-02429"></a>
  VUID-VkGeometryTrianglesNV-vertexOffset-02429  
  `vertexOffset` **must** be a multiple of the component size of
  `vertexFormat`

- <a href="#VUID-VkGeometryTrianglesNV-vertexFormat-02430"
  id="VUID-VkGeometryTrianglesNV-vertexFormat-02430"></a>
  VUID-VkGeometryTrianglesNV-vertexFormat-02430  
  `vertexFormat` **must** be one of `VK_FORMAT_R32G32B32_SFLOAT`,
  `VK_FORMAT_R32G32_SFLOAT`, `VK_FORMAT_R16G16B16_SFLOAT`,
  `VK_FORMAT_R16G16_SFLOAT`, `VK_FORMAT_R16G16_SNORM`, or
  `VK_FORMAT_R16G16B16_SNORM`

- <a href="#VUID-VkGeometryTrianglesNV-vertexStride-03818"
  id="VUID-VkGeometryTrianglesNV-vertexStride-03818"></a>
  VUID-VkGeometryTrianglesNV-vertexStride-03818  
  `vertexStride` **must** be less than or equal to 2<sup>32</sup>-1

- <a href="#VUID-VkGeometryTrianglesNV-indexOffset-02431"
  id="VUID-VkGeometryTrianglesNV-indexOffset-02431"></a>
  VUID-VkGeometryTrianglesNV-indexOffset-02431  
  `indexOffset` **must** be less than the size of `indexData`

- <a href="#VUID-VkGeometryTrianglesNV-indexOffset-02432"
  id="VUID-VkGeometryTrianglesNV-indexOffset-02432"></a>
  VUID-VkGeometryTrianglesNV-indexOffset-02432  
  `indexOffset` **must** be a multiple of the element size of
  `indexType`

- <a href="#VUID-VkGeometryTrianglesNV-indexType-02433"
  id="VUID-VkGeometryTrianglesNV-indexType-02433"></a>
  VUID-VkGeometryTrianglesNV-indexType-02433  
  `indexType` **must** be `VK_INDEX_TYPE_UINT16`,
  `VK_INDEX_TYPE_UINT32`, or `VK_INDEX_TYPE_NONE_NV`

- <a href="#VUID-VkGeometryTrianglesNV-indexData-02434"
  id="VUID-VkGeometryTrianglesNV-indexData-02434"></a>
  VUID-VkGeometryTrianglesNV-indexData-02434  
  `indexData` **must** be [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) if
  `indexType` is `VK_INDEX_TYPE_NONE_NV`

- <a href="#VUID-VkGeometryTrianglesNV-indexData-02435"
  id="VUID-VkGeometryTrianglesNV-indexData-02435"></a>
  VUID-VkGeometryTrianglesNV-indexData-02435  
  `indexData` **must** be a valid `VkBuffer` handle if `indexType` is
  not `VK_INDEX_TYPE_NONE_NV`

- <a href="#VUID-VkGeometryTrianglesNV-indexCount-02436"
  id="VUID-VkGeometryTrianglesNV-indexCount-02436"></a>
  VUID-VkGeometryTrianglesNV-indexCount-02436  
  `indexCount` **must** be `0` if `indexType` is `VK_INDEX_TYPE_NONE_NV`

- <a href="#VUID-VkGeometryTrianglesNV-transformOffset-02437"
  id="VUID-VkGeometryTrianglesNV-transformOffset-02437"></a>
  VUID-VkGeometryTrianglesNV-transformOffset-02437  
  `transformOffset` **must** be less than the size of `transformData`

- <a href="#VUID-VkGeometryTrianglesNV-transformOffset-02438"
  id="VUID-VkGeometryTrianglesNV-transformOffset-02438"></a>
  VUID-VkGeometryTrianglesNV-transformOffset-02438  
  `transformOffset` **must** be a multiple of `16`

Valid Usage (Implicit)

- <a href="#VUID-VkGeometryTrianglesNV-sType-sType"
  id="VUID-VkGeometryTrianglesNV-sType-sType"></a>
  VUID-VkGeometryTrianglesNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_GEOMETRY_TRIANGLES_NV`

- <a href="#VUID-VkGeometryTrianglesNV-pNext-pNext"
  id="VUID-VkGeometryTrianglesNV-pNext-pNext"></a>
  VUID-VkGeometryTrianglesNV-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkGeometryTrianglesNV-vertexData-parameter"
  id="VUID-VkGeometryTrianglesNV-vertexData-parameter"></a>
  VUID-VkGeometryTrianglesNV-vertexData-parameter  
  If `vertexData` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `vertexData` **must** be a valid [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html) handle

- <a href="#VUID-VkGeometryTrianglesNV-vertexFormat-parameter"
  id="VUID-VkGeometryTrianglesNV-vertexFormat-parameter"></a>
  VUID-VkGeometryTrianglesNV-vertexFormat-parameter  
  `vertexFormat` **must** be a valid [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) value

- <a href="#VUID-VkGeometryTrianglesNV-indexData-parameter"
  id="VUID-VkGeometryTrianglesNV-indexData-parameter"></a>
  VUID-VkGeometryTrianglesNV-indexData-parameter  
  If `indexData` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `indexData` **must** be a valid [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html) handle

- <a href="#VUID-VkGeometryTrianglesNV-indexType-parameter"
  id="VUID-VkGeometryTrianglesNV-indexType-parameter"></a>
  VUID-VkGeometryTrianglesNV-indexType-parameter  
  `indexType` **must** be a valid [VkIndexType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkIndexType.html) value

- <a href="#VUID-VkGeometryTrianglesNV-transformData-parameter"
  id="VUID-VkGeometryTrianglesNV-transformData-parameter"></a>
  VUID-VkGeometryTrianglesNV-transformData-parameter  
  If `transformData` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `transformData` **must** be a valid [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html) handle

- <a href="#VUID-VkGeometryTrianglesNV-commonparent"
  id="VUID-VkGeometryTrianglesNV-commonparent"></a>
  VUID-VkGeometryTrianglesNV-commonparent  
  Each of `indexData`, `transformData`, and `vertexData` that are valid
  handles of non-ignored parameters **must** have been created,
  allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_ray_tracing](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_ray_tracing.html), [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html),
[VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html), [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html),
[VkGeometryDataNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGeometryDataNV.html),
[VkIndexType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkIndexType.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkGeometryTrianglesNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
