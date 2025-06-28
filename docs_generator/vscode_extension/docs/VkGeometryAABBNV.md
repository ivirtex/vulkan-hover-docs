# VkGeometryAABBNV(3) Manual Page

## Name

VkGeometryAABBNV - Structure specifying axis-aligned bounding box geometry in a bottom-level acceleration structure



## [](#_c_specification)C Specification

The `VkGeometryAABBNV` structure specifies axis-aligned bounding box geometry in a bottom-level acceleration structure, and is defined as:

```c++
// Provided by VK_NV_ray_tracing
typedef struct VkGeometryAABBNV {
    VkStructureType    sType;
    const void*        pNext;
    VkBuffer           aabbData;
    uint32_t           numAABBs;
    uint32_t           stride;
    VkDeviceSize       offset;
} VkGeometryAABBNV;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `aabbData` is the buffer containing axis-aligned bounding box data.
- `numAABBs` is the number of AABBs in this geometry.
- `stride` is the stride in bytes between AABBs in `aabbData`.
- `offset` is the offset in bytes of the first AABB in `aabbData`.

## [](#_description)Description

The AABB data in memory is six 32-bit floats consisting of the minimum x, y, and z values followed by the maximum x, y, and z values.

Valid Usage

- [](#VUID-VkGeometryAABBNV-offset-02439)VUID-VkGeometryAABBNV-offset-02439  
  `offset` **must** be less than the size of `aabbData`
- [](#VUID-VkGeometryAABBNV-offset-02440)VUID-VkGeometryAABBNV-offset-02440  
  `offset` **must** be a multiple of `8`
- [](#VUID-VkGeometryAABBNV-stride-02441)VUID-VkGeometryAABBNV-stride-02441  
  `stride` **must** be a multiple of `8`

Valid Usage (Implicit)

- [](#VUID-VkGeometryAABBNV-sType-sType)VUID-VkGeometryAABBNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_GEOMETRY_AABB_NV`
- [](#VUID-VkGeometryAABBNV-pNext-pNext)VUID-VkGeometryAABBNV-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkGeometryAABBNV-aabbData-parameter)VUID-VkGeometryAABBNV-aabbData-parameter  
  If `aabbData` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `aabbData` **must** be a valid [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html) handle

## [](#_see_also)See Also

[VK\_NV\_ray\_tracing](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_ray_tracing.html), [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html), [VkGeometryDataNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGeometryDataNV.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkGeometryAABBNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0