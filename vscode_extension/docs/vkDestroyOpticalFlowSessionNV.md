# vkDestroyOpticalFlowSessionNV(3) Manual Page

## Name

vkDestroyOpticalFlowSessionNV - Destroy optical flow session object



## [](#_c_specification)C Specification

To destroy an optical flow session object, call:

```c++
// Provided by VK_NV_optical_flow
void vkDestroyOpticalFlowSessionNV(
    VkDevice                                    device,
    VkOpticalFlowSessionNV                      session,
    const VkAllocationCallbacks*                pAllocator);
```

## [](#_parameters)Parameters

- `device` is the device that was used for the creation of the optical flow session.
- `session` is the optical flow session to be destroyed.
- `pAllocator` controls host memory allocation as described in the [Memory Allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-allocation) chapter.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-vkDestroyOpticalFlowSessionNV-device-parameter)VUID-vkDestroyOpticalFlowSessionNV-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkDestroyOpticalFlowSessionNV-session-parameter)VUID-vkDestroyOpticalFlowSessionNV-session-parameter  
  `session` **must** be a valid [VkOpticalFlowSessionNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpticalFlowSessionNV.html) handle
- [](#VUID-vkDestroyOpticalFlowSessionNV-pAllocator-parameter)VUID-vkDestroyOpticalFlowSessionNV-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkDestroyOpticalFlowSessionNV-session-parent)VUID-vkDestroyOpticalFlowSessionNV-session-parent  
  `session` **must** have been created, allocated, or retrieved from `device`

## [](#_see_also)See Also

[VK\_NV\_optical\_flow](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_optical_flow.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkOpticalFlowSessionNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpticalFlowSessionNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkDestroyOpticalFlowSessionNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0