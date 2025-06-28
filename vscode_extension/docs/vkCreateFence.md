# vkCreateFence(3) Manual Page

## Name

vkCreateFence - Create a new fence object



## [](#_c_specification)C Specification

To create a fence, call:

```c++
// Provided by VK_VERSION_1_0
VkResult vkCreateFence(
    VkDevice                                    device,
    const VkFenceCreateInfo*                    pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkFence*                                    pFence);
```

## [](#_parameters)Parameters

- `device` is the logical device that creates the fence.
- `pCreateInfo` is a pointer to a [VkFenceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFenceCreateInfo.html) structure containing information about how the fence is to be created.
- `pAllocator` controls host memory allocation as described in the [Memory Allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-allocation) chapter.
- `pFence` is a pointer to a handle in which the resulting fence object is returned.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-vkCreateFence-device-parameter)VUID-vkCreateFence-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkCreateFence-pCreateInfo-parameter)VUID-vkCreateFence-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid [VkFenceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFenceCreateInfo.html) structure
- [](#VUID-vkCreateFence-pAllocator-parameter)VUID-vkCreateFence-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkCreateFence-pFence-parameter)VUID-vkCreateFence-pFence-parameter  
  `pFence` **must** be a valid pointer to a [VkFence](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFence.html) handle
- [](#VUID-vkCreateFence-device-queuecount)VUID-vkCreateFence-device-queuecount  
  The device **must** have been created with at least `1` queue

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkFence](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFence.html), [VkFenceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFenceCreateInfo.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCreateFence)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0