# VkGeometryAABBNV(3) Manual Page

## Name

VkGeometryAABBNV - Structure specifying axis-aligned bounding box
geometry in a bottom-level acceleration structure



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkGeometryAABBNV` structure specifies axis-aligned bounding box
geometry in a bottom-level acceleration structure, and is defined as:

``` c
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

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `aabbData` is the buffer containing axis-aligned bounding box data.

- `numAABBs` is the number of AABBs in this geometry.

- `stride` is the stride in bytes between AABBs in `aabbData`.

- `offset` is the offset in bytes of the first AABB in `aabbData`.

## <a href="#_description" class="anchor"></a>Description

The AABB data in memory is six 32-bit floats consisting of the minimum
x, y, and z values followed by the maximum x, y, and z values.

Valid Usage

- <a href="#VUID-VkGeometryAABBNV-offset-02439"
  id="VUID-VkGeometryAABBNV-offset-02439"></a>
  VUID-VkGeometryAABBNV-offset-02439  
  `offset` **must** be less than the size of `aabbData`

- <a href="#VUID-VkGeometryAABBNV-offset-02440"
  id="VUID-VkGeometryAABBNV-offset-02440"></a>
  VUID-VkGeometryAABBNV-offset-02440  
  `offset` **must** be a multiple of `8`

- <a href="#VUID-VkGeometryAABBNV-stride-02441"
  id="VUID-VkGeometryAABBNV-stride-02441"></a>
  VUID-VkGeometryAABBNV-stride-02441  
  `stride` **must** be a multiple of `8`

Valid Usage (Implicit)

- <a href="#VUID-VkGeometryAABBNV-sType-sType"
  id="VUID-VkGeometryAABBNV-sType-sType"></a>
  VUID-VkGeometryAABBNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_GEOMETRY_AABB_NV`

- <a href="#VUID-VkGeometryAABBNV-pNext-pNext"
  id="VUID-VkGeometryAABBNV-pNext-pNext"></a>
  VUID-VkGeometryAABBNV-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkGeometryAABBNV-aabbData-parameter"
  id="VUID-VkGeometryAABBNV-aabbData-parameter"></a>
  VUID-VkGeometryAABBNV-aabbData-parameter  
  If `aabbData` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), `aabbData`
  **must** be a valid [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html) handle

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_ray_tracing](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_ray_tracing.html), [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html),
[VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html),
[VkGeometryDataNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGeometryDataNV.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkGeometryAABBNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
