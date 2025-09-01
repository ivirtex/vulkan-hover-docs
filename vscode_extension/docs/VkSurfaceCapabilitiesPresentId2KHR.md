# VkSurfaceCapabilitiesPresentId2KHR(3) Manual Page

## Name

VkSurfaceCapabilitiesPresentId2KHR - Structure describing presentation-ID capabilities of a surface



## [](#_c_specification)C Specification

The `VkSurfaceCapabilitiesPresentId2KHR` structure is defined as:

```c++
// Provided by VK_KHR_present_id2
typedef struct VkSurfaceCapabilitiesPresentId2KHR {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           presentId2Supported;
} VkSurfaceCapabilitiesPresentId2KHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `presentId2Supported` is a boolean describing whether the surface is able to support the present-ID extension

## [](#_description)Description

This structure **can** be included in the `pNext` chain of [VkSurfaceCapabilities2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceCapabilities2KHR.html) to determine support for present-wait. If `presentId2Supported` is `VK_FALSE`, it indicates that attaching an ID to presentation requests is not possible for this surface.

Applications **must** not attempt to include [VkPresentId2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentId2KHR.html) in the `pNext` chain of a [VkPresentInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentInfoKHR.html) if `presentId2Supported` is `VK_FALSE`.

Valid Usage (Implicit)

- [](#VUID-VkSurfaceCapabilitiesPresentId2KHR-sType-sType)VUID-VkSurfaceCapabilitiesPresentId2KHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SURFACE_CAPABILITIES_PRESENT_ID_2_KHR`

## [](#_see_also)See Also

[VK\_KHR\_present\_id2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_present_id2.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkSurfaceCapabilitiesPresentId2KHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0