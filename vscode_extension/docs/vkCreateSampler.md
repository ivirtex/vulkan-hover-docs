# vkCreateSampler(3) Manual Page

## Name

vkCreateSampler - Create a new sampler object



## [](#_c_specification)C Specification

To create a sampler object, call:

```c++
// Provided by VK_VERSION_1_0
VkResult vkCreateSampler(
    VkDevice                                    device,
    const VkSamplerCreateInfo*                  pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkSampler*                                  pSampler);
```

## [](#_parameters)Parameters

- `device` is the logical device that creates the sampler.
- `pCreateInfo` is a pointer to a [VkSamplerCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerCreateInfo.html) structure specifying the state of the sampler object.
- `pAllocator` controls host memory allocation as described in the [Memory Allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-allocation) chapter.
- `pSampler` is a pointer to a [VkSampler](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampler.html) handle in which the resulting sampler object is returned.

## [](#_description)Description

Valid Usage

- [](#VUID-vkCreateSampler-device-09668)VUID-vkCreateSampler-device-09668  
  `device` **must** support at least one queue family with one of the `VK_QUEUE_COMPUTE_BIT` or `VK_QUEUE_GRAPHICS_BIT` capabilities
- [](#VUID-vkCreateSampler-maxSamplerAllocationCount-04110)VUID-vkCreateSampler-maxSamplerAllocationCount-04110  
  There **must** be less than [VkPhysicalDeviceLimits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceLimits.html)::`maxSamplerAllocationCount` [VkSampler](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampler.html) objects currently created on the device

Valid Usage (Implicit)

- [](#VUID-vkCreateSampler-device-parameter)VUID-vkCreateSampler-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkCreateSampler-pCreateInfo-parameter)VUID-vkCreateSampler-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid [VkSamplerCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerCreateInfo.html) structure
- [](#VUID-vkCreateSampler-pAllocator-parameter)VUID-vkCreateSampler-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkCreateSampler-pSampler-parameter)VUID-vkCreateSampler-pSampler-parameter  
  `pSampler` **must** be a valid pointer to a [VkSampler](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampler.html) handle

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_INVALID_OPAQUE_CAPTURE_ADDRESS_KHR`
- `VK_ERROR_OUT_OF_DEVICE_MEMORY`
- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkSampler](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampler.html), [VkSamplerCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerCreateInfo.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCreateSampler).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0