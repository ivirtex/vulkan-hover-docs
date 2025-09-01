# vkGetPhysicalDeviceFragmentShadingRatesKHR(3) Manual Page

## Name

vkGetPhysicalDeviceFragmentShadingRatesKHR - Get available shading rates for a physical device



## [](#_c_specification)C Specification

To query available shading rates, call:

```c++
// Provided by VK_KHR_fragment_shading_rate
VkResult vkGetPhysicalDeviceFragmentShadingRatesKHR(
    VkPhysicalDevice                            physicalDevice,
    uint32_t*                                   pFragmentShadingRateCount,
    VkPhysicalDeviceFragmentShadingRateKHR*     pFragmentShadingRates);
```

## [](#_parameters)Parameters

- `physicalDevice` is the handle to the physical device whose properties will be queried.
- `pFragmentShadingRateCount` is a pointer to an integer related to the number of fragment shading rates available or queried, as described below.
- `pFragmentShadingRates` is either `NULL` or a pointer to an array of [VkPhysicalDeviceFragmentShadingRateKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFragmentShadingRateKHR.html) structures.

## [](#_description)Description

If `pFragmentShadingRates` is `NULL`, then the number of fragment shading rates available is returned in `pFragmentShadingRateCount`. Otherwise, `pFragmentShadingRateCount` **must** point to a variable set by the application to the number of elements in the `pFragmentShadingRates` array, and on return the variable is overwritten with the number of structures actually written to `pFragmentShadingRates`. If `pFragmentShadingRateCount` is less than the number of fragment shading rates available, at most `pFragmentShadingRateCount` structures will be written, and `VK_INCOMPLETE` will be returned instead of `VK_SUCCESS`, to indicate that not all the available fragment shading rates were returned.

The returned array of fragment shading rates **must** be ordered from largest `fragmentSize.width` value to smallest, and each set of fragment shading rates with the same `fragmentSize.width` value **must** be ordered from largest `fragmentSize.height` to smallest. Any two entries in the array **must** not have the same `fragmentSize` values.

For any entry in the array, the following rules also apply:

- The value of `fragmentSize.width` **must** be less than or equal to [`maxFragmentSize.width`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-maxFragmentSize).
- The value of `fragmentSize.width` **must** be greater than or equal to `1`.
- The value of `fragmentSize.width` **must** be a power-of-two.
- The value of `fragmentSize.height` **must** be less than or equal to [`maxFragmentSize.height`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-maxFragmentSize).
- The value of `fragmentSize.height` **must** be greater than or equal to `1`.
- The value of `fragmentSize.height` **must** be a power-of-two.
- The highest sample count in `sampleCounts` **must** be less than or equal to [`maxFragmentShadingRateRasterizationSamples`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-maxFragmentShadingRateRasterizationSamples).
- The product of `fragmentSize.width`, `fragmentSize.height`, and the highest sample count in `sampleCounts` **must** be less than or equal to [`maxFragmentShadingRateCoverageSamples`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-maxFragmentShadingRateCoverageSamples).

Implementations **must** support at least the following shading rates:

  `sampleCounts` `fragmentSize`

`VK_SAMPLE_COUNT_1_BIT` | `VK_SAMPLE_COUNT_4_BIT`

{2,2}

`VK_SAMPLE_COUNT_1_BIT` | `VK_SAMPLE_COUNT_4_BIT`

{2,1}

~0

{1,1}

If [`framebufferColorSampleCounts`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-framebufferColorSampleCounts), includes `VK_SAMPLE_COUNT_2_BIT`, the required rates **must** also include `VK_SAMPLE_COUNT_2_BIT`.

Note

Including the {1,1} fragment size is done for completeness; it has no actual effect on the support of rendering without setting the fragment size. All sample counts and render pass transforms are supported for this rate.

The returned set of fragment shading rates **must** be returned in the native (rotated) coordinate system. For rasterization using render pass `transform` not equal to `VK_SURFACE_TRANSFORM_IDENTITY_BIT_KHR`, the application **must** transform the returned fragment shading rates into the current (unrotated) coordinate system to get the supported rates for that transform.

Note

For example, consider an implementation returning support for 4x2, but not 2x4 in the set of supported fragment shading rates. This means that for transforms `VK_SURFACE_TRANSFORM_ROTATE_90_BIT_KHR` and `VK_SURFACE_TRANSFORM_ROTATE_270_BIT_KHR`, 2x4 is a supported rate, but 4x2 is an unsupported rate.

Valid Usage (Implicit)

- [](#VUID-vkGetPhysicalDeviceFragmentShadingRatesKHR-physicalDevice-parameter)VUID-vkGetPhysicalDeviceFragmentShadingRatesKHR-physicalDevice-parameter  
  `physicalDevice` **must** be a valid [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html) handle
- [](#VUID-vkGetPhysicalDeviceFragmentShadingRatesKHR-pFragmentShadingRateCount-parameter)VUID-vkGetPhysicalDeviceFragmentShadingRatesKHR-pFragmentShadingRateCount-parameter  
  `pFragmentShadingRateCount` **must** be a valid pointer to a `uint32_t` value
- [](#VUID-vkGetPhysicalDeviceFragmentShadingRatesKHR-pFragmentShadingRates-parameter)VUID-vkGetPhysicalDeviceFragmentShadingRatesKHR-pFragmentShadingRates-parameter  
  If the value referenced by `pFragmentShadingRateCount` is not `0`, and `pFragmentShadingRates` is not `NULL`, `pFragmentShadingRates` **must** be a valid pointer to an array of `pFragmentShadingRateCount` [VkPhysicalDeviceFragmentShadingRateKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFragmentShadingRateKHR.html) structures

Return Codes

On success, this command returns

- `VK_INCOMPLETE`
- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_KHR\_fragment\_shading\_rate](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_fragment_shading_rate.html), [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html), [VkPhysicalDeviceFragmentShadingRateKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFragmentShadingRateKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetPhysicalDeviceFragmentShadingRatesKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0