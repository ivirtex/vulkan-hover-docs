# VkGeometryTrianglesNV(3) Manual Page

## Name

VkGeometryTrianglesNV - Structure specifying a triangle geometry in a bottom-level acceleration structure



## [](#_c_specification)C Specification

The `VkGeometryTrianglesNV` structure specifies triangle geometry in a bottom-level acceleration structure and is defined as:

```c++
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

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `vertexData` is the buffer containing vertex data for this geometry.
- `vertexOffset` is the offset in bytes within `vertexData` containing vertex data for this geometry.
- `vertexCount` is the number of valid vertices.
- `vertexStride` is the stride in bytes between each vertex.
- `vertexFormat` is a [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) describing the format of each vertex element.
- `indexData` is the buffer containing index data for this geometry.
- `indexOffset` is the offset in bytes within `indexData` containing index data for this geometry.
- `indexCount` is the number of indices to include in this geometry.
- `indexType` is a [VkIndexType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndexType.html) describing the format of each index.
- `transformData` is an optional buffer containing an [VkTransformMatrixNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTransformMatrixNV.html) structure defining a transformation to be applied to this geometry.
- `transformOffset` is the offset in bytes in `transformData` of the transform information described above.

## [](#_description)Description

If `indexType` is `VK_INDEX_TYPE_NONE_NV`, then this structure describes a set of triangles determined by `vertexCount`. Otherwise, this structure describes a set of indexed triangles determined by `indexCount`.

Valid Usage

- [](#VUID-VkGeometryTrianglesNV-vertexOffset-02428)VUID-VkGeometryTrianglesNV-vertexOffset-02428  
  `vertexOffset` **must** be less than the size of `vertexData`
- [](#VUID-VkGeometryTrianglesNV-vertexOffset-02429)VUID-VkGeometryTrianglesNV-vertexOffset-02429  
  `vertexOffset` **must** be a multiple of the component size of `vertexFormat`
- [](#VUID-VkGeometryTrianglesNV-vertexFormat-02430)VUID-VkGeometryTrianglesNV-vertexFormat-02430  
  `vertexFormat` **must** be one of `VK_FORMAT_R32G32B32_SFLOAT`, `VK_FORMAT_R32G32_SFLOAT`, `VK_FORMAT_R16G16B16_SFLOAT`, `VK_FORMAT_R16G16_SFLOAT`, `VK_FORMAT_R16G16_SNORM`, or `VK_FORMAT_R16G16B16_SNORM`
- [](#VUID-VkGeometryTrianglesNV-vertexStride-03818)VUID-VkGeometryTrianglesNV-vertexStride-03818  
  `vertexStride` **must** be less than or equal to 232-1
- [](#VUID-VkGeometryTrianglesNV-indexOffset-02431)VUID-VkGeometryTrianglesNV-indexOffset-02431  
  `indexOffset` **must** be less than the size of `indexData`
- [](#VUID-VkGeometryTrianglesNV-indexOffset-02432)VUID-VkGeometryTrianglesNV-indexOffset-02432  
  `indexOffset` **must** be a multiple of the element size of `indexType`
- [](#VUID-VkGeometryTrianglesNV-indexType-02433)VUID-VkGeometryTrianglesNV-indexType-02433  
  `indexType` **must** be `VK_INDEX_TYPE_UINT16`, `VK_INDEX_TYPE_UINT32`, or `VK_INDEX_TYPE_NONE_NV`
- [](#VUID-VkGeometryTrianglesNV-indexData-02434)VUID-VkGeometryTrianglesNV-indexData-02434  
  `indexData` **must** be [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) if `indexType` is `VK_INDEX_TYPE_NONE_NV`
- [](#VUID-VkGeometryTrianglesNV-indexData-02435)VUID-VkGeometryTrianglesNV-indexData-02435  
  `indexData` **must** be a valid `VkBuffer` handle if `indexType` is not `VK_INDEX_TYPE_NONE_NV`
- [](#VUID-VkGeometryTrianglesNV-indexCount-02436)VUID-VkGeometryTrianglesNV-indexCount-02436  
  `indexCount` **must** be `0` if `indexType` is `VK_INDEX_TYPE_NONE_NV`
- [](#VUID-VkGeometryTrianglesNV-transformOffset-02437)VUID-VkGeometryTrianglesNV-transformOffset-02437  
  `transformOffset` **must** be less than the size of `transformData`
- [](#VUID-VkGeometryTrianglesNV-transformOffset-02438)VUID-VkGeometryTrianglesNV-transformOffset-02438  
  `transformOffset` **must** be a multiple of `16`

Valid Usage (Implicit)

- [](#VUID-VkGeometryTrianglesNV-sType-sType)VUID-VkGeometryTrianglesNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_GEOMETRY_TRIANGLES_NV`
- [](#VUID-VkGeometryTrianglesNV-pNext-pNext)VUID-VkGeometryTrianglesNV-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkGeometryTrianglesNV-vertexData-parameter)VUID-VkGeometryTrianglesNV-vertexData-parameter  
  If `vertexData` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `vertexData` **must** be a valid [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html) handle
- [](#VUID-VkGeometryTrianglesNV-vertexFormat-parameter)VUID-VkGeometryTrianglesNV-vertexFormat-parameter  
  `vertexFormat` **must** be a valid [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) value
- [](#VUID-VkGeometryTrianglesNV-indexData-parameter)VUID-VkGeometryTrianglesNV-indexData-parameter  
  If `indexData` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `indexData` **must** be a valid [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html) handle
- [](#VUID-VkGeometryTrianglesNV-indexType-parameter)VUID-VkGeometryTrianglesNV-indexType-parameter  
  `indexType` **must** be a valid [VkIndexType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndexType.html) value
- [](#VUID-VkGeometryTrianglesNV-transformData-parameter)VUID-VkGeometryTrianglesNV-transformData-parameter  
  If `transformData` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `transformData` **must** be a valid [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html) handle
- [](#VUID-VkGeometryTrianglesNV-commonparent)VUID-VkGeometryTrianglesNV-commonparent  
  Each of `indexData`, `transformData`, and `vertexData` that are valid handles of non-ignored parameters **must** have been created, allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_see_also)See Also

[VK\_NV\_ray\_tracing](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_ray_tracing.html), [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html), [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html), [VkGeometryDataNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGeometryDataNV.html), [VkIndexType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkIndexType.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkGeometryTrianglesNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0