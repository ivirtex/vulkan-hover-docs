# VkSwapchainLatencyCreateInfoNV(3) Manual Page

## Name

VkSwapchainLatencyCreateInfoNV - Specify that a swapchain will use low
latency mode



## <a href="#_c_specification" class="anchor"></a>C Specification

To allow low latency mode to be used by a swapchain, add a
`VkSwapchainLatencyCreateInfoNV` structure to the `pNext` chain of
[VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainCreateInfoKHR.html).

The `VkSwapchainLatencyCreateInfoNV` structure is defined as:

``` c
// Provided by VK_NV_low_latency2
typedef struct VkSwapchainLatencyCreateInfoNV {
    VkStructureType    sType;
    const void*        pNext;
    VkBool32           latencyModeEnable;
} VkSwapchainLatencyCreateInfoNV;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `lowLatencyModeEnable` indicates if the swapchain created will utilize
  low latency mode.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkSwapchainLatencyCreateInfoNV-sType-sType"
  id="VUID-VkSwapchainLatencyCreateInfoNV-sType-sType"></a>
  VUID-VkSwapchainLatencyCreateInfoNV-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_SWAPCHAIN_LATENCY_CREATE_INFO_NV`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_low_latency2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_low_latency2.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkSwapchainLatencyCreateInfoNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
