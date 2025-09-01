# vkDestroySampler(3) Manual Page

## Name

vkDestroySampler - Destroy a sampler object



## [](#_c_specification)C Specification

To destroy a sampler, call:

```c++
// Provided by VK_VERSION_1_0
void vkDestroySampler(
    VkDevice                                    device,
    VkSampler                                   sampler,
    const VkAllocationCallbacks*                pAllocator);
```

## [](#_parameters)Parameters

- `device` is the logical device that destroys the sampler.
- `sampler` is the sampler to destroy.
- `pAllocator` controls host memory allocation as described in the [Memory Allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-allocation) chapter.

## [](#_description)Description

Valid Usage

- [](#VUID-vkDestroySampler-sampler-01082)VUID-vkDestroySampler-sampler-01082  
  All submitted commands that refer to `sampler` **must** have completed execution
- [](#VUID-vkDestroySampler-sampler-01083)VUID-vkDestroySampler-sampler-01083  
  If `VkAllocationCallbacks` were provided when `sampler` was created, a compatible set of callbacks **must** be provided here
- [](#VUID-vkDestroySampler-sampler-01084)VUID-vkDestroySampler-sampler-01084  
  If no `VkAllocationCallbacks` were provided when `sampler` was created, `pAllocator` **must** be `NULL`

Valid Usage (Implicit)

- [](#VUID-vkDestroySampler-device-parameter)VUID-vkDestroySampler-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkDestroySampler-sampler-parameter)VUID-vkDestroySampler-sampler-parameter  
  If `sampler` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `sampler` **must** be a valid [VkSampler](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampler.html) handle
- [](#VUID-vkDestroySampler-pAllocator-parameter)VUID-vkDestroySampler-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkDestroySampler-sampler-parent)VUID-vkDestroySampler-sampler-parent  
  If `sampler` is a valid handle, it **must** have been created, allocated, or retrieved from `device`

Host Synchronization

- Host access to `sampler` **must** be externally synchronized

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkSampler](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampler.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkDestroySampler).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0