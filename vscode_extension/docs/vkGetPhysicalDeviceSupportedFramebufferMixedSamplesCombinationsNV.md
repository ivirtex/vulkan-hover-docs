# vkGetPhysicalDeviceSupportedFramebufferMixedSamplesCombinationsNV(3) Manual Page

## Name

vkGetPhysicalDeviceSupportedFramebufferMixedSamplesCombinationsNV - Query supported sample count combinations



## [](#_c_specification)C Specification

To query the set of mixed sample combinations of coverage reduction mode, rasterization samples and color, depth, stencil attachment sample counts that are supported by a physical device, call:

```c++
// Provided by VK_NV_coverage_reduction_mode
VkResult vkGetPhysicalDeviceSupportedFramebufferMixedSamplesCombinationsNV(
    VkPhysicalDevice                            physicalDevice,
    uint32_t*                                   pCombinationCount,
    VkFramebufferMixedSamplesCombinationNV*     pCombinations);
```

## [](#_parameters)Parameters

- `physicalDevice` is the physical device from which to query the set of combinations.
- `pCombinationCount` is a pointer to an integer related to the number of combinations available or queried, as described below.
- `pCombinations` is either `NULL` or a pointer to an array of [VkFramebufferMixedSamplesCombinationNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFramebufferMixedSamplesCombinationNV.html) values, indicating the supported combinations of coverage reduction mode, rasterization samples, and color, depth, stencil attachment sample counts.

## [](#_description)Description

If `pCombinations` is `NULL`, then the number of supported combinations for the given `physicalDevice` is returned in `pCombinationCount`. Otherwise, `pCombinationCount` **must** point to a variable set by the application to the number of elements in the `pCombinations` array, and on return the variable is overwritten with the number of values actually written to `pCombinations`. If the value of `pCombinationCount` is less than the number of combinations supported for the given `physicalDevice`, at most `pCombinationCount` values will be written to `pCombinations`, and `VK_INCOMPLETE` will be returned instead of `VK_SUCCESS`, to indicate that not all the supported values were returned.

Valid Usage (Implicit)

- [](#VUID-vkGetPhysicalDeviceSupportedFramebufferMixedSamplesCombinationsNV-physicalDevice-parameter)VUID-vkGetPhysicalDeviceSupportedFramebufferMixedSamplesCombinationsNV-physicalDevice-parameter  
  `physicalDevice` **must** be a valid [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html) handle
- [](#VUID-vkGetPhysicalDeviceSupportedFramebufferMixedSamplesCombinationsNV-pCombinationCount-parameter)VUID-vkGetPhysicalDeviceSupportedFramebufferMixedSamplesCombinationsNV-pCombinationCount-parameter  
  `pCombinationCount` **must** be a valid pointer to a `uint32_t` value
- [](#VUID-vkGetPhysicalDeviceSupportedFramebufferMixedSamplesCombinationsNV-pCombinations-parameter)VUID-vkGetPhysicalDeviceSupportedFramebufferMixedSamplesCombinationsNV-pCombinations-parameter  
  If the value referenced by `pCombinationCount` is not `0`, and `pCombinations` is not `NULL`, `pCombinations` **must** be a valid pointer to an array of `pCombinationCount` [VkFramebufferMixedSamplesCombinationNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFramebufferMixedSamplesCombinationNV.html) structures

Return Codes

On success, this command returns

- `VK_INCOMPLETE`
- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`
- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_NV\_coverage\_reduction\_mode](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_coverage_reduction_mode.html), [VkFramebufferMixedSamplesCombinationNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFramebufferMixedSamplesCombinationNV.html), [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetPhysicalDeviceSupportedFramebufferMixedSamplesCombinationsNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0