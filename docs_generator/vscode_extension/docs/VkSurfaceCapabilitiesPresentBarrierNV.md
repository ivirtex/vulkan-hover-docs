# VkSurfaceCapabilitiesPresentBarrierNV(3) Manual Page

## Name

VkSurfaceCapabilitiesPresentBarrierNV - Structure describing present barrier capabilities of a surface



## [](#_c_specification)C Specification

The `VkSurfaceCapabilitiesPresentBarrierNV` structure is defined as:

```c++
// Provided by VK_NV_present_barrier
typedef struct VkSurfaceCapabilitiesPresentBarrierNV {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           presentBarrierSupported;
} VkSurfaceCapabilitiesPresentBarrierNV;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `presentBarrierSupported` is a boolean describing whether the surface is able to make use of the present barrier feature.

## [](#_description)Description

This structure **can** be included in the `pNext` chain of [VkSurfaceCapabilities2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceCapabilities2KHR.html) to determine support for present barrier access. If `presentBarrierSupported` is `VK_FALSE`, it indicates that the present barrier feature is not obtainable for this surface.

Valid Usage (Implicit)

- [](#VUID-VkSurfaceCapabilitiesPresentBarrierNV-sType-sType)VUID-VkSurfaceCapabilitiesPresentBarrierNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SURFACE_CAPABILITIES_PRESENT_BARRIER_NV`

## [](#_see_also)See Also

[VK\_NV\_present\_barrier](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_present_barrier.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkSurfaceCapabilitiesPresentBarrierNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0