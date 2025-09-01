# vkGetPhysicalDeviceOpticalFlowImageFormatsNV(3) Manual Page

## Name

vkGetPhysicalDeviceOpticalFlowImageFormatsNV - Query image formats for optical flow



## [](#_c_specification)C Specification

To enumerate the supported image formats for a specific optical flow usage, call:

```c++
// Provided by VK_NV_optical_flow
VkResult vkGetPhysicalDeviceOpticalFlowImageFormatsNV(
    VkPhysicalDevice                            physicalDevice,
    const VkOpticalFlowImageFormatInfoNV*       pOpticalFlowImageFormatInfo,
    uint32_t*                                   pFormatCount,
    VkOpticalFlowImageFormatPropertiesNV*       pImageFormatProperties);
```

## [](#_parameters)Parameters

- `physicalDevice` is the physical device being queried.
- []()`pOpticalFlowImageFormatInfo` is a pointer to a [VkOpticalFlowImageFormatInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpticalFlowImageFormatInfoNV.html) structure specifying the optical flow usage for which information is returned.
- []()`pFormatCount` is a pointer to an integer related to the number of optical flow properties available or queried, as described below.
- []()`pImageFormatProperties` is a pointer to an array of [VkOpticalFlowImageFormatPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpticalFlowImageFormatPropertiesNV.html) structures in which supported formats and image parameters are returned.

## [](#_description)Description

If `pImageFormatProperties` is `NULL`, then the number of optical flow properties supported for the given `physicalDevice` is returned in `pFormatCount`. Otherwise, `pFormatCount` **must** point to a variable set by the application to the number of elements in the `pImageFormatProperties` array, and on return the variable is overwritten with the number of values actually written to `pImageFormatProperties`. If the value of `pFormatCount` is less than the number of optical flow properties supported, at most `pFormatCount` values will be written to `pImageFormatProperties`, and `VK_INCOMPLETE` will be returned instead of `VK_SUCCESS`, to indicate that not all the available values were returned.

Before creating an image to be used as an optical flow frame, obtain the supported image creation parameters by querying with [vkGetPhysicalDeviceFormatProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFormatProperties2.html) and [vkGetPhysicalDeviceImageFormatProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceImageFormatProperties2.html) using one of the reported formats and adding [VkOpticalFlowImageFormatInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpticalFlowImageFormatInfoNV.html) to the `pNext` chain of [VkPhysicalDeviceImageFormatInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceImageFormatInfo2.html).

When querying the parameters with [vkGetPhysicalDeviceImageFormatProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceImageFormatProperties2.html) for images used for optical flow operations, the [VkOpticalFlowImageFormatInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpticalFlowImageFormatInfoNV.html)::`usage` field **must** contain one or more of the bits defined in [VkOpticalFlowUsageFlagBitsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpticalFlowUsageFlagBitsNV.html).

Valid Usage (Implicit)

- [](#VUID-vkGetPhysicalDeviceOpticalFlowImageFormatsNV-physicalDevice-parameter)VUID-vkGetPhysicalDeviceOpticalFlowImageFormatsNV-physicalDevice-parameter  
  `physicalDevice` **must** be a valid [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html) handle
- [](#VUID-vkGetPhysicalDeviceOpticalFlowImageFormatsNV-pOpticalFlowImageFormatInfo-parameter)VUID-vkGetPhysicalDeviceOpticalFlowImageFormatsNV-pOpticalFlowImageFormatInfo-parameter  
  `pOpticalFlowImageFormatInfo` **must** be a valid pointer to a valid [VkOpticalFlowImageFormatInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpticalFlowImageFormatInfoNV.html) structure
- [](#VUID-vkGetPhysicalDeviceOpticalFlowImageFormatsNV-pFormatCount-parameter)VUID-vkGetPhysicalDeviceOpticalFlowImageFormatsNV-pFormatCount-parameter  
  `pFormatCount` **must** be a valid pointer to a `uint32_t` value
- [](#VUID-vkGetPhysicalDeviceOpticalFlowImageFormatsNV-pImageFormatProperties-parameter)VUID-vkGetPhysicalDeviceOpticalFlowImageFormatsNV-pImageFormatProperties-parameter  
  If the value referenced by `pFormatCount` is not `0`, and `pImageFormatProperties` is not `NULL`, `pImageFormatProperties` **must** be a valid pointer to an array of `pFormatCount` [VkOpticalFlowImageFormatPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpticalFlowImageFormatPropertiesNV.html) structures

Return Codes

On success, this command returns

- `VK_INCOMPLETE`
- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_EXTENSION_NOT_PRESENT`
- `VK_ERROR_FORMAT_NOT_SUPPORTED`
- `VK_ERROR_INITIALIZATION_FAILED`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

Note

`VK_FORMAT_B8G8R8A8_UNORM`, `VK_FORMAT_R8_UNORM` and `VK_FORMAT_G8_B8R8_2PLANE_420_UNORM` are initially supported for images with [optical usage](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#opticalflow-usage) `VK_OPTICAL_FLOW_USAGE_INPUT_BIT_NV`.

`VK_FORMAT_R16G16_SFIXED5_NV` is initially supported for images with [optical flow usage](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#opticalflow-usage) `VK_OPTICAL_FLOW_USAGE_OUTPUT_BIT_NV`, `VK_OPTICAL_FLOW_USAGE_HINT_BIT_NV` and `VK_OPTICAL_FLOW_USAGE_GLOBAL_FLOW_BIT_NV`.

`VK_FORMAT_R8_UINT` and `VK_FORMAT_R32_UINT` are initially supported for images with [optical flow usage](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#opticalflow-usage) `VK_OPTICAL_FLOW_USAGE_COST_BIT_NV`. It is recommended to use `VK_FORMAT_R8_UINT` because of the lower bandwidth.

## [](#_see_also)See Also

[VK\_NV\_optical\_flow](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_optical_flow.html), [VkOpticalFlowImageFormatInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpticalFlowImageFormatInfoNV.html), [VkOpticalFlowImageFormatPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpticalFlowImageFormatPropertiesNV.html), [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetPhysicalDeviceOpticalFlowImageFormatsNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0