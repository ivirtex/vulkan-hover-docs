# VkSurfaceCapabilitiesPresentWait2KHR(3) Manual Page

## Name

VkSurfaceCapabilitiesPresentWait2KHR - Structure describing presentation-wait capabilities of a surface



## [](#_c_specification)C Specification

The `VkSurfaceCapabilitiesPresentWait2KHR` structure is defined as:

```c++
// Provided by VK_KHR_present_wait2
typedef struct VkSurfaceCapabilitiesPresentWait2KHR {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           presentWait2Supported;
} VkSurfaceCapabilitiesPresentWait2KHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `presentWait2Supported` is a boolean describing whether the surface is able to support the present-wait extension

## [](#_description)Description

This structure **can** be included in the `pNext` chain of [VkSurfaceCapabilities2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceCapabilities2KHR.html) to determine support for present-wait. If `presentWait2Supported` is `VK_FALSE`, it indicates that waiting for presentation is not possible for this surface.

Applications **must** not attempt to call [vkWaitForPresent2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkWaitForPresent2KHR.html) on a swapchain if `presentWait2Supported` is `VK_FALSE`.

Valid Usage (Implicit)

- [](#VUID-VkSurfaceCapabilitiesPresentWait2KHR-sType-sType)VUID-VkSurfaceCapabilitiesPresentWait2KHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SURFACE_CAPABILITIES_PRESENT_WAIT_2_KHR`

## [](#_see_also)See Also

[VK\_KHR\_present\_wait2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_present_wait2.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkSurfaceCapabilitiesPresentWait2KHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0