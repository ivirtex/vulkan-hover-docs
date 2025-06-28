# VkSwapchainLatencyCreateInfoNV(3) Manual Page

## Name

VkSwapchainLatencyCreateInfoNV - Specify that a swapchain will use low latency mode



## [](#_c_specification)C Specification

To allow low latency mode to be used by a swapchain, add a `VkSwapchainLatencyCreateInfoNV` structure to the `pNext` chain of [VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainCreateInfoKHR.html).

The `VkSwapchainLatencyCreateInfoNV` structure is defined as:

```c++
// Provided by VK_NV_low_latency2
typedef struct VkSwapchainLatencyCreateInfoNV {
    VkStructureType    sType;
    const void*        pNext;
    VkBool32           latencyModeEnable;
} VkSwapchainLatencyCreateInfoNV;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `latencyModeEnable` is `VK_TRUE` if the created swapchain will utilize low latency mode, `VK_FALSE` otherwise.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkSwapchainLatencyCreateInfoNV-sType-sType)VUID-VkSwapchainLatencyCreateInfoNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SWAPCHAIN_LATENCY_CREATE_INFO_NV`

## [](#_see_also)See Also

[VK\_NV\_low\_latency2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_low_latency2.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkSwapchainLatencyCreateInfoNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0