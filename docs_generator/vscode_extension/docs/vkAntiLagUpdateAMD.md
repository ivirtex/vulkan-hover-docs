# vkAntiLagUpdateAMD(3) Manual Page

## Name

vkAntiLagUpdateAMD - Provide information to reduce latency



## [](#_c_specification)C Specification

To lower latency, call:

```c++
// Provided by VK_AMD_anti_lag
void vkAntiLagUpdateAMD(
    VkDevice                                    device,
    const VkAntiLagDataAMD*                     pData);
```

## [](#_parameters)Parameters

- `device` is the logical device
- `pData` is a pointer to a [VkAntiLagDataAMD](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAntiLagDataAMD.html) structure containing latency reduction parameters.

## [](#_description)Description

This command should be executed immediately before the application processes user input. If `pData` is not `NULL` and [VkAntiLagDataAMD](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAntiLagDataAMD.html)::`pPresentationInfo` is not `NULL`, this command **should** be executed again before [vkQueuePresentKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkQueuePresentKHR.html), with `pPresentationInfo` set to matching values.

Valid Usage

- [](#VUID-vkAntiLagUpdateAMD-antiLag-10061)VUID-vkAntiLagUpdateAMD-antiLag-10061  
  The [`antiLag`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-antiLag) feature **must** be enabled

Valid Usage (Implicit)

- [](#VUID-vkAntiLagUpdateAMD-device-parameter)VUID-vkAntiLagUpdateAMD-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkAntiLagUpdateAMD-pData-parameter)VUID-vkAntiLagUpdateAMD-pData-parameter  
  `pData` **must** be a valid pointer to a valid [VkAntiLagDataAMD](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAntiLagDataAMD.html) structure

## [](#_see_also)See Also

[VK\_AMD\_anti\_lag](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_AMD_anti_lag.html), [VkAntiLagDataAMD](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAntiLagDataAMD.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkAntiLagUpdateAMD)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0