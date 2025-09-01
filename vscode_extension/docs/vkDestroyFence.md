# vkDestroyFence(3) Manual Page

## Name

vkDestroyFence - Destroy a fence object



## [](#_c_specification)C Specification

To destroy a fence, call:

```c++
// Provided by VK_VERSION_1_0
void vkDestroyFence(
    VkDevice                                    device,
    VkFence                                     fence,
    const VkAllocationCallbacks*                pAllocator);
```

## [](#_parameters)Parameters

- `device` is the logical device that destroys the fence.
- `fence` is the handle of the fence to destroy.
- `pAllocator` controls host memory allocation as described in the [Memory Allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-allocation) chapter.

## [](#_description)Description

Valid Usage

- [](#VUID-vkDestroyFence-fence-01120)VUID-vkDestroyFence-fence-01120  
  All [queue submission](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#devsandqueues-submission) commands that refer to `fence` **must** have completed execution
- [](#VUID-vkDestroyFence-fence-01121)VUID-vkDestroyFence-fence-01121  
  If `VkAllocationCallbacks` were provided when `fence` was created, a compatible set of callbacks **must** be provided here
- [](#VUID-vkDestroyFence-fence-01122)VUID-vkDestroyFence-fence-01122  
  If no `VkAllocationCallbacks` were provided when `fence` was created, `pAllocator` **must** be `NULL`

Valid Usage (Implicit)

- [](#VUID-vkDestroyFence-device-parameter)VUID-vkDestroyFence-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkDestroyFence-fence-parameter)VUID-vkDestroyFence-fence-parameter  
  If `fence` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `fence` **must** be a valid [VkFence](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFence.html) handle
- [](#VUID-vkDestroyFence-pAllocator-parameter)VUID-vkDestroyFence-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkDestroyFence-fence-parent)VUID-vkDestroyFence-fence-parent  
  If `fence` is a valid handle, it **must** have been created, allocated, or retrieved from `device`

Host Synchronization

- Host access to `fence` **must** be externally synchronized

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkFence](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFence.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkDestroyFence).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0