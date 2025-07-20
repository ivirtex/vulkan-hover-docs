# vkCreateSamplerYcbcrConversion(3) Manual Page

## Name

vkCreateSamplerYcbcrConversion - Create a new Y′CBCR conversion



## [](#_c_specification)C Specification

To create a [VkSamplerYcbcrConversion](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerYcbcrConversion.html), call:

```c++
// Provided by VK_VERSION_1_1
VkResult vkCreateSamplerYcbcrConversion(
    VkDevice                                    device,
    const VkSamplerYcbcrConversionCreateInfo*   pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkSamplerYcbcrConversion*                   pYcbcrConversion);
```

or the equivalent command

```c++
// Provided by VK_KHR_sampler_ycbcr_conversion
VkResult vkCreateSamplerYcbcrConversionKHR(
    VkDevice                                    device,
    const VkSamplerYcbcrConversionCreateInfo*   pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkSamplerYcbcrConversion*                   pYcbcrConversion);
```

## [](#_parameters)Parameters

- `device` is the logical device that creates the sampler Y′CBCR conversion.
- `pCreateInfo` is a pointer to a [VkSamplerYcbcrConversionCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerYcbcrConversionCreateInfo.html) structure specifying the requested sampler Y′CBCR conversion.
- `pAllocator` controls host memory allocation as described in the [Memory Allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-allocation) chapter.
- `pYcbcrConversion` is a pointer to a [VkSamplerYcbcrConversion](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerYcbcrConversion.html) handle in which the resulting sampler Y′CBCR conversion is returned.

## [](#_description)Description

The interpretation of the configured sampler Y′CBCR conversion is described in more detail in [the description of sampler Y′CBCR conversion](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#textures-sampler-YCbCr-conversion) in the [Image Operations](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#textures) chapter.

Valid Usage

- [](#VUID-vkCreateSamplerYcbcrConversion-None-01648)VUID-vkCreateSamplerYcbcrConversion-None-01648  
  The [`samplerYcbcrConversion`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-samplerYcbcrConversion) feature **must** be enabled

Valid Usage (Implicit)

- [](#VUID-vkCreateSamplerYcbcrConversion-device-parameter)VUID-vkCreateSamplerYcbcrConversion-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkCreateSamplerYcbcrConversion-pCreateInfo-parameter)VUID-vkCreateSamplerYcbcrConversion-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid [VkSamplerYcbcrConversionCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerYcbcrConversionCreateInfo.html) structure
- [](#VUID-vkCreateSamplerYcbcrConversion-pAllocator-parameter)VUID-vkCreateSamplerYcbcrConversion-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkCreateSamplerYcbcrConversion-pYcbcrConversion-parameter)VUID-vkCreateSamplerYcbcrConversion-pYcbcrConversion-parameter  
  `pYcbcrConversion` **must** be a valid pointer to a [VkSamplerYcbcrConversion](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerYcbcrConversion.html) handle
- [](#VUID-vkCreateSamplerYcbcrConversion-device-queuecount)VUID-vkCreateSamplerYcbcrConversion-device-queuecount  
  The device **must** have been created with at least `1` queue

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`
- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_VERSION\_1\_1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_1.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkSamplerYcbcrConversion](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerYcbcrConversion.html), [VkSamplerYcbcrConversionCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerYcbcrConversionCreateInfo.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCreateSamplerYcbcrConversion)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0