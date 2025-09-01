# VkSwapchainCounterCreateInfoEXT(3) Manual Page

## Name

VkSwapchainCounterCreateInfoEXT - Specify the surface counters desired



## [](#_c_specification)C Specification

To enable surface counters when creating a swapchain, add a `VkSwapchainCounterCreateInfoEXT` structure to the `pNext` chain of [VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainCreateInfoKHR.html). `VkSwapchainCounterCreateInfoEXT` is defined as:

```c++
// Provided by VK_EXT_display_control
typedef struct VkSwapchainCounterCreateInfoEXT {
    VkStructureType             sType;
    const void*                 pNext;
    VkSurfaceCounterFlagsEXT    surfaceCounters;
} VkSwapchainCounterCreateInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `surfaceCounters` is a bitmask of [VkSurfaceCounterFlagBitsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceCounterFlagBitsEXT.html) specifying surface counters to enable for the swapchain.

## [](#_description)Description

Valid Usage

- [](#VUID-VkSwapchainCounterCreateInfoEXT-surfaceCounters-01244)VUID-VkSwapchainCounterCreateInfoEXT-surfaceCounters-01244  
  The bits in `surfaceCounters` **must** be supported by [VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainCreateInfoKHR.html)::`surface`, as reported by [vkGetPhysicalDeviceSurfaceCapabilities2EXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceSurfaceCapabilities2EXT.html)

Valid Usage (Implicit)

- [](#VUID-VkSwapchainCounterCreateInfoEXT-sType-sType)VUID-VkSwapchainCounterCreateInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SWAPCHAIN_COUNTER_CREATE_INFO_EXT`
- [](#VUID-VkSwapchainCounterCreateInfoEXT-surfaceCounters-parameter)VUID-VkSwapchainCounterCreateInfoEXT-surfaceCounters-parameter  
  `surfaceCounters` **must** be a valid combination of [VkSurfaceCounterFlagBitsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceCounterFlagBitsEXT.html) values

## [](#_see_also)See Also

[VK\_EXT\_display\_control](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_display_control.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkSurfaceCounterFlagsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSurfaceCounterFlagsEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkSwapchainCounterCreateInfoEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0