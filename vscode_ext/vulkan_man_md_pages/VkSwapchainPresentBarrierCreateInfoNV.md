# VkSwapchainPresentBarrierCreateInfoNV(3) Manual Page

## Name

VkSwapchainPresentBarrierCreateInfoNV - specify the present barrier
membership of this swapchain



## <a href="#_c_specification" class="anchor"></a>C Specification

The
[VkSwapchainPresentBarrierCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainPresentBarrierCreateInfoNV.html)
structure is defined as:

``` c
// Provided by VK_NV_present_barrier
typedef struct VkSwapchainPresentBarrierCreateInfoNV {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           presentBarrierEnable;
} VkSwapchainPresentBarrierCreateInfoNV;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `presentBarrierEnable` is a boolean value indicating a request for
  using the *present barrier*.

## <a href="#_description" class="anchor"></a>Description

If the `pNext` chain of
[VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainCreateInfoKHR.html) does not
include this structure, the default value for `presentBarrierEnable` is
`VK_FALSE`, meaning the swapchain does not request to use the present
barrier. Additionally, when recreating a swapchain that was using the
present barrier, and the `pNext` chain of
[VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainCreateInfoKHR.html) does not
include this structure, it means the swapchain will stop using the
present barrier.

Valid Usage (Implicit)

- <a href="#VUID-VkSwapchainPresentBarrierCreateInfoNV-sType-sType"
  id="VUID-VkSwapchainPresentBarrierCreateInfoNV-sType-sType"></a>
  VUID-VkSwapchainPresentBarrierCreateInfoNV-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_SWAPCHAIN_PRESENT_BARRIER_CREATE_INFO_NV`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_present_barrier](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_present_barrier.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkSwapchainPresentBarrierCreateInfoNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
