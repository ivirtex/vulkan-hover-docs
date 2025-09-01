# VkSwapchainPresentBarrierCreateInfoNV(3) Manual Page

## Name

VkSwapchainPresentBarrierCreateInfoNV - specify the present barrier membership of this swapchain



## [](#_c_specification)C Specification

The [VkSwapchainPresentBarrierCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainPresentBarrierCreateInfoNV.html) structure is defined as:

```c++
// Provided by VK_NV_present_barrier
typedef struct VkSwapchainPresentBarrierCreateInfoNV {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           presentBarrierEnable;
} VkSwapchainPresentBarrierCreateInfoNV;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `presentBarrierEnable` is a boolean value indicating a request for using the *present barrier*.

## [](#_description)Description

If the `pNext` chain of [VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainCreateInfoKHR.html) does not include this structure, the default value for `presentBarrierEnable` is `VK_FALSE`, meaning the swapchain does not request to use the present barrier. Additionally, when recreating a swapchain that was using the present barrier, and the `pNext` chain of [VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainCreateInfoKHR.html) does not include this structure, it means the swapchain will stop using the present barrier.

Valid Usage (Implicit)

- [](#VUID-VkSwapchainPresentBarrierCreateInfoNV-sType-sType)VUID-VkSwapchainPresentBarrierCreateInfoNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SWAPCHAIN_PRESENT_BARRIER_CREATE_INFO_NV`

## [](#_see_also)See Also

[VK\_NV\_present\_barrier](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_present_barrier.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkSwapchainPresentBarrierCreateInfoNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0