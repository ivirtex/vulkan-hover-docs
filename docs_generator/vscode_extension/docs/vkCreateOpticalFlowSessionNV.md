# vkCreateOpticalFlowSessionNV(3) Manual Page

## Name

vkCreateOpticalFlowSessionNV - Creates an optical flow session object



## [](#_c_specification)C Specification

To create an optical flow session object, call:

```c++
// Provided by VK_NV_optical_flow
VkResult vkCreateOpticalFlowSessionNV(
    VkDevice                                    device,
    const VkOpticalFlowSessionCreateInfoNV*     pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkOpticalFlowSessionNV*                     pSession);
```

## [](#_parameters)Parameters

- `device` is the logical device that creates the optical flow session object.
- `pCreateInfo` is a pointer to a [VkOpticalFlowSessionCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpticalFlowSessionCreateInfoNV.html) structure containing parameters specifying the creation of the optical flow session.
- `pAllocator` controls host memory allocation as described in the [Memory Allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-allocation) chapter.
- `pSession` is a pointer to a [VkOpticalFlowSessionNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpticalFlowSessionNV.html) handle specifying the optical flow session object which will be created by this function when it returns `VK_SUCCESS`

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-vkCreateOpticalFlowSessionNV-device-parameter)VUID-vkCreateOpticalFlowSessionNV-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkCreateOpticalFlowSessionNV-pCreateInfo-parameter)VUID-vkCreateOpticalFlowSessionNV-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid [VkOpticalFlowSessionCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpticalFlowSessionCreateInfoNV.html) structure
- [](#VUID-vkCreateOpticalFlowSessionNV-pAllocator-parameter)VUID-vkCreateOpticalFlowSessionNV-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkCreateOpticalFlowSessionNV-pSession-parameter)VUID-vkCreateOpticalFlowSessionNV-pSession-parameter  
  `pSession` **must** be a valid pointer to a [VkOpticalFlowSessionNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpticalFlowSessionNV.html) handle
- [](#VUID-vkCreateOpticalFlowSessionNV-device-queuecount)VUID-vkCreateOpticalFlowSessionNV-device-queuecount  
  The device **must** have been created with at least `1` queue

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_INITIALIZATION_FAILED`

## [](#_see_also)See Also

[VK\_NV\_optical\_flow](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_optical_flow.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkOpticalFlowSessionCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpticalFlowSessionCreateInfoNV.html), [VkOpticalFlowSessionNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpticalFlowSessionNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCreateOpticalFlowSessionNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0