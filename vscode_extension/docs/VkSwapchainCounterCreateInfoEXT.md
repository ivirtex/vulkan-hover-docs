# VkSwapchainCounterCreateInfoEXT(3) Manual Page

## Name

VkSwapchainCounterCreateInfoEXT - Specify the surface counters desired



## <a href="#_c_specification" class="anchor"></a>C Specification

To enable surface counters when creating a swapchain, add a
`VkSwapchainCounterCreateInfoEXT` structure to the `pNext` chain of
[VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainCreateInfoKHR.html).
`VkSwapchainCounterCreateInfoEXT` is defined as:

``` c
// Provided by VK_EXT_display_control
typedef struct VkSwapchainCounterCreateInfoEXT {
    VkStructureType             sType;
    const void*                 pNext;
    VkSurfaceCounterFlagsEXT    surfaceCounters;
} VkSwapchainCounterCreateInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `surfaceCounters` is a bitmask of
  [VkSurfaceCounterFlagBitsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceCounterFlagBitsEXT.html)
  specifying surface counters to enable for the swapchain.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkSwapchainCounterCreateInfoEXT-surfaceCounters-01244"
  id="VUID-VkSwapchainCounterCreateInfoEXT-surfaceCounters-01244"></a>
  VUID-VkSwapchainCounterCreateInfoEXT-surfaceCounters-01244  
  The bits in `surfaceCounters` **must** be supported by
  [VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainCreateInfoKHR.html)::`surface`,
  as reported by
  [vkGetPhysicalDeviceSurfaceCapabilities2EXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceSurfaceCapabilities2EXT.html)

Valid Usage (Implicit)

- <a href="#VUID-VkSwapchainCounterCreateInfoEXT-sType-sType"
  id="VUID-VkSwapchainCounterCreateInfoEXT-sType-sType"></a>
  VUID-VkSwapchainCounterCreateInfoEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_SWAPCHAIN_COUNTER_CREATE_INFO_EXT`

- <a
  href="#VUID-VkSwapchainCounterCreateInfoEXT-surfaceCounters-parameter"
  id="VUID-VkSwapchainCounterCreateInfoEXT-surfaceCounters-parameter"></a>
  VUID-VkSwapchainCounterCreateInfoEXT-surfaceCounters-parameter  
  `surfaceCounters` **must** be a valid combination of
  [VkSurfaceCounterFlagBitsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceCounterFlagBitsEXT.html) values

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_display_control](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_display_control.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[VkSurfaceCounterFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceCounterFlagsEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkSwapchainCounterCreateInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
